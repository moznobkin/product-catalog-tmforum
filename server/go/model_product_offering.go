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

// Represents entities that are orderable from the provider of the catalog, this resource includes pricing information.
type ProductOffering struct {

	// Unique identifier of the productOffering
	Id string `json:"id,omitempty"`

	// Reference of the ProductOffering
	Href string `json:"href,omitempty"`

	// Description of the productOffering
	Description string `json:"description,omitempty"`

	// isBundle determines whether a productOffering represents a single productOffering (false), or a bundle of productOfferings (true).
	IsBundle bool `json:"isBundle,omitempty"`

	// A flag indicating if this product offer can be sold stand-alone for sale or not. If this flag is false it indicates that the offer can only be sold within a bundle.
	IsSellable bool `json:"isSellable,omitempty"`

	// Date and time of the last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`

	// Name of the productOffering
	Name string `json:"name,omitempty"`

	// A string providing a complementary information on the value of the lifecycle status attribute.
	StatusReason string `json:"statusReason,omitempty"`

	// ProductOffering version
	Version string `json:"version,omitempty"`

	// An agreement represents a contract or arrangement, either written or verbal and sometimes enforceable by law, such as a service level agreement or a customer price agreement. An agreement involves a number of other business entities, such as products, services, and resources and/or their specifications.
	Agreement []AgreementRef `json:"agreement,omitempty"`

	// Complements the description of an element (for instance a product) through video, pictures...
	Attachment []AttachmentRefOrValue `json:"attachment,omitempty"`

	// A type of ProductOffering that belongs to a grouping of ProductOfferings made available to the market. It inherits of all attributes of ProductOffering.
	BundledProductOffering []BundledProductOffering `json:"bundledProductOffering,omitempty"`

	// The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other categories and/or product offerings, resource or service candidates.
	Category []CategoryRef `json:"category,omitempty"`

	// The channel defines the channel for selling product offerings.
	Channel []ChannelRef `json:"channel,omitempty"`

	// provides references to the corresponding market segment as target of product offerings. A market segment is grouping of Parties, GeographicAreas, SalesChannels, and so forth.
	MarketSegment []MarketSegmentRef `json:"marketSegment,omitempty"`

	// Place defines the places where the products are sold or delivered.
	Place []PlaceRef `json:"place,omitempty"`

	// A use of the ProductSpecificationCharacteristicValue by a ProductOffering to which additional properties (attributes) apply or override the properties of similar properties contained in ProductSpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding product specification. The available characteristic values for a ProductSpecificationCharacteristic in a Product specification can be modified at the ProductOffering level. For example, a characteristic 'Color' might have values White, Blue, Green, and Red. But, the list of values can be restricted to e.g. White and Blue in an associated product offering. It should be noted that the list of values in 'ProductSpecificationCharacteristicValueUse' is a strict subset of the list of values as defined in the corresponding product specification characteristics.
	ProdSpecCharValueUse []ProductSpecificationCharacteristicValueUse `json:"prodSpecCharValueUse,omitempty"`

	// An amount, usually of money, that is asked for or allowed when a ProductOffering is bought, rented, or leased. The price is valid for a defined period of time and may not represent the actual price paid by a customer.
	ProductOfferingPrice []ProductOfferingPriceRef `json:"productOfferingPrice,omitempty"`

	// A condition under which a ProductOffering is made available to Customers. For instance, a productOffering can be offered with multiple commitment periods.
	ProductOfferingTerm []ProductOfferingTerm `json:"productOfferingTerm,omitempty"`

	// A ProductSpecification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role.
	ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty"`

	// A resource candidate is an entity that makes a ResourceSpecification available to a catalog.
	ResourceCandidate *ResourceCandidateRef `json:"resourceCandidate,omitempty"`

	// ServiceCandidate is an entity that makes a ServiceSpecification available to a catalog.
	ServiceCandidate *ServiceCandidateRef `json:"serviceCandidate,omitempty"`

	// A service level agreement (SLA) is a type of agreement that represents a formal negotiated agreement between two parties designed to create a common understanding about products, services, priorities, responsibilities, and so forth. The SLA is a set of appropriate procedures and targets formally or informally agreed between parties in order to achieve and maintain specified Quality of Service.
	ServiceLevelAgreement *SlaRef `json:"serviceLevelAgreement,omitempty"`

	// The period for which the productOffering is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
}
