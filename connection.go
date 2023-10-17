package PasetoprojectBackend

import (
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func MongoCreateConnection(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func GetAllUser(MongoConn *mongo.Database, colname string) []User {
	data := atdb.GetAllDoc[[]User](MongoConn, colname)
	return data
}

func GetOneUser(MongoConn *mongo.Database, colname string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	data := atdb.GetOneDoc[User](MongoConn, colname, filter)
	return data
}

func PasswordValidator(MongoConn *mongo.Database, colname string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	data := atdb.GetOneDoc[User](MongoConn, colname, filter)
	hashChecker := CompareHashPass(userdata.Password, data.Password)
	return hashChecker
}
