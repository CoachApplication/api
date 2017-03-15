package api

type API interface {
	Validate() api_result.Result
	Operations() api_operations.Operations
}
