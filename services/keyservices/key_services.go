package keyservices

import (
	"github.com/hariprathap-hp/tinyURL_KGS_API/domain"
	"github.com/hariprathap-hp/utils_repo/errors"
)

func PopulateKeys() *errors.RestErr {
	var key domain.Key
	if err := key.Populate(true, ""); err != nil {
		return err
	}
	return nil
}

func GetKeys() ([]string, *errors.RestErr) {
	var key domain.Key
	results, err := key.Get(50)
	if err != nil {
		return nil, err
	}
	return results, nil
}
