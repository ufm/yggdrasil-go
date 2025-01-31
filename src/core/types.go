package core

// In-band packet types
const (
	typeSessionDummy = iota // nolint:deadcode,varcheck
	typeSessionTraffic
	typeSessionProto
	typeSessionEncTraffic
)

// Protocol packet types
const (
	typeProtoDummy = iota
	typeProtoNodeInfoRequest
	typeProtoNodeInfoResponse
	typeProtoDebug = 255
)
