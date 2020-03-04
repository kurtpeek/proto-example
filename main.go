package main

import (
	"fmt"

	"github.com/kurtpeek/proto-example/result"
)

func main() {
	res := result.Result{Status: result.Result_STATUS_FAILED}
	fmt.Printf("%+v\n", res)
}
