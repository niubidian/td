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

// GetStickersRequest represents TL type `getStickers#a0ef757c`.
type GetStickersRequest struct {
	// String representation of emoji. If empty, returns all known installed stickers
	Emoji string
	// The maximum number of stickers to be returned
	Limit int32
}

// GetStickersRequestTypeID is TL type id of GetStickersRequest.
const GetStickersRequestTypeID = 0xa0ef757c

// Ensuring interfaces in compile-time for GetStickersRequest.
var (
	_ bin.Encoder     = &GetStickersRequest{}
	_ bin.Decoder     = &GetStickersRequest{}
	_ bin.BareEncoder = &GetStickersRequest{}
	_ bin.BareDecoder = &GetStickersRequest{}
)

func (g *GetStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Emoji == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStickersRequest) String() string {
	if g == nil {
		return "GetStickersRequest(nil)"
	}
	type Alias GetStickersRequest
	return fmt.Sprintf("GetStickersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStickersRequest) TypeID() uint32 {
	return GetStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStickersRequest) TypeName() string {
	return "getStickers"
}

// TypeInfo returns info about TL type.
func (g *GetStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStickers",
		ID:   GetStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emoji",
			SchemaName: "emoji",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStickers#a0ef757c as nil")
	}
	b.PutID(GetStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStickers#a0ef757c as nil")
	}
	b.PutString(g.Emoji)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStickers#a0ef757c to nil")
	}
	if err := b.ConsumeID(GetStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStickers#a0ef757c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStickers#a0ef757c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getStickers#a0ef757c: field emoji: %w", err)
		}
		g.Emoji = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getStickers#a0ef757c: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetStickersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getStickers#a0ef757c as nil")
	}
	b.ObjStart()
	b.PutID("getStickers")
	b.Comma()
	b.FieldStart("emoji")
	b.PutString(g.Emoji)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetStickersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getStickers#a0ef757c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getStickers"); err != nil {
				return fmt.Errorf("unable to decode getStickers#a0ef757c: %w", err)
			}
		case "emoji":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getStickers#a0ef757c: field emoji: %w", err)
			}
			g.Emoji = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getStickers#a0ef757c: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmoji returns value of Emoji field.
func (g *GetStickersRequest) GetEmoji() (value string) {
	if g == nil {
		return
	}
	return g.Emoji
}

// GetLimit returns value of Limit field.
func (g *GetStickersRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetStickers invokes method getStickers#a0ef757c returning error if any.
func (c *Client) GetStickers(ctx context.Context, request *GetStickersRequest) (*Stickers, error) {
	var result Stickers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
