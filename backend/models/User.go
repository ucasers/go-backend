package models

import (
	"github.com/badoux/checkmail"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeSave(*gorm.DB) error {
	u.Prepare()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate(action string) map[string]string {
	errorMessages := make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update", "login", "forgotpassword":
		if u.Email == "" {
			errorMessages["Required_email"] = "Required Email"
		} else if err = checkmail.ValidateFormat(u.Email); err != nil {
			errorMessages["Invalid_email"] = "Invalid Email"
		}
		if action != "forgotpassword" && u.Password == "" {
			errorMessages["Required_password"] = "Required Password"
		}
	default:
		if u.Username == "" {
			errorMessages["Required_username"] = "Required Username"
		}
		if u.Password == "" {
			errorMessages["Required_password"] = "Required Password"
		} else if len(u.Password) < 6 {
			errorMessages["Invalid_password"] = "Password should be at least 6 characters"
		}
		if u.Email == "" {
			errorMessages["Required_email"] = "Required Email"
		} else if err = checkmail.ValidateFormat(u.Email); err != nil {
			errorMessages["Invalid_email"] = "Invalid Email"
		}
	}
	return errorMessages
}
