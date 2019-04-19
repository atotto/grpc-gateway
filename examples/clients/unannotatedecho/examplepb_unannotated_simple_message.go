/*
 * examples/proto/examplepb/unannotated_echo_service.proto
 *
 * Unannotated Echo Service Similar to echo_service.proto but without annotations. See unannotated_echo_service.yaml for the equivalent of the annotations in gRPC API configuration format.  Echo Service API consists of a single service which returns a message.
 *
 * OpenAPI spec version: version not set
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package unannotatedecho

// UnannotatedSimpleMessage represents a simple message sent to the unannotated Echo service.
type ExamplepbUnannotatedSimpleMessage struct {

	// Id represents the message identifier.
	Id string `json:"id,omitempty"`

	Num string `json:"num,omitempty"`

	Duration string `json:"duration,omitempty"`
}
