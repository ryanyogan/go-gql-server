// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type User struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	UserID      string     `json:"userID"`
	Name        *string    `json:"name"`
	FirstName   *string    `json:"firstName"`
	LastName    *string    `json:"lastName"`
	NickName    *string    `json:"nickName"`
	Description *string    `json:"description"`
	Location    *string    `json:"location"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type UserInput struct {
	Email       *string `json:"email"`
	UserID      *string `json:"userID"`
	DisplayName *string `json:"displayName"`
	Name        *string `json:"name"`
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	NickName    *string `json:"nickName"`
	Description *string `json:"description"`
	Location    *string `json:"location"`
}

type Users struct {
	Count *int    `json:"count"`
	List  []*User `json:"list"`
}
