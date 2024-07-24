package domain

type Role int

type roles struct {
	Unknown  Role
	Villager Role
	Wolf     Role
}

var Roles = roles{
	Unknown:  0,
	Villager: 1,
	Wolf:     2,
}

func (role Role) String() string {
	return [2]string{"villager", "wolf"}[role]
}
