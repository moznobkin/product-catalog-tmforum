/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 19.0 - June 2019  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Complements the description of an element (for instance a product) through video, pictures...
type Attachment struct {

	// Unique identifier for this particular attachment
	Id string `json:"id,omitempty"`

	// URI for this Attachment
	Href string `json:"href,omitempty"`

	// Attachment type such as video, picture
	AttachmentType string `json:"attachmentType,omitempty"`

	// The actual contents of the attachment object, if embedded, encoded as base64
	Content string `json:"content,omitempty"`

	// A narrative text describing the content of the attachment
	Description string `json:"description,omitempty"`

	// Attachment mime type such as extension file for video, picture and document
	MimeType string `json:"mimeType,omitempty"`

	// The name of the attachment
	Name string `json:"name,omitempty"`

	// Uniform Resource Locator, is a web page address (a subset of URI)
	Url string `json:"url,omitempty"`

	// The size of the attachment.
	Size *Quantity `json:"size,omitempty"`

	// The period of time for which the attachment is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
}
