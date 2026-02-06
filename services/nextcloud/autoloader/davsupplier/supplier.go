package davsupplier

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/emersion/go-webdav"
	"github.com/rs/zerolog"
)

type Suppler struct {
	baseFolder string

	client *webdav.Client
}

func NewSupplier(httpClient *http.Client, host, username, password, baseFolder string) (*Suppler, error) {
	davClient, err := webdav.NewClient(
		webdav.HTTPClientWithBasicAuth(httpClient, username, password),
		fmt.Sprintf("%s/remote.php/dav/files/%s/", host, username),
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating webdav Client")
	}

	return &Suppler{client: davClient, baseFolder: baseFolder}, nil
}

func HTTPErrorCode(err error) int {
	if err == nil {
		return 0
	}

	str := err.Error()

	fields := strings.Split(str, ": ")
	for _, field := range fields {
		if len(field) < 3 {
			continue
		}

		code, parseErr := strconv.Atoi(field[:3])
		if parseErr == nil {
			return code
		}
	}

	return 0
}

func (r *Suppler) pathJoin(elem ...string) string {
	elem = append([]string{r.baseFolder}, elem...)
	return filepath.Join(elem...)
}

const lastImgHashFile = ".last-img.md5"

func (r *Suppler) GetLastHash(ctx context.Context, cameraID, folder string) (string, error) {
	path := r.pathJoin(cameraID, folder, lastImgHashFile)

	file, err := r.client.Open(ctx, path)
	if err != nil {
		return "", errors.Wrap(err, "open")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", errors.Wrap(err, "read")
	}

	if len(content) == 0 {
		return "", errors.New("empty content")
	}

	return string(content), nil
}

func (r *Suppler) SetLastHash(ctx context.Context, cameraID, folder, hash string) error {
	path := r.pathJoin(cameraID, folder, lastImgHashFile)

	file, err := r.client.Create(ctx, path)
	if err != nil {
		return errors.Wrap(err, "create")
	}

	_, err = file.Write([]byte(hash))
	if err != nil {
		return errors.Wrap(err, "write")
	}

	err = file.Close()
	if err != nil {
		return errors.Wrap(err, "close")
	}

	return nil
}

func (r *Suppler) Stat(ctx context.Context, name string) (*webdav.FileInfo, error) {
	return r.client.Stat(ctx, r.pathJoin(name))
}

func (r *Suppler) Create(ctx context.Context, name string) (io.WriteCloser, error) {
	return r.client.Create(ctx, r.pathJoin(name))
}

func (r *Suppler) MkdirAll(ctx context.Context, name string) error {
	_, err := r.Stat(ctx, name)
	if err == nil {
		return nil
	}

	var curPath = ""

	for _, path := range strings.Split(r.pathJoin(name), "/") {
		curPath = filepath.Join(curPath, path)

		err := r.client.Mkdir(ctx, curPath)
		if err != nil && HTTPErrorCode(err) != http.StatusMethodNotAllowed {
			return errors.Wrap(err, "mkdir")
		}

		if err == nil {
			zerolog.Ctx(ctx).Info().
				Str("path", curPath).
				Msg("folder.created")
		}
	}

	return nil
}
