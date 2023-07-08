package api

import (
	"github.com/gin-gonic/gin"
)

type ApiCreate interface {
	EtcdPath(c *gin.Context)
	Cloud(c *gin.Context)
	Vpn(c *gin.Context)
}

type Create struct{}

func (create *Create) EtcdPath(c *gin.Context) {}

func (create *Create) Vpn(c *gin.Context) {}

func (create *Create) Cloud(c *gin.Context) {}
