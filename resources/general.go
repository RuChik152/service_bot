package resources

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFile(name string) (string, error) {
	if cwd, err := os.Getwd(); err != nil {
		return "", fmt.Errorf("ошибка при получении текущего рабочего каталога: %s", err)
	} else {
		return filepath.Join(cwd, "assets", name), nil
	}
}
