package greet

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in", name)
		}
	}
}
