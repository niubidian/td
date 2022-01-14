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

// SetBioRequest represents TL type `setBio#9f772354`.
type SetBioRequest struct {
	// The new value of the user bio; 0-70 characters without line feeds
	Bio string
}

// SetBioRequestTypeID is TL type id of SetBioRequest.
const SetBioRequestTypeID = 0x9f772354

// Ensuring interfaces in compile-time for SetBioRequest.
var (
	_ bin.Encoder     = &SetBioRequest{}
	_ bin.Decoder     = &SetBioRequest{}
	_ bin.BareEncoder = &SetBioRequest{}
	_ bin.BareDecoder = &SetBioRequest{}
)

func (s *SetBioRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Bio == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBioRequest) String() string {
	if s == nil {
		return "SetBioRequest(nil)"
	}
	type Alias SetBioRequest
	return fmt.Sprintf("SetBioRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBioRequest) TypeID() uint32 {
	return SetBioRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBioRequest) TypeName() string {
	return "setBio"
}

// TypeInfo returns info about TL type.
func (s *SetBioRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBio",
		ID:   SetBioRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bio",
			SchemaName: "bio",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBioRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBio#9f772354 as nil")
	}
	b.PutID(SetBioRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBioRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBio#9f772354 as nil")
	}
	b.PutString(s.Bio)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBioRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBio#9f772354 to nil")
	}
	if err := b.ConsumeID(SetBioRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBio#9f772354: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBioRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBio#9f772354 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBio#9f772354: field bio: %w", err)
		}
		s.Bio = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBioRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBio#9f772354 as nil")
	}
	b.ObjStart()
	b.PutID("setBio")
	b.Comma()
	b.FieldStart("bio")
	b.PutString(s.Bio)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBioRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBio#9f772354 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBio"); err != nil {
				return fmt.Errorf("unable to decode setBio#9f772354: %w", err)
			}
		case "bio":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setBio#9f772354: field bio: %w", err)
			}
			s.Bio = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBio returns value of Bio field.
func (s *SetBioRequest) GetBio() (value string) {
	if s == nil {
		return
	}
	return s.Bio
}

// SetBio invokes method setBio#9f772354 returning error if any.
func (c *Client) SetBio(ctx context.Context, bio string) error {
	var ok Ok

	request := &SetBioRequest{
		Bio: bio,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
