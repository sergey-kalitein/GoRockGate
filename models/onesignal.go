package models

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/tbalthazar/onesignal-go"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	NotificationTypeAPN = "apn"
	NotificationTypeGCM = "gcm"
	NotificationTypeWeb = "web"
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

// The REST Auth Key must be set for every call related to an App.
// See also the `BasicAuthKey` property of the `onesignal.App` struct.
func (o *OneSignalService) SetAppRestAuthKey(appRestAuthKey string) {
	o.client.AppKey = appRestAuthKey
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
	appName := StripDomainName(webOriginDomain)
	log.Printf("Finding App: '%s'\n", appName)
	app, isFound := o.apps[appName]
	if !isFound {
		log.Print(color.New(color.BgYellow).Sprintf("[WARNING] App %s not found, creating a new one..\n", appName))
		// Let's try to create an app if we are not able
		// to find it in our app "registry"
		createdApp, _, err := o.CreateApp(webOriginDomain)
		if err != nil {
			return nil, err
		}
		return createdApp, nil
	} else {
		log.Printf("App '%s' has been found\n", appName)
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
		identity := app.ChromeWebOrigin
		if identity == "" {
			// All the apps created by this tool have the same name
			// as the domain itself
			identity = app.Name
		}
		o.apps[StripDomainName(identity)] = app
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
	appName := StripDomainName(domainOrigin)
	appRequest := &onesignal.AppRequest{
		// TODO: add missing field values?
		Name:                             appName,
		GCMKey:                           o.config.GCMKey,
		ChromeKey:                        o.config.ChromeKey,
		ChromeWebKey:                     o.config.ChromeWebKey,
		ChromeWebOrigin:                  domainOrigin,
		ChromeWebGCMSenderID:             o.config.ChromeWebGCMSenderID,
		ChromeWebDefaultNotificationIcon: "",
		ChromeWebSubDomain:               "",
		APNSEnv:                          o.config.APNSEnv,
		APNSP12:                          o.config.APNSP12,
		APNSP12Password:                  o.config.APNSP12Password,
		SafariAPNSP12:                    o.config.SafariAPNSP12,
		SafariAPNSP12Password:            o.config.SafariAPNSP12Password,
		SafariSiteOrigin:                 domainOrigin,
		SafariIcon1616:                   "public/safari_packages/469c9b1c-86f2-43ec-98d5-35dd74c10f80/icons/16x16.png",
		SafariIcon3232:                   "public/safari_packages/469c9b1c-86f2-43ec-98d5-35dd74c10f80/icons/16x16@2x.png",
		SafariIcon6464:                   "public/safari_packages/469c9b1c-86f2-43ec-98d5-35dd74c10f80/icons/32x32@2x.png",
		SafariIcon128128:                 "public/safari_packages/469c9b1c-86f2-43ec-98d5-35dd74c10f80/icons/128x128.png",
		SafariIcon256256:                 "public/safari_packages/469c9b1c-86f2-43ec-98d5-35dd74c10f80/icons/128x128@2x.png",
		SiteName:                         "'" + appName + "' website",
	}
	jsonText, _ := json.MarshalIndent(*appRequest, "", "    ")
	log.Println(color.New(color.FgHiGreen).Sprintf("Creating a new app:\n %s", string(jsonText)))

	// TODO: implement direct API call to pass the `Android App ID`
	return o.client.Apps.Create(appRequest)
}

func (o *OneSignalService) SendNotification(notificationRequest *onesignal.NotificationRequest) (*onesignal.NotificationCreateResponse, error) {
	response, _, err := o.client.Notifications.Create(notificationRequest)
	return response, err
}
