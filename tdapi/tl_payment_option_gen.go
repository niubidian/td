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

// PaymentOption represents TL type `paymentOption#ee79989b`.
type PaymentOption struct {
	// Title for the payment option
	Title string
	// Payment form URL to be opened in a web view
	URL string
}

// PaymentOptionTypeID is TL type id of PaymentOption.
const PaymentOptionTypeID = 0xee79989b

// Ensuring interfaces in compile-time for PaymentOption.
var (
	_ bin.Encoder     = &PaymentOption{}
	_ bin.Decoder     = &PaymentOption{}
	_ bin.BareEncoder = &PaymentOption{}
	_ bin.BareDecoder = &PaymentOption{}
)

func (p *PaymentOption) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Title == "") {
		return false
	}
	if !(p.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentOption) String() string {
	if p == nil {
		return "PaymentOption(nil)"
	}
	type Alias PaymentOption
	return fmt.Sprintf("PaymentOption%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentOption) TypeID() uint32 {
	return PaymentOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentOption) TypeName() string {
	return "paymentOption"
}

// TypeInfo returns info about TL type.
func (p *PaymentOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentOption",
		ID:   PaymentOptionTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentOption) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentOption#ee79989b as nil")
	}
	b.PutID(PaymentOptionTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentOption) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentOption#ee79989b as nil")
	}
	b.PutString(p.Title)
	b.PutString(p.URL)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentOption) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentOption#ee79989b to nil")
	}
	if err := b.ConsumeID(PaymentOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentOption#ee79989b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentOption) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentOption#ee79989b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentOption#ee79989b: field title: %w", err)
		}
		p.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentOption#ee79989b: field url: %w", err)
		}
		p.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentOption) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentOption#ee79989b as nil")
	}
	b.ObjStart()
	b.PutID("paymentOption")
	b.Comma()
	b.FieldStart("title")
	b.PutString(p.Title)
	b.Comma()
	b.FieldStart("url")
	b.PutString(p.URL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentOption) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentOption#ee79989b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentOption"); err != nil {
				return fmt.Errorf("unable to decode paymentOption#ee79989b: %w", err)
			}
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentOption#ee79989b: field title: %w", err)
			}
			p.Title = value
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentOption#ee79989b: field url: %w", err)
			}
			p.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTitle returns value of Title field.
func (p *PaymentOption) GetTitle() (value string) {
	if p == nil {
		return
	}
	return p.Title
}

// GetURL returns value of URL field.
func (p *PaymentOption) GetURL() (value string) {
	if p == nil {
		return
	}
	return p.URL
}
