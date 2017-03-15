package operation

type Operations interface {
	Add(Operation)
	Get(string) (Operation, error)
	Order() []string
}
