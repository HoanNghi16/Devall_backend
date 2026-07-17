package validator

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
)

func ValidateFile(fileHeader *multipart.FileHeader, maxSize int64, allowedMimes []string)( error){
	if fileHeader == nil{
		return errors.New("File không tồn tại!")
	}
	if fileHeader.Size > maxSize {
		return errors.New("Dung lượng file quá lớn!")
	}

	file, err:= fileHeader.Open()
	if err != nil{
		return errors.New("Mở file thất bại!")
	}

	buf := make([]byte, 512)

	_, err = file.Read(buf)

	mime := http.DetectContentType(buf)

	log.Println("Detected MIME:", mime)

	for _,allowedMime := range allowedMimes{
		if mime == allowedMime{
			return nil
		}
	}
	defer file.Close()
	
	return errors.New("Định dạng file này không được hỗ trợ!")
}