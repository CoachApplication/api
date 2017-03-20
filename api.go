package api

/**
 * The API instance is the end goal of any application which uses the coach api.  The usable deliverable of the entire
 * coach application is an API instance, which can give the application access to a list of operations which can be
 * executed.
 */

// API an application consumable object that will give access to all of the api Operation objects
type API interface {
	// Validate tells the consumer whether or the API is properly configured and ready to provide Operations
	Validate() Result
	// Operations provides a set of Operation
	Operations() Operations
}
