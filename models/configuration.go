package models

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
