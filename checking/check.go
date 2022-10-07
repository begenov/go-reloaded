package checing

import (
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {

		fmt.Println("Incorrect values")
		os.Exit(0)

	}
}
