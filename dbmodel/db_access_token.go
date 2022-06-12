package dbmodel

import "gorm.io/gorm"

type AccessToken struct{
  gorm.Model 
  UserID uint  `json:"userId"`
  Token string `json:"token"`
}
