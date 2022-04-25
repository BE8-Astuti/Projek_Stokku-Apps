package user

import (
	"errors"
	"projek/be8/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (ur *UserRepo) InsertUser(newUser entities.User) (entities.User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa insert data")
	}

	log.Info()
	return newUser, nil
}

// func (ur *UserRepo) GetAllUser() ([]entities.User, error) {
// 	arrUser := []entities.User{}

// 	if err := ur.Db.Find(&arrUser).Error; err != nil {
// 		log.Warn(err)
// 		return nil, errors.New("tidak bisa select data")
// 	}

// 	if len(arrUser) == 0 {
// 		log.Warn("tidak ada data")
// 		return nil, errors.New("tidak ada data")
// 	}

// 	log.Info()
// 	return arrUser, nil
// }

func (ur *UserRepo) Login(name string, password string) (entities.User, error) {
	users := []entities.User{}

	if err := ur.Db.Where("name = ? AND password = ?", name, password).First(&users).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa select data")
	}

	return users[0], nil
}
