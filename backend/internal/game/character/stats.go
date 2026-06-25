package character

type Stats struct {
	MaxHP   int `json:"maxHp"`
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	Speed   int `json:"speed"`
}
