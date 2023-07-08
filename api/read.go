package api

import (
	"github.com/gin-gonic/gin"
)

type ApiRead interface {
	EtcdPath(c *gin.Context)
	Cloud(c *gin.Context)
	Vpn(c *gin.Context)
}

type Read struct{}

func (read *Read) EtcdPath(c *gin.Context) {}

func (read *Read) Vpn(c *gin.Context) {}

func (read *Read) Cloud(c *gin.Context) {}
