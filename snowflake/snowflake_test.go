package snowflake_test

import (
	"fmt"
	"github.com/sendya/pkg/snowflake"
	"testing"
)

func TestNewNode(t *testing.T) {
	_, err := snowflake.NewNode(1)
	if err != nil {
		t.Error(err)
	}
}

// lazy check if Generate will create duplicate IDs
// would be good to later enhance this with more smarts
func TestGenerateDuplicateID(t *testing.T) {
	node, _ := snowflake.NewNode(1)

	var x, y snowflake.ID
	for i := 0; i < 1000000; i++ {
		y = node.Generate()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}

// Converters/Parsers Test funcs
// We should have funcs here to test conversion both ways for everything

func TestPrintAll(t *testing.T) {
	node, err := snowflake.NewNode(0)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	id := node.Generate()

	t.Logf("Int64    : %#v", id.Int64())
	t.Logf("String   : %#v", id.String())
	t.Logf("Base2    : %#v", id.Base2())
	t.Logf("Base32   : %#v", id.Base32())
	t.Logf("Base36   : %#v", id.Base36())
	t.Logf("Base58   : %#v", id.Base58())
	t.Logf("Base64   : %#v", id.Base64())
	t.Logf("Bytes    : %#v", id.Bytes())
	t.Logf("IntBytes : %#v", id.IntBytes())
}

func BenchmarkGenerate(b *testing.B) {
	node, _ := snowflake.NewNode(1)
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = node.Generate()
	}
}

func BenchmarkGenerateMaxSequence(b *testing.B) {
	snowflake.NodeBits = 1
	snowflake.StepBits = 21
	node, _ := snowflake.NewNode(1)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ID := node.Generate()
		fmt.Println(ID)
	}
}
