package web

import "time"

type ResponseCreateCourse struct {
	Name           string    `json:"name" form:"name"`
	Price          int       `json:"price" form:"price"`
	Category       string    `json:"category" form:"category"`
	Description    string    `json:"desc" form:"desc"`
	Waktu          time.Time `json:"time" form:"time"`
	MaxParticipant int       `json:"max_participant" form:"max_participant"`
}

type ResponseFindCourse struct {
	Name           string            `json:"name" form:"name"`
	Price          int               `json:"price" form:"price"`
	Category       string            `json:"category" form:"category"`
	Description    string            `json:"desc" form:"desc"`
	Waktu          time.Time         `json:"time" form:"time"`
	Participant    []ListParticipant `json:"participant" form:"participant"`
	MaxParticipant int               `json:"max_participant" form:"max_participant"`
}

type ResponseCreateUser struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Avatar  string `json:"avatar" form:"avatar"`
	Address string `json:"address" form:"address"`
	Notelp  int    `json:"no_telp" form:"no_telp"`
}

type ResponseCreateCategory struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"desc" form:"desc"`
}

type ResponseFindCategory struct {
	Name        string       `json:"name" form:"name"`
	Description string       `json:"desc" form:"desc"`
	Course      []ListCourse `json:"course" form:"course"`
}

type ResponseCreateOrder struct {
	User     string `json:"user_name" form:"user_name"`
	Course   string `json:"course_name" form:"course_name"`
	Price    int    `json:"course_price" form:"course_price"`
	Category string `json:"course_category" form:"course_category"`
	Status   bool   `json:"status_paid" form:"status_paid"`
}

type ListCourse struct {
	Name           string    `json:"name" form:"name"`
	Price          int       `json:"price" form:"price"`
	Description    string    `json:"desc" form:"desc"`
	Waktu          time.Time `json:"time" form:"time"`
	MaxParticipant int       `json:"max_participant" form:"max_participant"`
}

type ListParticipant struct {
	UserName  string `json:"user_name" form:"user_name"`
	UserEmail string `json:"user_email" form:"user_email"`
}
