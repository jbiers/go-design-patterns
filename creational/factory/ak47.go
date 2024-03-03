package factory

type ak47 struct {
	Gun
}

func newAK47() iGun {
	return &ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
