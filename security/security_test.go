package security

import (
	"strconv"
	"testing"
)

var tests = []struct {
	plainPwd string
	hashPwd  string
}{
	{"example1", "$2a$10$7n6u96U2UD5polWnnhW7.eBXDrGt.VgKpQDARtiqhgjelsy0sCN/S"},
	{"p@ssword", "$2a$10$C3ybIDyzsAhx1MgA33yVguC5oR0zyHP/M2ExaKTv9wioNqWz9WfDq"},
	{"abcdefghi", "$2a$10$29ZeqsSWFo1kQ1BfU3Cot.d7tOftergFVyZCK2e/.QBfnx3WaCZ3y"},
	{"123456789", "$2a$10$aQOUsb6QbLdP6qB8uVt6Xug8jiA.ZIRtE1hxhXgH/7mBfGKtROhcq"},
	{"111111111111", "$2a$10$dEWSP91xmDdRODB/Mvkh3OJV3xzo4ZrxVHgVjeq3ZPc/geDhJjZai"},
	{"ILoveAxth", "$2a$10$kzl8SAEkfYCl8g8.wRxYn.TQ4iRe6qnVUW5qzxI6oZcS7/H.POHZG"},
}

func TestGeneratePwd(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := GeneratePwd(tt.plainPwd)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestComparePwd(t *testing.T) {
	for i, example := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := ComparePwd(example.hashPwd, example.plainPwd)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
