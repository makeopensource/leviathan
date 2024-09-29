package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DockerAPI struct {
}

// DockerContainerIdDelete Delete /docker/:containerId
func (dk DockerAPI) DockerContainerIdDelete(c *gin.Context) {

}

// DockerContainerIdGet Get /docker/:containerId
func (dk DockerAPI) DockerContainerIdGet(c *gin.Context) {

}

// DockerContainerIdStartGet Get /docker/:containerId/start
func (dk DockerAPI) DockerContainerIdStartGet(c *gin.Context) {

}

// DockerContainerIdStopGet Get /docker/:containerId/stop
func (dk DockerAPI) DockerContainerIdStopGet(c *gin.Context) {

}

// DockerImagesCreatePost Post /docker/images/create
func (dk DockerAPI) DockerImagesCreatePost(c *gin.Context) {
	log.Debug().Msgf("Recived create image request")
	c.Status(200)
}

// DockerImagesGet Get /docker/images
func (dk DockerAPI) DockerImagesGet(c *gin.Context) {
	log.Debug().Msgf("Recived list image request")
	c.Status(200)
}
