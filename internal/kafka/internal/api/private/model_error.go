/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager APIs that are used by internal services e.g kas-fleetshard operators.
 *
 * API version: 1.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// Error struct for Error
type Error struct {
	// Human-readable description of the error. Intended for human consumption
	Reason string `json:"reason"`
	// Relatively unique operation ID of the request associated to the error
	OperationId string `json:"operation_id,omitempty"`
	// The unique and immutable identifier of the resource
	Id string `json:"id"`
	// The kind of the resource
	Kind string `json:"kind"`
	// The absolute path of the resource
	Href string `json:"href"`
	// Code of the error
	Code string `json:"code"`
}
