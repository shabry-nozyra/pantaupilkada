package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/pantaupilkada/handlers"
	"github.com/shabry-nozyra/pantaupilkada/helpers/env"
	"github.com/shabry-nozyra/pantaupilkada/helpers/log"
	"github.com/shabry-nozyra/pantaupilkada/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)


func main()  {
	l := log.NewLog("Log", "")
	defer l.Close()

	cs := env.Get().ConnectionString()
	db, err := gorm.Open(postgres.Open(cs), &gorm.Config{})

	if err != nil{
		l.Fatal(err)

		return
	}

	err = models.MigrateModel(db)
	if err != nil{
		l.Error("AutoMigrate Gagal")
		return
	}
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()

	c := cors.DefaultConfig()
	c.AllowWildcard = true
	c.AllowCredentials = true
	c.AllowOrigins = []string{"https://pantaufront.azurewebsites.net","http://localhost:8001"}
	c.AddAllowHeaders("Authorization", "Content-Type", "Access-Control-Allow-Credentials")
	c.AddExposeHeaders("Authorization", "Content-Type", "Access-Control-Allow-Credentials")
	g.Use(cors.New(c))

	h := handlers.Context{Gin: g, DB: db, Log: l}
	h.Register("")
	port := os.Getenv("AppPort")

	l.Infof("start listen and serve at %v", port)
	s := &http.Server{Addr: "0.0.0.0:" + port, Handler: g}
	err = s.ListenAndServe()
	if err != nil {
		l.Fatal("failed to connect to serv")
		return
	}
}
