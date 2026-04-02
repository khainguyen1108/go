package models

import (
	"time"
)

type LoginRequest struct {
	UserId    string `json:"userId" validate:"required" code:"1001"`
	Password  string `json:"password" validate:"required,min=6" code:"1002"`
	UserAgent string
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required" code:"1003"`
	UserAgent    string
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	Id                     int        `db:"id"`
	UserId                 string     `db:"user_id"`
	Email                  string     `db:"email"`
	Name                   string     `db:"name"`
	NickName               string     `db:"nick_name"`
	Phone                  string     `db:"phone"`
	Avatar                 *string    `db:"avatar"`
	Password               string     `db:"password" json:"-"`
	UpdateId               *string    `db:"update_id"`
	CreateId               *string    `db:"create_id"`
	UpdateDateTime         *time.Time `db:"update_date_time"`
	CreateDateTime         *time.Time `db:"create_date_time"`
	AccessType             int        `db:"access_type"`
	UserType               int        `db:"user_type"`
	UserKind               int        `db:"user_kind"`
	SnsType                int        `db:"sns_type"`
	SnsId                  *string    `db:"sns_id"`
	Sex                    int        `db:"sex"`
	Birthday               *string    `db:"birthday"`
	SkillLevel             int        `db:"skill_level"`
	AveragePar             float64    `db:"average_par"`
	Handicap               float64    `db:"handicap"`
	Ci                     *string    `db:"ci"`
	UserStatus             int        `db:"user_status"`
	PublicProfile          int        `db:"public_profile"`
	Introduction           *string    `db:"introduction"`
	KcbDateTime            *time.Time `db:"kcb_date_time"`
	JoinDateTime           *time.Time `db:"join_date_time"`
	LeftDateTime           *time.Time `db:"left_date_time"`
	LeftReason             *string    `db:"left_reason"`
	LeftType               string     `db:"left_type"`
	LastLogin              string     `db:"last_login"`
	LoginCnt               string     `db:"login_cnt"`
	PushToken              *string    `db:"push_token"`
	HandicapIssuedDateTime *time.Time `db:"handicap_issued_date_time"`
	HandicapExpireDateTime *time.Time `db:"handicap_expired_date_time"`
	Score                  int        `db:"score"`
	ScoreCnt               int        `db:"score_cnt"`
	UseStatus              int        `db:"use_status"`
}

type AccountSession struct {
	Id           int       `db:"id"`
	UserId       int       `db:"user_id"`
	SessionId    string    `db:"session_id"`
	RefreshToken string    `db:"refresh_token"`
	IsUsed       int       `db:"is_used"`
	IsRevoked    int       `db:"is_revoked"`
	DeviceInfo   string    `db:"device_info"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
