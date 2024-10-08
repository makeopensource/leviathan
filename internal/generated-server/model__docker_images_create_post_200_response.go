/*
 * Leviathan api
 *
 * OpenAPI spec for leviathan
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package serverstub

type DockerImagesCreatePost200Response struct {

	// Logs generated when image was built
	Logs string `json:"logs,omitempty"`

	TagName string `json:"tagName,omitempty"`
}
