package rpctype

// Server is an interface for RPC servers to use
type Server interface {
	Start(address string) error
}
