package application

import "fmt"

type BuilderNotFoundError struct {
	key string
}

func (bnf *BuilderNotFoundError) Error() string {
	return fmt.Sprintf("Builder not found : %s", bnf.key)
}
