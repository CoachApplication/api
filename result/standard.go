package result

import (
	api_property
)

// StandardResult is as default Result implementation
type StandardResult struct {
	errors []error
	success bool
	finished []chan bool
	properties api_property.Properties
}

// NewStandardResult constructs a new StandardResult
func NewStandardResult() *StandardResult {
	return &StandardResult{
		errors: []error{},
		success: true, // default to successful
		finished: []chan bool{},
		properties: api_property.NewSimplePropertiesEmpty().Properties(),
	}
}

// Result explicitly converts this struct to the Result interface (for clarity and validation)
func (sr *StandardResult) Result() api_result.Result {
	return api_result.Result(sr)
}

/**
 * Result interface 
 */

// Return a slice of any errors that occured
func (sr *StandardResult) Errors() []error {
	return sr.errors
}

// Finished returns a tracking bool channel that can be used to mark when the operation is completed
func (sr *StandardResult) Finished() chan bool {
	finished = make(chan bool)
	sr.finished = append(sr.finished, finished)
	return finished
}

// Success returns a boolean success value
func (sr *StandardResult) Success() bool) {
	return sr.success
}

// Properties returns an ordered list of property values for the result
func (sr *StandardResult) Properties() api_property.Properties {
	return sr.properties
}

/**
 * Methods for creating the result data
 */

// AddError adds an Error to the result
func (sr *StandardResult) AddError(err Error) {
	sr.errors = append(sr.errors, err)
}

// MarkFailed marks this result as failed
func (sr *StandardResult) MarkFailed() {
	sr.success = false
}

// MarkSucceeded marks this result as succeeded
func (sr *StandardResult) MarkSucceeded() {
	sr.success = true
}

// MarkFinished marks this result operations as completed
func (sr *StandardResult) MarkFinished() {
	go func(finishedList []chan bool) {
		for _, eachFinished := range finishedList {
			eachFinished <- true
			close(eachFinished)
		}
	}(sr.finished)
	sr.finished = []chan bool{}
}

// Merge a result into this result
func (sr *StandardResult) Merge(merge Result) {
	if !merge.Success() {
		sr.MarkFailed()
	}
	for _, err := merge.Errors() {
		sr.AddError(err)
	}
}
