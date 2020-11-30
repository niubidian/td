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

// AccountInitTakeoutSessionRequest represents TL type `account.initTakeoutSession#f05b4804`.
type AccountInitTakeoutSessionRequest struct {
	// Flags field of AccountInitTakeoutSessionRequest.
	Flags bin.Fields
	// Contacts field of AccountInitTakeoutSessionRequest.
	Contacts bool
	// MessageUsers field of AccountInitTakeoutSessionRequest.
	MessageUsers bool
	// MessageChats field of AccountInitTakeoutSessionRequest.
	MessageChats bool
	// MessageMegagroups field of AccountInitTakeoutSessionRequest.
	MessageMegagroups bool
	// MessageChannels field of AccountInitTakeoutSessionRequest.
	MessageChannels bool
	// Files field of AccountInitTakeoutSessionRequest.
	Files bool
	// FileMaxSize field of AccountInitTakeoutSessionRequest.
	//
	// Use SetFileMaxSize and GetFileMaxSize helpers.
	FileMaxSize int
}

// AccountInitTakeoutSessionRequestTypeID is TL type id of AccountInitTakeoutSessionRequest.
const AccountInitTakeoutSessionRequestTypeID = 0xf05b4804

// Encode implements bin.Encoder.
func (i *AccountInitTakeoutSessionRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.initTakeoutSession#f05b4804 as nil")
	}
	b.PutID(AccountInitTakeoutSessionRequestTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.initTakeoutSession#f05b4804: field flags: %w", err)
	}
	if i.Flags.Has(5) {
		b.PutInt(i.FileMaxSize)
	}
	return nil
}

// SetContacts sets value of Contacts conditional field.
func (i *AccountInitTakeoutSessionRequest) SetContacts(value bool) {
	if value {
		i.Flags.Set(0)
	} else {
		i.Flags.Unset(0)
	}
}

// SetMessageUsers sets value of MessageUsers conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageUsers(value bool) {
	if value {
		i.Flags.Set(1)
	} else {
		i.Flags.Unset(1)
	}
}

// SetMessageChats sets value of MessageChats conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageChats(value bool) {
	if value {
		i.Flags.Set(2)
	} else {
		i.Flags.Unset(2)
	}
}

// SetMessageMegagroups sets value of MessageMegagroups conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageMegagroups(value bool) {
	if value {
		i.Flags.Set(3)
	} else {
		i.Flags.Unset(3)
	}
}

// SetMessageChannels sets value of MessageChannels conditional field.
func (i *AccountInitTakeoutSessionRequest) SetMessageChannels(value bool) {
	if value {
		i.Flags.Set(4)
	} else {
		i.Flags.Unset(4)
	}
}

// SetFiles sets value of Files conditional field.
func (i *AccountInitTakeoutSessionRequest) SetFiles(value bool) {
	if value {
		i.Flags.Set(5)
	} else {
		i.Flags.Unset(5)
	}
}

// SetFileMaxSize sets value of FileMaxSize conditional field.
func (i *AccountInitTakeoutSessionRequest) SetFileMaxSize(value int) {
	i.Flags.Set(5)
	i.FileMaxSize = value
}

// GetFileMaxSize returns value of FileMaxSize conditional field and
// boolean which is true if field was set.
func (i *AccountInitTakeoutSessionRequest) GetFileMaxSize() (value int, ok bool) {
	if !i.Flags.Has(5) {
		return value, false
	}
	return i.FileMaxSize, true
}

// Decode implements bin.Decoder.
func (i *AccountInitTakeoutSessionRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.initTakeoutSession#f05b4804 to nil")
	}
	if err := b.ConsumeID(AccountInitTakeoutSessionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.initTakeoutSession#f05b4804: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.initTakeoutSession#f05b4804: field flags: %w", err)
		}
	}
	i.Contacts = i.Flags.Has(0)
	i.MessageUsers = i.Flags.Has(1)
	i.MessageChats = i.Flags.Has(2)
	i.MessageMegagroups = i.Flags.Has(3)
	i.MessageChannels = i.Flags.Has(4)
	i.Files = i.Flags.Has(5)
	if i.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.initTakeoutSession#f05b4804: field file_max_size: %w", err)
		}
		i.FileMaxSize = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountInitTakeoutSessionRequest.
var (
	_ bin.Encoder = &AccountInitTakeoutSessionRequest{}
	_ bin.Decoder = &AccountInitTakeoutSessionRequest{}
)