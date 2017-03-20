package api

/**
 * Builder is a tool provided by a handler (a 3rd party coach package) which can be used to build up lists of operations
 * based on activating pieces using the app approach.
 *
 * The Builder is the handler part of the APP builder, where the app allows a separation of the Handler library from its
 * configuration.  This is usefull because Handlers may be used in different ways, which can then be configured at or at
 * least closer to runtime.
 * The Activation "Implemntation" is a string key that the Handler itself should recognize internally.  The App would
 * have to provide only implementation keys that make sense, but the Handler can ignore invalid values witout a panic.
 */
type Builder interface {
	// Id provides a unique machine name for the Builder
	Id() string
	// SetParent Provides the API reference to the Builder which may use it's operations internally
	SetParent(API)
	// Activate Enable keyed implementations, providing settings for those handler implementations
	Activate([]string, SettingsProvider) error
	// Validates Ask the builder if it is happy and willing to provide operations
	Validate() Result
	// Operations provide any Builder user with a set of Operation objects
	Operations() Operations
}
