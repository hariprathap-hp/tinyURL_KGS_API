package keyservices

import (
	keymodels "test3/hariprathap-hp/DesignTinyURL/tinyURL_KGS_API/domain"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
)

func PopulateKeys() *errors.RestErr {
	var key keymodels.Key
	if err := key.Populate(true, ""); err != nil {
		return err
	}
	return nil
}

func GetKeys() ([]string, *errors.RestErr) {
	var key keymodels.Key
	results, err := key.Get(50)
	if err != nil {
		return nil, err
	}
	return results, nil
}
