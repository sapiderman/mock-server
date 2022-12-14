package config

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defCfg      map[string]string
	initialized = false
)

// LoadConfig loads configuration file
func LoadConfig() error {
	logf := log.WithField("fn", "LoadConfig")
	logf.Info("loading config...")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("env")
	viper.SetConfigFile("./.env") // define config .env file to read
	viper.AutomaticEnv()          // read also from environment variables
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		logf.Warn("config file not found, got err:", err)
	} else {
		logf.Info("config file found and loaded.")
	}

	defCfg["app.id"] = "mock-server-go"
	defCfg["app.version"] = "0.0.0"
	defCfg["app.env"] = "production" // set to production / development

	defCfg["server.host"] = "localhost"
	defCfg["server.port"] = "7000"
	defCfg["server.log.level"] = "debug" // valid values are trace, debug, info, warn, error, fatal
	defCfg["server.timeout.write"] = "15 seconds"
	defCfg["server.timeout.read"] = "15 seconds"
	defCfg["server.timeout.idle"] = "60 seconds"
	defCfg["server.timeout.graceshut"] = "15 seconds"

	defCfg["server.context.timeout"] = "30" // seconds

	return nil
}

// SetConfig put configuration key value
func SetConfig(key, value string) {
	viper.Set(key, value)
}

// Get fetch configuration as string value
func Get(key string) string {
	if !initialized {
		LoadConfig()
	}
	ret := viper.GetString(key)
	if len(ret) == 0 {
		if ret, ok := defCfg[key]; ok {
			return ret
		}
		log.Debugf("%s config key not found", key)
	}
	return ret
}

// GetBoolean fetch configuration as boolean value
func GetBoolean(key string) bool {
	if len(Get(key)) == 0 {
		return false
	}
	b, err := strconv.ParseBool(Get(key))
	if err != nil {
		panic(err)
	}
	return b
}

// GetInt fetch configuration as integer value
func GetInt(key string) int {
	if len(Get(key)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(Get(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// GetFloat fetch configuration as float value
func GetFloat(key string) float64 {
	if len(Get(key)) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(Get(key), 64)
	if err != nil {
		panic(err)
	}
	return f
}

// Set configuration key value
func Set(key, value string) {
	defCfg[key] = value
}
