package entity

import (
	"MyGramHacktiv8/pkg/errs"
	"MyGramHacktiv8/pkg/helpers"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

var (
	SECRET_KEY = helpers.GodotEnv("SECRET_KEY")
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique;type:varchar(191)" form:"username" json:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;unique;type:varchar(191)" form:"email" json:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `gorm:"not null;type:varchar(191)" form:"password" json:"password" valid:"required~Password is required, minstringlength(6)~Password must be at least 6 characters"`
	Age      uint8  `gorm:"not null" form:"age" json:"age" valid:"required~Age is required, range(8|100)~Age must be between 8 and 100"`
}

// ! TODO: Validator Age not trigerred, also hash pashing not working in hooks
func (u *User) HashPass() errs.MessageErr {
	salt := 8
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		return errs.NewInternalServerErrorr("error hashing password")
	}

	u.Password = string(hash)

	return nil
}

// * Verify encrypted password with bcrypt.
func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

// * Generate token with jwt.
func (u *User) GenerateToken() string {
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(SECRET_KEY))

	return signedToken
}

// * Verify token with jwt for auth purposes.
func (u *User) VerifyToken(tokenStr string) error {
	fmt.Println("tokenStr => ", tokenStr)
	if bearer := strings.HasPrefix(tokenStr, "Bearer "); !bearer {
		return errors.New("login to proceed")
	}

	stringToken := strings.Split(tokenStr, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("login to proceed")
		}

		return []byte(SECRET_KEY), nil
	})

	var mapClaims jwt.MapClaims

	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errors.New("login to proceed")
	} else {
		mapClaims = v
	}

	if exp, ok := mapClaims["exp"].(float64); !ok || !token.Valid {
		return errors.New("login to proceed")
	} else {
		if int64(exp)-time.Now().Unix() < 0 {
			return errors.New("login to proceed")
		}
	}

	if v, ok := mapClaims["id"].(float64); !ok || !token.Valid {
		return errors.New("login to proceed")
	} else {
		u.ID = uint(v)
	}

	if v, ok := mapClaims["email"].(string); !ok || !token.Valid {
		return errors.New("login to proceed")
	} else {
		u.Email = v
	}

	return nil
}
