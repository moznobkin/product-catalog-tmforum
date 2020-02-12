/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 19.0 - June 2019  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// The notification data structure
type CatalogDeleteEvent struct {

	// Identifier of the Process flow
	Id string `json:"id,omitempty"`

	// Reference of the ProcessFlow
	Href string `json:"href,omitempty"`

	// The identifier of the notification.
	EventId string `json:"eventId,omitempty"`

	// Time of the event occurrence.
	EventTime time.Time `json:"eventTime,omitempty"`

	// The type of the notification.
	EventType string `json:"eventType,omitempty"`

	// The correlation id for this event.
	CorrelationId string `json:"correlationId,omitempty"`

	// The domain of the event.
	Domain string `json:"domain,omitempty"`

	// The title of the event.
	Title string `json:"title,omitempty"`

	// An explnatory of the event.
	Description string `json:"description,omitempty"`

	// A priority.
	Priority string `json:"priority,omitempty"`

	// The time the event occured.
	TimeOcurred time.Time `json:"timeOcurred,omitempty"`

	// The event payload linked to the involved resource object
	Event *CatalogDeleteEventPayload `json:"event,omitempty"`
}
