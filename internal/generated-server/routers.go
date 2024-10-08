/*
 * Leviathan api
 *
 * OpenAPI spec for leviathan
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package serverstub

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the CoursesAPI part of the API
	CoursesAPI CoursesAPI
	// Routes for the DockerAPI part of the API
	DockerAPI DockerAPI
	// Routes for the StatsAPI part of the API
	StatsAPI StatsAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"CourseCourseIdDelete",
			http.MethodDelete,
			"/course/:courseId",
			handleFunctions.CoursesAPI.CourseCourseIdDelete,
		},
		{
			"CourseCourseIdGet",
			http.MethodGet,
			"/course/:courseId",
			handleFunctions.CoursesAPI.CourseCourseIdGet,
		},
		{
			"CourseCourseIdPatch",
			http.MethodPatch,
			"/course/:courseId",
			handleFunctions.CoursesAPI.CourseCourseIdPatch,
		},
		{
			"CoursePost",
			http.MethodPost,
			"/course",
			handleFunctions.CoursesAPI.CoursePost,
		},
		{
			"DockerContainerIdDelete",
			http.MethodDelete,
			"/docker/:containerId",
			handleFunctions.DockerAPI.DockerContainerIdDelete,
		},
		{
			"DockerContainerIdGet",
			http.MethodGet,
			"/docker/:containerId",
			handleFunctions.DockerAPI.DockerContainerIdGet,
		},
		{
			"DockerContainerIdStartGet",
			http.MethodGet,
			"/docker/:containerId/start",
			handleFunctions.DockerAPI.DockerContainerIdStartGet,
		},
		{
			"DockerContainerIdStopGet",
			http.MethodGet,
			"/docker/:containerId/stop",
			handleFunctions.DockerAPI.DockerContainerIdStopGet,
		},
		{
			"DockerImagesCreatePost",
			http.MethodPost,
			"/docker/images/create",
			handleFunctions.DockerAPI.DockerImagesCreatePost,
		},
		{
			"DockerImagesGet",
			http.MethodGet,
			"/docker/images",
			handleFunctions.DockerAPI.DockerImagesGet,
		},
		{
			"StatsGet",
			http.MethodGet,
			"/stats",
			handleFunctions.StatsAPI.StatsGet,
		},
	}
}
