package api

/**
 * Ui defines User Interface metadata for a described object
 */
type Ui interface {
	// Id Provide a unique machine name string value
	Id() string
	// Label Provide a string human readable short Label, perhaps a few words
	Label() string
	// Description Provide a longer human readable description, perhaps a paragraph
	Description() string
	// Help Provide man like Help instructions
	Help() string
}
