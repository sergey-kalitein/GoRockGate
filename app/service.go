package app

import "rockgate/models"

// The Instance List of Application Services
var Services struct {
	OneSignalService *models.OneSignalService
}

func InitServices() {
	// Init the OneSignal service and load all available apps
	Services.OneSignalService = models.NewOneSignalService(Config())
	LoadOneSignalApps()
}

func LoadOneSignalApps() {
	_, err := Services.OneSignalService.LoadAppsList()
	if err != nil {
		Fatal("[OneSignal::ListApps]" + err.Error())
	}
}
