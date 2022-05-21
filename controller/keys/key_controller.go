package keys

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hariprathap-hp/tinyURL_KGS_API/domain"
	"github.com/hariprathap-hp/tinyURL_KGS_API/services/keyservices"
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
	keys := domain.UniqKeys{
		Keys: results,
	}
	c.JSON(http.StatusOK, keys)
}
