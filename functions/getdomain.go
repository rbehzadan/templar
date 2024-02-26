package functions

import (
	"os"
	"strings"
)

func GetDomainName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "exmaple.ir"
	}
	parts := strings.Split(hostname, ".")
	if len(parts) < 2 {
		return "exmaple.ir"
	}
	return strings.Join(parts[1:], ".")
}
