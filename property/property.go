package property

import (
	api_ui "github.com/james-nesbitt/coach-api/ui"
	api_usage "github.com/james-nesbitt/coach-api/usage"
)

/**
 * Property is a single variable used to configure an operation during it's Exec.
 *
 * There is a Properties interface which provides a protable ordered collection of properties.  This collection is meant
 * to be passed to the Operation Exec.  Property defaults can be retreived from the Operation.
 *
 * Property, like Operation, is meant to be usable with any API consumer UI, and therefore provides
 */
type Property interface {
	// Id provides a machine name string for the property.  This should be uniquein an operation.
	Id() string

	// Ui Provude UI metadata for the Property
	Ui() api_ui.Metadata

	// Usage Provide Usage information about the element
	Usage() api_usage.Usage

	// Get retrieve a value from the Property
	Get() interface{}
	// Set assign a value to the Property
	Set(interface{}) error
}
