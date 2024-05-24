package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"github.com/ucasers/go-backend/backend/security"
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

func (u *User) BeforeSave() error {
	u.Prepare()
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.UpdatedAt = time.Now()
}

func (u *User) AfterFind() error {
	// Don't return the user password
	u.Password = ""
	return nil
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

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	if err := u.BeforeSave(); err != nil {
		return &User{}, err
	}
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.hashPasswordIfExists()
	if err != nil {
		return &User{}, err
	}

	updateData := map[string]interface{}{
		"email":      u.Email,
		"updated_at": time.Now(),
	}

	if u.Password != "" {
		updateData["password"] = u.Password
	}

	err = db.Debug().Model(&User{}).Where("id = ?", uid).Updates(updateData).Error
	if err != nil {
		return &User{}, err
	}

	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *User) UpdatePassword(db *gorm.DB) error {
	err := u.BeforeSave()
	if err != nil {
		return err
	}
	err = db.Debug().Model(&User{}).Where("email = ?", u.Email).Updates(map[string]interface{}{
		"password":   u.Password,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// Helper function to hash password if it exists
func (u *User) hashPasswordIfExists() error {
	if u.Password != "" {
		return u.BeforeSave()
	}
	return nil
}
