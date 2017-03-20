package api

// Operations An ordered key list of Operations
type Operations interface {
	// Add Add a new Operation to the list
	Add(Operation)
	// Get Retrieve an Operation from the list by key
	Get(string) (Operation, error)
	// Order Retrieve the ordered key list of Operation objects
	Order() []string
}
