package types

type healthcheck struct {
	status bool
}

func NewHealthcheck() healthcheck {
	return healthcheck{
		status: true,
	}
}
