package application

type SettingsProvider interface {
	unMarshall(interface{}) error
}
