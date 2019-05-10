package databases

import (
	"github.com/gocql/gocql"
	"time"
)

func ConnectToDatabase() (error , *gocql.Session,*gocql.KeyspaceMetadata) {
	cluster := gocql.NewCluster("127.0.0.1")
	// cluster.Authenticator = gocql.PasswordAuthenticator{
    //     Username: "cassandra",
    //     Password: "cassandra",
	// }
	cluster.Keyspace = "hoan"
	cluster.Timeout = 5*time.Second
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum
	session,err := cluster.CreateSession()
	keySpaceMeta,_ := session.KeyspaceMetadata("hoan")
	return err,session,keySpaceMeta
}