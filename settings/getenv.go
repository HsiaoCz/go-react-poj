package settings

import "os"

func GetPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		port = ":3001"
	}
	return port
}

func GetMongoUri(key string) string {
	mongoUri := os.Getenv(key)
	if mongoUri == "" {
		mongoUri = "mongodb://localhost:27017"
	}
	return mongoUri
}

func GetDBname(key string) string {
	dbname := os.Getenv(key)
	if dbname == "" {
		dbname = "gorct"
	}
	return dbname
}

func GetUserColl(key string) string {
	userColl := os.Getenv(key)
	if userColl == "" {
		userColl = "user"
	}
	return userColl
}
