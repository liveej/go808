package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go808"
	"go808/protocol"
	"go808/protocol/extra"
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
		"纬度":    entity.Latitude,
		"经度":    entity.Longitude,
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

// 处理上传媒体
func handleUploadMediaPacket(session *go808.Session, message *protocol.Message) {
	entity := message.Body.(*protocol.T808_0x0801)

	// 读取完整数据包
	fullPacket := make([]byte, 1024*1024)
	n, _ := entity.Packet.Read(fullPacket[:])
	fmt.Println(n)

	session.Send(&protocol.T808_0x8800{
		MediaID: entity.MediaID,
	})
}

func main() {
	server := go808.NewServer(go808.Options{
		Keepalive:       60,
		AutoMergePacket: true,
		CloseHandler:    nil,
	})
	server.AddHandler(protocol.MsgT808_0x0102, handleAuthentication)
	server.AddHandler(protocol.MsgT808_0x0200, handleReportLocation)
	server.AddHandler(protocol.MsgT808_0x0801, handleUploadMediaPacket)
	server.Run("tcp", 8808)
}
