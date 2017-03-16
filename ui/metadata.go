package ui

/**
 * @TODO this will need some translation process in the future.
 *
 * We will pivot this interface to provide all information that may be needed in UI elements.  THis way any translation
 * functionality needed can be focus3ed on here.
 */

// Interface provides UI elements that can be used by API consumers to relate to users
type Metadata interface {

	// Id Provide a string machine name that is unique for the element in it's scope
	Id() string

	// Provide a short readable name for the element, to be used for UI labels for the element
	Label() string

	// Description Provide a long, perhaps multi-paragraph string description of the element
	Description() string

	// Help provide a long, multi-paragraph help.usage string for the element
	Help() string

	/**
	 * SetLanguage() ?
	 */
}
