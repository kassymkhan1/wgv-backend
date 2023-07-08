package main

import (
	"github.com/gin-gonic/gin"
	"wgv/api"
	"wgv/internal/db"
	"wgv/internal/etcd"
	"wgv/internal/firebase"
	"wgv/internal/traefik"
)

func main() {
	var auth firebase.FIREBASE
	auth = &firebase.Fire{
		Auth: "./serviceAccount.json",
	}

	router := gin.Default()

	var routing etcd.Route
	routing = etcd.Etcd{}

	var traffic traefik.Rout
	traffic = traefik.Traefik{}

	postgres := db.Postgres{}

	router.Use(func(c *gin.Context) {
		c.Set("db", postgres)
		c.Set("traefik", traffic)
		c.Set("etcd", routing)
		c.Set("firebase", auth.SetupFirebase)
	})

	// Authorization and configuration Cors
	router.Use(CorsMiddleware())
	router.Use(auth.AuthMiddleware)

	// Creat API
	var create api.ApiCreate
	create = &api.Create{}

	create_router := router.Group("/create")
	{
		create_router.POST("/cloud", create.Cloud)
		create_router.POST("/etcd", create.EtcdPath)
		create_router.POST("/vpn", create.Vpn)
	}

	// Update Api
	var update api.ApiUpdate
	update = &api.Update{}

	updateRouter := router.Group("/create")
	{
		updateRouter.POST("/cloud", update.Cloud)
		updateRouter.POST("/etcd", update.EtcdPath)
		updateRouter.POST("/vpn", update.Vpn)
	}

	// Read Api
	var read api.ApiRead
	read = &api.Read{}

	readRouter := router.Group("/create")
	{
		readRouter.POST("/cloud", read.Cloud)
		readRouter.POST("/etcd", read.EtcdPath)
		readRouter.POST("/vpn", read.Vpn)
	}

}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
