package files

import "os"

func ReadFile(fileName string) ([]byte, error) {
	fileData, err := os.ReadFile(fileName + ".json")
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

func WriteFile(fileName string, content []byte) error {
	file, err := os.Create(fileName + ".json") //mb it's ok creating new file, think its gonna rewrite file O_O
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
