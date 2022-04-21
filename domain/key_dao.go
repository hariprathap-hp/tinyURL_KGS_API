package keymodels

import (
	"fmt"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_KGS_API/dataSource/cassandra"
	"test3/hariprathap-hp/system_design/utils_repo/errors"

	"github.com/gocql/gocql"
)

const (
	queryPopulateTableAvail = "INSERT INTO keys_avail (token_id) values (?)"
	queryPopulateTableUsed  = "INSERT INTO keys_used (token_id) values (?)"
	queryDeleteKeys         = "DELETE FROM keys_avail WHERE token_id IN ?"
	queryGetKeys            = "SELECT token_id from keys_avail limit ?"
)

var (
	session *gocql.Session
)

func init() {
	var err error
	session, err = cassandra.GetSession()
	if err != nil {
		fmt.Println("creating cassandra session failed")
		panic(err)
	}
}

func (k *Key) Populate(keys_avail bool, key string) *errors.RestErr {
	//insert keys into the table keys_available as well as keys_used table
	if keys_avail {
		for i := 0; i < 25000; i++ {
			id := getID()
			if err := session.Query(queryPopulateTableAvail, id).Exec(); err != nil {
				return errors.NewInternalServerError(err.Error())
			}
		}
	} else {
		if err := session.Query(queryPopulateTableUsed, key).Exec(); err != nil {
			return errors.NewInternalServerError(err.Error())
		}
	}
	return nil
}

func (k *Key) Get(limit int) ([]string, *errors.RestErr) {
	//provides keys to the URL service from KGS
	scanner := session.Query(queryGetKeys, limit).Iter().Scanner()

	var results []string
	var token string

	for scanner.Next() {
		if err := scanner.Scan(&token); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, token)
	}

	if err := delete(results); err != nil {
		return nil, err
	}

	for i := 0; i < limit; i++ {
		if err := k.Populate(false, results[i]); err != nil {
			return nil, err
		}
	}
	return results, nil
}

func delete(keys []string) *errors.RestErr {
	//To remove the used keys from keys_available table
	if err := session.Query(queryDeleteKeys, keys).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
