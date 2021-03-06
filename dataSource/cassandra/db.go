package cassandra

import "github.com/gocql/gocql"

var (
	cluster *gocql.ClusterConfig
)

func init() {
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "keystore"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
