package app

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"rockgate/models"
)

const (
	EnvironmentPrefix = "RG"
	ConfigName        = "rockgate"
	ConfigType        = "yaml"
	DefaultConfigPath = "./conf"
)

var config *models.Configuration

var configWarnings []string

func LoadConfiguration() ([]string, error) {
	configFriendlyName := "config/rockgate.yml"

	err := viperReadConfig()
	if err != nil {
		return []string{}, err
	}

	config = &models.Configuration{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[%s] %s", configFriendlyName, err.Error()))
	}

	configWarnings := collectConfigWarnings(configFriendlyName)
	// Errors must be at the very end
	if config.OneSignalUserKey == "" {
		return configWarnings, errors.New(fmt.Sprintf("[%s] the OneSignal User Key is undefined "+
			"(see the 'OneSignalUserKey' config parameter)", configFriendlyName))
	}

	if len(configWarnings) > 0 {
		return configWarnings, nil
	} else {
		return []string{}, nil
	}
}

func viperReadConfig() error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(EnvironmentPrefix)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(DefaultConfigPath)
	err := viper.ReadInConfig()
	return err
}

func collectConfigWarnings(configFriendlyName string) []string {
	if configWarnings == nil {
		configWarnings = make([]string, 0)
	}
	warningTemplate := "[" + configFriendlyName + "]" + " '%s' configuration parameter is missing"

	if config.GCMKey == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "GCMKey"))
	}
	if config.AndroidGCMSenderID == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "AndroidGCMSenderID"))
	}
	if config.ChromeWebKey == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "ChromeWebKey"))
	}
	//if config.ChromeWebOrigin == "" {
	//	configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "ChromeWebOrigin"))
	//}
	if config.ChromeWebGCMSenderID == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "ChromeWebGCMSenderID"))
	}
	if config.APNSEnv == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "APNSEnv"))
	}
	if config.APNSP12 == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "APNSP12"))
	}
	if config.APNSP12Password == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "APNSP12Password"))
	}
	if config.SafariAPNSP12 == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "SafariAPNSP12"))
	}
	if config.SafariAPNSP12Password == "" {
		configWarnings = append(configWarnings, fmt.Sprintf(warningTemplate, "SafariAPNSP12Password"))
	}

	return configWarnings
}

// Returns the Gateway configuration.
// If any of these parameters are undefined then certain functions of the gateway
// may NOT work properly, let's say creation of new Apps etc.
func Config() models.Configuration {
	if config == nil {
		LoadConfiguration()
	}
	return *config
}

func ConfigWarnings() []string {
	if configWarnings == nil {
		LoadConfiguration()
	}
	return configWarnings
}

func Configure() {
	configWarnings, err := LoadConfiguration()
	if err != nil {
		Fatal(err.Error())
	}
	if len(configWarnings) > 0 {
		for _, warningText := range configWarnings {
			color.New(color.FgBlack, color.BgYellow).Printf("[CONFIG WARNING] %s", warningText)
			fmt.Println()
		}
	}
}

// Whether we need to log the incoming push messages
// and the output response payload?
func IsLoggingPayloadEnabled() bool {
	return viper.GetBool("LOG_PAYLOAD_ENABLED")
}
