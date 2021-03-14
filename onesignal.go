package main

import (
	"github.com/tbalthazar/onesignal-go"
	"net/http"
	"regexp"
	"strings"
)

type AppsBySite map[string]onesignal.App

type OneSignalService struct {
	client *onesignal.Client
	config Configuration
	apps   AppsBySite
}

func NewOneSignalService(conf Configuration) *OneSignalService {
	client := onesignal.NewClient(nil)
	client.UserKey = conf.OneSignalUserKey
	return &OneSignalService{client: client, config: conf}
}

// Strips protocol from the domain name
func StripDomainName(webOriginDomain string) string {
	domainStripped := strings.TrimSpace(strings.ToLower(webOriginDomain))
	return regexp.MustCompile(`^(?:http|https)://(.*?)`).ReplaceAllString(domainStripped, `$1`)
}

func (o *OneSignalService) ListApps() ([]onesignal.App, *http.Response, error) {
	return o.client.Apps.List()
}

func (o *OneSignalService) GetApps() (*AppsBySite, error) {
	if o.apps == nil {
		_, err := o.LoadAppsList()
		if err != nil {
			return nil, err
		}
	}
	return &o.apps, nil
}

func (o *OneSignalService) FindAppOrCreate(webOriginDomain string) (*onesignal.App, error) {
	app, isFound := o.apps[StripDomainName(webOriginDomain)]
	if !isFound {
		// Let's try to create an app if we are not able
		// to find it in our app "registry"
		createdApp, _, err := o.CreateApp(webOriginDomain)
		if err != nil {
			return nil, err
		}
		return createdApp, nil
	}
	return &app, nil
}

// Loads a list of available OneSignal apps
func (o *OneSignalService) LoadAppsList() (AppsBySite, error) {
	apps, _, err := o.ListApps()
	if err != nil {
		return nil, err
	}
	o.apps = make(AppsBySite, 0)
	for _, app := range apps {
		// `ChromeWebOrigin` is the only domain-related identity
		o.apps[StripDomainName(app.ChromeWebOrigin)] = app
	}

	return o.apps, nil
}

func (o *OneSignalService) GetApp(appID string) (*onesignal.App, *http.Response, error) {
	return o.client.Apps.Get(appID)
}

func (o *OneSignalService) SetCurrentAppKey(appKey string) {
	o.client.AppKey = appKey
}

// The new App is created based on the domain of origin.
// All other settings are retrieved from the configuration.
func (o *OneSignalService) CreateApp(domainOrigin string) (*onesignal.App, *http.Response, error) {
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

func (o *OneSignalService) SendNotification(notificationRequest *onesignal.NotificationRequest) (*onesignal.NotificationCreateResponse, error) {
	response, _, err := o.client.Notifications.Create(notificationRequest)
	return response, err
}
