package types

type healthcheck struct {
	Status bool
}

func NewHealthcheck() healthcheck {
	return healthcheck{
		Status: true,
	}
}
