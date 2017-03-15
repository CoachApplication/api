package application

// Builders provides an ordered set of keyed builders
type Builders interace {
	// Add a new builder to the ordered set
	Add(Buider)
	// Get a builder that matches a key from the set
	Get(string) (Builder, error)
	// Order the builder keys from the set
	Order() []string
}
