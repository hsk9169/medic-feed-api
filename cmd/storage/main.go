package main

import (
	_ "log"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"

	"github.com/medic-basic/s3-test/Config"
	"github.com/medic-basic/s3-test/Routes"
	"github.com/medic-basic/s3-test/Services"
)

func main() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error .env")
	//}

	//Config.CACHE = redis.NewClient(Config.CacheOptions())
	//if _, err := Config.CACHE.Ping().Result(); err != nil {
	//  errMsg := fmt.Sprintf("Failed to connect to cache : %s", err)
	//  panic(errMsg)
	//}
	//fmt.Println("Cache connection established")

	//Config.DB, err = gorm.Open(mysql.Open(Config.DbURL()))
	//if err != nil {
	//  errMsg := fmt.Sprintf("Failed to connect to database : %s", err)
	//  panic(errMsg)
	//}
	//fmt.Println("Database connection established")

	//Config.DB.AutoMigrate(
	//  &Schema.User{},
	//)

	Config.InitStorageCfg()
	Services.InitStorage()
	Services.InitBasic()

	r := Routes.SetupRouter()
	_ = r.Run(":8080")
}
