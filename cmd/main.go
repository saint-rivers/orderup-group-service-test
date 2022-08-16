package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saint-rivers/test-api/pkg/common/db"
	"github.com/saint-rivers/test-api/pkg/groups"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	groups.RegisterRoutes(r, h)
	r.Run(port)
}
