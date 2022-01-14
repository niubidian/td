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

// SynchronizeLanguagePackRequest represents TL type `synchronizeLanguagePack#84e5e72e`.
type SynchronizeLanguagePackRequest struct {
	// Language pack identifier
	LanguagePackID string
}

// SynchronizeLanguagePackRequestTypeID is TL type id of SynchronizeLanguagePackRequest.
const SynchronizeLanguagePackRequestTypeID = 0x84e5e72e

// Ensuring interfaces in compile-time for SynchronizeLanguagePackRequest.
var (
	_ bin.Encoder     = &SynchronizeLanguagePackRequest{}
	_ bin.Decoder     = &SynchronizeLanguagePackRequest{}
	_ bin.BareEncoder = &SynchronizeLanguagePackRequest{}
	_ bin.BareDecoder = &SynchronizeLanguagePackRequest{}
)

func (s *SynchronizeLanguagePackRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.LanguagePackID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SynchronizeLanguagePackRequest) String() string {
	if s == nil {
		return "SynchronizeLanguagePackRequest(nil)"
	}
	type Alias SynchronizeLanguagePackRequest
	return fmt.Sprintf("SynchronizeLanguagePackRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SynchronizeLanguagePackRequest) TypeID() uint32 {
	return SynchronizeLanguagePackRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SynchronizeLanguagePackRequest) TypeName() string {
	return "synchronizeLanguagePack"
}

// TypeInfo returns info about TL type.
func (s *SynchronizeLanguagePackRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "synchronizeLanguagePack",
		ID:   SynchronizeLanguagePackRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LanguagePackID",
			SchemaName: "language_pack_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SynchronizeLanguagePackRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode synchronizeLanguagePack#84e5e72e as nil")
	}
	b.PutID(SynchronizeLanguagePackRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SynchronizeLanguagePackRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode synchronizeLanguagePack#84e5e72e as nil")
	}
	b.PutString(s.LanguagePackID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SynchronizeLanguagePackRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode synchronizeLanguagePack#84e5e72e to nil")
	}
	if err := b.ConsumeID(SynchronizeLanguagePackRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode synchronizeLanguagePack#84e5e72e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SynchronizeLanguagePackRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode synchronizeLanguagePack#84e5e72e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode synchronizeLanguagePack#84e5e72e: field language_pack_id: %w", err)
		}
		s.LanguagePackID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SynchronizeLanguagePackRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode synchronizeLanguagePack#84e5e72e as nil")
	}
	b.ObjStart()
	b.PutID("synchronizeLanguagePack")
	b.Comma()
	b.FieldStart("language_pack_id")
	b.PutString(s.LanguagePackID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SynchronizeLanguagePackRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode synchronizeLanguagePack#84e5e72e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("synchronizeLanguagePack"); err != nil {
				return fmt.Errorf("unable to decode synchronizeLanguagePack#84e5e72e: %w", err)
			}
		case "language_pack_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode synchronizeLanguagePack#84e5e72e: field language_pack_id: %w", err)
			}
			s.LanguagePackID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLanguagePackID returns value of LanguagePackID field.
func (s *SynchronizeLanguagePackRequest) GetLanguagePackID() (value string) {
	if s == nil {
		return
	}
	return s.LanguagePackID
}

// SynchronizeLanguagePack invokes method synchronizeLanguagePack#84e5e72e returning error if any.
func (c *Client) SynchronizeLanguagePack(ctx context.Context, languagepackid string) error {
	var ok Ok

	request := &SynchronizeLanguagePackRequest{
		LanguagePackID: languagepackid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
