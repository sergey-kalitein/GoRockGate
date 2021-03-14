package main

import (
	"github.com/tbalthazar/onesignal-go"
	"net/http"
)

type OneSignalService struct {
	client *onesignal.Client
	config Configuration
}

func NewOneSignalService(conf Configuration) *OneSignalService {
	client := onesignal.NewClient(nil)
	client.UserKey = conf.OneSignalUserKey
	return &OneSignalService{client: client, config: conf}
}

func (o OneSignalService) ListApps() ([]onesignal.App, *http.Response, error) {
	return o.client.Apps.List()
}

func (o OneSignalService) GetApp(appID string) (*onesignal.App, *http.Response, error) {
	return o.client.Apps.Get(appID)
}

func (o OneSignalService) SetCurrentAppKey(appKey string) {
	o.client.AppKey = appKey
}

// The new App is created based on the domain of origin.
// All other settings are retrieved from the configuration.
func (o OneSignalService) CreateApp(domainOrigin string) (*onesignal.App, *http.Response, error) {
	appRequest := &onesignal.AppRequest{
		// TODO: implement the App Creation
		Name:                             "",
		GCMKey:                           "",
		ChromeKey:                        "",
		ChromeWebKey:                     "",
		ChromeWebOrigin:                  "",
		ChromeWebGCMSenderID:             "",
		ChromeWebDefaultNotificationIcon: "",
		ChromeWebSubDomain:               "",
		APNSEnv:                          "",
		APNSP12:                          "",
		APNSP12Password:                  "",
		SafariAPNSP12:                    "",
		SafariAPNSP12Password:            "",
		SafariSiteOrigin:                 "",
		SafariIcon1616:                   "",
		SafariIcon3232:                   "",
		SafariIcon6464:                   "",
		SafariIcon128128:                 "",
		SafariIcon256256:                 "",
		SiteName:                         "",
	}
	return o.client.Apps.Create(appRequest)
}
