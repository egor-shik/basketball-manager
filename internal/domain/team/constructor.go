package team

func NewTeam(name string, coach *coach.Coach, players []*player.Player) *Team {
    return &Team{
        Name: name,
        Players: players,
        Coach: coach,
        Budget: Budget{},
        Morale: TeamMorale{},
    }