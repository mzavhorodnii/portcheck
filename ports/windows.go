//go:build windows

package ports

import (
	"os/exec"
	"strings"
)

func getPlatformPorts() ([]PortInfo, error) {
	out, err := exec.Command("netstat", "-ano").Output()
	if err != nil {
		return nil, err
	}

	return parseWindows(string(out)), nil
}

func parseWindows(output string) []PortInfo {
	var result []PortInfo

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 4 {
			continue
		}

		var protocol string
		switch fields[0] {
		case "TCP":
			protocol = "TCP"
		case "UDP":
			protocol = "UDP"
		default:
			continue
		}

		port := extractPort(fields[1])
		if port == 0 {
			continue
		}

		info := PortInfo{
			Protocol: protocol,
			Port:     port,
			Address:  fields[1],
		}

		if protocol == "TCP" && len(fields) >= 4 {
			info.Status = fields[3]
			if len(fields) >= 5 {
				info.PID = atoiSafe(fields[4])
			}
		} else if protocol == "UDP" && len(fields) >= 3 {
			info.PID = atoiSafe(fields[3])
		}

		result = append(result, info)
	}

	return result
}

func atoiSafe(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c-'0')
	}
	return n
}
