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

// MessagesGetSavedGifsRequest represents TL type `messages.getSavedGifs#83bf3d52`.
type MessagesGetSavedGifsRequest struct {
	// Hash field of MessagesGetSavedGifsRequest.
	Hash int
}

// MessagesGetSavedGifsRequestTypeID is TL type id of MessagesGetSavedGifsRequest.
const MessagesGetSavedGifsRequestTypeID = 0x83bf3d52

// Encode implements bin.Encoder.
func (g *MessagesGetSavedGifsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSavedGifs#83bf3d52 as nil")
	}
	b.PutID(MessagesGetSavedGifsRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetSavedGifsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSavedGifs#83bf3d52 to nil")
	}
	if err := b.ConsumeID(MessagesGetSavedGifsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSavedGifs#83bf3d52: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedGifs#83bf3d52: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetSavedGifsRequest.
var (
	_ bin.Encoder = &MessagesGetSavedGifsRequest{}
	_ bin.Decoder = &MessagesGetSavedGifsRequest{}
)