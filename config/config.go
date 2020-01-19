package config

type StatusType uint
const (
	Status StatusType = 1
	DeleteStatus StatusType = 2
)

func (f StatusType) String() string  {
	switch f {
	case Status:
		return "未删除"
	case DeleteStatus:
		return "已删除"
	default:
		return "未知状态"
	}
}
