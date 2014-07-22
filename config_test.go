package jsonconfig

import (
	"log"
	"testing"
)

func TestLoadFile(t *testing.T) {
	filename := "./test/config.json"
	config := LoadFile(filename)
	if config == nil {
		t.Error("Expected non-nil config got ", config)
	}
	log.Printf("Config is: %v", config)
	if str, ok := config.GetStr("hello"); !ok {
		t.Error("Expected str value from config with key: str, got ", str)
	} else {
		log.Printf("string passed %s", str)
	}

	if mapStr, ok := config.GetMapStr("mapStr"); !ok {
		t.Error("Expected str value from config with key: mapStr, got ", mapStr)
	} else {
		log.Printf("Map[string]string passed %v", mapStr)
	}

	if strSlice, ok := config.GetStrSlice("strSlice"); !ok {
		t.Error("Expected value from config with key: strSlice, got ", strSlice)
	} else {
		log.Printf("Str Slice passed %v", strSlice)
	}

	log.Printf("Test SetEnvs")
	SetEnvs([]string{"Hello", "World", "Test"})
	config = LoadFile(filename)
	if config != nil {
		t.Error("Expect to get nil config ", config)
	} else {
		log.Printf("SetEnvs passed")
	}
}
