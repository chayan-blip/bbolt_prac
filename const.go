package bolt

const Version = 1

const magic int32 = 0xBEEFCODE

const (
	MaxKeySize = 511
	MaxDataSize = 0xffffffff
)

const (
	DefaultMapSize = 1048576
	DefaultReaderCount = 126
)