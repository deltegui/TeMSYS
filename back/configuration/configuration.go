package configuration

import (
	"github.com/deltegui/configloader"
)

//Configuration representation of json config file
type Configuration struct {
	ListenURL      string `paramName:"url"`
	JWTKey         string `paramName:"jwtkey"`
	DatabaseDriver string `paramName:"dbdriver"`
	Database       string `paramName:"dbname"`
	TLSCRT         string `paramName:"tlscrt"`
	TLSKEY         string `paramName:"tlskey"`
	TLSEnabled     bool   `paramName:"tlsenabled"`
	CORS           string `paramName:"CORS"`
}

// Load from dev.json file and overwrite
// default values if console params are provided
func Load() Configuration {
	return *configloader.NewConfigLoaderFor(&Configuration{}).
		AddHook(configloader.CreateFileHook("./dev.json")).
		AddHook(configloader.CreateParamsHook()).
		AddHook(configloader.CreateEnvHook()).
		Retrieve().(*Configuration)
}
