package configuration

import "github.com/deltegui/configloader"

//Configuration representation of json config file
type Configuration struct {
	ListenURL      string `paramName:"url"`
	DatabaseDriver string `paramName:"dbdriver"`
	Database       string `paramName:"dbname"`
	RabbitMQ       string `paramName:"rabbit"`
}

//Load configuration from config.json file and overwrite
//default values if console params are provided
func Load() Configuration {
	return *configloader.NewConfigLoaderFor(&Configuration{}).
		AddHook(configloader.CreateFileHook("./config.json")).
		AddHook(configloader.CreateParamsHook()).
		Retrieve().(*Configuration)
}
