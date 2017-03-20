package api

// Builders provides an ordered set of keyed builders
type Builders interface {
	// Add a new builder to the ordered set
	Add(Builder)
	// Get a builder that matches a key from the set
	Get(string) (Builder, error)
	// Order the builder keys from the set
	Order() []string
}
