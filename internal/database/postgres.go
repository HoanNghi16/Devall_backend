package database

import (
	"fmt"
	"os"

	"github.com/HoanNghi16/Devall_backend/internal/algorithm"
	"github.com/HoanNghi16/Devall_backend/internal/course"
	"github.com/HoanNghi16/Devall_backend/internal/media"
	"github.com/HoanNghi16/Devall_backend/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")

	dsn := "host=" + db_host +
			" user=" + db_user +
			" password=" + db_password + 
			" dbname=" + db_name + 
			" port=" + db_port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//Nếu error khi kết nối -> trả về nil, error
	if err != nil{
		return nil, err
	}

	//Nếu error khi auto migrate -> trả về nil, error
	if err := db.AutoMigrate(
		&user.User{}, 
		&user.Profile{},
		&course.Course{},
		&course.Lesson{},
		&course.ContentBlock{},
		&course.Topic{},
		&course.TopicCourse{},
		&algorithm.Algorithm{},
		&algorithm.SolvingHistory{},
		&algorithm.Tag{},
		&media.Media{},
	) ; err != nil{
		return nil, fmt.Errorf("auto migrate failed: %w", err)
	}
	return db, err
}