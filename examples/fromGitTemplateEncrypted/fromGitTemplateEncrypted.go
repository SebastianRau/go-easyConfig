package main

import (
	"flag"
	"fmt"

	"github.com/sebastianRau/go-easyConfig/pkg/demo"
	"github.com/sebastianRau/go-easyConfig/pkg/gitTools"

	easyconfig "github.com/sebastianRau/go-easyConfig/pkg/easyConfig"
)

func main() {

	const (
		configTemplateName = "easyConfigDemo.template.yaml"
		configDataName     = "easyConfigDemo.data.encrypted.yaml"
	)

	var (
		demoCfg           demo.DemoConfig
		encryptionKeyPath = flag.String("k", "", "Encryption Key File")
	)
	flag.Parse()

	err := easyconfig.TemplateFromGit(
		&gitTools.GitConfig{
			ProviderUrl: "github.com",
			RepoUrl:     "https://github.com/SebastianRau/go-easyConfigDemo/",
			Branch:      "master",
			PemBytes:    []byte{},
			PemPassword: ""},
		configTemplateName,
		configDataName,
		&demoCfg)
	demo.CheckError(err)

	err = easyconfig.EncryptFromFile(*encryptionKeyPath, &demoCfg)
	demo.CheckError(err)

	fmt.Println(demoCfg.String())
}
