module github.com/ashwinshirva/quickpicktools/to-jpg-worker

go 1.19

require (
	github.com/ashwinshirva/quickpicktools/proto-gen v0.0.0-20230317140104-a73abb31dc9d
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/grpc v1.53.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923 // indirect
	google.golang.org/protobuf v1.29.1 // indirect
)

//replace github.com/quickpicktools/proto-gen => ../proto-gen
