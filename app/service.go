package app

import "rockgate/models"

var OneSignalService *models.OneSignalService

func InitServices() {
	// Init the OneSignal service and load all available apps
	OneSignalService = models.NewOneSignalService(Config())
	LoadOneSignalApps()
}

func LoadOneSignalApps() {
	_, err := OneSignalService.LoadAppsList()
	if err != nil {
		Fatal("[OneSignal::ListApps]" + err.Error())
	}
}
