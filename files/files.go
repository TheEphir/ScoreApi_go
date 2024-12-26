package files

import (
	"os"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	fileData, err := os.ReadFile("./db/" + fileName + ".json")
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

func WriteFile(fileName string, content []byte) error {
	file, err := os.Create("./db/" + fileName + ".json") //mb it's ok creating new file, think its gonna rewrite file O_O
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func ReadDb() ([]string, error) {
	files, err := os.ReadDir("./db")
	if err != nil {
		return nil, err
	}

	fList := []string{}

	for _, item := range files {
		fName := strings.Split(item.Name(), ".")
		fList = append(fList, fName[0])
	}

	return fList, nil
}
