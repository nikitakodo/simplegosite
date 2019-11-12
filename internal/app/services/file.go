package services

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
)

type FileService struct {
	UploadDir     string
	MaxUploadSize int64
	Repository    store.RepositoryInterface
	Logger        *logrus.Logger
}

func NewFileService(uploadDir string, maxUploadSize int64, repo store.RepositoryInterface, logger *logrus.Logger) *FileService {
	if maxUploadSize == 0 {
		maxUploadSize = 10 * 1024
	}
	return &FileService{
		UploadDir:     uploadDir,
		MaxUploadSize: maxUploadSize,
		Repository:    repo,
		Logger:        logger,
	}
}

func (service *FileService) Upload(keyName string, r *http.Request) (*model.File, error) {
	file, handler, err := r.FormFile(keyName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileName := patternName(handler.Filename)
	tempFile, err := ioutil.TempFile(service.UploadDir, fileName)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()
	newFile := &model.File{
		Name:      fileName,
		Path:      service.UploadDir + string(os.PathSeparator) + fileName,
		Size:      handler.Size,
		Extension: filepath.Ext(fileName),
	}

	//todo save to db

	return newFile, nil
}

func (service *FileService) Delete(file *model.File) error {
	//todo remove file from db

	return os.Remove(file.Path + string(os.PathSeparator) + file.Name)
}

func patternName(fileName string) string {
	return "upload-*" + fileName
}
