package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"green/internal/models"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type response struct {
	Message interface{} `json:"message"`
}

func InitConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func ErrorResponse(w http.ResponseWriter, err error, statusCode, errorCode int) {
	message := models.ErrorResponse{Error: err.Error(), ErrorCode: errorCode}
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(data)
	return
}

func Response(w http.ResponseWriter, data interface{}) {
	result := &response{Message: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SaveImageFromBase64(base64Image, filePath string) error {
	base64Image = strings.TrimPrefix(base64Image, "data:image/jpeg;base64,")
	base64Image = strings.TrimPrefix(base64Image, "data:image/png;base64,")
	base64Image = strings.TrimPrefix(base64Image, "data:image/gif;base64,")
	base64Image = strings.TrimPrefix(base64Image, "data:image/bmp;base64,")
	base64Image = strings.TrimPrefix(base64Image, "data:image/webp;base64,")
	imgData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return fmt.Errorf("failed to decode base64 string: %v", err)
	}
	dir := filepath.Dir(filePath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()
	if _, err = io.Copy(file, strings.NewReader(string(imgData))); err != nil {
		return fmt.Errorf("failed to write image to file: %v", err)
	}

	return nil
}

func ConvertImageToBase64(dir string, fileName string) (string, error) {
	filePath := filepath.Join(dir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %v", fileName, err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)
	imageData, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %v", fileName, err)
	}
	base64String := base64.StdEncoding.EncodeToString(imageData)
	return base64String, nil
}

func RemoveFile(dir, filename string) error {
	filePath := filepath.Join(dir, filename)
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
