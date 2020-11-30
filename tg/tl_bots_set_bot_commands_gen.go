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

// BotsSetBotCommandsRequest represents TL type `bots.setBotCommands#805d46f6`.
type BotsSetBotCommandsRequest struct {
	// Commands field of BotsSetBotCommandsRequest.
	Commands []BotCommand
}

// BotsSetBotCommandsRequestTypeID is TL type id of BotsSetBotCommandsRequest.
const BotsSetBotCommandsRequestTypeID = 0x805d46f6

// Encode implements bin.Encoder.
func (s *BotsSetBotCommandsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode bots.setBotCommands#805d46f6 as nil")
	}
	b.PutID(BotsSetBotCommandsRequestTypeID)
	b.PutVectorHeader(len(s.Commands))
	for idx, v := range s.Commands {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode bots.setBotCommands#805d46f6: field commands element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *BotsSetBotCommandsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode bots.setBotCommands#805d46f6 to nil")
	}
	if err := b.ConsumeID(BotsSetBotCommandsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.setBotCommands#805d46f6: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode bots.setBotCommands#805d46f6: field commands: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommand
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode bots.setBotCommands#805d46f6: field commands: %w", err)
			}
			s.Commands = append(s.Commands, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for BotsSetBotCommandsRequest.
var (
	_ bin.Encoder = &BotsSetBotCommandsRequest{}
	_ bin.Decoder = &BotsSetBotCommandsRequest{}
)