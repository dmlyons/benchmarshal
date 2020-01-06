package benchm

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/ugorji/go/codec"
)

var data = createCollection(1000)

func loadSlice() RDataCollection {
	return data
}

func checkEqual(expected, actual RDataCollection) bool {
	js1, _ := json.Marshal(expected)
	js2, _ := json.Marshal(expected)
	return bytes.Equal(js1, js2)
}

func BenchmarkAll(b *testing.B) {
	for i := 1; i <= 100000; i *= 10 {
		r := createCollection(i)
		b.Run(fmt.Sprintf("JSON %d", i), func(b *testing.B) {
			js := []byte{}
			b.Run("marshal", func(b *testing.B) {
				// bench marshalling the slice
				for n := 0; n < b.N; n++ {
					js = []byte{}
					js, _ = json.Marshal(r)
				}
			})
			b.Logf("JSON Marshal Size: %d", len(js))
			r = make(RDataCollection, 0)

			b.Run("unmarshal", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					json.Unmarshal(js, &r)
				}
			})
		})
		b.Run(fmt.Sprintf("Gob %d", i), func(b *testing.B) {
			var buf bytes.Buffer
			b.Run("encode", func(b *testing.B) {
				// bench marshalling the slice
				for n := 0; n < b.N; n++ {
					buf = bytes.Buffer{}
					e := gob.NewEncoder(&buf)
					e.Encode(r)
				}
			})
			b.Logf("Gob Encode Size: %d", buf.Len())
			r = make(RDataCollection, 0)

			g := buf.Bytes()
			b.Run("decode", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					bf := bytes.NewBuffer(g)
					dec := gob.NewDecoder(bf)
					dec.Decode(&r)
				}
			})
		})

		b.Run(fmt.Sprintf("Msgpack %d", i), func(b *testing.B) {
			var mh codec.MsgpackHandle
			js := []byte{}
			b.Run("marshal", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					enc := codec.NewEncoderBytes(&js, &mh)
					_ = enc.Encode(r)
				}
			})
			b.Logf("codec Msgpack Marshal Size: %d", len(js))
			r = make(RDataCollection, 0)

			b.Run("unmarshal", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					dec := codec.NewDecoderBytes(js, &mh)
					_ = dec.Decode(&r)
				}
			})
		})

		// quick sanity check once a round
		if !checkEqual(data, r) {
			b.Fatal("NOT EQUAL!")
		}
	}
}

func createCollection(num int) RDataCollection {
	r := make(RDataCollection, 0, num)
	for i := 1; i <= num; i++ {
		now := time.Now()
		nid := randomdata.Number(1, 100000)
		rid := randomdata.Number(1, 100000)
		tid := randomdata.Number(1, 100000)
		val := randomdata.Alphanumeric(12)
		r = append(r, &RData{
			Date:  &now,
			NID:   &nid,
			RID:   &rid,
			Tid:   &tid,
			Value: &val,
		})
	}
	return r
}

type RData struct {
	Date  *time.Time
	NID   *int
	RID   *int
	Tid   *int
	Value interface{}
}

type RDataCollection []*RData
