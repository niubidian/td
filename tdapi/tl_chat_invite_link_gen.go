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

// ChatInviteLink represents TL type `chatInviteLink#f3bb8d04`.
type ChatInviteLink struct {
	// Chat invite link
	InviteLink string
	// Name of the link
	Name string
	// User identifier of an administrator created the link
	CreatorUserID int64
	// Point in time (Unix timestamp) when the link was created
	Date int32
	// Point in time (Unix timestamp) when the link was last edited; 0 if never or unknown
	EditDate int32
	// Point in time (Unix timestamp) when the link will expire; 0 if never
	ExpirationDate int32
	// The maximum number of members, which can join the chat using the link simultaneously;
	// 0 if not limited. Always 0 if the link requires approval
	MemberLimit int32
	// Number of chat members, which joined the chat using the link
	MemberCount int32
	// Number of pending join requests created using this link
	PendingJoinRequestCount int32
	// True, if the link only creates join request. If true, total number of joining members
	// will be unlimited
	CreatesJoinRequest bool
	// True, if the link is primary. Primary invite link can't have name, expiration date, or
	// usage limit. There is exactly one primary invite link for each administrator with
	// can_invite_users right at a given time
	IsPrimary bool
	// True, if the link was revoked
	IsRevoked bool
}

// ChatInviteLinkTypeID is TL type id of ChatInviteLink.
const ChatInviteLinkTypeID = 0xf3bb8d04

// Ensuring interfaces in compile-time for ChatInviteLink.
var (
	_ bin.Encoder     = &ChatInviteLink{}
	_ bin.Decoder     = &ChatInviteLink{}
	_ bin.BareEncoder = &ChatInviteLink{}
	_ bin.BareDecoder = &ChatInviteLink{}
)

func (c *ChatInviteLink) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.InviteLink == "") {
		return false
	}
	if !(c.Name == "") {
		return false
	}
	if !(c.CreatorUserID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.EditDate == 0) {
		return false
	}
	if !(c.ExpirationDate == 0) {
		return false
	}
	if !(c.MemberLimit == 0) {
		return false
	}
	if !(c.MemberCount == 0) {
		return false
	}
	if !(c.PendingJoinRequestCount == 0) {
		return false
	}
	if !(c.CreatesJoinRequest == false) {
		return false
	}
	if !(c.IsPrimary == false) {
		return false
	}
	if !(c.IsRevoked == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteLink) String() string {
	if c == nil {
		return "ChatInviteLink(nil)"
	}
	type Alias ChatInviteLink
	return fmt.Sprintf("ChatInviteLink%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteLink) TypeID() uint32 {
	return ChatInviteLinkTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteLink) TypeName() string {
	return "chatInviteLink"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteLink) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteLink",
		ID:   ChatInviteLinkTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "CreatorUserID",
			SchemaName: "creator_user_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "EditDate",
			SchemaName: "edit_date",
		},
		{
			Name:       "ExpirationDate",
			SchemaName: "expiration_date",
		},
		{
			Name:       "MemberLimit",
			SchemaName: "member_limit",
		},
		{
			Name:       "MemberCount",
			SchemaName: "member_count",
		},
		{
			Name:       "PendingJoinRequestCount",
			SchemaName: "pending_join_request_count",
		},
		{
			Name:       "CreatesJoinRequest",
			SchemaName: "creates_join_request",
		},
		{
			Name:       "IsPrimary",
			SchemaName: "is_primary",
		},
		{
			Name:       "IsRevoked",
			SchemaName: "is_revoked",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteLink) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLink#f3bb8d04 as nil")
	}
	b.PutID(ChatInviteLinkTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteLink) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLink#f3bb8d04 as nil")
	}
	b.PutString(c.InviteLink)
	b.PutString(c.Name)
	b.PutInt53(c.CreatorUserID)
	b.PutInt32(c.Date)
	b.PutInt32(c.EditDate)
	b.PutInt32(c.ExpirationDate)
	b.PutInt32(c.MemberLimit)
	b.PutInt32(c.MemberCount)
	b.PutInt32(c.PendingJoinRequestCount)
	b.PutBool(c.CreatesJoinRequest)
	b.PutBool(c.IsPrimary)
	b.PutBool(c.IsRevoked)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteLink) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLink#f3bb8d04 to nil")
	}
	if err := b.ConsumeID(ChatInviteLinkTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteLink) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLink#f3bb8d04 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field invite_link: %w", err)
		}
		c.InviteLink = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field name: %w", err)
		}
		c.Name = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field creator_user_id: %w", err)
		}
		c.CreatorUserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field edit_date: %w", err)
		}
		c.EditDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field expiration_date: %w", err)
		}
		c.ExpirationDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field member_limit: %w", err)
		}
		c.MemberLimit = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field member_count: %w", err)
		}
		c.MemberCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field pending_join_request_count: %w", err)
		}
		c.PendingJoinRequestCount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field creates_join_request: %w", err)
		}
		c.CreatesJoinRequest = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field is_primary: %w", err)
		}
		c.IsPrimary = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field is_revoked: %w", err)
		}
		c.IsRevoked = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatInviteLink) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLink#f3bb8d04 as nil")
	}
	b.ObjStart()
	b.PutID("chatInviteLink")
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(c.InviteLink)
	b.Comma()
	b.FieldStart("name")
	b.PutString(c.Name)
	b.Comma()
	b.FieldStart("creator_user_id")
	b.PutInt53(c.CreatorUserID)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(c.Date)
	b.Comma()
	b.FieldStart("edit_date")
	b.PutInt32(c.EditDate)
	b.Comma()
	b.FieldStart("expiration_date")
	b.PutInt32(c.ExpirationDate)
	b.Comma()
	b.FieldStart("member_limit")
	b.PutInt32(c.MemberLimit)
	b.Comma()
	b.FieldStart("member_count")
	b.PutInt32(c.MemberCount)
	b.Comma()
	b.FieldStart("pending_join_request_count")
	b.PutInt32(c.PendingJoinRequestCount)
	b.Comma()
	b.FieldStart("creates_join_request")
	b.PutBool(c.CreatesJoinRequest)
	b.Comma()
	b.FieldStart("is_primary")
	b.PutBool(c.IsPrimary)
	b.Comma()
	b.FieldStart("is_revoked")
	b.PutBool(c.IsRevoked)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatInviteLink) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLink#f3bb8d04 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatInviteLink"); err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: %w", err)
			}
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field invite_link: %w", err)
			}
			c.InviteLink = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field name: %w", err)
			}
			c.Name = value
		case "creator_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field creator_user_id: %w", err)
			}
			c.CreatorUserID = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field date: %w", err)
			}
			c.Date = value
		case "edit_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field edit_date: %w", err)
			}
			c.EditDate = value
		case "expiration_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field expiration_date: %w", err)
			}
			c.ExpirationDate = value
		case "member_limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field member_limit: %w", err)
			}
			c.MemberLimit = value
		case "member_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field member_count: %w", err)
			}
			c.MemberCount = value
		case "pending_join_request_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field pending_join_request_count: %w", err)
			}
			c.PendingJoinRequestCount = value
		case "creates_join_request":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field creates_join_request: %w", err)
			}
			c.CreatesJoinRequest = value
		case "is_primary":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field is_primary: %w", err)
			}
			c.IsPrimary = value
		case "is_revoked":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLink#f3bb8d04: field is_revoked: %w", err)
			}
			c.IsRevoked = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInviteLink returns value of InviteLink field.
func (c *ChatInviteLink) GetInviteLink() (value string) {
	if c == nil {
		return
	}
	return c.InviteLink
}

// GetName returns value of Name field.
func (c *ChatInviteLink) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

// GetCreatorUserID returns value of CreatorUserID field.
func (c *ChatInviteLink) GetCreatorUserID() (value int64) {
	if c == nil {
		return
	}
	return c.CreatorUserID
}

// GetDate returns value of Date field.
func (c *ChatInviteLink) GetDate() (value int32) {
	if c == nil {
		return
	}
	return c.Date
}

// GetEditDate returns value of EditDate field.
func (c *ChatInviteLink) GetEditDate() (value int32) {
	if c == nil {
		return
	}
	return c.EditDate
}

// GetExpirationDate returns value of ExpirationDate field.
func (c *ChatInviteLink) GetExpirationDate() (value int32) {
	if c == nil {
		return
	}
	return c.ExpirationDate
}

// GetMemberLimit returns value of MemberLimit field.
func (c *ChatInviteLink) GetMemberLimit() (value int32) {
	if c == nil {
		return
	}
	return c.MemberLimit
}

// GetMemberCount returns value of MemberCount field.
func (c *ChatInviteLink) GetMemberCount() (value int32) {
	if c == nil {
		return
	}
	return c.MemberCount
}

// GetPendingJoinRequestCount returns value of PendingJoinRequestCount field.
func (c *ChatInviteLink) GetPendingJoinRequestCount() (value int32) {
	if c == nil {
		return
	}
	return c.PendingJoinRequestCount
}

// GetCreatesJoinRequest returns value of CreatesJoinRequest field.
func (c *ChatInviteLink) GetCreatesJoinRequest() (value bool) {
	if c == nil {
		return
	}
	return c.CreatesJoinRequest
}

// GetIsPrimary returns value of IsPrimary field.
func (c *ChatInviteLink) GetIsPrimary() (value bool) {
	if c == nil {
		return
	}
	return c.IsPrimary
}

// GetIsRevoked returns value of IsRevoked field.
func (c *ChatInviteLink) GetIsRevoked() (value bool) {
	if c == nil {
		return
	}
	return c.IsRevoked
}
