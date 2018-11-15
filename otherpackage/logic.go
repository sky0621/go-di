package otherpackage

import (
	app "github.com/sky0621/go-di"
)

// Logic ...
func Logic(dicon *app.DIContainer) {
	cloudSQLAccessor := dicon.GetAccessor(app.CloudSQLAccessor)
	cloudSQLAccessor.Duck()

	cloudPubSubAccessor := dicon.GetAccessor(app.CloudPubSubAccessor)
	cloudPubSubAccessor.Duck()
}
