package components

import (
	"local/installer/global"
	"os"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"gopkg.in/src-d/go-git.v4"
	x "gopkg.in/src-d/go-git.v4/_examples"
)

func DownloadRepoSrc(url string, name string, privateKey string, password string) {

	privateKeyData := []byte(privateKey)

	// Clone the given repository to the given directory
	x.Info("git clone %s %s --recursive", url, name)

	publicKeys, keyErr := ssh.NewPublicKeys("git", privateKeyData, password)

	if keyErr != nil {
		x.Warning("generate publickeys failed: %s\n", keyErr)
		return
	}
	_, cloneErr := git.PlainClone(global.SrcRootDir+"/"+name, false, &git.CloneOptions{
		URL:               url,
		Auth:              publicKeys,
		Progress:          os.Stdout,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	x.CheckIfError(cloneErr)

}

func DownloadPublicRepoSrc(url string, name string) {

	// Clone the given repository to the given directory
	x.Info("git clone %s %s --recursive", url, name)

	_, cloneErr := git.PlainClone(global.SrcRootDir+"/"+name, false, &git.CloneOptions{
		URL:               url,
		Progress:          os.Stdout,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	x.CheckIfError(cloneErr)

}
