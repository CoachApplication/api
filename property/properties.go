package property

// Properties an ordered map of properties
type Properties interface {
	// Add adds a new property to the list, keyed by property id
	Add(Property)
	// Get retrieves a keyed property from the list
	Get(string) (Property, error)
	// Order returns the ordered Property key list
	Order() []string
}
