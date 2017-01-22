package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mvdan/xurls"
)

func main() {
	// Get Data from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, tld := range xurls.TLDs {
			fmt.Println(fmt.Sprintf("%s.%s", scanner.Text(), tld))
		}
	}

	os.Exit(0)
}
