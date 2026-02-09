package loaderservice

import (
	"autoloader/davsupplier"
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/service_utils/logger_utils"
)

type Service struct {
	davSupplier *davsupplier.Suppler
	et          *exiftool.Exiftool

	mountPath     string
	tmpDir        string
	foldersToLoad []string
}

func NewService(
	davSupplier *davsupplier.Suppler,
	et *exiftool.Exiftool,
	mountPath string,
	tmpDir string,
	foldersToLoad []string,
) *Service {
	return &Service{davSupplier: davSupplier, et: et, mountPath: mountPath, tmpDir: tmpDir, foldersToLoad: foldersToLoad}
}

func (r *Service) Run(ctx context.Context) error {
	mountDir, err := os.ReadDir(r.mountPath)
	if err != nil {
		return errors.Wrap(err, "read mount dir")
	}

	for _, entry := range mountDir {
		logger := zerolog.Ctx(ctx).With().Str("folders", entry.Name()).Logger()
		logger.Debug().Msg("checking.folder")
		if !entry.IsDir() {
			logger.Debug().Msg("not.dir")
			continue
		}

		cameraDir, err := os.ReadDir(filepath.Join(r.mountPath, entry.Name()))
		if err != nil {
			return errors.Wrap(err, "read camera dir")
		}

		logger.Debug().Interface("dirs", cameraDir).Msg("checking.camera")

		for _, cameraEntry := range cameraDir {
			logger.Debug().Str("entry", cameraEntry.Name()).Msg("checking.camera.dir")
			if !slices.Contains(r.foldersToLoad, cameraEntry.Name()) {
				logger.Debug().Str("entry", cameraEntry.Name()).Msg("wrong.name")
				continue
			}

			err = r.uploadFolder(ctx, entry.Name(), cameraEntry.Name(), filepath.Join(r.mountPath, entry.Name(), cameraEntry.Name()))
			if err != nil {
				return errors.Wrap(err, "upload folder")
			}
		}
	}

	return nil
}

func (r *Service) uploadFolder(ctx context.Context, cameraID string, photosFolder, fullpath string) error {
	ctx = logger_utils.WithValue(ctx, "path", fullpath)

	zerolog.Ctx(ctx).Debug().Msg("uploading.folder")

	files, err := os.ReadDir(fullpath)
	if err != nil {
		return errors.Wrap(err, "read dir")
	}

	if len(files) == 0 {
		zerolog.Ctx(ctx).Info().Msg("no files to upload")
		return nil
	}

	slices.Reverse(files)

	lastFileHash, err := getLastFileHash(filepath.Join(fullpath, files[0].Name()))
	if err != nil {
		return errors.Wrap(err, "get last file hash")
	}

	zerolog.Ctx(ctx).Debug().Str("hash", lastFileHash).Msg("last.file.hash")

	lastUploadedHash, err := r.davSupplier.GetLastHash(ctx, cameraID, photosFolder)
	if err != nil && davsupplier.HTTPErrorCode(err) != http.StatusNotFound {
		return errors.Wrap(err, "get last uploaded hash")
	}

	var count int

	for _, file := range files {
		next, err := r.uploadFile(ctx, cameraID, photosFolder, fullpath, file, lastUploadedHash)
		if err != nil {
			return errors.Wrap(err, "upload file")
		}

		if !next {
			zerolog.Ctx(ctx).Debug().Msg("found.already.uploaded")
			break
		}

		count++
	}

	err = r.davSupplier.SetLastHash(ctx, cameraID, photosFolder, lastFileHash)
	if err != nil {
		return errors.Wrap(err, "set last file hash")
	}

	if count == 0 {
		zerolog.Ctx(ctx).
			Info().
			Msg("no.files.changed")
	} else {
		zerolog.Ctx(ctx).
			Info().
			Int("count", count).
			Msg("files.uploaded")
	}

	return nil
}

func (r *Service) uploadFile(ctx context.Context, cameraID string, photosFolder, path string, file os.DirEntry, lastUploadedHash string) (bool, error) {
	fd, err := os.Open(filepath.Join(path, file.Name()))
	if err != nil {
		return false, errors.Wrap(err, "open file")
	}
	defer fd.Close()

	fileContent, err := io.ReadAll(fd)
	if err != nil {
		return false, errors.Wrap(err, "read file")
	}

	if lastUploadedHash != "" && lastUploadedHash == getContentHash(fileContent) {
		zerolog.Ctx(ctx).Debug().
			Str("file", file.Name()).
			Msg("found.uploaded.file")

		return false, nil
	}

	createdAt, err := r.getImageDate(filepath.Join(path, file.Name()))
	if err != nil {
		zerolog.Ctx(ctx).Warn().
			Err(err).
			Str("file", file.Name()).
			Msg("failed.to.get.image.date")
	}

	dstFolder := compileDSTFolder(cameraID, photosFolder, createdAt)

	err = r.davSupplier.MkdirAll(ctx, dstFolder)
	if err != nil {
		return false, errors.Wrap(err, "mkdir")
	}

	writer, err := r.davSupplier.Create(ctx, filepath.Join(dstFolder, file.Name()))
	if err != nil {
		return false, errors.Wrap(err, "create file")
	}

	_, err = writer.Write(fileContent)
	if err != nil {
		return false, errors.Wrap(err, "copying file")
	}

	err = writer.Close()
	if err != nil {
		return false, errors.Wrap(err, "close writer")
	}

	zerolog.Ctx(ctx).Info().Str("file", file.Name()).Msg("uploaded.file")

	return true, nil
}

func compileDSTFolder(cameraID, photosFolder string, createdAt time.Time) string {
	if createdAt.IsZero() {
		return filepath.Join(cameraID, photosFolder)
	}

	return filepath.Join(cameraID, photosFolder, createdAt.Format("2006-01"))
}

func getLastFileHash(name string) (string, error) {
	byteArr, err := os.ReadFile(name)
	if err != nil {
		return "", errors.Wrap(err, "read file")
	}

	return getContentHash(byteArr), nil
}

func getContentHash(content []byte) string {
	hash := md5.Sum(content)
	return hex.EncodeToString(hash[:])
}
