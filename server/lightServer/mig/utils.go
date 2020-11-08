package mig

import (
	"strings"
)

func StrsContain(testees []string, tester string) bool {
	for _, testee := range testees {
		if tester == testee {
			return true
		}
	}
	return false
}

func MimeFile(tester string) string {
	return strings.ReplaceAll(tester, "/", "_")
}
