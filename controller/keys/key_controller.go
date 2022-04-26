package keyscontroller

import (
	"fmt"
	"net/http"
	keymodels "test3/hariprathap-hp/DesignTinyURL/tinyURL_KGS_API/domain"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_KGS_API/services/keyservices"

	"github.com/gin-gonic/gin"
)

func PopulateKeys(c *gin.Context) {
	err := keyservices.PopulateKeys()
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, "keys are loaded into cassandra")
}

func GetKeys(c *gin.Context) {
	results, err := keyservices.GetKeys()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var keys keymodels.Results
	keys.Uniq_keys = results
	fmt.Printf("%T\n", keys.Uniq_keys)
	c.JSON(http.StatusOK, keys)
}
