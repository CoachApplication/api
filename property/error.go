package property

/**
 * Errors 
 */

// PropertyNotFoundError an error for when a requested property is not in the list
type PropertyNotFoundError struct {
	key string
}

// Error returns an error string (interace: Error)
func (pnfe *PropertyNotFoundError) Error() string {
	return "Property not found: "+pnfe.key
}
