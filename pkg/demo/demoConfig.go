package demo

import (
	"fmt"
	"os"
)

type DemoConfig struct {
	Hello  string
	World  string
	Secret string
}

func (d *DemoConfig) String() string {
	return fmt.Sprintf("%s %s! My Secret is: %s", d.Hello, d.World, d.Secret)
}

func CheckError(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}
