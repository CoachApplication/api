package api

// Result indicates status and properties of the completion of a process such as an operation
type Result interface {
	// Return a slice of any errors that occured
	Errors() []error

	/**
	 * Finished returns a tracking bool channel that can be used to mark when the operation is completed
	 *
	 * The chan should be written to each time the result is marked finished (which typically only happens once).
	 * This is only to make the Finished Marking safe to be repeated multiple times.
	 * @TODO is the above behaviour appropriate?
	 *
	 * Note that if a Result is considered finished already when this chan is retrieved, it should still receive
	 * an initial write, so that the chan can be read without any concern about chan read timeout.
	 */
	Finished() chan bool
	// Success returns a boolean success value
	Success() bool
	// Properties returns an ordered list of property values for the result
	Properties() Properties
}
