package display

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mzavhorodnii/portcheck/ports"
)

func Render(results []ports.PortInfo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "PROTO\tPORT\tPID\tPROCESS\tSTATUS\tADDRESS")
	fmt.Fprintln(w, "-----\t----\t---\t-------\t------\t-------")

	for _, r := range results {
		fmt.Fprintf(w, "%s\t%d\t%d\t%s\t%s\t%s\n",
			r.Protocol,
			r.Port,
			r.PID,
			r.Process,
			r.Status,
			r.Address,
		)
	}

	w.Flush()
}
