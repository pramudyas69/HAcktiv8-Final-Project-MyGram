package user_pg

import (
	"MyGramHacktiv8/entity"
	"MyGramHacktiv8/pkg/errs"
	"MyGramHacktiv8/repository/userRepository"
	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) userRepository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) GetUserByIDAndEmail(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	err := u.db.Where("email = ? AND id = ?", userPayload.Email, userPayload.ID).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) Login(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	err := u.db.Where("email = ?", userPayload.Email).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) Register(userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	if checkEmail := u.db.Model(user).Where("email = ?", userPayload.Email).Find(&user); checkEmail.RowsAffected > 0 {
		return nil, errs.NewBadRequest("Email is exist")
	}

	if checkUsername := u.db.Model(user).Where("username = ?", userPayload.Username).Find(&user); checkUsername.RowsAffected > 0 {
		return nil, errs.NewBadRequest("Username is exist")
	}

	if err := u.db.Create(userPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	if err := u.db.Where("email = ?", userPayload.Email).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) UpdateUserData(userId uint, userPayload *entity.User) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	err := u.db.Model(user).Where("id = ?", userId).Updates(userPayload).Take(&user).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) DeleteUser(userId uint) errs.MessageErr {
	user := entity.User{}

	err := u.db.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
