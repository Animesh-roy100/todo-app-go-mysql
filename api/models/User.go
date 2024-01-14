package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"todolist/api/security"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

// User struct is used to store user information in the database
   type User struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username   string    `gorm:"size:255;not null;unique" json:"username"`
	Email      string    `gorm:"size:100;not null;unique" json:"email"`
	Password   string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
   }
   
   // BeforeSave is a hook that is called before a user is saved to the database
   // It hashes the user's password before saving it
   func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
	 return err
	}
	u.Password = string(hashedPassword)
	return nil
   }
   
   // Prepare is a function that is called before a user is saved to the database
   // It escapes any HTML characters and trims any whitespace
   func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
   }
   
   // Validate is a function that is used to validate a user before saving it to the database
   // It takes an action as an argument, which is used to determine which validation to perform
   func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error
	switch strings.ToLower(action) {
	case "update":
	 if u.Email == "" {
	  err = errors.New("Required Email")
	  errorMessages["Required_email"] = err.Error()
	 }
	 if u.Email != "" {
	  if err = checkmail.ValidateFormat(u.Email); err != nil {
	   err = errors.New("Invalid Email")
	   errorMessages["Invalid_email"] = err.Error()
	  }
	 }
	case "login":
	 if u.Password == "" {
	  err = errors.New("Required Password")
	  errorMessages["Required_password"] = err.Error()
	 }
	 if u.Email == "" {
	  err = errors.New("Required Email")
	  errorMessages["Required_email"] = err.Error()
	 }
	 if u.Email != "" {
	  if err = checkmail.ValidateFormat(u.Email); err != nil {
	   err = errors.New("Invalid Email")
	   errorMessages["Invalid_email"] = err.Error()
	  }
	 }
	default:
	 if u.Username == "" {
	  err = errors.New("Required Username")
	  errorMessages["Required_username"] = err.Error()
	 }
	 if u.Password == "" {
	  err = errors.New("Required Password")
	  errorMessages["Required_password"] = err.Error()
	 }
	 if u.Password != "" && len(u.Password) < 6 {
	  err = errors.New("Password should be atleast 6 characters")
	  errorMessages["Invalid_password"] = err.Error()
	 }
	 if u.Email == "" {
	  err = errors.New("Required Email")
	   errorMessages["Required_email"] = err.Error()
	 }
	 if u.Email != "" {
	  if err = checkmail.ValidateFormat(u.Email); err != nil {
	   err = errors.New("Invalid Email")
	   errorMessages["Invalid_email"] = err.Error()
	  }
	 }
	}
	return errorMessages
   }
   
   // SaveUser is a function that is used to save a user to the database
   // It takes a pointer to a gorm.DB as an argument
   func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
	 return &User{}, err
	}
	return u, nil
   }