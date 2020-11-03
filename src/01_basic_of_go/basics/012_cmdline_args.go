package basics

import (
	"fmt"
	"os"
)

/* reading cmdline args and accessing it */

func MainCmdLineArgs() {
	arguements := os.Args
	for place, value := range arguements {
		fmt.Println(place, value)
	}
}
