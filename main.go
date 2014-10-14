package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func main() {
	routes, err := ioutil.ReadFile("/proc/net/route")

	if err != nil {
		die(err)
	}

	lines := strings.Split(string(routes), "\n")

	for _, line := range lines[1 : len(lines)-1] {
		fields := strings.Fields(line)

		if fields[0] == "eth0" && fields[1] == "00000000" {
			quads, err := hex.DecodeString(fields[2])

			if err != nil {
				die(err)
			}

			fmt.Printf("%d.%d.%d.%d\n", quads[3], quads[2], quads[1], quads[0])
			return
		}
	}
}
