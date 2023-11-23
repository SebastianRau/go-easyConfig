package easyconfig

import (
	"fmt"
	"io"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/sebastianrau/go-easyConfig/pkg/configParser"
	"github.com/sebastianrau/go-easyConfig/pkg/encryption"
	"github.com/sebastianrau/go-easyConfig/pkg/gitTools"
	"github.com/sebastianrau/go-easyConfig/pkg/templating"
)

func FromRaw(config []byte, out interface{}) error {
	return configParser.ReadString(config, out)
}
func FromFile(filePath string, out interface{}) error {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return FromRaw(buf, out)
}
func FromGit(gitConfig *gitTools.GitConfig, fileName string, out interface{}) error {
	fs := memfs.New()

	err := RawGit(gitConfig, &fs)
	if err != nil {
		return err
	}

	buf, err := memFsReadFile(&fs, fileName)
	if err != nil {
		return err
	}

	return FromRaw(buf, out)
}

func RawGit(gitConfig *gitTools.GitConfig, fs *billy.Filesystem) error {
	return gitTools.CloneGitRepoSsh(gitConfig, fs)
}

func TemplateFromRaw(template []byte, data []byte, out interface{}) error {

	parsed, err := templating.ParseTemplate(template, data)
	if err != nil {
		return err
	}

	err = configParser.ReadString(parsed, out)
	if err != nil {
		return err
	}

	return nil
}
func TemplateFromFile(templateFilePath string, dataFilePath string, out interface{}) error {

	templatebuffer, err := os.ReadFile(templateFilePath)
	if err != nil {
		return err
	}

	dataBuffer, err := os.ReadFile(dataFilePath)
	if err != nil {
		return err
	}

	return TemplateFromRaw(templatebuffer, dataBuffer, out)
}
func TemplateFromGit(gitConfig *gitTools.GitConfig, templateFile string, dataFile string, out interface{}) error {
	fs := memfs.New()

	err := RawGit(gitConfig, &fs)
	if err != nil {
		return err
	}

	templatebuffer, err := memFsReadFile(&fs, templateFile)
	if err != nil {
		return err
	}

	dataBuffer, err := memFsReadFile(&fs, dataFile)
	if err != nil {
		return err
	}

	return TemplateFromRaw(templatebuffer, dataBuffer, out)
}

func DecryptFromRaw(key []byte, cfg interface{}) error {
	err := encryption.DecryptConfig(cfg, key)
	if err != nil {
		return err
	}
	return nil
}
func DecryptFromFile(encryptionKeyPath string, cfg interface{}) error {
	if encryptionKeyPath == "" {
		return fmt.Errorf("no encryption key found")
	}

	key, err := os.ReadFile(encryptionKeyPath)
	if err != nil {
		return err
	}

	return DecryptFromRaw(key, cfg)
}

func memFsReadFile(fs *billy.Filesystem, file string) ([]byte, error) {
	f, err := (*fs).Open(file)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return []byte{}, err
	}
	return buf, nil
}
