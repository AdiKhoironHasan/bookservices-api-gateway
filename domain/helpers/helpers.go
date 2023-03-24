package helpers

import (
	"fmt"
	"os"
)

func Debug(data ...interface{}) {
	fmt.Println("====== START DEBUG ======")
	fmt.Println()

	for _, v := range data {
		fmt.Println(fmt.Sprintf("%+v : %+v", v, v))
	}

	fmt.Println()
	fmt.Println("====== END DEBUG ======")
	os.Exit(0)
}
