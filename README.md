jsonconfig
=========

Golang JSON config file reader + parser

This simple package read config out of json file.
It'll iterate through different possible environments and merge config from env with lower priority with the one exist in higher priority 

main.go:
~~~ go
package main
import (
	"github.com/ke-da/jsonconfig"
)

func main() {
	config := jsonconfig.LoadFile("config.json")
	config.GetStr("Hello") //"Dev World"
	config.GetMapStr("mapStr") //{"Key" => "Value"} map[string]string
	config.GetStrSlice("strSlice") //["Str1", "Str2"} []string
}
~~~

config.json:
~~~ json
{
	"dev": {
		"hello": "Dev World",
		"arr": ["dev1","dev2"],
		"mapStr": {
			"Key": "Value"
		},
		"strSlice": [
			"Str1",
			"Str2"
		]
	},
	"test":{
		"hello": "Test World",
		"arr": ["test1","test2"]
	}
}
~~~

