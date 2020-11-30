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

// FoldersEditPeerFoldersRequest represents TL type `folders.editPeerFolders#6847d0ab`.
type FoldersEditPeerFoldersRequest struct {
	// FolderPeers field of FoldersEditPeerFoldersRequest.
	FolderPeers []InputFolderPeer
}

// FoldersEditPeerFoldersRequestTypeID is TL type id of FoldersEditPeerFoldersRequest.
const FoldersEditPeerFoldersRequestTypeID = 0x6847d0ab

// Encode implements bin.Encoder.
func (e *FoldersEditPeerFoldersRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode folders.editPeerFolders#6847d0ab as nil")
	}
	b.PutID(FoldersEditPeerFoldersRequestTypeID)
	b.PutVectorHeader(len(e.FolderPeers))
	for idx, v := range e.FolderPeers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode folders.editPeerFolders#6847d0ab: field folder_peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *FoldersEditPeerFoldersRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode folders.editPeerFolders#6847d0ab to nil")
	}
	if err := b.ConsumeID(FoldersEditPeerFoldersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode folders.editPeerFolders#6847d0ab: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode folders.editPeerFolders#6847d0ab: field folder_peers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputFolderPeer
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode folders.editPeerFolders#6847d0ab: field folder_peers: %w", err)
			}
			e.FolderPeers = append(e.FolderPeers, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for FoldersEditPeerFoldersRequest.
var (
	_ bin.Encoder = &FoldersEditPeerFoldersRequest{}
	_ bin.Decoder = &FoldersEditPeerFoldersRequest{}
)