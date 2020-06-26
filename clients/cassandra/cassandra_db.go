package cassandra

import "github.com/gocql/gocql"

var (
	session *gocql.Session
)

func init() {
	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "pincode"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
