package common

import (
	"bytes"
	"testing"
)

func TestParseColonLine(t *testing.T) {
	tcs := []struct{
		file []byte
		parseColon int
		entries [][]byte
		remain  []byte
	}{
		{
			[]byte{},
			2,
			[][]byte{},
			[]byte{},
		},
		{
			[]byte("1:2:3:4\n"),
			4,
			[][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4")},
			[]byte{},
		},
		{
			[]byte("1:2:3:4\n"),
			3,
			[][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4")},
			[]byte{},
		},
		{
			[]byte("1:2:3:4"),
			4,
			[][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4")},
			[]byte{},
		},
		{
			[]byte("1:2:3:4\n"),
			2,
			[][]byte{[]byte("1"), []byte("2"), []byte("3")},
			[]byte("4\n"),
		},
	}
	for _, tc := range tcs {
		resEntries, resRemain := ParseColonLine(tc.file, tc.parseColon)
		for i := range resEntries {
			if i >= len(tc.entries) {
				t.Errorf("entries length are not equal, res: %d, expect: %d", len(resEntries), len(tc.entries))
			}

			if !bytes.Equal(resEntries[i], tc.entries[i]) {
				t.Errorf("entries are not equal, res: %v, expect: %v", resEntries[i], tc.entries[i])
			}
		}
		if !bytes.Equal(resRemain, tc.remain) {
			t.Errorf("remains are not equal, res: %v, expect: %v", resRemain, tc.remain)
		}
	}
}
