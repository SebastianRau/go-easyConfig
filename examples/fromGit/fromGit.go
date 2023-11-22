package main

import (
	"fmt"

	"github.com/sebastianrau/go-easyConfig/pkg/demo"
	"github.com/sebastianrau/go-easyConfig/pkg/gitTools"

	easyconfig "github.com/sebastianrau/go-easyConfig/pkg/easyConfig"
)

func main() {

	const (
		configFileName = "easyConfigDemo.yaml"
	)

	var (
		demoCfg demo.DemoConfig
	)

	err := easyconfig.FromGit(
		&gitTools.GitConfig{
			ProviderUrl: "github.com",
			RepoUrl:     "https://github.com/SebastianRau/go-easyConfigDemo/",
			Branch:      "master",
			PemBytes:    []byte{},
			PemPassword: ""},
		configFileName,
		&demoCfg)

	demo.CheckError(err)
	fmt.Println(demoCfg.String())
}
