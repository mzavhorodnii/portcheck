package ports

import "fmt"

type PortInfo struct {
	Protocol string
	Port     int
	PID      int
	Process  string
	Status   string
	Address  string
}
type Filter struct {
	Port   int
	TCP    bool
	UDP    bool
	Status string
}

func GetPorts(filter Filter) ([]PortInfo, error) {
	all, err := getPlatformPorts()
	if err != nil {
		return nil, err
	}
	return applyFilter(all, filter), nil
}

func applyFilter(ports []PortInfo, filter Filter) []PortInfo {
	var result []PortInfo
	for _, p := range ports {
		if filter.Port != 0 && p.Port != filter.Port {
			continue
		}
		if p.Protocol == "TCP" && !filter.TCP {
			continue
		}
		if p.Protocol == "UDP" && !filter.UDP {
			continue
		}
		if filter.Status != "" && p.Status != filter.Status {
			continue
		}
		result = append(result, p)
	}
	result = deduplicate(result)
	return result
}

func deduplicate(ports []PortInfo) []PortInfo {
	seen := make(map[string]bool)
	var result []PortInfo

	for _, p := range ports {
		key := fmt.Sprintf("%d-%d-%s", p.Port, p.PID, p.Status)
		if !seen[key] {
			seen[key] = true
			result = append(result, p)
		}
	}

	return result
}
