package api

import (
	"github.com/gin-gonic/gin"
	spec "github.com/makeopensource/leviathan/internal/generated-server"
)

func SetupPaths() *gin.Engine {
	courseApi := CourseAPI{}
	dkApi := DockerAPI{}
	statApi := StatsAPI{}

	registerHandlers := spec.ApiHandleFunctions{
		CoursesAPI: courseApi,
		DockerAPI:  dkApi,
		StatsAPI:   statApi,
	}

	router := spec.NewRouter(registerHandlers)
	return router
}
