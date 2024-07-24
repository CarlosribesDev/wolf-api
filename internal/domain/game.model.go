package domain

import "time"

type NewGame struct {
	HostId    int       `json:"hostId"`
	CreatedAt time.Time `json:"createdAt"`
}

type Game struct {
	Id        int       `json:"id"`
	HostId    int       `json:"hostId"`
	CreatedAt time.Time `json:"createdAt"`
}
