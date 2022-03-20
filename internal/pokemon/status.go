package pokemon

type Status struct {
	HP             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"special_attack"`
	SpecialDefense int `json:"special_defense"`
	Speed          int `json:"speed"`
	Total          int `json:"total"`
}

func (s *Status) CalculateTotal() {
	s.Total = s.HP + s.Attack + s.Defense + s.SpecialAttack + s.SpecialDefense + s.Speed
}
