package main

import (
	"fmt"

	"github.com/sebastianRau/go-easyConfig/pkg/demo"
	"github.com/sebastianRau/go-easyConfig/pkg/gitTools"

	easyconfig "github.com/sebastianRau/go-easyConfig/pkg/easyConfig"
)

func main() {

	const (
		configTemplateName = "easyConfigDemo.template.yaml"
		configDataName     = "easyConfigDemo.data.yaml"
	)

	var (
		demoCfg demo.DemoConfig
	)

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
	fmt.Println(demoCfg.String())
}
