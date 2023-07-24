package domain

type TagKey int32

const (
	ClientId TagKey = iota
	Carrier
	User
)

func (k TagKey) String() string {
	switch k {
	case ClientId:
		return "ClientId"
	case Carrier:
		return "Carrier"
	}
	return "unknown"
}
