/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ImageCreate struct for ImageCreate
type ImageCreate struct {
	Name        string            `json:"name"`
	Vm          string            `json:"vm,omitempty"`
	Replica     string            `json:"replica,omitempty"`
	Service     string            `json:"service,omitempty"`
	Description string            `json:"description,omitempty"`
	Tag         map[string]string `json:"tag,omitempty"`
}
