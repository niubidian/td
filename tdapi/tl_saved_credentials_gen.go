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

// SavedCredentials represents TL type `savedCredentials#e9ee14dc`.
type SavedCredentials struct {
	// Unique identifier of the saved credentials
	ID string
	// Title of the saved credentials
	Title string
}

// SavedCredentialsTypeID is TL type id of SavedCredentials.
const SavedCredentialsTypeID = 0xe9ee14dc

// Ensuring interfaces in compile-time for SavedCredentials.
var (
	_ bin.Encoder     = &SavedCredentials{}
	_ bin.Decoder     = &SavedCredentials{}
	_ bin.BareEncoder = &SavedCredentials{}
	_ bin.BareDecoder = &SavedCredentials{}
)

func (s *SavedCredentials) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.Title == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SavedCredentials) String() string {
	if s == nil {
		return "SavedCredentials(nil)"
	}
	type Alias SavedCredentials
	return fmt.Sprintf("SavedCredentials%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SavedCredentials) TypeID() uint32 {
	return SavedCredentialsTypeID
}

// TypeName returns name of type in TL schema.
func (*SavedCredentials) TypeName() string {
	return "savedCredentials"
}

// TypeInfo returns info about TL type.
func (s *SavedCredentials) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "savedCredentials",
		ID:   SavedCredentialsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SavedCredentials) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode savedCredentials#e9ee14dc as nil")
	}
	b.PutID(SavedCredentialsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SavedCredentials) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode savedCredentials#e9ee14dc as nil")
	}
	b.PutString(s.ID)
	b.PutString(s.Title)
	return nil
}

// Decode implements bin.Decoder.
func (s *SavedCredentials) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode savedCredentials#e9ee14dc to nil")
	}
	if err := b.ConsumeID(SavedCredentialsTypeID); err != nil {
		return fmt.Errorf("unable to decode savedCredentials#e9ee14dc: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SavedCredentials) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode savedCredentials#e9ee14dc to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedCredentials#e9ee14dc: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedCredentials#e9ee14dc: field title: %w", err)
		}
		s.Title = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SavedCredentials) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode savedCredentials#e9ee14dc as nil")
	}
	b.ObjStart()
	b.PutID("savedCredentials")
	b.Comma()
	b.FieldStart("id")
	b.PutString(s.ID)
	b.Comma()
	b.FieldStart("title")
	b.PutString(s.Title)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SavedCredentials) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode savedCredentials#e9ee14dc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("savedCredentials"); err != nil {
				return fmt.Errorf("unable to decode savedCredentials#e9ee14dc: %w", err)
			}
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode savedCredentials#e9ee14dc: field id: %w", err)
			}
			s.ID = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode savedCredentials#e9ee14dc: field title: %w", err)
			}
			s.Title = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *SavedCredentials) GetID() (value string) {
	if s == nil {
		return
	}
	return s.ID
}

// GetTitle returns value of Title field.
func (s *SavedCredentials) GetTitle() (value string) {
	if s == nil {
		return
	}
	return s.Title
}
