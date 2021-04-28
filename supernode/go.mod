module github.com/pastelnetwork/gonode/supernode

go 1.16

require (
	github.com/nats-io/nats-server/v2 v2.2.1
	github.com/pastelnetwork/gonode/common v0.0.0
	github.com/pastelnetwork/gonode/pastel-client v0.0.0
	github.com/pastelnetwork/gonode/proto v0.0.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.37.0
)

replace github.com/pastelnetwork/gonode/common => ../common
replace github.com/pastelnetwork/gonode/proto => ../proto
replace github.com/pastelnetwork/gonode/pastel-client => ../pastel-client