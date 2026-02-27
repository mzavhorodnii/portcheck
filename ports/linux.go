//go:build linux

package ports

import (
	"os/exec"
	"strings"
)

func getPlatformPorts() ([]PortInfo, error) {
	tcpOut, err := exec.Command("netstat", "-tnlp").Output()
	if err != nil {
		return nil, err
	}

	udpOut, err := exec.Command("netstat", "-unlp").Output()
	if err != nil {
		return nil, err
	}

	var result []PortInfo
	result = append(result, parseLinux(string(tcpOut), "TCP")...)
	result = append(result, parseLinux(string(udpOut), "UDP")...)

	return result, nil
}

func parseLinux(output string, protocol string) []PortInfo {
	var result []PortInfo

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 6 {
			continue
		}

		if fields[0] != "tcp" && fields[0] != "tcp6" &&
			fields[0] != "udp" && fields[0] != "udp6" {
			continue
		}

		port := extractPort(fields[3])
		if port == 0 {
			continue
		}

		info := PortInfo{
			Protocol: protocol,
			Port:     port,
			Address:  fields[3],
			Status:   fields[5],
		}

		if len(fields) >= 7 {
			info.Process, info.PID = parseProcess(fields[6])
		}

		result = append(result, info)
	}

	return result
}
