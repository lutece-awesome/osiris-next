package constant

import "flag"

var (
	Port = flag.Int("port", 9876, "The gRPC server port")
	Work = flag.String("work", "work", "The working directory")
)