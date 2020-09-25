package main

import (
	_ "config-server/server/core"
	"config-server/server/global"
	"config-server/server/initialize"
	"fmt"
)

func main()  {

	initialize.Gorm()

	fmt.Println(global.GVA_CONFIG.Email.EmailHost)
}
