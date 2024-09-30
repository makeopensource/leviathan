package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DockerAPI struct {
}

// DockerContainerIdDelete Delete /docker/:containerId
func (dk DockerAPI) DockerContainerIdDelete(c *gin.Context) {
	log.Debug().Msgf("Recived Delete container request with id: %s", c.Param("containerId"))
	c.Status(200)
}

// DockerContainerIdGet Get /docker/:containerId
func (dk DockerAPI) DockerContainerIdGet(c *gin.Context) {
	log.Debug().Msgf("Recived container info request with id: %s", c.Param("containerId"))
	c.Status(200)
}

// DockerContainerIdStartGet Get /docker/:containerId/start
func (dk DockerAPI) DockerContainerIdStartGet(c *gin.Context) {
	log.Debug().Msgf("Recived start container request with id: %s", c.Param("containerId"))
	c.Status(200)
}

// DockerContainerIdStopGet Get /docker/:containerId/stop
func (dk DockerAPI) DockerContainerIdStopGet(c *gin.Context) {
	log.Debug().Msgf("Recived stop container request with id: %s", c.Param("containerId"))
	c.Status(200)
}

// DockerImagesCreatePost Post /docker/images/create
func (dk DockerAPI) DockerImagesCreatePost(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Error().Err(err).Msgf("Could not parse multipart form")
		return
	}
	tagName := form.Value["tagName"]

	log.Debug().Any("Tagname", tagName).Msgf("Recived create image request")
	c.Status(200)
}

// DockerImagesGet Get /docker/images
func (dk DockerAPI) DockerImagesGet(c *gin.Context) {
	log.Debug().Msgf("Recived list image request")
	c.Status(200)
}
