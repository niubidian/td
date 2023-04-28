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

// ProcessChatFolderNewChatsRequest represents TL type `processChatFolderNewChats#109f8a8b`.
type ProcessChatFolderNewChatsRequest struct {
	// Chat folder identifier
	ChatFolderID int32
	// Identifiers of the new chats, which are added to the chat folder. The chats are
	// automatically joined if they aren't joined yet
	AddedChatIDs []int64
}

// ProcessChatFolderNewChatsRequestTypeID is TL type id of ProcessChatFolderNewChatsRequest.
const ProcessChatFolderNewChatsRequestTypeID = 0x109f8a8b

// Ensuring interfaces in compile-time for ProcessChatFolderNewChatsRequest.
var (
	_ bin.Encoder     = &ProcessChatFolderNewChatsRequest{}
	_ bin.Decoder     = &ProcessChatFolderNewChatsRequest{}
	_ bin.BareEncoder = &ProcessChatFolderNewChatsRequest{}
	_ bin.BareDecoder = &ProcessChatFolderNewChatsRequest{}
)

func (p *ProcessChatFolderNewChatsRequest) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ChatFolderID == 0) {
		return false
	}
	if !(p.AddedChatIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ProcessChatFolderNewChatsRequest) String() string {
	if p == nil {
		return "ProcessChatFolderNewChatsRequest(nil)"
	}
	type Alias ProcessChatFolderNewChatsRequest
	return fmt.Sprintf("ProcessChatFolderNewChatsRequest%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ProcessChatFolderNewChatsRequest) TypeID() uint32 {
	return ProcessChatFolderNewChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ProcessChatFolderNewChatsRequest) TypeName() string {
	return "processChatFolderNewChats"
}

// TypeInfo returns info about TL type.
func (p *ProcessChatFolderNewChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "processChatFolderNewChats",
		ID:   ProcessChatFolderNewChatsRequestTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatFolderID",
			SchemaName: "chat_folder_id",
		},
		{
			Name:       "AddedChatIDs",
			SchemaName: "added_chat_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ProcessChatFolderNewChatsRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode processChatFolderNewChats#109f8a8b as nil")
	}
	b.PutID(ProcessChatFolderNewChatsRequestTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ProcessChatFolderNewChatsRequest) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode processChatFolderNewChats#109f8a8b as nil")
	}
	b.PutInt32(p.ChatFolderID)
	b.PutInt(len(p.AddedChatIDs))
	for _, v := range p.AddedChatIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *ProcessChatFolderNewChatsRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode processChatFolderNewChats#109f8a8b to nil")
	}
	if err := b.ConsumeID(ProcessChatFolderNewChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ProcessChatFolderNewChatsRequest) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode processChatFolderNewChats#109f8a8b to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: field chat_folder_id: %w", err)
		}
		p.ChatFolderID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: field added_chat_ids: %w", err)
		}

		if headerLen > 0 {
			p.AddedChatIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: field added_chat_ids: %w", err)
			}
			p.AddedChatIDs = append(p.AddedChatIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ProcessChatFolderNewChatsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode processChatFolderNewChats#109f8a8b as nil")
	}
	b.ObjStart()
	b.PutID("processChatFolderNewChats")
	b.Comma()
	b.FieldStart("chat_folder_id")
	b.PutInt32(p.ChatFolderID)
	b.Comma()
	b.FieldStart("added_chat_ids")
	b.ArrStart()
	for _, v := range p.AddedChatIDs {
		b.PutInt53(v)
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
func (p *ProcessChatFolderNewChatsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode processChatFolderNewChats#109f8a8b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("processChatFolderNewChats"); err != nil {
				return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: %w", err)
			}
		case "chat_folder_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: field chat_folder_id: %w", err)
			}
			p.ChatFolderID = value
		case "added_chat_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: field added_chat_ids: %w", err)
				}
				p.AddedChatIDs = append(p.AddedChatIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode processChatFolderNewChats#109f8a8b: field added_chat_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatFolderID returns value of ChatFolderID field.
func (p *ProcessChatFolderNewChatsRequest) GetChatFolderID() (value int32) {
	if p == nil {
		return
	}
	return p.ChatFolderID
}

// GetAddedChatIDs returns value of AddedChatIDs field.
func (p *ProcessChatFolderNewChatsRequest) GetAddedChatIDs() (value []int64) {
	if p == nil {
		return
	}
	return p.AddedChatIDs
}

// ProcessChatFolderNewChats invokes method processChatFolderNewChats#109f8a8b returning error if any.
func (c *Client) ProcessChatFolderNewChats(ctx context.Context, request *ProcessChatFolderNewChatsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}