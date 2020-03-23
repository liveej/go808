package protocol

// 消息类型
type Type uint16

const (
	// 终端应答
	TypeT808_0x0001 Type = 0x0001
	// 终端心跳
	TypeT808_0x0002 Type = 0x0002
	// 终端注销
	TypeT808_0x0003 Type = 0x0003
	// 终端注册
	TypeT808_0x0100 Type = 0x0100
	// 终端鉴权
	TypeT808_0x0102 Type = 0x0102
	// 查询终端参数应答
	TypeT808_0x0104 Type = 0x0104
	// 查询终端属性应答
	TypeT808_0x0107 Type = 0x0107
	// 汇报位置
	TypeT808_0x0200 Type = 0x0200
	// 位置信息查询应答
	TypeT808_0x0201 Type = 0x0201
	// 数据上行透传
	TypeT808_0x0900 Type = 0x0900
	// 平台通用应答
	TypeT808_0x8001 Type = 0x8001
	// 终端注册应答
	TypeT808_0x8100 Type = 0x8100
	// 设置终端参数
	TypeT808_0x8103 Type = 0x8103
	// 查询终端参数
	TypeT808_0x8104 Type = 0x8104
	// 终端控制
	TypeT808_0x8105 Type = 0x8105
	// 查询指定参数
	TypeT808_0x8106 Type = 0x8106
	// 查询终端属性
	TypeT808_0x8107 Type = 0x8107
	// 查询车辆位置
	TypeT808_0x8201 Type = 0x8201
	// 文本信息下发
	TypeT808_0x8300 Type = 0x8300
	// 设置圆形区域
	TypeT808_0x8600 Type = 0x8600
	// 删除圆形区域
	TypeT808_0x8601 Type = 0x8601
	// 设置矩形区域
	TypeT808_0x8602 Type = 0x8602
	// 删除矩形区域
	TypeT808_0x8603 Type = 0x8603
	// 设置多边形区域
	TypeT808_0x8604 Type = 0x8604
	// 删除多边形区域
	TypeT808_0x8605 Type = 0x8605
	// 数据下行透传
	TypeT808_0x8900 Type = 0x8900
)

// 消息实体映射
var entityMapper = map[uint16]func() Entity{
	uint16(TypeT808_0x0001): func() Entity {
		return new(T808_0x0001)
	},
	uint16(TypeT808_0x0002): func() Entity {
		return new(T808_0x0002)
	},
	uint16(TypeT808_0x0003): func() Entity {
		return new(T808_0x0003)
	},
	uint16(TypeT808_0x0100): func() Entity {
		return new(T808_0x0100)
	},
	uint16(TypeT808_0x0102): func() Entity {
		return new(T808_0x0102)
	},
	uint16(TypeT808_0x0104): func() Entity {
		return new(T808_0x0104)
	},
	uint16(TypeT808_0x0107): func() Entity {
		return new(T808_0x0107)
	},
	uint16(TypeT808_0x0200): func() Entity {
		return new(T808_0x0200)
	},
	uint16(TypeT808_0x0201): func() Entity {
		return new(T808_0x0201)
	},
	uint16(TypeT808_0x8001): func() Entity {
		return new(T808_0x8001)
	},
	uint16(TypeT808_0x8100): func() Entity {
		return new(T808_0x8100)
	},
	uint16(TypeT808_0x8103): func() Entity {
		return new(T808_0x8103)
	},
	uint16(TypeT808_0x8104): func() Entity {
		return new(T808_0x8104)
	},
	uint16(TypeT808_0x8105): func() Entity {
		return new(T808_0x8105)
	},
	uint16(TypeT808_0x8106): func() Entity {
		return new(T808_0x8106)
	},
	uint16(TypeT808_0x8107): func() Entity {
		return new(T808_0x8107)
	},
	uint16(TypeT808_0x8300): func() Entity {
		return new(T808_0x8300)
	},
	uint16(TypeT808_0x8600): func() Entity {
		return new(T808_0x8600)
	},
	uint16(TypeT808_0x8601): func() Entity {
		return new(T808_0x8601)
	},
	uint16(TypeT808_0x8602): func() Entity {
		return new(T808_0x8602)
	},
	uint16(TypeT808_0x8603): func() Entity {
		return new(T808_0x8603)
	},
	uint16(TypeT808_0x8604): func() Entity {
		return new(T808_0x8604)
	},
	uint16(TypeT808_0x8605): func() Entity {
		return new(T808_0x8605)
	},
	uint16(TypeT808_0x8900): func() Entity {
		return new(T808_0x8900)
	},
	uint16(TypeT808_0x0900): func() Entity {
		return new(T808_0x0900)
	},
}

// 消息注册
func Register(typ uint16, creator func() Entity) {
	entityMapper[typ] = creator
}

// 消息解码
func Decode(typ uint16, data []byte) (Entity, int, error) {
	creator, ok := entityMapper[typ]
	if !ok {
		return nil, 0, ErrMessageNotRegistered
	}

	entity := creator()
	count, err := entity.Decode(data)
	if err != nil {
		return nil, 0, err
	}
	return entity, count, nil
}
