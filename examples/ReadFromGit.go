package main

import (
	"fmt"

	"github.com/sebastianRau/go-easyConfig/pkg/demo"
	"github.com/sebastianRau/go-easyConfig/pkg/gitTools"

	easyconfig "github.com/sebastianRau/go-easyConfig/pkg/easyConfig"
)

func main() {

	const (
		configFileName = "test.yaml"
	)

	var demoCfg demo.DemoConfig

	err := easyconfig.ReadConfigFromGit(
		&gitTools.GitConfig{
			ProviderUrl: "github.com",
			RepoUrl:     "",
			Branch:      "",
			PemBytes:    []byte{},
			PemPassword: ""},
		configFileName,
		demoCfg)

	demo.CheckError(err)
	fmt.Println(demoCfg.String())
}
