package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// This configuration items must be defined in order to be able
// to create new One Signal apps
type Configuration struct {
	GCMKey                string `json:"gcm_key,omitempty"`
	ChromeKey             string `json:"chrome_key,omitempty"`
	ChromeWebKey          string `json:"chrome_web_key,omitempty"`
	ChromeWebOrigin       string `json:"chrome_web_origin,omitempty"`
	ChromeWebGCMSenderID  string `json:"chrome_web_gcm_sender_id,omitempty"`
	APNSEnv               string `json:"apns_env,omitempty"`
	APNSP12               string `json:"apns_p12,omitempty"`
	APNSP12Password       string `json:"apns_p12_password,omitempty"`
	SafariAPNSP12         string `json:"safari_apns_p12,omitempty"`
	SafariAPNSP12Password string `json:"safari_apns_p12_password,omitempty"`
}

var config *Configuration

var configWarnings []string

func LoadConfiguration() ([]string, error) {
	viper.SetConfigName("rockgate")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return []string{}, err
	}

	config = &Configuration{
		// Android
		GCMKey: viper.GetString("GCMKey"),
		// Chrome Web Push
		ChromeKey:            viper.GetString("ChromeKey"),
		ChromeWebKey:         viper.GetString("ChromeWebKey"),
		ChromeWebOrigin:      viper.GetString("ChromeWebOrigin"),
		ChromeWebGCMSenderID: viper.GetString("ChromeWebGCMSenderID"),
		// Apple Push Notifications
		APNSEnv:         viper.GetString("APNSEnv"),
		APNSP12:         viper.GetString("APNSP12"),
		APNSP12Password: viper.GetString("APNSP12Password"),
		// Safari Web Push
		SafariAPNSP12:         viper.GetString("SafariAPNSP12"),
		SafariAPNSP12Password: viper.GetString("SafariAPNSP12Password"),
	}

	if configWarnings == nil {
		configWarnings = make([]string, 0)
	}
	errorTemplate := "'%s' configuration parameter is missing"
	switch {
	case config.GCMKey == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "GCMKey"))

	case config.ChromeKey == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "ChromeKey"))

	case config.ChromeWebKey == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "ChromeWebKey"))

	case config.ChromeWebOrigin == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "ChromeWebOrigin"))

	case config.ChromeWebGCMSenderID == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "ChromeWebGCMSenderID"))

	case config.APNSEnv == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "APNSEnv"))

	case config.APNSP12 == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "APNSP12"))

	case config.APNSP12Password == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "APNSP12Password"))

	case config.SafariAPNSP12 == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "SafariAPNSP12"))

	case config.SafariAPNSP12Password == "":
		configWarnings = append(configWarnings, fmt.Sprintf(errorTemplate, "SafariAPNSP12Password"))

	}

	if len(configWarnings) > 0 {
		return configWarnings, nil
		//errors.New("inconsistent configuration, " +
		//	"certain functions may be unavailable, " +
		//	"e.g. OneSignal App Creation etc")
	} else {
		return []string{}, nil
	}
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
