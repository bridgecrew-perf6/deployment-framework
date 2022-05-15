package components

import (
	"fmt"
	"log"
	"os/exec"
)

func Build(folderPath string) {
	cmd := exec.Command("go", "build -ldflags -s -w")
	cmd.Dir = folderPath
	log.Printf("Running command and waiting for it to finish...")
	out, err := cmd.Output()
	log.Printf("Command finished with error: %v", err)
	fmt.Println(string(out))
}
