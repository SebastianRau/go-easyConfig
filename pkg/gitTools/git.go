package gitTools

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-billy/v5"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
)

type GitConfig struct {
	ProviderUrl string
	RepoUrl     string
	Branch      string
	PemBytes    []byte
	PemPassword string
}

func CloneGitRepoSsh(cfg *GitConfig, worktree *billy.Filesystem) error {

	CheckKnownHosts(cfg)

	gitCloneOptions := &git.CloneOptions{
		URL:               cfg.RepoUrl,
		RemoteName:        cfg.Branch,
		SingleBranch:      true,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}

	if cfg.PemBytes != nil && len(cfg.PemBytes) > 0 {
		publicKeys, err := ssh.NewPublicKeys("git", cfg.PemBytes, cfg.PemPassword)
		if err != nil {
			return err
		}
		gitCloneOptions.Auth = publicKeys
	}

	_, err := git.Clone(memory.NewStorage(), *worktree, gitCloneOptions)
	if err != nil {
		return err
	}

	return nil
}

func CheckKnownHosts(cfg *GitConfig) error {

	homeDir, _ := os.UserHomeDir()
	knownHosts := homeDir + "/.ssh/known_hosts"
	bytes, err := os.ReadFile(knownHosts)
	if err != nil {
		return err
	}

	if strings.Contains(string(bytes), cfg.ProviderUrl) {
		return nil
	}

	return fmt.Errorf("%s not found in %s", cfg.ProviderUrl, knownHosts)

}
