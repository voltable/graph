package triples_test

import (
	"encoding/binary"
	fmt "fmt"
	"math"
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.Graph/container/table"
	"github.com/RossMerr/Caudex.Graph/container/triples"
	proto "github.com/golang/protobuf/proto"
)

func TestNewTripleFromTable(t *testing.T) {

	type args struct {
		t  func(string) []*triples.Triple
		r  []string
		c  []string
		in string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Table Triples",
			args: args{
				in: `log_id src_ip server_ip
001 128.0.0.1 208.29.69.138
002 192.168.1.2 157.166.255.18
003 128.0.0.1 74.125.224.72`,
				t: func(in string) []*triples.Triple {
					table := table.NewTableFromReader(3, 5, strings.NewReader(in))
					return triples.NewTriplesFromTable(table)
				},
				r: []string{"log_id|001", "log_id|001", "log_id|002", "log_id|002", "log_id|003", "log_id|003"},
				c: []string{"src_ip|128.0.0.1", "server_ip|208.29.69.138", "src_ip|192.168.1.2", "server_ip|157.166.255.18", "src_ip|128.0.0.1", "server_ip|74.125.224.72"},
			},
		},
		{
			name: "Table Transpose Triples",
			args: args{
				in: `log_id src_ip server_ip
001 128.0.0.1 208.29.69.138
002 192.168.1.2 157.166.255.18
003 128.0.0.1 74.125.224.72`,
				t: func(in string) []*triples.Triple {
					table := table.NewTableFromReader(3, 5, strings.NewReader(in))
					return triples.NewTripleTransposeFromTable(table)
				},
				r: []string{"src_ip|128.0.0.1", "server_ip|208.29.69.138", "src_ip|192.168.1.2", "server_ip|157.166.255.18", "src_ip|128.0.0.1", "server_ip|74.125.224.72"},
				c: []string{"log_id|001", "log_id|001", "log_id|002", "log_id|002", "log_id|003", "log_id|003"},
			},
		},
		{
			name: "Transpose",
			args: args{
				in: `log_id src_ip server_ip
001 128.0.0.1 208.29.69.138
002 192.168.1.2 157.166.255.18
003 128.0.0.1 74.125.224.72`,
				t: func(in string) []*triples.Triple {
					table := table.NewTableFromReader(3, 5, strings.NewReader(in))
					return triples.Transpose(triples.NewTriplesFromTable(table))
				},
				r: []string{"src_ip|128.0.0.1", "server_ip|208.29.69.138", "src_ip|192.168.1.2", "server_ip|157.166.255.18", "src_ip|128.0.0.1", "server_ip|74.125.224.72"},
				c: []string{"log_id|001", "log_id|001", "log_id|002", "log_id|002", "log_id|003", "log_id|003"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triples := tt.args.t(tt.args.in)
			for i, triple := range triples {
				if tt.args.r[i] != triple.Row {
					t.Errorf("%+v got %+v, want %+v", tt.name, triple.Row, tt.args.r[i])
				}

				if tt.args.c[i] != triple.Column {
					t.Errorf("%+v got %+v, want %+v", tt.name, triple.Column, tt.args.c[i])
				}
			}
		})
	}
}

func Test_Proto(t *testing.T) {
	any, _ := triples.NewAny(float64(5))

	triple := &triples.Triple{Row: "1", Column: "2", Value: any}

	in, err := proto.Marshal(triple)
	if err != nil {
		t.Errorf("Failed to encode address book: %+v", err)
	}

	out := &triples.Triple{}
	if err := proto.Unmarshal(in, out); err != nil {
		t.Errorf("Failed to parse address book: %+v", err)
	}

	fmt.Printf("tripe: %+v", out)
	i, err := out.Value.Interface()

	fmt.Printf("\nfloat64: %+v", i)

}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}
