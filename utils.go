package rirs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func downloadFile(dirPath string, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("получен неожиданный статус: %s", resp.Status)
	}

	fileName := filepath.Base(url)
	if fileName == "." || fileName == "/" {
		fileName = fmt.Sprintf("index-%s", uuid.New().String())
	}

	filePath := filepath.Join(dirPath, fileName)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", fmt.Errorf("не удалось создать каталог: %v", err)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при записи файла: %v", err)
	}

	return filePath, nil
}
