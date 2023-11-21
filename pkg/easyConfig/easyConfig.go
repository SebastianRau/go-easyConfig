package easyconfig

import (
	"io"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/sebastianRau/go-easyConfig/pkg/configParser"
	"github.com/sebastianRau/go-easyConfig/pkg/gitTools"
	"github.com/sebastianRau/go-easyConfig/pkg/templating"
)

func ConfigFromGit(gitConfig *gitTools.GitConfig, fileName string, out interface{}) error {
	fs := memfs.New()

	err := gitTools.CloneGitRepoSsh(gitConfig, &fs)
	if err != nil {
		return err
	}

	f, err := fs.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = configParser.ReadString(buf, out)
	if err != nil {
		return err
	}
	return nil
}

func TemplateFromGit(gitConfig *gitTools.GitConfig, templateFile string, dataFile string, out interface{}) error {
	fs := memfs.New()

	err := gitTools.CloneGitRepoSsh(gitConfig, &fs)
	if err != nil {
		return err
	}

	ft, err := fs.Open(templateFile)
	if err != nil {
		return err
	}
	defer ft.Close()

	templatebuffer, err := io.ReadAll(ft)
	if err != nil {
		return err
	}

	fd, err := fs.Open(dataFile)
	if err != nil {
		return err
	}
	defer fd.Close()

	dataBuffer, err := io.ReadAll(fd)
	if err != nil {
		return err
	}

	parsed, err := templating.ParseTemplate(templatebuffer, dataBuffer)
	if err != nil {
		return err
	}

	err = configParser.ReadString(parsed, out)
	if err != nil {
		return err
	}

	return nil
}
