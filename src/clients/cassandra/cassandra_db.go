package cassandra

import (
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var session *gocql.Session

func init() {
	// Connect to Cassandra cluster:
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cluster := gocql.NewCluster(os.Getenv("DB_HOST"))
	cluster.Keyspace = os.Getenv("DB_KEYSPACE")
	cluster.Consistency = gocql.Quorum

	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
	fmt.Println("cassandra connected successfully")
	defer session.Close()
}

func GetSession() *gocql.Session {
	return session
}
