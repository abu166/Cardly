package main

import (
	"strings"
)

func DisplayCardInfo(number string, brands, issuers map[string]string) (string, string, bool) {
	valid := LuhnCheck(number)

	brand, issuer := "-", "-"

	// Find the brand by prefix
	for code, name := range brands {
		if code == "" {
			continue
		}
		if strings.HasPrefix(number, code) {
			brand = name
			break
		} else {
			brand = "-"
		}
	}

	// Find the issuer by prefix
	for code, name := range issuers {
		if code == "" {
			continue
		}
		if strings.HasPrefix(number, code) {
			issuer = name
			break
		} else {
			issuer = "-"
		}
	}

	return brand, issuer, valid
}
