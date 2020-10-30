package client

import "encoding/xml"

// Defines structs describing netconf configuration.

// Config defines properties that configure netconf session behaviour.
type Config struct {
	// Defines the time in seconds that the client will wait to receive a hello message from the server.
	SetupTimeoutSecs int

	// Attrs is used to set the attributes of the <rpc> element in RPCMessage.
	Attrs []xml.MarshalerAttr
}

var DefaultConfig = &Config{
	SetupTimeoutSecs: 5,
}
