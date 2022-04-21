package application

import (
	keyscontroller "test3/hariprathap-hp/DesignTinyURL/tinyURL_KGS_API/controller/keys"
	pingcontroller "test3/hariprathap-hp/DesignTinyURL/tinyURL_KGS_API/controller/ping"
)

func mapUrls() {
	router.GET("/ping", pingcontroller.Ping)
	router.GET("/populate", keyscontroller.PopulateKeys)
	router.GET("/getKeys", keyscontroller.GetKeys)
}
