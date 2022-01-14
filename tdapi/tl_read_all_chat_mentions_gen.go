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

// ReadAllChatMentionsRequest represents TL type `readAllChatMentions#50eab2b5`.
type ReadAllChatMentionsRequest struct {
	// Chat identifier
	ChatID int64
}

// ReadAllChatMentionsRequestTypeID is TL type id of ReadAllChatMentionsRequest.
const ReadAllChatMentionsRequestTypeID = 0x50eab2b5

// Ensuring interfaces in compile-time for ReadAllChatMentionsRequest.
var (
	_ bin.Encoder     = &ReadAllChatMentionsRequest{}
	_ bin.Decoder     = &ReadAllChatMentionsRequest{}
	_ bin.BareEncoder = &ReadAllChatMentionsRequest{}
	_ bin.BareDecoder = &ReadAllChatMentionsRequest{}
)

func (r *ReadAllChatMentionsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReadAllChatMentionsRequest) String() string {
	if r == nil {
		return "ReadAllChatMentionsRequest(nil)"
	}
	type Alias ReadAllChatMentionsRequest
	return fmt.Sprintf("ReadAllChatMentionsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReadAllChatMentionsRequest) TypeID() uint32 {
	return ReadAllChatMentionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReadAllChatMentionsRequest) TypeName() string {
	return "readAllChatMentions"
}

// TypeInfo returns info about TL type.
func (r *ReadAllChatMentionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "readAllChatMentions",
		ID:   ReadAllChatMentionsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReadAllChatMentionsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readAllChatMentions#50eab2b5 as nil")
	}
	b.PutID(ReadAllChatMentionsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReadAllChatMentionsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode readAllChatMentions#50eab2b5 as nil")
	}
	b.PutInt53(r.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReadAllChatMentionsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readAllChatMentions#50eab2b5 to nil")
	}
	if err := b.ConsumeID(ReadAllChatMentionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode readAllChatMentions#50eab2b5: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReadAllChatMentionsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode readAllChatMentions#50eab2b5 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode readAllChatMentions#50eab2b5: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReadAllChatMentionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode readAllChatMentions#50eab2b5 as nil")
	}
	b.ObjStart()
	b.PutID("readAllChatMentions")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(r.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReadAllChatMentionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode readAllChatMentions#50eab2b5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("readAllChatMentions"); err != nil {
				return fmt.Errorf("unable to decode readAllChatMentions#50eab2b5: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode readAllChatMentions#50eab2b5: field chat_id: %w", err)
			}
			r.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (r *ReadAllChatMentionsRequest) GetChatID() (value int64) {
	if r == nil {
		return
	}
	return r.ChatID
}

// ReadAllChatMentions invokes method readAllChatMentions#50eab2b5 returning error if any.
func (c *Client) ReadAllChatMentions(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &ReadAllChatMentionsRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
