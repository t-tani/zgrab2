// Package quasar provides a zgrab2 module that scans for Quasar C2 server.
// Default port: 4782 (TCP)
//
// Connects, and reads the initial tcp data payload.

package quasar

import (
	log "github.com/sirupsen/logrus"
	"github.com/zmap/zgrab2"
)

// Flags holds the command-line configuration for the Quasar scan module.
// Populated by the framework.
type Flags struct {
	zgrab2.BaseFlags
	Verbose     bool `long:"verbose" description:"More verbose logging, include debug fields in the scan results"`
}

// Module implements the zgrab2.Module interface.
type Module struct {
}

// Scanner implements the zgrab2.Scanner interface.
type Scanner struct {
	config *Flags
}

// RegisterModule registers the zgrab2 module.
func RegisterModule() {
	var module Module
	_, err := zgrab2.AddCommand("quasar", "quasar", module.Description(), 4782, &module)
	if err != nil {
		log.Fatal(err)
	}
}

// NewFlags returns a default Flags object.
func (module *Module) NewFlags() interface{} {
	return new(Flags)
}

// NewScanner returns a new Scanner instance.
func (module *Module) NewScanner() zgrab2.Scanner {
	return new(Scanner)
}

// Description returns an overview of this module.
func (module *Module) Description() string {
	return "Probe for Quasar C2 Server"
}

// Validate checks that the flags are valid.
// On success, returns nil.
// On failure, returns an error instance describing the error.
func (flags *Flags) Validate(args []string) error {
	return nil
}

// Help returns the module's help string.
func (flags *Flags) Help() string {
	return ""
}

// Init initializes the Scanner.
func (scanner *Scanner) Init(flags zgrab2.ScanFlags) error {
	f, _ := flags.(*Flags)
	scanner.config = f
	return nil
}

// InitPerSender initializes the scanner for a given sender.
func (scanner *Scanner) InitPerSender(senderID int) error {
	return nil
}

// GetName returns the Scanner name defined in the Flags.
func (scanner *Scanner) GetName() string {
	return scanner.config.Name
}

// GetTrigger returns the Trigger defined in the Flags.
func (scanner *Scanner) GetTrigger() string {
	return scanner.config.Trigger
}

// Protocol returns the protocol identifier of the scan.
func (scanner *Scanner) Protocol() string {
	return "quasar"
}

// Scan probes for a  Quasar C2 server.
// 1. Opens a TCP connection to the configured port (default 4782)
// 2. Read server response
// 3. Check response size
// 4. Check tcp payload header.
func (scanner *Scanner) Scan(target zgrab2.ScanTarget) (zgrab2.ScanStatus, interface{}, error) {
	conn, err := target.Open(&scanner.config.BaseFlags)
	if err != nil {
		return zgrab2.TryGetScanStatus(err), nil, err
	}
	defer conn.Close()
	result := new(QuasarLog)

	err = checkInitResponse(result, conn)
	if !result.IsQuasar {
		result = nil
	}
	return zgrab2.TryGetScanStatus(err), result, err
}
