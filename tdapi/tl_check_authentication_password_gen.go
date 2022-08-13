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

// CheckAuthenticationPasswordRequest represents TL type `checkAuthenticationPassword#87424ba0`.
type CheckAuthenticationPasswordRequest struct {
	// The 2-step verification password to check
	Password string
}

// CheckAuthenticationPasswordRequestTypeID is TL type id of CheckAuthenticationPasswordRequest.
const CheckAuthenticationPasswordRequestTypeID = 0x87424ba0

// Ensuring interfaces in compile-time for CheckAuthenticationPasswordRequest.
var (
	_ bin.Encoder     = &CheckAuthenticationPasswordRequest{}
	_ bin.Decoder     = &CheckAuthenticationPasswordRequest{}
	_ bin.BareEncoder = &CheckAuthenticationPasswordRequest{}
	_ bin.BareDecoder = &CheckAuthenticationPasswordRequest{}
)

func (c *CheckAuthenticationPasswordRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckAuthenticationPasswordRequest) String() string {
	if c == nil {
		return "CheckAuthenticationPasswordRequest(nil)"
	}
	type Alias CheckAuthenticationPasswordRequest
	return fmt.Sprintf("CheckAuthenticationPasswordRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckAuthenticationPasswordRequest) TypeID() uint32 {
	return CheckAuthenticationPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckAuthenticationPasswordRequest) TypeName() string {
	return "checkAuthenticationPassword"
}

// TypeInfo returns info about TL type.
func (c *CheckAuthenticationPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkAuthenticationPassword",
		ID:   CheckAuthenticationPasswordRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckAuthenticationPasswordRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationPassword#87424ba0 as nil")
	}
	b.PutID(CheckAuthenticationPasswordRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckAuthenticationPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationPassword#87424ba0 as nil")
	}
	b.PutString(c.Password)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckAuthenticationPasswordRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationPassword#87424ba0 to nil")
	}
	if err := b.ConsumeID(CheckAuthenticationPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkAuthenticationPassword#87424ba0: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckAuthenticationPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationPassword#87424ba0 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkAuthenticationPassword#87424ba0: field password: %w", err)
		}
		c.Password = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckAuthenticationPasswordRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationPassword#87424ba0 as nil")
	}
	b.ObjStart()
	b.PutID("checkAuthenticationPassword")
	b.Comma()
	b.FieldStart("password")
	b.PutString(c.Password)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckAuthenticationPasswordRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationPassword#87424ba0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkAuthenticationPassword"); err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationPassword#87424ba0: %w", err)
			}
		case "password":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationPassword#87424ba0: field password: %w", err)
			}
			c.Password = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPassword returns value of Password field.
func (c *CheckAuthenticationPasswordRequest) GetPassword() (value string) {
	if c == nil {
		return
	}
	return c.Password
}

// CheckAuthenticationPassword invokes method checkAuthenticationPassword#87424ba0 returning error if any.
func (c *Client) CheckAuthenticationPassword(ctx context.Context, password string) error {
	var ok Ok

	request := &CheckAuthenticationPasswordRequest{
		Password: password,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
