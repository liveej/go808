package protocol

// 消息ID枚举
type MsgID uint16

const (
	// 终端应答
	MsgT808_0x0001 MsgID = 0x0001
	// 终端心跳
	MsgT808_0x0002 MsgID = 0x0002
	// 终端注销
	MsgT808_0x0003 MsgID = 0x0003
	// 终端注册
	MsgT808_0x0100 MsgID = 0x0100
	// 终端鉴权
	MsgT808_0x0102 MsgID = 0x0102
	// 查询终端参数应答
	MsgT808_0x0104 MsgID = 0x0104
	// 查询终端属性应答
	MsgT808_0x0107 MsgID = 0x0107
	// 终端升级结果通知
	MsgT808_0x0108 MsgID = 0x0108
	// 汇报位置
	MsgT808_0x0200 MsgID = 0x0200
	// 位置信息查询应答
	MsgT808_0x0201 MsgID = 0x0201
	// 事件报告
	MsgT808_0x0301 MsgID = 0x0301
	// 提问答案
	MsgT808_0x0302 MsgID = 0x0302
	// 信息点播/取消
	MsgT808_0x0303 MsgID = 0x0303
	// 车辆控制
	MsgT808_0x0500 MsgID = 0x0500
	// 多媒体数据上传
	MsgT808_0x0801 MsgID = 0x0801
	// 平台通用应答
	MsgT808_0x8001 MsgID = 0x8001
	// 补传分包请求
	MsgT808_0x8003 MsgID = 0x8003
	// 终端注册应答
	MsgT808_0x8100 MsgID = 0x8100
	// 设置终端参数
	MsgT808_0x8103 MsgID = 0x8103
	// 查询终端参数
	MsgT808_0x8104 MsgID = 0x8104
	// 终端控制
	MsgT808_0x8105 MsgID = 0x8105
	// 查询指定参数
	MsgT808_0x8106 MsgID = 0x8106
	// 查询终端属性
	MsgT808_0x8107 MsgID = 0x8107
	// 下发终端升级包
	MsgT808_0x8108 MsgID = 0x8108
	// 查询车辆位置
	MsgT808_0x8201 MsgID = 0x8201
	// 临时位置跟踪控制
	MsgT808_0x8202 MsgID = 0x8202
	// 人工确认报警消息
	MsgT808_0x8203 MsgID = 0x8203
	// 文本信息下发
	MsgT808_0x8300 MsgID = 0x8300
	// 事件设置
	MsgT808_0x8301 MsgID = 0x8301
	// 提问下发
	MsgT808_0x8302 MsgID = 0x8302
	// 位置点播菜单设置
	MsgT808_0x8303 MsgID = 0x8303
	// 信息服务
	MsgT808_0x8304 MsgID = 0x8304
	// 电话回拨
	MsgT808_0x8400 MsgID = 0x8400
	// 设置电话本
	MsgT808_0x8401 MsgID = 0x8401
	// 车门控制
	MsgT808_0x8500 MsgID = 0x8500
	// 设置圆形区域
	MsgT808_0x8600 MsgID = 0x8600
	// 删除圆形区域
	MsgT808_0x8601 MsgID = 0x8601
	// 设置矩形区域
	MsgT808_0x8602 MsgID = 0x8602
	// 删除矩形区域
	MsgT808_0x8603 MsgID = 0x8603
	// 设置多边形区域
	MsgT808_0x8604 MsgID = 0x8604
	// 删除多边形区域
	MsgT808_0x8605 MsgID = 0x8605
	// 设置路线
	MsgT808_0x8606 MsgID = 0x8606
	// 删除路线
	MsgT808_0x8607 MsgID = 0x8607
	// 多媒体数据上传应答
	MsgT808_0x8800 MsgID = 0x8800
	// 数据下行透传
	MsgT808_0x8900 MsgID = 0x8900
	// 数据上行透传
	MsgT808_0x0900 MsgID = 0x0900
)

// 消息实体映射
var entityMapper = map[uint16]func() Entity{
	uint16(MsgT808_0x0001): func() Entity {
		return new(T808_0x0001)
	},
	uint16(MsgT808_0x0002): func() Entity {
		return new(T808_0x0002)
	},
	uint16(MsgT808_0x0003): func() Entity {
		return new(T808_0x0003)
	},
	uint16(MsgT808_0x0100): func() Entity {
		return new(T808_0x0100)
	},
	uint16(MsgT808_0x0102): func() Entity {
		return new(T808_0x0102)
	},
	uint16(MsgT808_0x0104): func() Entity {
		return new(T808_0x0104)
	},
	uint16(MsgT808_0x0107): func() Entity {
		return new(T808_0x0107)
	},
	uint16(MsgT808_0x0108): func() Entity {
		return new(T808_0x0108)
	},
	uint16(MsgT808_0x0200): func() Entity {
		return new(T808_0x0200)
	},
	uint16(MsgT808_0x0201): func() Entity {
		return new(T808_0x0201)
	},
	uint16(MsgT808_0x0301): func() Entity {
		return new(T808_0x0301)
	},
	uint16(MsgT808_0x0302): func() Entity {
		return new(T808_0x0302)
	},
	uint16(MsgT808_0x0303): func() Entity {
		return new(T808_0x0303)
	},
	uint16(MsgT808_0x0500): func() Entity {
		return new(T808_0x0500)
	},
	uint16(MsgT808_0x0801): func() Entity {
		return new(T808_0x0801)
	},
	uint16(MsgT808_0x8001): func() Entity {
		return new(T808_0x8001)
	},
	uint16(MsgT808_0x8003): func() Entity {
		return new(T808_0x8003)
	},
	uint16(MsgT808_0x8100): func() Entity {
		return new(T808_0x8100)
	},
	uint16(MsgT808_0x8103): func() Entity {
		return new(T808_0x8103)
	},
	uint16(MsgT808_0x8104): func() Entity {
		return new(T808_0x8104)
	},
	uint16(MsgT808_0x8105): func() Entity {
		return new(T808_0x8105)
	},
	uint16(MsgT808_0x8106): func() Entity {
		return new(T808_0x8106)
	},
	uint16(MsgT808_0x8107): func() Entity {
		return new(T808_0x8107)
	},
	uint16(MsgT808_0x8108): func() Entity {
		return new(T808_0x8108)
	},
	uint16(MsgT808_0x8201): func() Entity {
		return new(T808_0x8201)
	},
	uint16(MsgT808_0x8202): func() Entity {
		return new(T808_0x8202)
	},
	uint16(MsgT808_0x8203): func() Entity {
		return new(T808_0x8203)
	},
	uint16(MsgT808_0x8300): func() Entity {
		return new(T808_0x8300)
	},
	uint16(MsgT808_0x8301): func() Entity {
		return new(T808_0x8301)
	},
	uint16(MsgT808_0x8302): func() Entity {
		return new(T808_0x8302)
	},
	uint16(MsgT808_0x8303): func() Entity {
		return new(T808_0x8303)
	},
	uint16(MsgT808_0x8304): func() Entity {
		return new(T808_0x8304)
	},
	uint16(MsgT808_0x8400): func() Entity {
		return new(T808_0x8400)
	},
	uint16(MsgT808_0x8401): func() Entity {
		return new(T808_0x8401)
	},
	uint16(MsgT808_0x8500): func() Entity {
		return new(T808_0x8500)
	},
	uint16(MsgT808_0x8600): func() Entity {
		return new(T808_0x8600)
	},
	uint16(MsgT808_0x8601): func() Entity {
		return new(T808_0x8601)
	},
	uint16(MsgT808_0x8602): func() Entity {
		return new(T808_0x8602)
	},
	uint16(MsgT808_0x8603): func() Entity {
		return new(T808_0x8603)
	},
	uint16(MsgT808_0x8604): func() Entity {
		return new(T808_0x8604)
	},
	uint16(MsgT808_0x8605): func() Entity {
		return new(T808_0x8605)
	},
	uint16(MsgT808_0x8606): func() Entity {
		return new(T808_0x8606)
	},
	uint16(MsgT808_0x8607): func() Entity {
		return new(T808_0x8607)
	},
	uint16(MsgT808_0x8800): func() Entity {
		return new(T808_0x8800)
	},
	uint16(MsgT808_0x8900): func() Entity {
		return new(T808_0x8900)
	},
	uint16(MsgT808_0x0900): func() Entity {
		return new(T808_0x0900)
	},
}

// 类型注册
func Register(typ uint16, creator func() Entity) {
	entityMapper[typ] = creator
}
