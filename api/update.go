package api

import (
	"github.com/gin-gonic/gin"
)

type ApiUpdate interface {
	EtcdPath(c *gin.Context)
	Cloud(c *gin.Context)
	Vpn(c *gin.Context)
}

type Update struct{}

func (update *Update) EtcdPath(c *gin.Context) {}

func (update *Update) Vpn(c *gin.Context) {}

func (update *Update) Cloud(c *gin.Context) {}
