package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

// This configuration items must be defined in order to be able
// to create new One Signal apps
type Configuration struct {
	OneSignalUserKey      string
	RestApiKey            string
	GCMKey                string
	AndroidGCMSenderID    string
	ChromeKey             string
	ChromeWebKey          string
	ChromeWebOrigin       string
	ChromeWebGCMSenderID  string
	APNSEnv               string
	APNSP12               string
	APNSP12Password       string
	SafariAPNSP12         string
	SafariAPNSP12Password string
}

var config *Configuration

var configWarnings []string

// Whether we need to log the incoming push messages
// and the output response payload
var IsLoggingPayloadEnabled bool

func LoadConfiguration() ([]string, error) {
	configFriendlyName := "config/rockgate.yml"
	v := viper.New()
	v.SetConfigName("rockgate")
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf")
	err := v.ReadInConfig()
	if err != nil {
		return []string{}, err
	}

	config = &Configuration{}
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[%s] %s", configFriendlyName, err.Error()))
	}

	configWarnings := collectConfigWarnings(configFriendlyName)
	// Errors must be at the very end
	if config.OneSignalUserKey == "" {
		return configWarnings, errors.New(fmt.Sprintf("[%s] the OneSignal User Key is undefined "+
			"(see the 'OneSignalUserKey' config parameter)", configFriendlyName))
	}
	IsLoggingPayloadEnabled = v.GetBool("LoggingPayloadEnabled")

	if len(configWarnings) > 0 {
		return configWarnings, nil
	} else {
		return []string{}, nil
	}
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
func Config() Configuration {
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
