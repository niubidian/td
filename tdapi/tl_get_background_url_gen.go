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

// GetBackgroundURLRequest represents TL type `getBackgroundUrl#2bbc6fd2`.
type GetBackgroundURLRequest struct {
	// Background name
	Name string
	// Background type
	Type BackgroundTypeClass
}

// GetBackgroundURLRequestTypeID is TL type id of GetBackgroundURLRequest.
const GetBackgroundURLRequestTypeID = 0x2bbc6fd2

// Ensuring interfaces in compile-time for GetBackgroundURLRequest.
var (
	_ bin.Encoder     = &GetBackgroundURLRequest{}
	_ bin.Decoder     = &GetBackgroundURLRequest{}
	_ bin.BareEncoder = &GetBackgroundURLRequest{}
	_ bin.BareDecoder = &GetBackgroundURLRequest{}
)

func (g *GetBackgroundURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Name == "") {
		return false
	}
	if !(g.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetBackgroundURLRequest) String() string {
	if g == nil {
		return "GetBackgroundURLRequest(nil)"
	}
	type Alias GetBackgroundURLRequest
	return fmt.Sprintf("GetBackgroundURLRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetBackgroundURLRequest) TypeID() uint32 {
	return GetBackgroundURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetBackgroundURLRequest) TypeName() string {
	return "getBackgroundUrl"
}

// TypeInfo returns info about TL type.
func (g *GetBackgroundURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getBackgroundUrl",
		ID:   GetBackgroundURLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetBackgroundURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBackgroundUrl#2bbc6fd2 as nil")
	}
	b.PutID(GetBackgroundURLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetBackgroundURLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBackgroundUrl#2bbc6fd2 as nil")
	}
	b.PutString(g.Name)
	if g.Type == nil {
		return fmt.Errorf("unable to encode getBackgroundUrl#2bbc6fd2: field type is nil")
	}
	if err := g.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getBackgroundUrl#2bbc6fd2: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetBackgroundURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBackgroundUrl#2bbc6fd2 to nil")
	}
	if err := b.ConsumeID(GetBackgroundURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getBackgroundUrl#2bbc6fd2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetBackgroundURLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBackgroundUrl#2bbc6fd2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getBackgroundUrl#2bbc6fd2: field name: %w", err)
		}
		g.Name = value
	}
	{
		value, err := DecodeBackgroundType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getBackgroundUrl#2bbc6fd2: field type: %w", err)
		}
		g.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetBackgroundURLRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getBackgroundUrl#2bbc6fd2 as nil")
	}
	b.ObjStart()
	b.PutID("getBackgroundUrl")
	b.Comma()
	b.FieldStart("name")
	b.PutString(g.Name)
	b.Comma()
	b.FieldStart("type")
	if g.Type == nil {
		return fmt.Errorf("unable to encode getBackgroundUrl#2bbc6fd2: field type is nil")
	}
	if err := g.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getBackgroundUrl#2bbc6fd2: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetBackgroundURLRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getBackgroundUrl#2bbc6fd2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getBackgroundUrl"); err != nil {
				return fmt.Errorf("unable to decode getBackgroundUrl#2bbc6fd2: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getBackgroundUrl#2bbc6fd2: field name: %w", err)
			}
			g.Name = value
		case "type":
			value, err := DecodeTDLibJSONBackgroundType(b)
			if err != nil {
				return fmt.Errorf("unable to decode getBackgroundUrl#2bbc6fd2: field type: %w", err)
			}
			g.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (g *GetBackgroundURLRequest) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

// GetType returns value of Type field.
func (g *GetBackgroundURLRequest) GetType() (value BackgroundTypeClass) {
	if g == nil {
		return
	}
	return g.Type
}

// GetBackgroundURL invokes method getBackgroundUrl#2bbc6fd2 returning error if any.
func (c *Client) GetBackgroundURL(ctx context.Context, request *GetBackgroundURLRequest) (*HTTPURL, error) {
	var result HTTPURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
