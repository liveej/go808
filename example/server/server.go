package main

import (
	"fmt"

	"gitee.com/coco/go808"
	"gitee.com/coco/go808/protocol"
	"gitee.com/coco/go808/protocol/extra"
	log "github.com/sirupsen/logrus"
)

// 处理终端鉴权
func handleAuthentication(session *go808.Session, message *protocol.Message) {
	// 回复平台应答
	session.Reply(message, protocol.T808_0x8100_ResultSuccess)

	// 查询终端参数
	session.Request(new(protocol.T808_0x8104), func(answer *protocol.Message) {
		response := answer.Body.(*protocol.T808_0x0104)
		for _, param := range response.Params {
			fmt.Println("参数ID", param.ID())
		}
	})
}

// 处理上报位置
func handleReportLocation(session *go808.Session, message *protocol.Message) {
	// 打印消息
	entity := message.Body.(*protocol.T808_0x0200)
	fields := log.Fields{
		"IccID": message.Header.IccID,
		"警告":    fmt.Sprintf("0x%x", entity.Alarm),
		"状态":    fmt.Sprintf("0x%x", entity.Status),
		"纬度":    entity.Lat,
		"经度":    entity.Lng,
		"海拔":    entity.Altitude,
		"速度":    entity.Speed,
		"方向":    entity.Direction,
		"时间":    entity.Time,
	}

	for _, ext := range entity.Extras {
		switch ext.ID() {
		case extra.Extra_0x01{}.ID():
			fields["行驶里程"] = ext.(*extra.Extra_0x01).Value()
		case extra.Extra_0x02{}.ID():
			fields["剩余油量"] = ext.(*extra.Extra_0x02).Value()
		}
	}
	log.WithFields(fields).Info("上报终端位置信息")

	// 回复平台应答
	session.Reply(message, protocol.T808_0x8100_ResultSuccess)
}

// 处理批量上报位置
func handleBatchReportLocation(session *go808.Session, message *protocol.Message) {
	// 打印消息
	entity := message.Body.(*protocol.T808_0x0704)
	fields := log.Fields{
		"IccID": message.Header.IccID,
		"类型":    fmt.Sprintf("0x%x", entity.Type),
	}

	for _, rpt := range entity.Items {
		logRpt := log.Fields{
			"IccID": message.Header.IccID,
			"警告":    fmt.Sprintf("0x%x", rpt.Alarm),
			"状态":    fmt.Sprintf("0x%x", rpt.Status),
			"纬度":    rpt.Lat,
			"经度":    rpt.Lng,
			"海拔":    rpt.Altitude,
			"速度":    rpt.Speed,
			"方向":    rpt.Direction,
			"时间":    rpt.Time,
		}
		log.WithFields(logRpt).Info("批量上报位置items")
	}
	log.WithFields(fields).Info("批量上报终端位置信息")

	// 回复平台应答
	session.Reply(message, 0)
}

// 处理终端心跳
func handleHeartBeat(session *go808.Session, message *protocol.Message) {
	// 回复平台应答
	session.Reply(message, 0)
}

func main() {

	server, _ := go808.NewServer(go808.Options{
		Keepalive:       60,
		AutoMergePacket: true,
		CloseHandler:    nil,
		//PrivateKey:      privateKey,
	})
	server.AddHandler(protocol.MsgT808_0x0002, handleHeartBeat)
	server.AddHandler(protocol.MsgT808_0x0102, handleAuthentication)
	server.AddHandler(protocol.MsgT808_0x0200, handleReportLocation)
	server.AddHandler(protocol.MsgT808_0x0704, handleBatchReportLocation)
	server.Run("tcp", 8808)
}
