package data

type Team struct {
	ID   string `json:"idTeam"`
	Name string `json:"strTeam"`
}

type resp struct {
	Teams []Team `json:"teams"`
}
