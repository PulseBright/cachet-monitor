// Thanks go to https://github.com/Soontao/cachet-monitor/blob/master/tcp.go
package cachet

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

// Investigating template
var defaultTCPInvestigatingTpl = MessageTemplate{
	Subject: `{{ .Monitor.Name }} - {{ .SystemName }}`,
	Message: `{{ .Monitor.Name }} check **failed** (server time: {{ .now }})

{{ .FailReason }}`,
}

// Fixed template
var defaultTCPFixedTpl = MessageTemplate{
	Subject: `{{ .Monitor.Name }} - {{ .SystemName }}`,
	Message: `**Resolved** - {{ .now }}

Down seconds: {{ .downSeconds }}s`,
}

// TCPMonitor struct
type TCPMonitor struct {
	AbstractMonitor `mapstructure:",squash"`
	Port            int64
}

// CheckTCPPortAlive func
func CheckTCPPortAlive(ip string, port int64, timeout int64) (bool, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, strconv.FormatInt(port, 10)), time.Duration(timeout)*time.Second)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return false, err
	} else {
		return true, nil
	}

}

// test if it available
func (m *TCPMonitor) test() bool {
	if alive, e := CheckTCPPortAlive(m.Target, int64(m.Port), int64(m.Timeout)); alive {
		return true
	} else {
		m.lastFailReason = fmt.Sprintf("TCP check failed: %v", e)
		return false
	}
}

// Validate configuration
func (m *TCPMonitor) Validate() []string {

	// set incident temp
	m.Template.Investigating.SetDefault(defaultTCPInvestigatingTpl)
	m.Template.Fixed.SetDefault(defaultTCPFixedTpl)

	// super.Validate()
	errs := m.AbstractMonitor.Validate()

	if m.Target == "" {
		errs = append(errs, "Target is required")
	}

	return errs
}
