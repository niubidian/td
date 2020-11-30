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

// MessagesGetAllChatsRequest represents TL type `messages.getAllChats#eba80ff0`.
type MessagesGetAllChatsRequest struct {
	// ExceptIds field of MessagesGetAllChatsRequest.
	ExceptIds []int
}

// MessagesGetAllChatsRequestTypeID is TL type id of MessagesGetAllChatsRequest.
const MessagesGetAllChatsRequestTypeID = 0xeba80ff0

// Encode implements bin.Encoder.
func (g *MessagesGetAllChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAllChats#eba80ff0 as nil")
	}
	b.PutID(MessagesGetAllChatsRequestTypeID)
	b.PutVectorHeader(len(g.ExceptIds))
	for _, v := range g.ExceptIds {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAllChats#eba80ff0 to nil")
	}
	if err := b.ConsumeID(MessagesGetAllChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAllChats#eba80ff0: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAllChats#eba80ff0: field except_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getAllChats#eba80ff0: field except_ids: %w", err)
			}
			g.ExceptIds = append(g.ExceptIds, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllChatsRequest.
var (
	_ bin.Encoder = &MessagesGetAllChatsRequest{}
	_ bin.Decoder = &MessagesGetAllChatsRequest{}
)