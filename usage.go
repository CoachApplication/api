package api

const (
	/**
	 * Usage test key values
	 *
	 * These are typically used keys that are testd in an application
	 */
	UsageOperationPublicView    = "operation.public.view"
	UsageOperationPublicExecute = "operation.public.exec"

	UsagePropertyPublicView     = "property.public.view"
	UsagePropertyPublicWrite    = "property.public.write"
	UsagePropertyPublicRequired = "property.public.required"
)

// Usage
type Usage interface {
	// Return a bool if an associate object is allowed to be used for a string key purpose
	Allows(string) bool
}
