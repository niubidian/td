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

// ValidateOrderInfoRequest represents TL type `validateOrderInfo#90a9c4`.
type ValidateOrderInfoRequest struct {
	// Chat identifier of the Invoice message
	ChatID int64
	// Message identifier
	MessageID int64
	// The order information, provided by the user; pass null if empty
	OrderInfo OrderInfo
	// True, if the order information can be saved
	AllowSave bool
}

// ValidateOrderInfoRequestTypeID is TL type id of ValidateOrderInfoRequest.
const ValidateOrderInfoRequestTypeID = 0x90a9c4

// Ensuring interfaces in compile-time for ValidateOrderInfoRequest.
var (
	_ bin.Encoder     = &ValidateOrderInfoRequest{}
	_ bin.Decoder     = &ValidateOrderInfoRequest{}
	_ bin.BareEncoder = &ValidateOrderInfoRequest{}
	_ bin.BareDecoder = &ValidateOrderInfoRequest{}
)

func (v *ValidateOrderInfoRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.ChatID == 0) {
		return false
	}
	if !(v.MessageID == 0) {
		return false
	}
	if !(v.OrderInfo.Zero()) {
		return false
	}
	if !(v.AllowSave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *ValidateOrderInfoRequest) String() string {
	if v == nil {
		return "ValidateOrderInfoRequest(nil)"
	}
	type Alias ValidateOrderInfoRequest
	return fmt.Sprintf("ValidateOrderInfoRequest%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ValidateOrderInfoRequest) TypeID() uint32 {
	return ValidateOrderInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ValidateOrderInfoRequest) TypeName() string {
	return "validateOrderInfo"
}

// TypeInfo returns info about TL type.
func (v *ValidateOrderInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "validateOrderInfo",
		ID:   ValidateOrderInfoRequestTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "OrderInfo",
			SchemaName: "order_info",
		},
		{
			Name:       "AllowSave",
			SchemaName: "allow_save",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *ValidateOrderInfoRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode validateOrderInfo#90a9c4 as nil")
	}
	b.PutID(ValidateOrderInfoRequestTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *ValidateOrderInfoRequest) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode validateOrderInfo#90a9c4 as nil")
	}
	b.PutInt53(v.ChatID)
	b.PutInt53(v.MessageID)
	if err := v.OrderInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode validateOrderInfo#90a9c4: field order_info: %w", err)
	}
	b.PutBool(v.AllowSave)
	return nil
}

// Decode implements bin.Decoder.
func (v *ValidateOrderInfoRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode validateOrderInfo#90a9c4 to nil")
	}
	if err := b.ConsumeID(ValidateOrderInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *ValidateOrderInfoRequest) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode validateOrderInfo#90a9c4 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field chat_id: %w", err)
		}
		v.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field message_id: %w", err)
		}
		v.MessageID = value
	}
	{
		if err := v.OrderInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field order_info: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field allow_save: %w", err)
		}
		v.AllowSave = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *ValidateOrderInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode validateOrderInfo#90a9c4 as nil")
	}
	b.ObjStart()
	b.PutID("validateOrderInfo")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(v.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(v.MessageID)
	b.Comma()
	b.FieldStart("order_info")
	if err := v.OrderInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode validateOrderInfo#90a9c4: field order_info: %w", err)
	}
	b.Comma()
	b.FieldStart("allow_save")
	b.PutBool(v.AllowSave)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *ValidateOrderInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode validateOrderInfo#90a9c4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("validateOrderInfo"); err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field chat_id: %w", err)
			}
			v.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field message_id: %w", err)
			}
			v.MessageID = value
		case "order_info":
			if err := v.OrderInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field order_info: %w", err)
			}
		case "allow_save":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode validateOrderInfo#90a9c4: field allow_save: %w", err)
			}
			v.AllowSave = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (v *ValidateOrderInfoRequest) GetChatID() (value int64) {
	if v == nil {
		return
	}
	return v.ChatID
}

// GetMessageID returns value of MessageID field.
func (v *ValidateOrderInfoRequest) GetMessageID() (value int64) {
	if v == nil {
		return
	}
	return v.MessageID
}

// GetOrderInfo returns value of OrderInfo field.
func (v *ValidateOrderInfoRequest) GetOrderInfo() (value OrderInfo) {
	if v == nil {
		return
	}
	return v.OrderInfo
}

// GetAllowSave returns value of AllowSave field.
func (v *ValidateOrderInfoRequest) GetAllowSave() (value bool) {
	if v == nil {
		return
	}
	return v.AllowSave
}

// ValidateOrderInfo invokes method validateOrderInfo#90a9c4 returning error if any.
func (c *Client) ValidateOrderInfo(ctx context.Context, request *ValidateOrderInfoRequest) (*ValidatedOrderInfo, error) {
	var result ValidatedOrderInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
