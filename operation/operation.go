package operation

import (
	api_property "github.com/james-nesbitt/coach-api/property"
	api_result "github.com/james-nesbitt/coach-api/result"
	api_ui "github.com/james-nesbitt/coach-api/ui"
	api_usage "github.com/james-nesbitt/coach-api/usage"
)

/**
 * Operation individual API action
 *
 * Operation interfaces are the desired end product of the API.  The operations are then usable by any API consumer as
 * needed.
 */
type Operation interface {

	Id() string

	Metadata() api_ui.Metadata

	Usage() api_usage.Usage

	/**
	 * Exec runs the operation from a Properties set, and return a result
	 *
	 * Exec is expected to handle any forking needed internally, and to pass any response Property changes via the
	 * api_result.Result.Properties() method.
	 *
	 * Exec receives a Properties list that it should consider disposable.  This is by design so that any Operation
	 * consumer can reused the Properties object for subsequent calls, which may run in parrallel.
	 */
	Exec(api_property.Properties) api_result.Result

}
