package api

/**
 * Application provide the API part of a built application.
 *
 * The Application provides the API part of an application build which will hold, and Activate any Builder objects that
 * are used in the application.
 *
 * This allows a separation of the API and the builders, while still allowing runtime configuration of which Handler
 * Builder objects should be user, and provide settings to them.  In particular this separates the stage of creating
 * new Builder instances, and activating them, so that Activation can be based on runtime configuration.
 */
type Application interface {
	// AddBuilder Include a new Builder in the App
	AddBuilder(Builder)
	// Activate Enable any number of Implementations of a Handler by key, and provide settings
	Activate(string, []string, SettingsProvider) error

	// Validate Ask if the Application is ready to provide Operations
	Validate() Result
	// Operations Retrieve a list of Operations from all of the Builders
	Operations() Operations
}

type SettingsProvider interface {
	Apply(interface{}) error
}
