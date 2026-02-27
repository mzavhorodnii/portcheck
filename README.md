# portcheck

A simple CLI tool to inspect active network ports — see what's running, on which port, and which process owns it.

## Install
```bash
go install github.com/mzavhorodnii/portcheck@latest
```

## Usage
```bash
# Show all active ports
portcheck

# Check a specific port
portcheck 5432

# Filter by protocol
portcheck --tcp
portcheck --udp

# Filter by status
portcheck --status LISTEN
portcheck --status ESTABLISHED

# Combine filters (flags before port number)
portcheck --tcp --status LISTEN 8080
```

## Example output
```
PROTO   PORT    PID     PROCESS            STATUS        ADDRESS
-----   ----    ---     -------            ------        -------
TCP     5432    513     postgres           LISTEN        *.5432
TCP     8080    12453   node               LISTEN        127.0.0.1.8080
UDP     5353    500     mDNSResponder                    *.5353
```

## Platforms

- macOS
- Linux
- Windows

## Requirements

Go 1.21+

Make sure your Go bin directory is in PATH:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Add this to your `~/.zshrc` or `~/.bashrc` to make it permanent.