package settings

import "os"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3001"
	}
	return port
}

func GetMongoUri() string {
	mongoUri := os.Getenv("MONGOURI")
	if mongoUri == "" {
		mongoUri = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	}
	return mongoUri
}

func GetDBname() string {
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		dbname = "gorct"
	}
	return dbname
}

func GetUserColl() string {
	userColl := os.Getenv("USERCOLL")
	if userColl == "" {
		userColl = "user"
	}
	return userColl
}

func GetTodoColl() string {
	todoColl := os.Getenv("TODOCOLL")
	if todoColl == "" {
		todoColl = "todo"
	}
	return todoColl
}
