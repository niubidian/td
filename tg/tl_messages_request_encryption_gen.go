// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// MessagesRequestEncryptionRequest represents TL type `messages.requestEncryption#f64daf43`.
type MessagesRequestEncryptionRequest struct {
	// UserID field of MessagesRequestEncryptionRequest.
	UserID InputUserClass
	// RandomID field of MessagesRequestEncryptionRequest.
	RandomID int
	// GA field of MessagesRequestEncryptionRequest.
	GA []byte
}

// MessagesRequestEncryptionRequestTypeID is TL type id of MessagesRequestEncryptionRequest.
const MessagesRequestEncryptionRequestTypeID = 0xf64daf43

// Encode implements bin.Encoder.
func (r *MessagesRequestEncryptionRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestEncryption#f64daf43 as nil")
	}
	b.PutID(MessagesRequestEncryptionRequestTypeID)
	if r.UserID == nil {
		return fmt.Errorf("unable to encode messages.requestEncryption#f64daf43: field user_id is nil")
	}
	if err := r.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestEncryption#f64daf43: field user_id: %w", err)
	}
	b.PutInt(r.RandomID)
	b.PutBytes(r.GA)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestEncryptionRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestEncryption#f64daf43 to nil")
	}
	if err := b.ConsumeID(MessagesRequestEncryptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: field random_id: %w", err)
		}
		r.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: field g_a: %w", err)
		}
		r.GA = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesRequestEncryptionRequest.
var (
	_ bin.Encoder = &MessagesRequestEncryptionRequest{}
	_ bin.Decoder = &MessagesRequestEncryptionRequest{}
)