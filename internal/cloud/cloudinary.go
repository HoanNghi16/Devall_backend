package cloud

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func ConfigCloud()(*cloudinary.Cloudinary) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"),os.Getenv("CLOUD_API_KEY"), os.Getenv("CLOUD_API_SECRET")) 
	if err != nil{
		panic(err)
	}
	return cld
}