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

// LanguagePackString represents TL type `languagePackString#4df0e460`.
type LanguagePackString struct {
	// String key
	Key string
	// String value; pass null if the string needs to be taken from the built-in English
	// language pack
	Value LanguagePackStringValueClass
}

// LanguagePackStringTypeID is TL type id of LanguagePackString.
const LanguagePackStringTypeID = 0x4df0e460

// Ensuring interfaces in compile-time for LanguagePackString.
var (
	_ bin.Encoder     = &LanguagePackString{}
	_ bin.Decoder     = &LanguagePackString{}
	_ bin.BareEncoder = &LanguagePackString{}
	_ bin.BareDecoder = &LanguagePackString{}
)

func (l *LanguagePackString) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Key == "") {
		return false
	}
	if !(l.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LanguagePackString) String() string {
	if l == nil {
		return "LanguagePackString(nil)"
	}
	type Alias LanguagePackString
	return fmt.Sprintf("LanguagePackString%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LanguagePackString) TypeID() uint32 {
	return LanguagePackStringTypeID
}

// TypeName returns name of type in TL schema.
func (*LanguagePackString) TypeName() string {
	return "languagePackString"
}

// TypeInfo returns info about TL type.
func (l *LanguagePackString) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "languagePackString",
		ID:   LanguagePackStringTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Key",
			SchemaName: "key",
		},
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LanguagePackString) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackString#4df0e460 as nil")
	}
	b.PutID(LanguagePackStringTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LanguagePackString) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackString#4df0e460 as nil")
	}
	b.PutString(l.Key)
	if l.Value == nil {
		return fmt.Errorf("unable to encode languagePackString#4df0e460: field value is nil")
	}
	if err := l.Value.Encode(b); err != nil {
		return fmt.Errorf("unable to encode languagePackString#4df0e460: field value: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *LanguagePackString) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackString#4df0e460 to nil")
	}
	if err := b.ConsumeID(LanguagePackStringTypeID); err != nil {
		return fmt.Errorf("unable to decode languagePackString#4df0e460: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LanguagePackString) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackString#4df0e460 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackString#4df0e460: field key: %w", err)
		}
		l.Key = value
	}
	{
		value, err := DecodeLanguagePackStringValue(b)
		if err != nil {
			return fmt.Errorf("unable to decode languagePackString#4df0e460: field value: %w", err)
		}
		l.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LanguagePackString) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackString#4df0e460 as nil")
	}
	b.ObjStart()
	b.PutID("languagePackString")
	b.Comma()
	b.FieldStart("key")
	b.PutString(l.Key)
	b.Comma()
	b.FieldStart("value")
	if l.Value == nil {
		return fmt.Errorf("unable to encode languagePackString#4df0e460: field value is nil")
	}
	if err := l.Value.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode languagePackString#4df0e460: field value: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LanguagePackString) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackString#4df0e460 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("languagePackString"); err != nil {
				return fmt.Errorf("unable to decode languagePackString#4df0e460: %w", err)
			}
		case "key":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackString#4df0e460: field key: %w", err)
			}
			l.Key = value
		case "value":
			value, err := DecodeTDLibJSONLanguagePackStringValue(b)
			if err != nil {
				return fmt.Errorf("unable to decode languagePackString#4df0e460: field value: %w", err)
			}
			l.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetKey returns value of Key field.
func (l *LanguagePackString) GetKey() (value string) {
	if l == nil {
		return
	}
	return l.Key
}

// GetValue returns value of Value field.
func (l *LanguagePackString) GetValue() (value LanguagePackStringValueClass) {
	if l == nil {
		return
	}
	return l.Value
}
