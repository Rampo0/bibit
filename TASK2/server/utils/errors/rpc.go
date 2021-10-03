package errors

type rpcErr struct {
	message string
}

func NewRPCErr(message string) *rpcErr {
	return &rpcErr{message}
}

func (err *rpcErr) Error() string {
	return err.message
}
