package packet_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hiroshijp/mqtt_go/packet"
)

func TestToFixedHeader(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		args    args
		want    packet.FixedHeader
		wantErr bool
	}{
		{
			"Reserved",
			args{[]byte{
				0x00,
				0x00,
			}},
			packet.FixedHeader{PacketType: 0, Dup: 0, QoS1: 0, QoS2: 0, Retain: 0, RemainingLength: 0},
			false,
		},
		{
			"Connect",
			args{[]byte{
				0x1B,
				0x7F,
			}},
			packet.FixedHeader{PacketType: 1, Dup: 1, QoS1: 0, QoS2: 1, Retain: 1, RemainingLength: 127},
			false,
		},
		{
			"ConnAck",
			args{[]byte{
				0x24,
				0x80, 0x01,
			}},
			packet.FixedHeader{PacketType: 2, Dup: 0, QoS1: 1, QoS2: 0, Retain: 0, RemainingLength: 128},
			false,
		},
		{
			"Error nil",
			args{nil},
			packet.FixedHeader{},
			true,
		},
		{
			"Error 1byte",
			args{[]byte{
				0x24,
			}},
			packet.FixedHeader{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s : %#v", tt.name, tt.args.bs), func(t *testing.T) {
			got, err := packet.ToFixedHeader(tt.args.bs)

			// 返ってきたerrが欲しているerr出ない場合
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFixedHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFixedHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
