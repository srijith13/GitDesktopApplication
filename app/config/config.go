package config

import (
	"os"

	"github.com/joho/godotenv"
)

var GinMode string
var GinTZ string
var AppPort string
var DbHost string
var DbUser string
var DbPass string
var RedisHost string
var RedisPort string
var RedisPass string
var RedisUserRequest int
var RedisUserTime int
var RedisGuestRequest int
var RedisGuestTime int

// var DbNames string

func init() {
	godotenv.Load()

	GinMode = os.Getenv("GIN_MODE")
	GinTZ = os.Getenv("GIN_TZ")
	AppPort = os.Getenv("APP_PORT")

	if os.Getenv("DB_MODE") == "stage" {
		DbHost = os.Getenv("DB_STAGE_HOST")
		DbUser = os.Getenv("DB_STAGE_USER")
		DbPass = os.Getenv("DB_STAGE_PASS")
	} else if os.Getenv("DB_MODE") == "prod" {
		DbHost = os.Getenv("DB_PROD_HOST")
		DbUser = os.Getenv("DB_PROD_USER")
		DbPass = os.Getenv("DB_PROD_PASS")
	}

	// if os.Getenv("REDIS_MODE") == "stage" {
	// 	RedisHost = os.Getenv("REDIS_STAGE_HOST")
	// 	RedisPort = os.Getenv("REDIS_STAGE_PORT")
	// 	RedisPass = os.Getenv("REDIS_STAGE_PASS")
	// } else if os.Getenv("REDIS_MODE") == "prod" {
	// 	RedisHost = os.Getenv("REDIS_PROD_HOST")
	// 	RedisPort = os.Getenv("REDIS_PROD_PORT")
	// 	RedisPass = os.Getenv("REDIS_PROD_PASS")
	// }

}

// var connection = make(map[string]*sql.DB)
// var redisconn *redis.Client

// func CreateDbConPool(code string) *sql.DB {
// 	dbinfo := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s sslmode=disable`, DbHost, DbUser, DbPass, DbNames)
// 	if connection[code] == nil {
// 		connection[code], _ = sql.Open("postgres", dbinfo)
// 	}
// 	return connection[code]
// }

// func GetRedisClient() *redis.Client {
// 	addr := fmt.Sprintf(`%s:%s`, RedisHost, RedisPort)
// 	if redisconn == nil {
// 		redisconn = redis.NewClient(&redis.Options{
// 			Addr:     addr,
// 			Password: RedisPass, // no password set
// 			DB:       0,         // use default DB
// 		})
// 	}
// 	return redisconn
// }
