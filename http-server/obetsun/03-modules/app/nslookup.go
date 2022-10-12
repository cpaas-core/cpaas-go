package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
)

func NslookupHandler(w http.ResponseWriter, r *http.Request) {

	reg, _ := regexp.Compile(`^nslookup/((([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]))$`)
	match := reg.FindStringSubmatch(r.URL.Path[1:])

	if match != nil {
		fmt.Fprintf(w, "Server: %s\n", match[1])
		fmt.Fprintf(w, "\n")

		NSs, err := net.LookupNS(match[1])
		if err != nil {
			fmt.Fprintf(w, "Could not get NS record: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(w, "NS: \n")

		for _, ns := range NSs {
			fmt.Fprintf(w, "%s\n", ns.Host)
		}

		fmt.Fprintf(w, "\n")

		MXs, err := net.LookupMX(match[1])
		if err != nil {
			fmt.Fprintf(w, "Could not get NS record: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(w, "MX: \n")

		for _, mx := range MXs {
			fmt.Fprintf(w, "%s\n", mx.Host)
		}

		fmt.Fprintf(w, "\n")

		ips, err := net.LookupIP(match[1])
		if err != nil {
			fmt.Fprintf(w, "Could not get IPs: %v\n", err)
			os.Exit(1)
		}

		for _, ip := range ips {
			fmt.Fprintf(w, "Address: %s\n", ip.String())
		}
	} else {
		fmt.Fprintf(w, "Unknown parameter: %s!", r.URL.Path[1:])
	}
}
