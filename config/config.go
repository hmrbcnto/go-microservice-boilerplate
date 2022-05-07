package config

import "github.com/spf13/viper"

type Config struct {
	DbConfig DbConfig
}

// Exposed method

func LoadConfig() (*Config, error) {
	dbConf := loadDbConfig()

	err := dbConf.validate()

	if err != nil {
		return nil, err
	}

	return &Config{
		DbConfig: dbConf,
	}, nil
}

// Loads dbConfig
func loadDbConfig() DbConfig {
	viper.SetEnvPrefix("MONGO")
	viper.BindEnv("URI")

	return DbConfig{
		MongoURI: viper.GetString("URI"),
	}
}
