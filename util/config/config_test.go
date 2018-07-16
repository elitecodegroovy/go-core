package config

import (
	"testing"
	"os"
)


func TestEnvVars(t *testing.T){
	type Config struct {
		Id int
	}
	os.Setenv("Id", "abc")
	conf := Config{}
	setFromEnvVariables(&conf)

	if conf.Id != 0 {
		t.Error("Id should be ", conf.Id)
	}
}
