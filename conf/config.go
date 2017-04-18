package conf

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config the application's configuration
type Config struct {
	Port      int64         `json:"port"`
	Host      string        `json:"listen"`
	LogConfig LoggingConfig `json:"log_config"`
}

// LoadConfig loads the config from a file if specified, otherwise from the environment
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("echo-server")
		viper.AddConfigPath("./")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// config := new(Config)

	// ko := viper.Unmarshal(&config)
	// if ko != nil {
	// 	return nil, ko
	// }

	return populateConfig(new(Config))
}
