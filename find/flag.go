package find

import (
	"flag"
)

func Flag(Info *HostInfo) {
	flag.StringVar(&Info.Host, "h", "", "IP Address: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12")
	flag.StringVar(&Info.Ports, "p", "50000", "Ports Range: 50000 | 40000-65535 | 49440,51731,52841")
	flag.StringVar(&Info.Type, "t", "", "Choose Attack Type: scan | rce")
	flag.StringVar(&Info.Cmd, "c", "", "Input Cmd Command")
	flag.IntVar(&Threads, "n", 600, "Set Scan Threads")
	flag.Parse()
}
