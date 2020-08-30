package helper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Database, error) {
	config := ReadConfig()

	clientOptions := options.Client().ApplyURI("mongodb://" + config["host"])
	if config["password"] != "" && config["username"] != "" {
		clientOptions = options.Client().ApplyURI("mongodb://" + config["host"] + ":" + config["password"] + "@" + config["username"])
	}
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	db := client.Database(config["database"])

	return db, err
}

func CheckDatabaseConnection() error {
	WriteLog.Println("Checking Database Connection ... ")
	stats, err := ConnectDB()
	err = stats.Client().Ping(context.TODO(), nil)
	if err != nil {
		WriteLog.Println(err)
		return err
	}
	WriteLog.Println("Connected!")

	err = stats.Client().Disconnect(context.TODO())
	if err != nil {
		WriteLog.Println(err)
		return err
	}

	return nil
}

// func InsertData(param models.ShortURL) error {
// 	config := ReadConfig()
// 	client, err := ConnectDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection := client.Database(config["database"]).Collection(param.TableName())
// 	res, err := collection.InsertOne(context.TODO(), param)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Inserted Document", res)
// 	return err
// }
