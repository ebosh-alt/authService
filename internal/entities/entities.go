package entities

import (
	"database/sql"
)

func NewNullString(s string) *sql.NullString {
	if len(s) == 0 {
		return &sql.NullString{}
	} else {
		return &sql.NullString{
			String: s,
			Valid:  true,
		}
	}
}

type User struct {
	ID           int32
	TelegramID   int64
	FirstName    string
	AccessToken  string
	RefreshToken string
	CreateAt     string
	UpdatedAt    string
}

type UserDTO struct {
	ID           int32          `json:"id" db:"id"`
	TelegramID   int64          `json:"telegram_id" db:"telegram_id"`
	FirstName    sql.NullString `json:"first_name" db:"first_name"`
	AccessToken  sql.NullString `json:"access_token" db:"access_token"`
	RefreshToken sql.NullString `json:"refresh_token" db:"refresh_token"`
	CreateAt     sql.NullString `json:"create_at" db:"create_at"`
	UpdatedAt    sql.NullString `json:"updated_at" db:"updated_at"`
}

func (u *User) ConvertToDTO() *UserDTO {
	dto := new(UserDTO)
	dto.ID = u.ID
	dto.TelegramID = u.TelegramID
	dto.FirstName = *NewNullString(u.FirstName)
	dto.AccessToken = *NewNullString(u.AccessToken)
	dto.RefreshToken = *NewNullString(u.RefreshToken)
	dto.CreateAt = *NewNullString(u.CreateAt)
	dto.UpdatedAt = *NewNullString(u.UpdatedAt)
	return dto
}

func (dto *UserDTO) FromDTOConvert() *User {
	c := new(User)
	c.ID = dto.ID
	c.TelegramID = dto.TelegramID
	c.FirstName = dto.FirstName.String
	c.AccessToken = dto.AccessToken.String
	c.RefreshToken = dto.RefreshToken.String
	c.CreateAt = dto.CreateAt.String
	c.UpdatedAt = dto.UpdatedAt.String
	return c
}
