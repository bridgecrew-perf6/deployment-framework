package components

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/installer/global"
	"local/installer/models"
	"log"
	"os"
)

func PrepareDeployment() {
	err := os.RemoveAll(global.SrcRootDir)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFolderFile(fileName []byte, folderName string, nameOfFile string) {

	_, err := os.Stat(global.AppRoot + "/" + folderName)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(global.AppRoot+"/"+folderName, 0755)
		if errDir != nil {
			fmt.Println(err)
		}
	}
	e := ioutil.WriteFile(global.AppRoot+"/"+folderName+"/"+nameOfFile, fileName, 0644)
	if e != nil {
		panic(e)
	}
	fmt.Println("..file " + nameOfFile + " written to " + global.AppRoot + "/" + folderName)

}

func ReadConfigFile() []byte {
	fd, e := ioutil.ReadFile(global.RepositoriesFile)
	if e != nil {
		panic(e)
	}
	return fd
}

func GetRepoMetaData() models.RepositoryFile {
	fileInterface := models.RepositoryFile{}
	config := ReadConfigFile()
	json.Unmarshal(config, &fileInterface)
	return fileInterface
}
