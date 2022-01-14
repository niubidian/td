// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// TestBytes represents TL type `testBytes#a422c4de`.
type TestBytes struct {
	// Bytes
	Value []byte
}

// TestBytesTypeID is TL type id of TestBytes.
const TestBytesTypeID = 0xa422c4de

// Ensuring interfaces in compile-time for TestBytes.
var (
	_ bin.Encoder     = &TestBytes{}
	_ bin.Decoder     = &TestBytes{}
	_ bin.BareEncoder = &TestBytes{}
	_ bin.BareDecoder = &TestBytes{}
)

func (t *TestBytes) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestBytes) String() string {
	if t == nil {
		return "TestBytes(nil)"
	}
	type Alias TestBytes
	return fmt.Sprintf("TestBytes%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestBytes) TypeID() uint32 {
	return TestBytesTypeID
}

// TypeName returns name of type in TL schema.
func (*TestBytes) TypeName() string {
	return "testBytes"
}

// TypeInfo returns info about TL type.
func (t *TestBytes) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testBytes",
		ID:   TestBytesTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestBytes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testBytes#a422c4de as nil")
	}
	b.PutID(TestBytesTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestBytes) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testBytes#a422c4de as nil")
	}
	b.PutBytes(t.Value)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestBytes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testBytes#a422c4de to nil")
	}
	if err := b.ConsumeID(TestBytesTypeID); err != nil {
		return fmt.Errorf("unable to decode testBytes#a422c4de: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestBytes) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testBytes#a422c4de to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode testBytes#a422c4de: field value: %w", err)
		}
		t.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestBytes) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testBytes#a422c4de as nil")
	}
	b.ObjStart()
	b.PutID("testBytes")
	b.Comma()
	b.FieldStart("value")
	b.PutBytes(t.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestBytes) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testBytes#a422c4de to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testBytes"); err != nil {
				return fmt.Errorf("unable to decode testBytes#a422c4de: %w", err)
			}
		case "value":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode testBytes#a422c4de: field value: %w", err)
			}
			t.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (t *TestBytes) GetValue() (value []byte) {
	if t == nil {
		return
	}
	return t.Value
}
