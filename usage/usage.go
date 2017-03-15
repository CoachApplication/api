package usage

type Usage interface {
	Allows(string) bool
}
