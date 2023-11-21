package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sebastianRau/go-easyConfig/pkg/demo"

	easyconfig "github.com/sebastianRau/go-easyConfig/pkg/easyConfig"
)

func main() {

	const (
		configTemplateName = "../go-easyConfigDemo/easyConfigDemo.template.yaml"
		configDataName     = "../go-easyConfigDemo/easyConfigDemo.data.encrypted.yaml"
	)

	var (
		demoCfg demo.DemoConfig
	)

	var (
		encryptionKeyPath = flag.String("k", "", "Encryption Key File")
	)
	flag.Parse()

	err := easyconfig.TemplateFromFile(
		configTemplateName,
		configDataName,
		&demoCfg)

	demo.CheckError(err)

	if *encryptionKeyPath == "" {
		fmt.Println("no encryption key found")
		os.Exit(-1)
	}

	err = easyconfig.EncryptFromFile(*encryptionKeyPath, &demoCfg)
	if err != nil {
		demo.CheckError(err)
	}

	fmt.Println(demoCfg.String())
}
