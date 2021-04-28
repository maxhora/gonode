module github.com/pastelnetwork/walletnode

go 1.16

require (
	github.com/dimfeld/httptreemux/v5 v5.3.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/pastelnetwork/gonode/common v0.0.0
	github.com/pastelnetwork/gonode/pastel-client v0.0.0
	github.com/pastelnetwork/gonode/proto v0.0.0
	github.com/sergi/go-diff v1.2.0 // indirect
	goa.design/goa/v3 v3.3.1
	goa.design/plugins/v3 v3.3.1
	golang.org/x/net v0.0.0-20210410081132-afb366fc7cd1 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	google.golang.org/grpc v1.37.0
)

replace github.com/pastelnetwork/gonode/common => ../common

replace github.com/pastelnetwork/gonode/proto => ../proto

replace github.com/pastelnetwork/gonode/pastel-client => ../pastel-client