package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/mail"
	"personal-trainer-api/internal/entities/enums"
	"personal-trainer-api/internal/httpResponse"
)

type User struct {
	gorm.Model
	Name     string     `json:"name" gorm:"not_null"`
	Email    string     `json:"email" gorm:"unique"`
	Password string     `json:"password"`
	Role     enums.Role `json:"role"`
}

type Input struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewUser(input Input) (*User, []httpResponse.Cause) {
	err := input.validate()
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     enums.Role(input.Role),
	}, nil
}

func (i *Input) validate() []httpResponse.Cause {
	var causes []httpResponse.Cause

	if i.Name == "" {
		cause := httpResponse.Cause{
			Field:   "name",
			Message: "name is required",
		}
		causes = append(causes, cause)
	}

	if i.Email == "" {
		cause := httpResponse.Cause{
			Field:   "email",
			Message: "email is required",
		}
		causes = append(causes, cause)
	}

	if !IsvalidEmail(i.Email) {
		cause := httpResponse.Cause{
			Field:   "email",
			Message: "value informed is invalid",
		}
		causes = append(causes, cause)
	}

	if i.Password == "" {
		cause := httpResponse.Cause{
			Field:   "password",
			Message: "password is required",
		}
		causes = append(causes, cause)
	}

	if i.Role == "" {
		cause := httpResponse.Cause{
			Field:   "role",
			Message: "role is required",
		}
		causes = append(causes, cause)
	}

	if !enums.IsValidRole(i.Role) {
		cause := httpResponse.Cause{
			Field:   "role",
			Message: "value informed is invalid",
		}
		causes = append(causes, cause)
	}

	if len(causes) > 0 {
		return causes
	}
	return nil
}

func IsvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	return nil
}
