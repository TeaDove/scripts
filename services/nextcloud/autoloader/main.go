package main

import (
	"autoloader/davsupplier"
	"autoloader/loaderservice"
	"autoloader/settings"
	"net/http"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/service_utils/logger_utils"
)

func build() (*loaderservice.Service, error) {
	davSupplier, err := davsupplier.NewSupplier(
		http.DefaultClient,
		settings.Settings.DAVHostname,
		settings.Settings.DAVUsername,
		settings.Settings.DAVPassword,
		settings.Settings.DSTDir,
	)
	if err != nil {
		return nil, errors.Wrap(err, "davsupplier new supplier")
	}

	et, err := exiftool.NewExiftool()
	if err != nil {
		return nil, errors.Wrap(err, "new exiftool")
	}

	loaderService := loaderservice.NewService(
		davSupplier,
		et,
		settings.Settings.MountPath,
		settings.Settings.TMPDir,
		settings.Settings.FoldersToLoad,
	)

	return loaderService, nil
}

func main() {
	loaderService, err := build()
	if err != nil {
		panic(errors.Wrap(err, "build loaderservice"))
	}

	for {
		ctx := logger_utils.NewLoggedCtx()

		err = loaderService.Run(ctx)
		if err != nil {
			zerolog.Ctx(ctx).Error().
				Stack().Err(err).
				Msg("failed.to.run.loader")
		}

		time.Sleep(time.Minute)
	}
}
