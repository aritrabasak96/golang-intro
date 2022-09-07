/*
    ~ First we need to create a package or module.
	  go mod init <package name>

*/

package main

import (
	"fmt"

	"myproject/helper" // myproject is the module name that you have created: go mod init myproject
	"myproject/utils"
)

func main() {

	fmt.Println("main function")

	// to call those functions

	helper.MyHelperFunc()
	utils.MyUtil()
}
