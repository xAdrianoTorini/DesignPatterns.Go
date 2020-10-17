package main

type maverick struct {
	gun //Composition
}

func newMaverick() iGun { //Returns a Interface
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
