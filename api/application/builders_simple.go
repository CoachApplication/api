package application

// SimpleBuilders A simple Builders implementation that maintains an ordered list
type SimpleBuilders struct {
	bMap map[string]Builder 
	bOrder []string
}

// Add a new builder to the ordered set
func (sb *SimpleBuilders) Add(b Builder) {
	key := b.Id()
	if _, found := sb.bMap[key]; !found {
		sb.bOrder = append(sb.bOrder, key)
	}
	sb.bMap[key] = b
}

// Get a builder that matches a key from the set
func (sb *SimpleBuilders) Get(key string) (Builder, error) {
	sb.safe()
	if b, found := sb.bMap[key]; found {
		return b, nil
	} else {
		return b, error(&BuilderNotFoundError{key:key})
	}
}

// Order the builder keys from the set
func (sb *SimpleBuilders) Order() []string {
	sb.safe()
	return sb.bOrder
}

// safe initializer
func (sb *SimpleBuilders) safe() {
	if sb.bOrder == nil {
		sb.bMap = map[string]Builder{}
		sb.bOrder = []string{}
	}
}
