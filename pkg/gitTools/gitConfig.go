package gitTools

type GitConfig struct {
	ProviderUrl string
	RepoUrl     string
	Branch      string
	PemBytes    []byte
	PemPassword string
}
