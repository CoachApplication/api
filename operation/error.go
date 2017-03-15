package operation

import (
	"fmt"
)

// OperationNotFound a matching operation was not found Error
type OperationNotFound struct {
	Key string
}

// Error returns error string (interface: Error)
func (onf OperationNotFound) Error() string {
	return fmt.Sprintf("Operation was not found : '%s'", onf.Key)
}
