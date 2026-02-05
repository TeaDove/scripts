package main

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
	"github.com/emersion/go-webdav"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/service_utils/logger_utils"
)

const (
	path   = "https://pi.lan/remote.php/dav/files/%s/"
	dstDir = "autoloader"
)

func run(ctx context.Context, username, password string) error {
	client, err := webdav.NewClient(webdav.HTTPClientWithBasicAuth(&http.Client{}, username, password), fmt.Sprintf(path, username))
	if err != nil {
		return errors.Wrap(err, "creating webdav client")
	}

	_, err = client.Stat(context.Background(), dstDir)
	if err != nil {
		err = client.Mkdir(ctx, dstDir)
		if err != nil {
			return errors.Wrap(err, "creating directory")
		}
	}

	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info == nil || info.IsDir() || info.Size() == 0 {
			return nil
		}

		logger := zerolog.Ctx(ctx).With().Str("path", path).Logger()

		logger.Info().Msg("uploading.file")

		dst := filepath.Join(dstDir, info.Name())

		stat, err := client.Stat(ctx, dst)
		if err == nil && stat.Size != 0 {
			logger.Info().Msg("already.exists")
			return nil
		}

		writer, err := client.Create(ctx, dst)
		if err != nil {
			return errors.Wrap(err, "creating file")
		}

		file, err := os.Open(path)
		if err != nil {
			return errors.Wrap(err, "opening file")
		}

		_, err = file.WriteTo(writer)
		if err != nil {
			return errors.Wrap(err, "copying file")
		}

		err = writer.Close()
		if err != nil {
			return errors.Wrap(err, "closing file")
		}

		logger.Info().Msg("uploaded")

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "walking directory")
	}

	return nil
}

func main() {
	fmt.Printf("Enter username: ")
	var (
		username string
		password string
	)

	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Printf("%s\n", errors.Wrap(err, "failed to read username"))
	}

	fmt.Printf("Enter password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Printf("%s\n", errors.Wrap(err, "failed to read password"))
	}

	err = run(logger_utils.NewLoggedCtx(), username, password)
	if err != nil {
		fmt.Printf("%s\n", errors.Wrap(err, "failed to run autoloader"))
	}
}
