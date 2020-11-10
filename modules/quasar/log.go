package quasar

// QuasarLog is the struct returned to the caller.
type QuasarLog struct {
	// IsQuasar should always be true (otherwise, the result should have been nil).
	IsQuasar bool `json:"is_quasar"`

	// Version corresponds to the "Quasar.version" response field.
	Version string `json:"version"`

	// Id corresponds to the "id" response field, which is decoded as decimal integer
	// Id uint32 `json:"id"`
	
	// Hostname corresponds to the "hostName" field.
	// Hostname string `json:"hostname,omitempty"`

	// HostAddress corresponds to the "hostAddress" field.
	//HostAddress string `json:"host_address,omitempty"`
}