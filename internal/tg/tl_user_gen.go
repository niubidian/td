// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// UserEmpty represents TL type `userEmpty#200250ba`.
type UserEmpty struct {
	// ID field of UserEmpty.
	ID int
}

// UserEmptyTypeID is TL type id of UserEmpty.
const UserEmptyTypeID = 0x200250ba

// Encode implements bin.Encoder.
func (u *UserEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userEmpty#200250ba as nil")
	}
	b.PutID(UserEmptyTypeID)
	b.PutInt(u.ID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userEmpty#200250ba to nil")
	}
	if err := b.ConsumeID(UserEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userEmpty#200250ba: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userEmpty#200250ba: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// construct implements constructor of UserClass.
func (u UserEmpty) construct() UserClass { return &u }

// Ensuring interfaces in compile-time for UserEmpty.
var (
	_ bin.Encoder = &UserEmpty{}
	_ bin.Decoder = &UserEmpty{}

	_ UserClass = &UserEmpty{}
)

// User represents TL type `user#938458c1`.
type User struct {
	// Flags field of User.
	Flags bin.Fields
	// Self field of User.
	Self bool
	// Contact field of User.
	Contact bool
	// MutualContact field of User.
	MutualContact bool
	// Deleted field of User.
	Deleted bool
	// Bot field of User.
	Bot bool
	// BotChatHistory field of User.
	BotChatHistory bool
	// BotNochats field of User.
	BotNochats bool
	// Verified field of User.
	Verified bool
	// Restricted field of User.
	Restricted bool
	// Min field of User.
	Min bool
	// BotInlineGeo field of User.
	BotInlineGeo bool
	// Support field of User.
	Support bool
	// Scam field of User.
	Scam bool
	// ApplyMinPhoto field of User.
	ApplyMinPhoto bool
	// ID field of User.
	ID int
	// AccessHash field of User.
	//
	// Use SetAccessHash and GetAccessHash helpers.
	AccessHash int64
	// FirstName field of User.
	//
	// Use SetFirstName and GetFirstName helpers.
	FirstName string
	// LastName field of User.
	//
	// Use SetLastName and GetLastName helpers.
	LastName string
	// Username field of User.
	//
	// Use SetUsername and GetUsername helpers.
	Username string
	// Phone field of User.
	//
	// Use SetPhone and GetPhone helpers.
	Phone string
	// Photo field of User.
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo UserProfilePhotoClass
	// Status field of User.
	//
	// Use SetStatus and GetStatus helpers.
	Status UserStatusClass
	// BotInfoVersion field of User.
	//
	// Use SetBotInfoVersion and GetBotInfoVersion helpers.
	BotInfoVersion int
	// RestrictionReason field of User.
	//
	// Use SetRestrictionReason and GetRestrictionReason helpers.
	RestrictionReason []RestrictionReason
	// BotInlinePlaceholder field of User.
	//
	// Use SetBotInlinePlaceholder and GetBotInlinePlaceholder helpers.
	BotInlinePlaceholder string
	// LangCode field of User.
	//
	// Use SetLangCode and GetLangCode helpers.
	LangCode string
}

// UserTypeID is TL type id of User.
const UserTypeID = 0x938458c1

// Encode implements bin.Encoder.
func (u *User) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode user#938458c1 as nil")
	}
	b.PutID(UserTypeID)
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#938458c1: field flags: %w", err)
	}
	b.PutInt(u.ID)
	if u.Flags.Has(0) {
		b.PutLong(u.AccessHash)
	}
	if u.Flags.Has(1) {
		b.PutString(u.FirstName)
	}
	if u.Flags.Has(2) {
		b.PutString(u.LastName)
	}
	if u.Flags.Has(3) {
		b.PutString(u.Username)
	}
	if u.Flags.Has(4) {
		b.PutString(u.Phone)
	}
	if u.Flags.Has(5) {
		if u.Photo == nil {
			return fmt.Errorf("unable to encode user#938458c1: field photo is nil")
		}
		if err := u.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode user#938458c1: field photo: %w", err)
		}
	}
	if u.Flags.Has(6) {
		if u.Status == nil {
			return fmt.Errorf("unable to encode user#938458c1: field status is nil")
		}
		if err := u.Status.Encode(b); err != nil {
			return fmt.Errorf("unable to encode user#938458c1: field status: %w", err)
		}
	}
	if u.Flags.Has(14) {
		b.PutInt(u.BotInfoVersion)
	}
	if u.Flags.Has(18) {
		b.PutVectorHeader(len(u.RestrictionReason))
		for idx, v := range u.RestrictionReason {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode user#938458c1: field restriction_reason element with index %d: %w", idx, err)
			}
		}
	}
	if u.Flags.Has(19) {
		b.PutString(u.BotInlinePlaceholder)
	}
	if u.Flags.Has(22) {
		b.PutString(u.LangCode)
	}
	return nil
}

// SetSelf sets value of Self conditional field.
func (u *User) SetSelf(value bool) {
	if value {
		u.Flags.Set(10)
	} else {
		u.Flags.Unset(10)
	}
}

// SetContact sets value of Contact conditional field.
func (u *User) SetContact(value bool) {
	if value {
		u.Flags.Set(11)
	} else {
		u.Flags.Unset(11)
	}
}

// SetMutualContact sets value of MutualContact conditional field.
func (u *User) SetMutualContact(value bool) {
	if value {
		u.Flags.Set(12)
	} else {
		u.Flags.Unset(12)
	}
}

// SetDeleted sets value of Deleted conditional field.
func (u *User) SetDeleted(value bool) {
	if value {
		u.Flags.Set(13)
	} else {
		u.Flags.Unset(13)
	}
}

// SetBot sets value of Bot conditional field.
func (u *User) SetBot(value bool) {
	if value {
		u.Flags.Set(14)
	} else {
		u.Flags.Unset(14)
	}
}

// SetBotChatHistory sets value of BotChatHistory conditional field.
func (u *User) SetBotChatHistory(value bool) {
	if value {
		u.Flags.Set(15)
	} else {
		u.Flags.Unset(15)
	}
}

// SetBotNochats sets value of BotNochats conditional field.
func (u *User) SetBotNochats(value bool) {
	if value {
		u.Flags.Set(16)
	} else {
		u.Flags.Unset(16)
	}
}

// SetVerified sets value of Verified conditional field.
func (u *User) SetVerified(value bool) {
	if value {
		u.Flags.Set(17)
	} else {
		u.Flags.Unset(17)
	}
}

// SetRestricted sets value of Restricted conditional field.
func (u *User) SetRestricted(value bool) {
	if value {
		u.Flags.Set(18)
	} else {
		u.Flags.Unset(18)
	}
}

// SetMin sets value of Min conditional field.
func (u *User) SetMin(value bool) {
	if value {
		u.Flags.Set(20)
	} else {
		u.Flags.Unset(20)
	}
}

// SetBotInlineGeo sets value of BotInlineGeo conditional field.
func (u *User) SetBotInlineGeo(value bool) {
	if value {
		u.Flags.Set(21)
	} else {
		u.Flags.Unset(21)
	}
}

// SetSupport sets value of Support conditional field.
func (u *User) SetSupport(value bool) {
	if value {
		u.Flags.Set(23)
	} else {
		u.Flags.Unset(23)
	}
}

// SetScam sets value of Scam conditional field.
func (u *User) SetScam(value bool) {
	if value {
		u.Flags.Set(24)
	} else {
		u.Flags.Unset(24)
	}
}

// SetApplyMinPhoto sets value of ApplyMinPhoto conditional field.
func (u *User) SetApplyMinPhoto(value bool) {
	if value {
		u.Flags.Set(25)
	} else {
		u.Flags.Unset(25)
	}
}

// SetAccessHash sets value of AccessHash conditional field.
func (u *User) SetAccessHash(value int64) {
	u.Flags.Set(0)
	u.AccessHash = value
}

// GetAccessHash returns value of AccessHash conditional field and
// boolean which is true if field was set.
func (u *User) GetAccessHash() (value int64, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.AccessHash, true
}

// SetFirstName sets value of FirstName conditional field.
func (u *User) SetFirstName(value string) {
	u.Flags.Set(1)
	u.FirstName = value
}

// GetFirstName returns value of FirstName conditional field and
// boolean which is true if field was set.
func (u *User) GetFirstName() (value string, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.FirstName, true
}

// SetLastName sets value of LastName conditional field.
func (u *User) SetLastName(value string) {
	u.Flags.Set(2)
	u.LastName = value
}

// GetLastName returns value of LastName conditional field and
// boolean which is true if field was set.
func (u *User) GetLastName() (value string, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.LastName, true
}

// SetUsername sets value of Username conditional field.
func (u *User) SetUsername(value string) {
	u.Flags.Set(3)
	u.Username = value
}

// GetUsername returns value of Username conditional field and
// boolean which is true if field was set.
func (u *User) GetUsername() (value string, ok bool) {
	if !u.Flags.Has(3) {
		return value, false
	}
	return u.Username, true
}

// SetPhone sets value of Phone conditional field.
func (u *User) SetPhone(value string) {
	u.Flags.Set(4)
	u.Phone = value
}

// GetPhone returns value of Phone conditional field and
// boolean which is true if field was set.
func (u *User) GetPhone() (value string, ok bool) {
	if !u.Flags.Has(4) {
		return value, false
	}
	return u.Phone, true
}

// SetPhoto sets value of Photo conditional field.
func (u *User) SetPhoto(value UserProfilePhotoClass) {
	u.Flags.Set(5)
	u.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (u *User) GetPhoto() (value UserProfilePhotoClass, ok bool) {
	if !u.Flags.Has(5) {
		return value, false
	}
	return u.Photo, true
}

// SetStatus sets value of Status conditional field.
func (u *User) SetStatus(value UserStatusClass) {
	u.Flags.Set(6)
	u.Status = value
}

// GetStatus returns value of Status conditional field and
// boolean which is true if field was set.
func (u *User) GetStatus() (value UserStatusClass, ok bool) {
	if !u.Flags.Has(6) {
		return value, false
	}
	return u.Status, true
}

// SetBotInfoVersion sets value of BotInfoVersion conditional field.
func (u *User) SetBotInfoVersion(value int) {
	u.Flags.Set(14)
	u.BotInfoVersion = value
}

// GetBotInfoVersion returns value of BotInfoVersion conditional field and
// boolean which is true if field was set.
func (u *User) GetBotInfoVersion() (value int, ok bool) {
	if !u.Flags.Has(14) {
		return value, false
	}
	return u.BotInfoVersion, true
}

// SetRestrictionReason sets value of RestrictionReason conditional field.
func (u *User) SetRestrictionReason(value []RestrictionReason) {
	u.Flags.Set(18)
	u.RestrictionReason = value
}

// GetRestrictionReason returns value of RestrictionReason conditional field and
// boolean which is true if field was set.
func (u *User) GetRestrictionReason() (value []RestrictionReason, ok bool) {
	if !u.Flags.Has(18) {
		return value, false
	}
	return u.RestrictionReason, true
}

// SetBotInlinePlaceholder sets value of BotInlinePlaceholder conditional field.
func (u *User) SetBotInlinePlaceholder(value string) {
	u.Flags.Set(19)
	u.BotInlinePlaceholder = value
}

// GetBotInlinePlaceholder returns value of BotInlinePlaceholder conditional field and
// boolean which is true if field was set.
func (u *User) GetBotInlinePlaceholder() (value string, ok bool) {
	if !u.Flags.Has(19) {
		return value, false
	}
	return u.BotInlinePlaceholder, true
}

// SetLangCode sets value of LangCode conditional field.
func (u *User) SetLangCode(value string) {
	u.Flags.Set(22)
	u.LangCode = value
}

// GetLangCode returns value of LangCode conditional field and
// boolean which is true if field was set.
func (u *User) GetLangCode() (value string, ok bool) {
	if !u.Flags.Has(22) {
		return value, false
	}
	return u.LangCode, true
}

// Decode implements bin.Decoder.
func (u *User) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode user#938458c1 to nil")
	}
	if err := b.ConsumeID(UserTypeID); err != nil {
		return fmt.Errorf("unable to decode user#938458c1: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field flags: %w", err)
		}
	}
	u.Self = u.Flags.Has(10)
	u.Contact = u.Flags.Has(11)
	u.MutualContact = u.Flags.Has(12)
	u.Deleted = u.Flags.Has(13)
	u.Bot = u.Flags.Has(14)
	u.BotChatHistory = u.Flags.Has(15)
	u.BotNochats = u.Flags.Has(16)
	u.Verified = u.Flags.Has(17)
	u.Restricted = u.Flags.Has(18)
	u.Min = u.Flags.Has(20)
	u.BotInlineGeo = u.Flags.Has(21)
	u.Support = u.Flags.Has(23)
	u.Scam = u.Flags.Has(24)
	u.ApplyMinPhoto = u.Flags.Has(25)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field id: %w", err)
		}
		u.ID = value
	}
	if u.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field access_hash: %w", err)
		}
		u.AccessHash = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field first_name: %w", err)
		}
		u.FirstName = value
	}
	if u.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field last_name: %w", err)
		}
		u.LastName = value
	}
	if u.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field username: %w", err)
		}
		u.Username = value
	}
	if u.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field phone: %w", err)
		}
		u.Phone = value
	}
	if u.Flags.Has(5) {
		value, err := DecodeUserProfilePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field photo: %w", err)
		}
		u.Photo = value
	}
	if u.Flags.Has(6) {
		value, err := DecodeUserStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field status: %w", err)
		}
		u.Status = value
	}
	if u.Flags.Has(14) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field bot_info_version: %w", err)
		}
		u.BotInfoVersion = value
	}
	if u.Flags.Has(18) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field restriction_reason: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value RestrictionReason
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode user#938458c1: field restriction_reason: %w", err)
			}
			u.RestrictionReason = append(u.RestrictionReason, value)
		}
	}
	if u.Flags.Has(19) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field bot_inline_placeholder: %w", err)
		}
		u.BotInlinePlaceholder = value
	}
	if u.Flags.Has(22) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field lang_code: %w", err)
		}
		u.LangCode = value
	}
	return nil
}

// construct implements constructor of UserClass.
func (u User) construct() UserClass { return &u }

// Ensuring interfaces in compile-time for User.
var (
	_ bin.Encoder = &User{}
	_ bin.Decoder = &User{}

	_ UserClass = &User{}
)

// UserClass represents User generic type.
//
// Example:
//  g, err := DecodeUser(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UserEmpty: // userEmpty#200250ba
//  case *User: // user#938458c1
//  default: panic(v)
//  }
type UserClass interface {
	bin.Encoder
	bin.Decoder
	construct() UserClass
}

// DecodeUser implements binary de-serialization for UserClass.
func DecodeUser(buf *bin.Buffer) (UserClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserEmptyTypeID:
		// Decoding userEmpty#200250ba.
		v := UserEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserClass: %w", err)
		}
		return &v, nil
	case UserTypeID:
		// Decoding user#938458c1.
		v := User{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserClass: %w", bin.NewUnexpectedID(id))
	}
}