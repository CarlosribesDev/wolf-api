package domain

type Player struct {
	Id   int  `json:"id"`
	Role Role `json:"role"`
}

type NewPlayer struct {
}
