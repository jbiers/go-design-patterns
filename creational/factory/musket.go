package factory

type musket struct {
	Gun
}

func newMusket() iGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
