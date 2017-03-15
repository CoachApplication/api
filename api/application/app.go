package application

type Application interface {
	AddBuilder(Builder)
	Activate(string, []string, SettingsProvider) error

	Validate() api_result.Result
	Operations() api_operations.Operations	
}
