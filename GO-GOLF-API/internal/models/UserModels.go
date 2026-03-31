package models

import "time"

type LoginRequest struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
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
