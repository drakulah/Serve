package fs

import (
	"os"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadTextFile(path string) (string, error) {
	str := ""
	bytes, err := os.ReadFile(path)

	if err != nil {
		return str, err
	}

	str = string(bytes[:])

	return str, nil
}

type File struct {
	Name  string
	IsDir bool
}

func ReadDir(path string) (dirList []File) {
	dir, err := os.ReadDir(path)

	if err == nil {
		for _, e := range dir {
			dirList = append(dirList, File{
				Name:  e.Name(),
				IsDir: e.IsDir(),
			})
		}
	}

	return
}
