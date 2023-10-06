package db

import (
	"os"
    "context"
	// "go.mongodb.org/mongo-driver/bson" 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	err error
	MongoDb *mongo.Database
	client *mongo.Client
)

func Disconect_db(){

 client.Disconnect(context.TODO()) 
}

func Init_db()(error){

// DB Connections Paramters
	DBName := os.Getenv("MONGO_DB_NAME")
	DBUser := os.Getenv("MONGO_DB_USER")
	DBPass := os.Getenv("MONGO_DB_PASS")
	DBHost := os.Getenv("MONGO_DB_HOST")
	DBPort := os.Getenv("MONGO_DB_PORT")


	clientOpts := options.Client().ApplyURI("mongodb://"+DBUser+":"+DBPass+"@"+DBHost+":"+DBPort+"/?authSource=admin&authMechanism=SCRAM-SHA-256")
	cli,err := mongo.Connect(context.TODO(), clientOpts)
	client=cli
	if err!=nil{
		return err
	}
	// dbNames, err := client.ListDatabaseNames(context.TODO(), bson.M{}) 
	// if err != nil { 
	// return err
	// }  

	MongoDb = client.Database(DBName) 

	// fmt.Println("Available datatabases:") 
	// fmt.Println(dbNames)

 return nil
 }

