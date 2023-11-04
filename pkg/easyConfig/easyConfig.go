package easyconfig

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/sebastianRau/easyConfig/pkg/configParser"
	"github.com/sebastianRau/easyConfig/pkg/gitTools"
)

func ReadConfigFromGit(gitConfig *gitTools.GitConfig, fileName string, out interface{}) error {
	fs := memfs.New()

	err := gitTools.CloneGitRepoSsh(gitConfig, &fs)
	if err != nil {
		return err
	}

	configTemplate, err := fs.Open(fileName)
	if err != nil {
		return err
	}

	buf := make([]byte, 0)
	_, err = configTemplate.Read(buf)
	if err != nil {
		return err
	}

	configParser.ReadString(string(buf), out)
	return nil
}
