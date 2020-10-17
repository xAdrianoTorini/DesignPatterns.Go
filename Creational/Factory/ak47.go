package main

type ak47 struct {
	gun //Composition
}

func newAk47() iGun { //Returns a Interface
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
