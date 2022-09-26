package database

type Sort struct {
	Nom      string
	CoutMana int
	Degats   int
	BoostPv  int
}

func (s *Sort) Init(nom string, coutMana int, degats int, boostPv int) {
	s.Nom = nom
	s.CoutMana = coutMana
	s.Degats = degats
	s.BoostPv = boostPv
}

func (s *Sort) InitSort(number int) {
	//Sort 1
	if number == 1 {
		s.Init("Boule de feu", 10, 10, 0)
	}
	//Sort 2
	if number == 2 {
		s.Init("Eclair", 15, 15, 0)
	}
	//Sort 3
	if number == 3 {
		s.Init("Soin", 20, 0, 20)
	}
}
