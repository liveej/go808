package go808

import (
	"github.com/funny/link"
	log "github.com/sirupsen/logrus"
	"go808/protocol"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

// 请求上下文
type requestContext struct {
	msgID    uint16
	serialNo uint16
	callback func(answer *protocol.Message)
}

type Session struct {
	next    uint32
	iccID   uint64
	server  *Server
	session *link.Session

	mux      sync.Mutex
	requests []requestContext

	UserData interface{}
}

// 创建Session
func newSession(server *Server, sess *link.Session) *Session {
	return &Session{
		server:  server,
		session: sess,
	}
}

// 获取ID
func (session *Session) ID() uint64 {
	return session.session.ID()
}

// 获取服务实例
func (session *Session) GetServer() *Server {
	return session.server
}

// 发送消息
func (session *Session) Send(entity protocol.Entity) (uint16, error) {
	message := protocol.Message{
		Body: entity,
		Header: protocol.Header{
			MsgID:           entity.MsgID(),
			IccID:           atomic.LoadUint64(&session.iccID),
			MessageSerialNo: session.getNextID(),
		},
	}

	err := session.session.Send(message)
	if err != nil {
		return 0, err
	}
	return message.Header.MessageSerialNo, nil
}

// 回复消息
func (session *Session) Reply(msg *protocol.Message, result protocol.ResponseResult) (uint16, error) {
	entity := protocol.T808_0x8001{
		AnswerMessageSerialNo: msg.Header.MessageSerialNo,
		ResponseMsgID:         msg.Header.MsgID,
		Result:                result,
	}
	return session.Send(&entity)
}

// 发起请求
func (session *Session) Request(entity protocol.Entity, cb func(answer *protocol.Message)) (uint16, error) {
	serialNo, err := session.Send(entity)
	if err != nil {
		return 0, err
	}

	if cb != nil {
		session.addRequestContext(requestContext{
			msgID:    uint16(entity.MsgID()),
			serialNo: serialNo,
			callback: cb,
		})
	}
	return serialNo, nil
}

// 关闭连接
func (session *Session) Close() error {
	return session.session.Close()
}

// 获取消息ID
func (session *Session) getNextID() uint16 {
	var id uint32
	for {
		id = atomic.LoadUint32(&session.next)
		if id == 0xff {
			if atomic.CompareAndSwapUint32(&session.next, id, 1) {
				id = 1
				break
			}
		} else if atomic.CompareAndSwapUint32(&session.next, id, id+1) {
			id += 1
			break
		}
	}
	return uint16(id)
}

// 消息接收事件
func (session *Session) message(message *protocol.Message) {
	if message.Header.IccID > 0 {
		old := atomic.LoadUint64(&session.iccID)
		if old != 0 && old != message.Header.IccID {
			log.WithFields(log.Fields{
				"id":  session.ID(),
				"old": old,
				"new": message.Header.IccID,
			}).Warn("[JT/T 808] terminal IccID is inconsistent")
		}
		atomic.StoreUint64(&session.iccID, message.Header.IccID)
	}

	var messageSerialNo uint16
	switch message.Header.MsgID {
	case protocol.MsgT808_0x0001:
		// 终端通用应答
		messageSerialNo = message.Body.(*protocol.T808_0x0001).AnswerMessageSerialNo
	case protocol.MsgT808_0x0104:
		// 查询终端参数应答
		messageSerialNo = message.Body.(*protocol.T808_0x0104).AnswerMessageSerialNo
	case protocol.MsgT808_0x0201:
		// 位置信息查询应答
		messageSerialNo = message.Body.(*protocol.T808_0x0201).AnswerMessageSerialNo
	}
	if messageSerialNo == 0 {
		return
	}

	ctx, ok := session.takeRequestContext(messageSerialNo)
	if ok {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
			}
		}()
		ctx.callback(message)
	}
}

// 添加请求上下文
func (session *Session) addRequestContext(ctx requestContext) {
	session.mux.Lock()
	defer session.mux.Unlock()

	for idx, item := range session.requests {
		if item.msgID == ctx.msgID {
			session.requests[idx] = ctx
			return
		}
	}
	session.requests = append(session.requests, ctx)
}

// 取出请求上下文
func (session *Session) takeRequestContext(messageSerialNo uint16) (requestContext, bool) {
	session.mux.Lock()
	defer session.mux.Unlock()

	for idx, item := range session.requests {
		if item.serialNo == messageSerialNo {
			session.requests[idx] = session.requests[len(session.requests)-1]
			session.requests = session.requests[:len(session.requests)-1]
			return item, true
		}
	}
	return requestContext{}, false
}
