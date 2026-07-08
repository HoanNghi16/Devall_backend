package validator

import (
	"errors"
	"mime/multipart"
	"net/http"
)

func ValidateAndOpenFile(fileHeader *multipart.FileHeader, maxSize int64, allowedMimes []string)(multipart.File, error){
	if fileHeader == nil{
		return nil,errors.New("File không tồn tại!")
	}
	if fileHeader.Size > maxSize {
		return nil,errors.New("Dung lượng file quá lớn!")
	}

	file, err:= fileHeader.Open()
	if err != nil{
		return nil,errors.New("Mở file thất bại!")
	}
	defer file.Close()

	buf := make([]byte, 512)

	_, err = file.Read(buf)

	mime := http.DetectContentType(buf)

	for _,allowedMime := range allowedMimes{
		if mime == allowedMime{
			return file,nil
		}
	}
	
	return nil,errors.New("Định dạng file này không được hỗ trợ!")
}