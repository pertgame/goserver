package gate

import (
	"encoding/binary"
	"github.com/0990/goserver/util"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"reflect"
)

type Processor struct {
	littleEndian bool
	msgID2Info   map[uint16]*MsgInfo
}

type MsgInfo struct {
	msgType    reflect.Type
	msgHandler MsgHandler
}

type MsgHandler func(client *Client, msg proto.Message)

func NewProcessor() *Processor {
	p := new(Processor)
	p.littleEndian = false
	p.msgID2Info = make(map[uint16]*MsgInfo)
	return p
}

func (p *Processor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

func (p *Processor) Register(msg proto.Message) {
	msgID, msgType := util.ProtoHash(msg)
	if _, ok := p.msgID2Info[msgID]; ok {
		logrus.Errorf("message %s is already registered", msgType)
		return
	}

	msgInfo := new(MsgInfo)
	msgInfo.msgType = msgType
	p.msgID2Info[msgID] = msgInfo
	return
}

func (p *Processor) SetRouter(msg proto.Message) {
	msgID, msgType := util.ProtoHash(msg)
	_, ok := p.msgID2Info[msgID]
	if !ok {
		logrus.Errorf("message %s not registered", msgType)
		return
	}
	//TODO add router code

}

func (p *Processor) SetHandler(msg proto.Message, msgHandler MsgHandler) {
	msgID, msgType := util.ProtoHash(msg)
	msgInfo, ok := p.msgID2Info[msgID]
	if !ok {
		logrus.Errorf("message %s not registered", msgType)
		return
	}

	msgInfo.msgHandler = msgHandler
}

func (p *Processor) Route(msg proto.Message, client *Client) error {
	msgID, msgType := util.ProtoHash(msg)
	msgInfo, ok := p.msgID2Info[msgID]
	if !ok {
		logrus.Errorf("message %s not registered", msgType)
		return nil
	}

	if msgInfo.msgHandler != nil {
		msgInfo.msgHandler(client, msg)
	}

	return nil
}

func (p *Processor) Unmarshal(data []byte) (proto.Message, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short")
	}

	var msgID uint16
	if p.littleEndian {
		msgID = binary.LittleEndian.Uint16(data)
	} else {
		msgID = binary.BigEndian.Uint16(data)
	}

	msgInfo, exist := p.msgID2Info[msgID]
	if !exist {
		return nil, errors.New("msgID not registered")
	}

	msg := reflect.New(msgInfo.msgType.Elem()).Interface().(proto.Message)
	return msg, proto.Unmarshal(data[2:], msg.(proto.Message))
}

func (p *Processor) Marshal(msg proto.Message) ([]byte, error) {
	msgID, _ := util.ProtoHash(msg)

	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	msgIDData := make([]byte, 2)
	if p.littleEndian {
		binary.LittleEndian.PutUint16(msgIDData, msgID)
	} else {
		binary.BigEndian.PutUint16(msgIDData, msgID)
	}

	//TODO 性能优化
	ret := append(msgIDData, data...)
	return ret, nil
}
