package dbmodel

import "gorm.io/gorm"

type User struct{
  gorm.Model
  Name string `gorm:"Not Null" json:"name"`
  EmailID string `gorm:"Not Null" json:"emailId"`
  Password string `gorm:"Not Null" json:"password"`
  HighestScore uint `json:"highestScore"`
}
