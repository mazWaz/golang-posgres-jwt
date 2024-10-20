package auth

import "time"

type ResponseToken struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expire_time"`
}

type ResponseAuthToken struct {
	Access  ResponseToken `json:"access"`
	Refresh ResponseToken `json:"refresh"`
}
