package utils

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveFileFromForm(ctx *gin.Context, field string, path string) (filename string, ok bool) {
	file, err := ctx.FormFile(field)
	if err != nil {
		fmt.Println("Gagal mendapatkan file dari form: ", err)
		return "", false
	}

	extension := filepath.Base(file.Filename)

	newFileName := uuid.New().String() + extension

	if err := ctx.SaveUploadedFile(file, path+newFileName); err != nil {
		fmt.Println("Gagal menyimpan file ke folder: ", err)
		return "", false
	}

	return newFileName, true
}
