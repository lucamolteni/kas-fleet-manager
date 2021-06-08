/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances and connectors.
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// List struct for List
type List struct {
	Kind  string `json:"kind"`
	Page  int32  `json:"page"`
	Size  int32  `json:"size"`
	Total int32  `json:"total"`
}
