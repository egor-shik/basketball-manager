package coach

func NewCoach(id int, firstName, lastName string, exp int) *Coach {
	return &Coach{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		Experience: exp,
		Buffs:      CoachBuffs{},
	}
}
