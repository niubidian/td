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

// CreateSecretChatRequest represents TL type `createSecretChat#730dd22f`.
type CreateSecretChatRequest struct {
	// Secret chat identifier
	SecretChatID int32
}

// CreateSecretChatRequestTypeID is TL type id of CreateSecretChatRequest.
const CreateSecretChatRequestTypeID = 0x730dd22f

// Ensuring interfaces in compile-time for CreateSecretChatRequest.
var (
	_ bin.Encoder     = &CreateSecretChatRequest{}
	_ bin.Decoder     = &CreateSecretChatRequest{}
	_ bin.BareEncoder = &CreateSecretChatRequest{}
	_ bin.BareDecoder = &CreateSecretChatRequest{}
)

func (c *CreateSecretChatRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.SecretChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateSecretChatRequest) String() string {
	if c == nil {
		return "CreateSecretChatRequest(nil)"
	}
	type Alias CreateSecretChatRequest
	return fmt.Sprintf("CreateSecretChatRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateSecretChatRequest) TypeID() uint32 {
	return CreateSecretChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateSecretChatRequest) TypeName() string {
	return "createSecretChat"
}

// TypeInfo returns info about TL type.
func (c *CreateSecretChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createSecretChat",
		ID:   CreateSecretChatRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SecretChatID",
			SchemaName: "secret_chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateSecretChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createSecretChat#730dd22f as nil")
	}
	b.PutID(CreateSecretChatRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateSecretChatRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createSecretChat#730dd22f as nil")
	}
	b.PutInt32(c.SecretChatID)
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateSecretChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createSecretChat#730dd22f to nil")
	}
	if err := b.ConsumeID(CreateSecretChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createSecretChat#730dd22f: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateSecretChatRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createSecretChat#730dd22f to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode createSecretChat#730dd22f: field secret_chat_id: %w", err)
		}
		c.SecretChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreateSecretChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createSecretChat#730dd22f as nil")
	}
	b.ObjStart()
	b.PutID("createSecretChat")
	b.FieldStart("secret_chat_id")
	b.PutInt32(c.SecretChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreateSecretChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createSecretChat#730dd22f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createSecretChat"); err != nil {
				return fmt.Errorf("unable to decode createSecretChat#730dd22f: %w", err)
			}
		case "secret_chat_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode createSecretChat#730dd22f: field secret_chat_id: %w", err)
			}
			c.SecretChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSecretChatID returns value of SecretChatID field.
func (c *CreateSecretChatRequest) GetSecretChatID() (value int32) {
	return c.SecretChatID
}

// CreateSecretChat invokes method createSecretChat#730dd22f returning error if any.
func (c *Client) CreateSecretChat(ctx context.Context, secretchatid int32) (*Chat, error) {
	var result Chat

	request := &CreateSecretChatRequest{
		SecretChatID: secretchatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}