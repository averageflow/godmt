package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	correctSplitPackageLength = 3
)

const (
	correctFolderPermissions = 0744
)

func WriteResultToFile(result string, filename string, packageDeclaration []string) {
	if len(packageDeclaration) == correctSplitPackageLength {
		// If the package is from another folder then we will create the needed folder
		// else we simply don't need any packages
		_ = os.Mkdir(fmt.Sprintf("./result/%s", packageDeclaration[1]), os.FileMode(correctFolderPermissions))
	}

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	_, err = f.WriteString(result)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateResultFolder() {
	err := os.Mkdir("./result", os.FileMode(correctFolderPermissions))
	if errors.Is(err, os.ErrExist) {
		fmt.Println("Skipping folder creation since folder existed!")
	} else {
		log.Fatalln(err.Error())
	}
}
