package abstract

//IShort Interface
type IShort interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

//Short Abstract type
type Short struct {
	logo string
	size int
}

func (s *Short) setLogo(logo string) {
	s.logo = logo
}

//GetLogo - Beacause we are in diferents packags we will need make these methodes publics
func (s *Short) GetLogo() string {
	return s.logo
}

func (s *Short) setSize(size int) {
	s.size = size
}

//GetSize - Beacause we are in diferents packags we will need make these methodes publics
func (s *Short) GetSize() int {
	return s.size
}
