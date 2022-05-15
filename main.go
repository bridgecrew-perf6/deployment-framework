package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"local/installer/components"
	"os"
)

func setEncryptionKeyEnv(keyLength int, keyName string) {
	bytes := make([]byte, keyLength) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes) //encode key in bytes to string
	os.Setenv(keyName, key)
	fmt.Println("key :", os.Getenv(keyName))
}

func main() {
	setEncryptionKeyEnv(32, "AESKEY")
	setEncryptionKeyEnv(20, "ACCESS_SECRET")
	setEncryptionKeyEnv(20, "REFRESH_SECRET")
	descriptor := flag.String("descriptor", "latest", "docker image descriptor")
	hostPtr := flag.String("host", "localhost", "host / ip of the machine hosting the applications")
	domainPtr := flag.String("domain", "localhost", "issueing cookie domain.")
	idpUiPtr := flag.String("idpui", "https://localhost:8888", "idp ui url.")
	mysqlHostPtr := flag.String("mysqlhost", "localhost", "mysql server ip / name for collection service")
	mysqlUserPtr := flag.String("mysqluser", "root", "mysql user for collection service")
	mysqlPassPtr := flag.String("mysqlpass", "", "mysql password for collection service")
	idpSvcPtr := flag.String("idpsvc", "https://localhost:4000", "idp service url")
	releasePtr := flag.String("mode", "debug", "select mode either release or debug")
	redisURLPtr := flag.String("redis-host", "localhost", "redis ip or fqdn")
	securePtr := flag.String("secure-mode", "true", "enabled if services are configured with tls / ssl ")
	components.PrepareDeployment()
	flag.Parse()
	// const password = ""
	// enable privateKeyData for private repo via a pub/priv key
	// privateKeyData := components.GetPrivateKeyData()
	var repoFile = components.GetRepoMetaData()

	for _, r := range repoFile.Repositories {
		//enable this for private github repo access
		//components.DownloadRepoSrc(r.Repo, r.Name, privateKeyData, password)
		components.DownloadPublicRepoSrc(r.Repo, r.Name)

	}
	as := components.TokenAuthSrc(*releasePtr, *hostPtr, *domainPtr, *idpUiPtr, *securePtr, *redisURLPtr, *mysqlHostPtr, *mysqlUserPtr, *mysqlPassPtr)
	dc := components.DockerComposeSrc()
	bs := components.BuildScript(*descriptor)
	db := components.ComposeBuildScript()
	dp := components.WebIDPFileSrc(*idpSvcPtr)
	components.CreateFolderFile(dc, ".", "docker-compose.yml")
	components.CreateFolderFile(bs, "src", "source2ContainerBuild.sh")
	components.CreateFolderFile(as, "src/token-auth-svc", "Dockerfile")
	components.CreateFolderFile(dp, "src/idp", "Dockerfile")
	components.CreateFolderFile(db, ".", "deploy_build.sh")

}
