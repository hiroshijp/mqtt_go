package packet

import "errors"

type FixedHeader struct {
	PacketType      byte
	Dup             byte
	QoS1            byte
	QoS2            byte
	Retain          byte
	RemainingLength uint
}

func ToFixedHeader(bs []byte) (FixedHeader, error) {
	if len(bs) <= 1 {
		return FixedHeader{}, errors.New("error occured")
	}
	b := bs[0]
	packetType := b >> 4
	dup := refbit(b, 3)
	qos1 := refbit(b, 2)
	qos2 := refbit(b, 1)
	retain := refbit(b, 0)
	remainingLength := decodeRemainingLength(bs[1:])
	return FixedHeader{
		PacketType:      packetType,
		Dup:             dup,
		QoS1:            qos1,
		QoS2:            qos2,
		Retain:          retain,
		RemainingLength: remainingLength,
	}, nil
}

func refbit(b byte, n uint) byte {
	return (b >> n) & 1
}

// 何バイト残っているかを表すバイト、msbは予約されているため128以上の時は要素が2つ
func decodeRemainingLength(bs []byte) uint {
	multiplier := uint(1)
	var value uint
	for i := uint(0); i < 8; i++ {
		b := bs[i]
		digit := b
		value = value + uint(digit&127)*multiplier
		multiplier = multiplier * 128
		// msbを"1000 0000"でチェック、"0000 0000"になったら終了
		if (digit & 128) == 0 {
			break
		}
	}
	return value
}
