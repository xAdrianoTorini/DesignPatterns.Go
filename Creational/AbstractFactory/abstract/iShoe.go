package abstract

//IShoe interface
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

//Shoe Abstracrt Type
type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

//GetLogo - Beacause we are in diferents packags we will need make these methodes publics
func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

//GetSize - Beacause we are in diferents packags we will need make these methodes publics
func (s *Shoe) GetSize() int {
	return s.size
}
