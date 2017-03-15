package application

type SimpleApplication struct {
	builders Builders
}

func (sa *SimpleApplication) AddBuilder(builder Builder) {
	sa.builders.Add(builder)
}
func (sa *SimpleApplication) Activate(builderId string, implementations []string, settingsProvider SettingsProvider) error {
	if builder, err := sa.builders.Get(builderId); err == nil {
		return builder.Activate(implementations, settingsProvider)
	} else {
		return err
	}
}

func (sa *SimpleApplication) Validate() api_result.Result {

}
func (sa *SimpleApplication) Operations() api_operations.Operations {

}
