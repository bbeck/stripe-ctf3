// Code generated by protoc-gen-gogo.
// source: unrecognizedgroup.proto
// DO NOT EDIT!

package unrecognizedgroup

import testing "testing"
import math_rand "math/rand"
import time "time"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"
import testing1 "testing"
import math_rand1 "math/rand"
import time1 "time"
import encoding_json "encoding/json"
import testing2 "testing"
import math_rand2 "math/rand"
import time2 "time"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"
import math_rand3 "math/rand"
import time3 "time"
import testing3 "testing"
import fmt "fmt"
import math_rand4 "math/rand"
import time4 "time"
import testing4 "testing"
import code_google_com_p_gogoprotobuf_proto2 "code.google.com/p/gogoprotobuf/proto"
import math_rand5 "math/rand"
import time5 "time"
import testing5 "testing"
import fmt1 "fmt"
import go_parser "go/parser"
import math_rand6 "math/rand"
import time6 "time"
import testing6 "testing"
import code_google_com_p_gogoprotobuf_proto3 "code.google.com/p/gogoprotobuf/proto"
import testing7 "testing"

func TestNewNoGroupProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NewNoGroup{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNewNoGroupMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &NewNoGroup{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroupProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedOldWithGroup(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroup_Group1Proto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group1(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup_Group1{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroup_Group2Proto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group2(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup_Group2{}
	if err := code_google_com_p_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNewNoGroupJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NewNoGroup{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestAJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestOldWithGroupJSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedOldWithGroup(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestOldWithGroup_Group1JSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group1(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup_Group1{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestOldWithGroup_Group2JSON(t *testing1.T) {
	popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group2(popr, true)
	jsondata, err := encoding_json.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup_Group2{}
	err = encoding_json.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestNewNoGroupProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &NewNoGroup{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNewNoGroupProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &NewNoGroup{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestAProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroupProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOldWithGroup(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &OldWithGroup{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroupProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOldWithGroup(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &OldWithGroup{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroup_Group1ProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group1(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &OldWithGroup_Group1{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroup_Group1ProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group1(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &OldWithGroup_Group1{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroup_Group2ProtoText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group2(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.MarshalTextString(p)
	msg := &OldWithGroup_Group2{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestOldWithGroup_Group2ProtoCompactText(t *testing2.T) {
	popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group2(popr, true)
	data := code_google_com_p_gogoprotobuf_proto1.CompactTextString(p)
	msg := &OldWithGroup_Group2{}
	if err := code_google_com_p_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestNewNoGroupStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestAStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestOldWithGroupStringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedOldWithGroup(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestOldWithGroup_Group1Stringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group1(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestOldWithGroup_Group2Stringer(t *testing3.T) {
	popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group2(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestNewNoGroupSize(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func TestASize(t *testing4.T) {
	popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
	p := NewPopulatedA(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto2.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto2.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func TestNewNoGroupGoString(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestAGoString(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestOldWithGroupGoString(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedOldWithGroup(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestOldWithGroup_Group1GoString(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group1(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestOldWithGroup_Group2GoString(t *testing5.T) {
	popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group2(popr, false)
	s1 := p.GoString()
	s2 := fmt1.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestNewNoGroupVerboseEqual(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedNewNoGroup(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &NewNoGroup{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestAVerboseEqual(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedA(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &A{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestOldWithGroupVerboseEqual(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedOldWithGroup(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestOldWithGroup_Group1VerboseEqual(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group1(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup_Group1{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestOldWithGroup_Group2VerboseEqual(t *testing6.T) {
	popr := math_rand6.New(math_rand6.NewSource(time6.Now().UnixNano()))
	p := NewPopulatedOldWithGroup_Group2(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto3.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &OldWithGroup_Group2{}
	if err := code_google_com_p_gogoprotobuf_proto3.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestUnrecognizedgroupDescription(t *testing7.T) {
	UnrecognizedgroupDescription()
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
