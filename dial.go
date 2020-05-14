package kontext

import "google.golang.org/grpc"

func Dial(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithInsecure())
}
