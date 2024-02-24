package functions

import (
	"fmt"
	"net"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Title(s string) string {
	caser := cases.Title(language.AmericanEnglish)
	return caser.String(s)
}

func Split(sep, s string) []string {
	return strings.Split(s, sep)
}

// getDomainName attempts to find the domain name of the host.
// It returns the domain name if found, otherwise an error.
func getDomainName() (string, error) {
	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("error getting hostname: %v", err)
	}

	// Lookup IP addresses associated with the hostname
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return "", fmt.Errorf("error looking up IP addresses for hostname %s: %v", hostname, err)
	}

	for _, ip := range ips {
		// Perform a reverse lookup to find the FQDN
		names, err := net.LookupAddr(ip.String())
		if err != nil {
			// If there's an error, continue to the next IP
			continue
		}

		for _, name := range names {
			cleanName := strings.TrimSuffix(name, ".")
			return cleanName, nil
		}
	}

	return "", fmt.Errorf("domain name not found for host: %s", hostname)
}

func GetDomainName() string {
	domain, _ := getDomainName()
	return domain
}
