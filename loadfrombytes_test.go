package vocabtxt_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-vocabtxt"
)

func TestLoadFromBytes(t *testing.T) {
	tests := []struct{
		Bytes []byte
		Expected map[string]uint
	}{
		{
			Expected: map[string]uint{},
		},



		{
			Bytes:    []byte(
				"apple",
			),
			Expected: map[string]uint{
				"apple":0,
			},
		},
		{
			Bytes:    []byte(
				"apple"+"\n",
			),
			Expected: map[string]uint{
				"apple":0,
			},
		},
		{
			Bytes:    []byte(
				"apple"+"\r\n",
			),
			Expected: map[string]uint{
				"apple":0,
			},
		},



		{
			Bytes:    []byte(
				"apple"+"\n"+
				"banana",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
			},
		},
		{
			Bytes:    []byte(
				"apple"  +"\n"+
				"banana" +"\n",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
			},
		},



		{
			Bytes:    []byte(
				"apple"+"\r\n"+
				"banana",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
			},
		},
		{
			Bytes:    []byte(
				"apple"  +"\r\n"+
				"banana" +"\r\n",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
			},
		},



		{
			Bytes:    []byte(
				"apple"  +"\n"+
				"banana" +"\n"+
				"cherry",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
				"cherry":2,
			},
		},
		{
			Bytes:    []byte(
				"apple"  +"\n"+
				"banana" +"\n"+
				"cherry" +"\n",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
				"cherry":2,
			},
		},



		{
			Bytes:    []byte(
				"apple"  +"\r\n"+
				"banana" +"\r\n"+
				"cherry",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
				"cherry":2,
			},
		},
		{
			Bytes:    []byte(
				"apple"  +"\r\n"+
				"banana" +"\r\n"+
				"cherry" +"\r\n",
			),
			Expected: map[string]uint{
				"apple":0,
				"banana":1,
				"cherry":2,
			},
		},



		{
			Bytes:    []byte(
				"the"  +"\n"+
				"of"   +"\n"+
				"and"  +"\n"+
				"in"   +"\n"+
				"to"   +"\n"+
				"was"  +"\n"+
				"he"   +"\n"+
				"is"   +"\n"+
				"as"   +"\n"+
				"for"  +"\n"+
				"on"   +"\n"+
				"with" +"\n"+
				"that" +"\n"+
				"it"   +"\n"+
				"his"  +"\n"+
				"by"   +"\n",
			),
			Expected: map[string]uint{
				"the":0,
				"of":1,
				"and":2,
				"in":3,
				"to":4,
				"was":5,
				"he":6,
				"is":7,
				"as":8,
				"for":9,
				"on":10,
				"with":11,
				"that":12,
				"it":13,
				"his":14,
				"by":15,
			},
		},
	}

	for testNumber, test := range tests {
		var actual map[string]uint = map[string]uint{}

		err := vocabtxt.LoadFromBytes(&actual, test.Bytes)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual 'map' it not what was expected.", testNumber)
			t.Logf("EXPECTED\n%#v", expected)
			t.Logf("ACTUAL:\n%#v", actual)
			t.Logf("BYTES:\n%s", test.Bytes)
			continue
		}
	}
}

func TestLoadFromBytes_nilMap(t *testing.T) {
	var bytes []byte = []byte("apple\nbanana\ncherry\n")

	var actual map[string]uint = map[string]uint(nil)

	err := vocabtxt.LoadFromBytes(&actual, bytes)
	if nil != err {
		t.Errorf("Did not expect an error but actually got one.")
	}
}
