package models

type UserInfo struct {
	PrevGameAnswer *GameAnswer
	UsedCities     map[string]*string
}

type GameAnswer struct {
	LastChar rune
	NewCity  string
}
