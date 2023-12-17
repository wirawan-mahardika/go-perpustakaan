package app

import "github.com/spf13/viper"

func ViperConfig(name, configType, path string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(name)
	config.SetConfigType(configType)
	config.AddConfigPath(path)

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
