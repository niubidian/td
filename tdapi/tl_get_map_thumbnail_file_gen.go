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

// GetMapThumbnailFileRequest represents TL type `getMapThumbnailFile#f6e6979a`.
type GetMapThumbnailFileRequest struct {
	// Location of the map center
	Location Location
	// Map zoom level; 13-20
	Zoom int32
	// Map width in pixels before applying scale; 16-1024
	Width int32
	// Map height in pixels before applying scale; 16-1024
	Height int32
	// Map scale; 1-3
	Scale int32
	// Identifier of a chat, in which the thumbnail will be shown. Use 0 if unknown
	ChatID int64
}

// GetMapThumbnailFileRequestTypeID is TL type id of GetMapThumbnailFileRequest.
const GetMapThumbnailFileRequestTypeID = 0xf6e6979a

// Ensuring interfaces in compile-time for GetMapThumbnailFileRequest.
var (
	_ bin.Encoder     = &GetMapThumbnailFileRequest{}
	_ bin.Decoder     = &GetMapThumbnailFileRequest{}
	_ bin.BareEncoder = &GetMapThumbnailFileRequest{}
	_ bin.BareDecoder = &GetMapThumbnailFileRequest{}
)

func (g *GetMapThumbnailFileRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Location.Zero()) {
		return false
	}
	if !(g.Zoom == 0) {
		return false
	}
	if !(g.Width == 0) {
		return false
	}
	if !(g.Height == 0) {
		return false
	}
	if !(g.Scale == 0) {
		return false
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMapThumbnailFileRequest) String() string {
	if g == nil {
		return "GetMapThumbnailFileRequest(nil)"
	}
	type Alias GetMapThumbnailFileRequest
	return fmt.Sprintf("GetMapThumbnailFileRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMapThumbnailFileRequest) TypeID() uint32 {
	return GetMapThumbnailFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMapThumbnailFileRequest) TypeName() string {
	return "getMapThumbnailFile"
}

// TypeInfo returns info about TL type.
func (g *GetMapThumbnailFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMapThumbnailFile",
		ID:   GetMapThumbnailFileRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Zoom",
			SchemaName: "zoom",
		},
		{
			Name:       "Width",
			SchemaName: "width",
		},
		{
			Name:       "Height",
			SchemaName: "height",
		},
		{
			Name:       "Scale",
			SchemaName: "scale",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMapThumbnailFileRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMapThumbnailFile#f6e6979a as nil")
	}
	b.PutID(GetMapThumbnailFileRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMapThumbnailFileRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMapThumbnailFile#f6e6979a as nil")
	}
	if err := g.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getMapThumbnailFile#f6e6979a: field location: %w", err)
	}
	b.PutInt32(g.Zoom)
	b.PutInt32(g.Width)
	b.PutInt32(g.Height)
	b.PutInt32(g.Scale)
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMapThumbnailFileRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMapThumbnailFile#f6e6979a to nil")
	}
	if err := b.ConsumeID(GetMapThumbnailFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMapThumbnailFileRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMapThumbnailFile#f6e6979a to nil")
	}
	{
		if err := g.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field location: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field zoom: %w", err)
		}
		g.Zoom = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field width: %w", err)
		}
		g.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field height: %w", err)
		}
		g.Height = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field scale: %w", err)
		}
		g.Scale = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMapThumbnailFileRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMapThumbnailFile#f6e6979a as nil")
	}
	b.ObjStart()
	b.PutID("getMapThumbnailFile")
	b.Comma()
	b.FieldStart("location")
	if err := g.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getMapThumbnailFile#f6e6979a: field location: %w", err)
	}
	b.Comma()
	b.FieldStart("zoom")
	b.PutInt32(g.Zoom)
	b.Comma()
	b.FieldStart("width")
	b.PutInt32(g.Width)
	b.Comma()
	b.FieldStart("height")
	b.PutInt32(g.Height)
	b.Comma()
	b.FieldStart("scale")
	b.PutInt32(g.Scale)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMapThumbnailFileRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMapThumbnailFile#f6e6979a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMapThumbnailFile"); err != nil {
				return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: %w", err)
			}
		case "location":
			if err := g.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field location: %w", err)
			}
		case "zoom":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field zoom: %w", err)
			}
			g.Zoom = value
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field width: %w", err)
			}
			g.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field height: %w", err)
			}
			g.Height = value
		case "scale":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field scale: %w", err)
			}
			g.Scale = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMapThumbnailFile#f6e6979a: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLocation returns value of Location field.
func (g *GetMapThumbnailFileRequest) GetLocation() (value Location) {
	if g == nil {
		return
	}
	return g.Location
}

// GetZoom returns value of Zoom field.
func (g *GetMapThumbnailFileRequest) GetZoom() (value int32) {
	if g == nil {
		return
	}
	return g.Zoom
}

// GetWidth returns value of Width field.
func (g *GetMapThumbnailFileRequest) GetWidth() (value int32) {
	if g == nil {
		return
	}
	return g.Width
}

// GetHeight returns value of Height field.
func (g *GetMapThumbnailFileRequest) GetHeight() (value int32) {
	if g == nil {
		return
	}
	return g.Height
}

// GetScale returns value of Scale field.
func (g *GetMapThumbnailFileRequest) GetScale() (value int32) {
	if g == nil {
		return
	}
	return g.Scale
}

// GetChatID returns value of ChatID field.
func (g *GetMapThumbnailFileRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMapThumbnailFile invokes method getMapThumbnailFile#f6e6979a returning error if any.
func (c *Client) GetMapThumbnailFile(ctx context.Context, request *GetMapThumbnailFileRequest) (*File, error) {
	var result File

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
