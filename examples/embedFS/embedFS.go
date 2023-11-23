package main

import (
	"embed"
	"fmt"

	"github.com/sebastianrau/go-easyConfig/pkg/demo"
	"github.com/sebastianrau/go-easyConfig/pkg/gitTools"

	easyconfig "github.com/sebastianrau/go-easyConfig/pkg/easyConfig"
)

//go:embed keys/*
var keyFiles embed.FS

func main() {

	const (
		pemFile        = "keys/demoDeployKey"                 //from embedFS
		encryptionFile = "keys/testkey"                       //from embedFS
		dataName       = "easyConfigDemo.data.encrypted.yaml" //from Git
		templateName   = "easyConfigDemo.template.yaml"       //from Git
	)

	var (
		demoCfg demo.DemoConfig
	)

	pemBytes, err := keyFiles.ReadFile(pemFile)
	demo.CheckError(err)

	encryptionKey, err := keyFiles.ReadFile(encryptionFile)
	demo.CheckError(err)

	err = easyconfig.TemplateFromGit(
		&gitTools.GitConfig{
			ProviderUrl: "github.com",
			RepoUrl:     "git@github.com:SebastianRau/go-easyConfigDemo.git",
			Branch:      "master",
			PemBytes:    pemBytes,
			PemPassword: ""},
		templateName,
		dataName,
		&demoCfg)
	demo.CheckError(err)

	err = easyconfig.DecryptFromRaw(encryptionKey, &demoCfg)
	if err != nil {
		demo.CheckError(err)
	}

	fmt.Println(demoCfg.String())
}
