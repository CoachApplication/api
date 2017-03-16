package ui

// MergeMetaData Merge multiple Metadata objects into a single one, overriding any nonempty values
func MergeMetaData(source... Metadata) Metadata {
	var id, label, description, help string
	for _, each := range source {
		if n := each.Id(); n != "" {
			id = n
		}
		if n := each.Label(); n != "" {
			label = n
		}
		if n := each.Id(); n != "" {
			description = n
		}
		if n := each.Help(); n != "" {
			help = n
		}
	}
	return NewSimpleMetadata(id, label, description, help).Metadata()
}
