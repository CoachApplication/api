package property

import (
	"error"
)

// SimpleProperties is a simple Properties implementation that keeps an ordered list of Properties
type SimpleProperties struct {
	pMap map[string]Property
	pOrder []string
}

// NewSimpleProperties is a SimpleProperties constructor
func NewSimpleProperties() *SimpleProperties {
	return &SimpleProperties{}
}

// Properties converts this to the Properties interface explicitly
func (sp *SimpleProperties) Properties() Properties {
	return Properties(sp)
}

// Add adds a new property to the list, keyed by property id
func (sp *SimpleProperties) Add(prop Property) {
	sp.safe()
	key := prop.Id()
	if _, found := sp.pMap[key]; !found {
		sp.pOrder = append(sp.pOrder, key)
	}
	sp.pMap[key] = prop
}

// Get retrieves a keyed property from the list
func (sp *SimpleProperties) Get(key string) (Property, error) {
	sp.safe()
	if prop, found := sp.pMap[key]; found {
		return prop, nil
	} else {
		return prop, &PropertyNotFoundError{key: key}
	}
}

// Order returns the ordered Property key list
func (sp *SimpleProperties) Order() []string {
	sp.safe()
	return sp.pOrder
}

// Safe lazy initializer
func (sp *SimpleProperties) safe() {
	if sp.pMap == nil {
		sp.pMap = map[string]Property{}
		sp.pOrder = []string{}
	}
}
