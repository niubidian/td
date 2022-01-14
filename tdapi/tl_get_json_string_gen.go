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

// GetJSONStringRequest represents TL type `getJsonString#278b9421`.
type GetJSONStringRequest struct {
	// The JsonValue object
	JSONValue JSONValueClass
}

// GetJSONStringRequestTypeID is TL type id of GetJSONStringRequest.
const GetJSONStringRequestTypeID = 0x278b9421

// Ensuring interfaces in compile-time for GetJSONStringRequest.
var (
	_ bin.Encoder     = &GetJSONStringRequest{}
	_ bin.Decoder     = &GetJSONStringRequest{}
	_ bin.BareEncoder = &GetJSONStringRequest{}
	_ bin.BareDecoder = &GetJSONStringRequest{}
)

func (g *GetJSONStringRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.JSONValue == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetJSONStringRequest) String() string {
	if g == nil {
		return "GetJSONStringRequest(nil)"
	}
	type Alias GetJSONStringRequest
	return fmt.Sprintf("GetJSONStringRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetJSONStringRequest) TypeID() uint32 {
	return GetJSONStringRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetJSONStringRequest) TypeName() string {
	return "getJsonString"
}

// TypeInfo returns info about TL type.
func (g *GetJSONStringRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getJsonString",
		ID:   GetJSONStringRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "JSONValue",
			SchemaName: "json_value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetJSONStringRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getJsonString#278b9421 as nil")
	}
	b.PutID(GetJSONStringRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetJSONStringRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getJsonString#278b9421 as nil")
	}
	if g.JSONValue == nil {
		return fmt.Errorf("unable to encode getJsonString#278b9421: field json_value is nil")
	}
	if err := g.JSONValue.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getJsonString#278b9421: field json_value: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetJSONStringRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getJsonString#278b9421 to nil")
	}
	if err := b.ConsumeID(GetJSONStringRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getJsonString#278b9421: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetJSONStringRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getJsonString#278b9421 to nil")
	}
	{
		value, err := DecodeJSONValue(b)
		if err != nil {
			return fmt.Errorf("unable to decode getJsonString#278b9421: field json_value: %w", err)
		}
		g.JSONValue = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetJSONStringRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getJsonString#278b9421 as nil")
	}
	b.ObjStart()
	b.PutID("getJsonString")
	b.Comma()
	b.FieldStart("json_value")
	if g.JSONValue == nil {
		return fmt.Errorf("unable to encode getJsonString#278b9421: field json_value is nil")
	}
	if err := g.JSONValue.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getJsonString#278b9421: field json_value: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetJSONStringRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getJsonString#278b9421 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getJsonString"); err != nil {
				return fmt.Errorf("unable to decode getJsonString#278b9421: %w", err)
			}
		case "json_value":
			value, err := DecodeTDLibJSONJSONValue(b)
			if err != nil {
				return fmt.Errorf("unable to decode getJsonString#278b9421: field json_value: %w", err)
			}
			g.JSONValue = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetJSONValue returns value of JSONValue field.
func (g *GetJSONStringRequest) GetJSONValue() (value JSONValueClass) {
	if g == nil {
		return
	}
	return g.JSONValue
}

// GetJSONString invokes method getJsonString#278b9421 returning error if any.
func (c *Client) GetJSONString(ctx context.Context, jsonvalue JSONValueClass) (*Text, error) {
	var result Text

	request := &GetJSONStringRequest{
		JSONValue: jsonvalue,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
