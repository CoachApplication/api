package ui

// SimpleMetaData Metadata Implementation that parametrizes the fields, and offers itself as a wrapper
type SimpleMetadata struct {
	id string
	label string
	description string
	help string
}

// NewSimpleMetadata Creates a new SimpleMetadata from parametrized strings
func NewSimpleMetadata(id, label, description, help string) *SimpleMetadata {
	return &SimpleMetadata{
		id: id,
		label: label,
		description: description,
		help: help,
	}
}

// Id Provide a string machine name that is unique for the element in it's scope
func (smd *SimpleMetadata) Metadata() Metadata {
	return Metadata(smd)
}
// Id Provide a string machine name that is unique for the element in it's scope
func (smd *SimpleMetadata) Id() string {
	return smd.id
}

// Provide a short readable name for the element, to be used for UI labels for the element
func (smd *SimpleMetadata) Label() string {
	return smd.label
}

// Description Provide a long, perhaps multi-paragraph string description of the element
func (smd *SimpleMetadata) Description() string {
	return smd.description
}

// Help provide a long, multi-paragraph help.usage string for the element
func (smd *SimpleMetadata) Help() string {
	return smd.help
}
