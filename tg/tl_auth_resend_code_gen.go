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

// AuthResendCodeRequest represents TL type `auth.resendCode#3ef1a9bf`.
type AuthResendCodeRequest struct {
	// PhoneNumber field of AuthResendCodeRequest.
	PhoneNumber string
	// PhoneCodeHash field of AuthResendCodeRequest.
	PhoneCodeHash string
}

// AuthResendCodeRequestTypeID is TL type id of AuthResendCodeRequest.
const AuthResendCodeRequestTypeID = 0x3ef1a9bf

// Encode implements bin.Encoder.
func (r *AuthResendCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.resendCode#3ef1a9bf as nil")
	}
	b.PutID(AuthResendCodeRequestTypeID)
	b.PutString(r.PhoneNumber)
	b.PutString(r.PhoneCodeHash)
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthResendCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.resendCode#3ef1a9bf to nil")
	}
	if err := b.ConsumeID(AuthResendCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.resendCode#3ef1a9bf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#3ef1a9bf: field phone_number: %w", err)
		}
		r.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#3ef1a9bf: field phone_code_hash: %w", err)
		}
		r.PhoneCodeHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthResendCodeRequest.
var (
	_ bin.Encoder = &AuthResendCodeRequest{}
	_ bin.Decoder = &AuthResendCodeRequest{}
)