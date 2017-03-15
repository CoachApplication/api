package application

type SimpleBuilders struct {
	bMap map[string]Builder 
	bOrder []string
}

// Add a new builder to the ordered set
func (sb *SimpleBuilders) Add(b Builder) {
	key := b.Id*()
	if _, found := sb.bMap[key]; !found {
		sb.bOrder = append(sb.bOrder, key)
	}
	sb.bMap[key] = b
}

// Get a builder that matches a key from the set
func (sb *SimpleBuilders) Get(key string) (Builder, error) {
	sb.safe()
	if b, found := sb.bMap[key]; b {
		return b, nil
	} else {
		return b, error(&BuilderNotFound{key:key})
	}
}

// Order the builder keys from the set
func (sb *SimpleBuilders) Order() []string {
	sb.safe()
	return sb.bOrder
}

// safe intitializer
func (sb *SimpleBuilders) safe() {
	if builders.bOrder == nil {
		builders.bMap = map[string]Builder{}
		builders.bOrder = []string{}
	}
}

/**
 * Errors 
 */

// BuilderNotFoundError for when a builder key does not exist in the list
type BuilderNotFoundError struct {
	key string
}

// Error string return (interface: error)
func (bnf *BuilderNotFoundError) Error() string {
	return "Builder not found: "+bnf.key
}
