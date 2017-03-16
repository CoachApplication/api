package operation

// Operations provider of ordered sets of operations
type Operations interface {
	Add(Operation)
	Get(string) (Operation, error)
	Order() []string
}
