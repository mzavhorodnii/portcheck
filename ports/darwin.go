//go:build darwin

package ports

import (
	"os/exec"
	"strings"
)

func getPlatformPorts() ([]PortInfo, error) {
	tcpOut, err := exec.Command("netstat", "-anv", "-p", "tcp").Output()
	if err != nil {
		return nil, err
	}
	udpOut, err := exec.Command("netstat", "-anv", "-p", "udp").Output()
	if err != nil {
		return nil, err
	}
	var result []PortInfo
	result = append(result, parseNetstat(string(tcpOut), "TCP")...)
	result = append(result, parseNetstat(string(udpOut), "UDP")...)
	return result, nil
}

func parseNetstat(output string, protocol string) []PortInfo {
	var result []PortInfo

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 9 {
			continue
		}

		proto := fields[0]
		if !strings.HasPrefix(proto, "tcp") && !strings.HasPrefix(proto, "udp") {
			continue
		}

		port := extractPort(fields[3])
		if port == 0 {
			continue
		}

		var status string
		var processField int

		if strings.HasPrefix(proto, "tcp") {
			status = fields[5]
			processField = 10
		} else {
			status = ""
			processField = 9
		}

		if len(fields) <= processField {
			continue
		}

		info := PortInfo{
			Protocol: protocol,
			Port:     port,
			Address:  fields[3],
			Status:   status,
		}

		info.Process, info.PID = parseProcess(fields[processField])

		result = append(result, info)
	}

	return result
}

func extractPort(address string) int {
	parts := strings.Split(address, ".")
	if len(parts) < 2 {
		return 0
	}
	last := parts[len(parts)-1]
	port := 0
	for _, c := range last {
		if c < '0' || c > '9' {
			return 0
		}
		port = port*10 + int(c-'0')
	}
	return port
}

func parseProcess(field string) (string, int) {
	parts := strings.Split(field, ":")
	if len(parts) != 2 {
		return field, 0
	}
	pid := 0
	for _, c := range parts[1] {
		if c < '0' || c > '9' {
			return parts[0], 0
		}
		pid = pid*10 + int(c-'0')
	}
	return parts[0], pid
}
