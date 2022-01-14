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

// GetAccountTTLRequest represents TL type `getAccountTtl#e58a8b77`.
type GetAccountTTLRequest struct {
}

// GetAccountTTLRequestTypeID is TL type id of GetAccountTTLRequest.
const GetAccountTTLRequestTypeID = 0xe58a8b77

// Ensuring interfaces in compile-time for GetAccountTTLRequest.
var (
	_ bin.Encoder     = &GetAccountTTLRequest{}
	_ bin.Decoder     = &GetAccountTTLRequest{}
	_ bin.BareEncoder = &GetAccountTTLRequest{}
	_ bin.BareDecoder = &GetAccountTTLRequest{}
)

func (g *GetAccountTTLRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetAccountTTLRequest) String() string {
	if g == nil {
		return "GetAccountTTLRequest(nil)"
	}
	type Alias GetAccountTTLRequest
	return fmt.Sprintf("GetAccountTTLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetAccountTTLRequest) TypeID() uint32 {
	return GetAccountTTLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetAccountTTLRequest) TypeName() string {
	return "getAccountTtl"
}

// TypeInfo returns info about TL type.
func (g *GetAccountTTLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getAccountTtl",
		ID:   GetAccountTTLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetAccountTTLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAccountTtl#e58a8b77 as nil")
	}
	b.PutID(GetAccountTTLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetAccountTTLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAccountTtl#e58a8b77 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetAccountTTLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAccountTtl#e58a8b77 to nil")
	}
	if err := b.ConsumeID(GetAccountTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getAccountTtl#e58a8b77: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetAccountTTLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAccountTtl#e58a8b77 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetAccountTTLRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getAccountTtl#e58a8b77 as nil")
	}
	b.ObjStart()
	b.PutID("getAccountTtl")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetAccountTTLRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getAccountTtl#e58a8b77 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getAccountTtl"); err != nil {
				return fmt.Errorf("unable to decode getAccountTtl#e58a8b77: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAccountTTL invokes method getAccountTtl#e58a8b77 returning error if any.
func (c *Client) GetAccountTTL(ctx context.Context) (*AccountTTL, error) {
	var result AccountTTL

	request := &GetAccountTTLRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
