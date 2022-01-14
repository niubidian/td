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

// SendBotStartMessageRequest represents TL type `sendBotStartMessage#aa6a3eee`.
type SendBotStartMessageRequest struct {
	// Identifier of the bot
	BotUserID int64
	// Identifier of the target chat
	ChatID int64
	// A hidden parameter sent to the bot for deep linking purposes (https://core.telegram
	// org/bots#deep-linking)
	Parameter string
}

// SendBotStartMessageRequestTypeID is TL type id of SendBotStartMessageRequest.
const SendBotStartMessageRequestTypeID = 0xaa6a3eee

// Ensuring interfaces in compile-time for SendBotStartMessageRequest.
var (
	_ bin.Encoder     = &SendBotStartMessageRequest{}
	_ bin.Decoder     = &SendBotStartMessageRequest{}
	_ bin.BareEncoder = &SendBotStartMessageRequest{}
	_ bin.BareDecoder = &SendBotStartMessageRequest{}
)

func (s *SendBotStartMessageRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BotUserID == 0) {
		return false
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Parameter == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendBotStartMessageRequest) String() string {
	if s == nil {
		return "SendBotStartMessageRequest(nil)"
	}
	type Alias SendBotStartMessageRequest
	return fmt.Sprintf("SendBotStartMessageRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendBotStartMessageRequest) TypeID() uint32 {
	return SendBotStartMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendBotStartMessageRequest) TypeName() string {
	return "sendBotStartMessage"
}

// TypeInfo returns info about TL type.
func (s *SendBotStartMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendBotStartMessage",
		ID:   SendBotStartMessageRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Parameter",
			SchemaName: "parameter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendBotStartMessageRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendBotStartMessage#aa6a3eee as nil")
	}
	b.PutID(SendBotStartMessageRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendBotStartMessageRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendBotStartMessage#aa6a3eee as nil")
	}
	b.PutInt53(s.BotUserID)
	b.PutInt53(s.ChatID)
	b.PutString(s.Parameter)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendBotStartMessageRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendBotStartMessage#aa6a3eee to nil")
	}
	if err := b.ConsumeID(SendBotStartMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendBotStartMessageRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendBotStartMessage#aa6a3eee to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: field bot_user_id: %w", err)
		}
		s.BotUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: field parameter: %w", err)
		}
		s.Parameter = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendBotStartMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendBotStartMessage#aa6a3eee as nil")
	}
	b.ObjStart()
	b.PutID("sendBotStartMessage")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(s.BotUserID)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("parameter")
	b.PutString(s.Parameter)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendBotStartMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendBotStartMessage#aa6a3eee to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendBotStartMessage"); err != nil {
				return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: field bot_user_id: %w", err)
			}
			s.BotUserID = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: field chat_id: %w", err)
			}
			s.ChatID = value
		case "parameter":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendBotStartMessage#aa6a3eee: field parameter: %w", err)
			}
			s.Parameter = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (s *SendBotStartMessageRequest) GetBotUserID() (value int64) {
	if s == nil {
		return
	}
	return s.BotUserID
}

// GetChatID returns value of ChatID field.
func (s *SendBotStartMessageRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetParameter returns value of Parameter field.
func (s *SendBotStartMessageRequest) GetParameter() (value string) {
	if s == nil {
		return
	}
	return s.Parameter
}

// SendBotStartMessage invokes method sendBotStartMessage#aa6a3eee returning error if any.
func (c *Client) SendBotStartMessage(ctx context.Context, request *SendBotStartMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
