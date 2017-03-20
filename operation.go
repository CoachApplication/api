package api

/**
 * Operation individual API action
 *
 * Operation interfaces are the desired end product of the API.  The operations are then usable by any API consumer as
 * needed.
 */
type Operation interface {

	// Id Unique string machine name identifier for the Operation
	Id() string

	// UI Return a UI interaction definition for the Operation
	Ui() Ui

	// Usage Define how the Operation is intended to be executed.
	Usage() Usage

	// Properties provide the expected Operation with default values
	Properties() Properties

	// Validate Validate that the Operation can Execute if passed proper Property data
	Validate() Result

	/**
	 * Exec runs the operation from a Properties set, and return a result
	 *
	 * Exec is expected to handle any forking needed internally, and to pass any response Property changes via the
	 * api_result.Result.Properties() method.
	 *
	 * Exec receives a Properties list that it should consider disposable.  This is by design so that any Operation
	 * consumer can reused the Properties object for subsequent calls, which may run in parrallel.
	 */
	Exec(Properties) Result
}
