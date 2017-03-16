package application

import (
	api_operation "github.com/james-nesbitt/coach-api/operation"
	api_result "github.com/james-nesbitt/coach-api/result"
)

// SimpleApplication a simple Application implementation which maintains an ordered list of Builders
type SimpleApplication struct {
	builders Builders
}

// Application Explicitly converts this object to a Appication interface
func (sa *SimpleApplication) Application() Application {
	return Application(sa)
}

/**
 * Interface: Application
 */

// AddBuilder Include a new Builder in the App
func (sa *SimpleApplication) AddBuilder(builder Builder) {
	sa.builders.Add(builder)
}

// Activate Enable any number of Implementations of a Handler by key, and provide settings
func (sa *SimpleApplication) Activate(builderId string, implementations []string, settingsProvider SettingsProvider) error {
	if builder, err := sa.builders.Get(builderId); err == nil {
		return builder.Activate(implementations, settingsProvider)
	} else {
		return err
	}
}

// Validate Ask if the Application is ready to provide Operations
func (sa *SimpleApplication) Validate() api_result.Result {
	res := api_result.NewStandardResult()

	res.MarkFinished()
	return res.Result()
}

// Operations Retrieve a list of Operations from all of the Builders
func (sa *SimpleApplication) Operations() api_operation.Operations {
	ops := api_operation.NewSimpleOperations()
	for _, id := range sa.builders.Order() {
		b, _ := sa.builders.Get(id)
		ops.Merge(b.Operations())
	}
	return ops.Operations()
}
