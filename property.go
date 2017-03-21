package api

/**
 * Property is a single variable used to configure an operation during it's Exec.
 *
 * There is a Properties interface which provides a protable ordered collection of properties.  This collection is meant
 * to be passed to the Operation Exec.  Property defaults can be retreived from the Operation.
 *
 * Property, like Operation, is meant to be usable with any API consumer UI, and therefore provides
 */
type Property interface {
	// Id provides a machine name string for the property.  This should be unique in an operation.
	Id() string

	// Type string label for content type of property value
	Type() string

	// Ui Provide UI metadata for the Property
	Ui() Ui

	// Usage Provide Usage information about the element
	Usage() Usage

	// Validate Check that the property is properly configured
	Validate() Result

	// Get retrieve a value from the Property
	Get() interface{}
	// Set assign a value to the Property
	Set(interface{}) error
}
