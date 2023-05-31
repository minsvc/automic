package v1

import "github.com/gin-gonic/gin"

type Script struct{}

func NewScript() Script {
	return Script{}
}

func (s Script) Get(c *gin.Context)    {}
func (s Script) List(c *gin.Context)   {}
func (s Script) Create(c *gin.Context) {}
func (s Script) Update(c *gin.Context) {}
func (s Script) Delete(c *gin.Context) {}
