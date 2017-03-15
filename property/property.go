package property

type Property interface {
	Id() string

	Get() interface{}
	Set(interface{}) error
}
