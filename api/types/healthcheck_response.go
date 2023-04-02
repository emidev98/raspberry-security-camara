package types

type healthcheck struct {
	Status bool `json:"status"`
}

func NewHealthcheck() healthcheck {
	return healthcheck{
		Status: true,
	}
}
