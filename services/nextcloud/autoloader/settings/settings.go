package settings

import "github.com/teadove/teasutils/service_utils/settings_utils"

type AppSettings struct {
	MountPath     string   `env:"MOUNT_PATH" envDefault:"/media/teadove"`
	TMPDir        string   `env:"TMP_DIR" envDefault:"/tmp/autoloader-temporal-directory"`
	FoldersToLoad []string `env:"FOLDER_TO_LOAD" envDefault:"DCIM,SPIDCIM"`

	DAVUsername string `env:"DAV_USERNAME"`
	DAVPassword string `env:"DAV_PASSWORD"`
	DAVHostname string `env:"DAV_HOSTNAME" envDefault:"https://pi.lan"`

	DSTDir string `env:"DST_DIR" envDefault:"autoloader"`
}

var Settings = settings_utils.MustGetSetting[AppSettings]("AUTOLOADER_")
