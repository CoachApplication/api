package builder

type Builder interface {
	SetParent(api_api.API)
	Activate([]string, SettingsProvider) error
	Validate() api_result.Result
	Operations() api_operation.Operations
}
