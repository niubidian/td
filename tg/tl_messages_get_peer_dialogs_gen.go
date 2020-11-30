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

// MessagesGetPeerDialogsRequest represents TL type `messages.getPeerDialogs#e470bcfd`.
type MessagesGetPeerDialogsRequest struct {
	// Peers field of MessagesGetPeerDialogsRequest.
	Peers []InputDialogPeerClass
}

// MessagesGetPeerDialogsRequestTypeID is TL type id of MessagesGetPeerDialogsRequest.
const MessagesGetPeerDialogsRequestTypeID = 0xe470bcfd

// Encode implements bin.Encoder.
func (g *MessagesGetPeerDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getPeerDialogs#e470bcfd as nil")
	}
	b.PutID(MessagesGetPeerDialogsRequestTypeID)
	b.PutVectorHeader(len(g.Peers))
	for idx, v := range g.Peers {
		if v == nil {
			return fmt.Errorf("unable to encode messages.getPeerDialogs#e470bcfd: field peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getPeerDialogs#e470bcfd: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetPeerDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getPeerDialogs#e470bcfd to nil")
	}
	if err := b.ConsumeID(MessagesGetPeerDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getPeerDialogs#e470bcfd: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getPeerDialogs#e470bcfd: field peers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getPeerDialogs#e470bcfd: field peers: %w", err)
			}
			g.Peers = append(g.Peers, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetPeerDialogsRequest.
var (
	_ bin.Encoder = &MessagesGetPeerDialogsRequest{}
	_ bin.Decoder = &MessagesGetPeerDialogsRequest{}
)