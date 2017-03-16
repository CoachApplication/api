package operation

//SimpleOperations Operations implementation that maintains a simple ordered list
type SimpleOperations struct {
	oMap   map[string]Operation
	oOrder []string
}

// NewSimpleOperations SimpleOperations constructor
func NewSimpleOperations() *SimpleOperations {
	return &SimpleOperations{}
}

// Operations convert this object to an Operations interface explicitly
func (so *SimpleOperations) Operations() Operations {
	return Operations(so)
}

// safe lazy-initializer
func (so *SimpleOperations) safe() {
	if so.oMap == nil {
		so.oMap = map[string]Operation{}
		so.oOrder = []string{}
	}
}

// Merge merge one set of operations into this object
func (so *SimpleOperations) Merge(source Operations) {
	for _, id := range source.Order() {
		op, _ := source.Get(id)
		so.Add(op)
	}
}

/**
 * Interface: Operations
 */

// Add adds an operation to the list
func (so *SimpleOperations) Add(op Operation) error {
	so.safe()
	key := op.Id()
	if _, found := so.oMap[key]; !found {
		so.oOrder = append(so.oOrder, key)
	}
	so.oMap[key] = op
	return nil
}

// Get retrieves an operation from the list by string key
func (so *SimpleOperations) Get(key string) (Operation, error) {
	so.safe()
	if op, found := so.oMap[key]; found {
		return op, nil
	} else {
		return op, error(OperationNotFound{Key: key})
	}
}

// List retrieves an ordered string list of operation keys
func (so *SimpleOperations) List() []string {
	so.safe()
	return so.oOrder
}
