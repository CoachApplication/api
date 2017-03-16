package application

import (
	api_operation "github.com/james-nesbitt/coach-api/operation"
	api_result "github.com/james-nesbitt/coach-api/result"
)

// SimpleApplication a security oriented Application which makes all Operation authorization based
type SecuredApplication struct {
	SimpleApplication
}

// Application Explicitly converts this object to a Appication interface
func (sa *SecuredApplication) Application() Application {
	return Application(sa)
}

// Validate Ask if the Application is ready to provide Operations, which means that it has to be able to create
// authorizing wrapper operations.
func (sa *SecuredApplication) Validate() api_result.Result {
	res := api_result.NewStandardResult()

	ops := sa.SimpleApplication.Operations()
	if _, err := NewOpAuthorizeFromOperations(ops); err != nil {
		res.MarkFailed()
		res.AddError(err)
	}

	res.MarkFinished()
	return res.Result()
}

// Operations Retrieve a list of Operations from all of the Builders
func (sa *SecuredApplication) Operations() api_operation.Operations {
	ops := sa.SimpleApplication.Operations()

	sw, swErr := NewOpAuthorizeFromOperations(ops)

	for _, bid := range sa.builders.Order() {
		b, _ := sa.builders.Get(bid)
		bOps := b.Operations()
		for _, oid := range bOps.Order() {
			bOp, _ := bOps.Get(oid)

			if OpSkipsAuthorization(bOp) {
				ops.Add(bOp)
			} else if swErr == nil {
				bsOp := MakeSecuredOperation(*sw, bOp)
				ops.Add(bsOp)
			}
		}
	}

	return ops
}

type OpAuthorizer struct {
	op_User      api_operation.Operation
	op_Authorize api_operation.Operation
}

func NewOpAuthorizeFromOperations(ops api_operation.Operations) (*OpAuthorizer, error) {
	var userOp, authOp api_operation.Operation
	var err error

	if userOp, err = ops.Get("user"); err != nil {
		return nil, err
	}
	if authOp, err = ops.Get("auth"); err != nil {
		return nil, err
	}

	return &OpAuthorizer{
		op_User:      userOp,
		op_Authorize: authOp,
	}, nil
}

func NewOpAuthorize(userOp, authOp api_operation.Operation) *OpAuthorizer {
	return &OpAuthorizer{
		op_User:      userOp,
		op_Authorize: authOp,
	}
}

func OpSkipsAuthorization(op api_operation.Operation) bool {
	return false
}

// converts any operation to a secured operation,
func MakeSecuredOperation(sw OpAuthorizer, op api_operation.Operation) api_operation.Operation {
	return op
}
