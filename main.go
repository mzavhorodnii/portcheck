package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/mzavhorodnii/portcheck/display"
	"github.com/mzavhorodnii/portcheck/ports"
)

func main() {
	tcpOnly := flag.Bool("tcp", false, "Show only TCP ports")
	udpOnly := flag.Bool("udp", false, "Show only UDP ports")
	status := flag.String("status", "", "Filter by status (LISTEN, ESTABLISHED)")

	flag.Parse()

	if !*tcpOnly && !*udpOnly {
		*tcpOnly = true
		*udpOnly = true
	}

	filter := ports.Filter{
		TCP:    *tcpOnly,
		UDP:    *udpOnly,
		Status: *status,
	}

	args := flag.Args()
	if len(args) > 0 {
		port, err := strconv.Atoi(args[0])
		if err != nil || port < 1 || port > 65535 {
			fmt.Fprintf(os.Stderr, "Error: invalid port number %q\n", args[0])
			os.Exit(1)
		}
		filter.Port = port
	}

	results, err := ports.GetPorts(filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if len(results) == 0 {
		if filter.Port != 0 {
			fmt.Printf("Port %d is not in use\n", filter.Port)
		} else {
			fmt.Printf("No ports found\n")
		}
		return
	}
	display.Render(results)
	if filter.Port != 0 {
		fmt.Printf("\nPort %d is used by %s (PID %d)\n",
			results[0].Port,
			results[0].Process,
			results[0].PID,
		)
	}
}
