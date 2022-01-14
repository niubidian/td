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

// PageBlockListItem represents TL type `pageBlockListItem#5f521776`.
type PageBlockListItem struct {
	// Item label
	Label string
	// Item blocks
	PageBlocks []PageBlockClass
}

// PageBlockListItemTypeID is TL type id of PageBlockListItem.
const PageBlockListItemTypeID = 0x5f521776

// Ensuring interfaces in compile-time for PageBlockListItem.
var (
	_ bin.Encoder     = &PageBlockListItem{}
	_ bin.Decoder     = &PageBlockListItem{}
	_ bin.BareEncoder = &PageBlockListItem{}
	_ bin.BareDecoder = &PageBlockListItem{}
)

func (p *PageBlockListItem) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Label == "") {
		return false
	}
	if !(p.PageBlocks == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageBlockListItem) String() string {
	if p == nil {
		return "PageBlockListItem(nil)"
	}
	type Alias PageBlockListItem
	return fmt.Sprintf("PageBlockListItem%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageBlockListItem) TypeID() uint32 {
	return PageBlockListItemTypeID
}

// TypeName returns name of type in TL schema.
func (*PageBlockListItem) TypeName() string {
	return "pageBlockListItem"
}

// TypeInfo returns info about TL type.
func (p *PageBlockListItem) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageBlockListItem",
		ID:   PageBlockListItemTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Label",
			SchemaName: "label",
		},
		{
			Name:       "PageBlocks",
			SchemaName: "page_blocks",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageBlockListItem) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockListItem#5f521776 as nil")
	}
	b.PutID(PageBlockListItemTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageBlockListItem) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockListItem#5f521776 as nil")
	}
	b.PutString(p.Label)
	b.PutInt(len(p.PageBlocks))
	for idx, v := range p.PageBlocks {
		if v == nil {
			return fmt.Errorf("unable to encode pageBlockListItem#5f521776: field page_blocks element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare pageBlockListItem#5f521776: field page_blocks element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageBlockListItem) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockListItem#5f521776 to nil")
	}
	if err := b.ConsumeID(PageBlockListItemTypeID); err != nil {
		return fmt.Errorf("unable to decode pageBlockListItem#5f521776: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageBlockListItem) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockListItem#5f521776 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockListItem#5f521776: field label: %w", err)
		}
		p.Label = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockListItem#5f521776: field page_blocks: %w", err)
		}

		if headerLen > 0 {
			p.PageBlocks = make([]PageBlockClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePageBlock(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockListItem#5f521776: field page_blocks: %w", err)
			}
			p.PageBlocks = append(p.PageBlocks, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PageBlockListItem) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockListItem#5f521776 as nil")
	}
	b.ObjStart()
	b.PutID("pageBlockListItem")
	b.Comma()
	b.FieldStart("label")
	b.PutString(p.Label)
	b.Comma()
	b.FieldStart("page_blocks")
	b.ArrStart()
	for idx, v := range p.PageBlocks {
		if v == nil {
			return fmt.Errorf("unable to encode pageBlockListItem#5f521776: field page_blocks element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode pageBlockListItem#5f521776: field page_blocks element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PageBlockListItem) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockListItem#5f521776 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("pageBlockListItem"); err != nil {
				return fmt.Errorf("unable to decode pageBlockListItem#5f521776: %w", err)
			}
		case "label":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockListItem#5f521776: field label: %w", err)
			}
			p.Label = value
		case "page_blocks":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONPageBlock(b)
				if err != nil {
					return fmt.Errorf("unable to decode pageBlockListItem#5f521776: field page_blocks: %w", err)
				}
				p.PageBlocks = append(p.PageBlocks, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode pageBlockListItem#5f521776: field page_blocks: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLabel returns value of Label field.
func (p *PageBlockListItem) GetLabel() (value string) {
	if p == nil {
		return
	}
	return p.Label
}

// GetPageBlocks returns value of PageBlocks field.
func (p *PageBlockListItem) GetPageBlocks() (value []PageBlockClass) {
	if p == nil {
		return
	}
	return p.PageBlocks
}
