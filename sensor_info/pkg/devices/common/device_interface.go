package devices

type Device interface {
	GetID() string
	GetName() string
	GetRange() (float64, float64)
	GetStatus() string
	SetStatus(status string)
	FormatData() (string, error)
}
