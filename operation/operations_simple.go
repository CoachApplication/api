package operation

//SimpleOperations Operations implementation that maintains an ordered list
type SimpleOperations struct {
	oMap map[string]Operation
	oOrder []string
}

// safe lazy-initializer
func (so *SimpleOperations) safe() {
	if so.oMap == nil {
		so.oMap = map[string]Operation{}
		so.oOrder = []string{}
	}
}

// Add adds an operation to the list
func (so *SimpleOperations) Add(op api_operation.Operation) error {
	so.safe()
	key := op.Id()
	if _, found := so.oMap[key]; !found {
		so.oOrder = append(so.oOrder, key)
	}
	so.oMap[key] = op
}

// Get retrieves an operation from the list by string key
func (so *SimpleOperations) Get(key string) op api_operation.Operation, error) {
	so.safe()
	if op, found := so.oMap[key]; found {
		return op, nil
	} else {
		return op, Error(OperationNotFound{Key: key})
	}
}

// List retrieves an ordered string list of operation keys
func (so *SimpleOperations) List() []string {
	so.safe()
	return so.oOrder
}
