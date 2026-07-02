package player 


func NewPlayer(id int, firstName, lastName string, age int, position Position) *Player {
    return &Player{
        Personal: PersonalInfo{
            ID:        id,
            FirstName: firstName,
            LastName:  lastName,
            Age:       age,
            Pos:  position,
        },
        Ratings:   Ratings{},   // would be end later 
        Potential: Potential{},
        Stats:     Statistics{},
        Mood: Mood{},
        Contract:  Contract{},
    }
}