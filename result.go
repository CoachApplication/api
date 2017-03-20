package api

// Result indicates status and properties of the completion of a process such as an operation
type Result interface {
	// Return a slice of any errors that occured
	Errors() []error
	// Finished returns a tracking bool channel that can be used to mark when the operation is completed
	Finished() chan bool
	// Success returns a boolean success value
	Success() bool
	// Properties returns an ordered list of property values for the result
	Properties() Properties
}
