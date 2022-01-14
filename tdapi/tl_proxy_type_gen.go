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

// ProxyTypeSocks5 represents TL type `proxyTypeSocks5#caf342b3`.
type ProxyTypeSocks5 struct {
	// Username for logging in; may be empty
	Username string
	// Password for logging in; may be empty
	Password string
}

// ProxyTypeSocks5TypeID is TL type id of ProxyTypeSocks5.
const ProxyTypeSocks5TypeID = 0xcaf342b3

// construct implements constructor of ProxyTypeClass.
func (p ProxyTypeSocks5) construct() ProxyTypeClass { return &p }

// Ensuring interfaces in compile-time for ProxyTypeSocks5.
var (
	_ bin.Encoder     = &ProxyTypeSocks5{}
	_ bin.Decoder     = &ProxyTypeSocks5{}
	_ bin.BareEncoder = &ProxyTypeSocks5{}
	_ bin.BareDecoder = &ProxyTypeSocks5{}

	_ ProxyTypeClass = &ProxyTypeSocks5{}
)

func (p *ProxyTypeSocks5) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Username == "") {
		return false
	}
	if !(p.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ProxyTypeSocks5) String() string {
	if p == nil {
		return "ProxyTypeSocks5(nil)"
	}
	type Alias ProxyTypeSocks5
	return fmt.Sprintf("ProxyTypeSocks5%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ProxyTypeSocks5) TypeID() uint32 {
	return ProxyTypeSocks5TypeID
}

// TypeName returns name of type in TL schema.
func (*ProxyTypeSocks5) TypeName() string {
	return "proxyTypeSocks5"
}

// TypeInfo returns info about TL type.
func (p *ProxyTypeSocks5) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "proxyTypeSocks5",
		ID:   ProxyTypeSocks5TypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ProxyTypeSocks5) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeSocks5#caf342b3 as nil")
	}
	b.PutID(ProxyTypeSocks5TypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ProxyTypeSocks5) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeSocks5#caf342b3 as nil")
	}
	b.PutString(p.Username)
	b.PutString(p.Password)
	return nil
}

// Decode implements bin.Decoder.
func (p *ProxyTypeSocks5) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeSocks5#caf342b3 to nil")
	}
	if err := b.ConsumeID(ProxyTypeSocks5TypeID); err != nil {
		return fmt.Errorf("unable to decode proxyTypeSocks5#caf342b3: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ProxyTypeSocks5) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeSocks5#caf342b3 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode proxyTypeSocks5#caf342b3: field username: %w", err)
		}
		p.Username = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode proxyTypeSocks5#caf342b3: field password: %w", err)
		}
		p.Password = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ProxyTypeSocks5) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeSocks5#caf342b3 as nil")
	}
	b.ObjStart()
	b.PutID("proxyTypeSocks5")
	b.Comma()
	b.FieldStart("username")
	b.PutString(p.Username)
	b.Comma()
	b.FieldStart("password")
	b.PutString(p.Password)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *ProxyTypeSocks5) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeSocks5#caf342b3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("proxyTypeSocks5"); err != nil {
				return fmt.Errorf("unable to decode proxyTypeSocks5#caf342b3: %w", err)
			}
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode proxyTypeSocks5#caf342b3: field username: %w", err)
			}
			p.Username = value
		case "password":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode proxyTypeSocks5#caf342b3: field password: %w", err)
			}
			p.Password = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUsername returns value of Username field.
func (p *ProxyTypeSocks5) GetUsername() (value string) {
	if p == nil {
		return
	}
	return p.Username
}

// GetPassword returns value of Password field.
func (p *ProxyTypeSocks5) GetPassword() (value string) {
	if p == nil {
		return
	}
	return p.Password
}

// ProxyTypeHTTP represents TL type `proxyTypeHttp#a3c7c777`.
type ProxyTypeHTTP struct {
	// Username for logging in; may be empty
	Username string
	// Password for logging in; may be empty
	Password string
	// Pass true if the proxy supports only HTTP requests and doesn't support transparent TCP
	// connections via HTTP CONNECT method
	HTTPOnly bool
}

// ProxyTypeHTTPTypeID is TL type id of ProxyTypeHTTP.
const ProxyTypeHTTPTypeID = 0xa3c7c777

// construct implements constructor of ProxyTypeClass.
func (p ProxyTypeHTTP) construct() ProxyTypeClass { return &p }

// Ensuring interfaces in compile-time for ProxyTypeHTTP.
var (
	_ bin.Encoder     = &ProxyTypeHTTP{}
	_ bin.Decoder     = &ProxyTypeHTTP{}
	_ bin.BareEncoder = &ProxyTypeHTTP{}
	_ bin.BareDecoder = &ProxyTypeHTTP{}

	_ ProxyTypeClass = &ProxyTypeHTTP{}
)

func (p *ProxyTypeHTTP) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Username == "") {
		return false
	}
	if !(p.Password == "") {
		return false
	}
	if !(p.HTTPOnly == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ProxyTypeHTTP) String() string {
	if p == nil {
		return "ProxyTypeHTTP(nil)"
	}
	type Alias ProxyTypeHTTP
	return fmt.Sprintf("ProxyTypeHTTP%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ProxyTypeHTTP) TypeID() uint32 {
	return ProxyTypeHTTPTypeID
}

// TypeName returns name of type in TL schema.
func (*ProxyTypeHTTP) TypeName() string {
	return "proxyTypeHttp"
}

// TypeInfo returns info about TL type.
func (p *ProxyTypeHTTP) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "proxyTypeHttp",
		ID:   ProxyTypeHTTPTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "HTTPOnly",
			SchemaName: "http_only",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ProxyTypeHTTP) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeHttp#a3c7c777 as nil")
	}
	b.PutID(ProxyTypeHTTPTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ProxyTypeHTTP) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeHttp#a3c7c777 as nil")
	}
	b.PutString(p.Username)
	b.PutString(p.Password)
	b.PutBool(p.HTTPOnly)
	return nil
}

// Decode implements bin.Decoder.
func (p *ProxyTypeHTTP) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeHttp#a3c7c777 to nil")
	}
	if err := b.ConsumeID(ProxyTypeHTTPTypeID); err != nil {
		return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ProxyTypeHTTP) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeHttp#a3c7c777 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: field username: %w", err)
		}
		p.Username = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: field password: %w", err)
		}
		p.Password = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: field http_only: %w", err)
		}
		p.HTTPOnly = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ProxyTypeHTTP) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeHttp#a3c7c777 as nil")
	}
	b.ObjStart()
	b.PutID("proxyTypeHttp")
	b.Comma()
	b.FieldStart("username")
	b.PutString(p.Username)
	b.Comma()
	b.FieldStart("password")
	b.PutString(p.Password)
	b.Comma()
	b.FieldStart("http_only")
	b.PutBool(p.HTTPOnly)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *ProxyTypeHTTP) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeHttp#a3c7c777 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("proxyTypeHttp"); err != nil {
				return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: %w", err)
			}
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: field username: %w", err)
			}
			p.Username = value
		case "password":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: field password: %w", err)
			}
			p.Password = value
		case "http_only":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode proxyTypeHttp#a3c7c777: field http_only: %w", err)
			}
			p.HTTPOnly = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUsername returns value of Username field.
func (p *ProxyTypeHTTP) GetUsername() (value string) {
	if p == nil {
		return
	}
	return p.Username
}

// GetPassword returns value of Password field.
func (p *ProxyTypeHTTP) GetPassword() (value string) {
	if p == nil {
		return
	}
	return p.Password
}

// GetHTTPOnly returns value of HTTPOnly field.
func (p *ProxyTypeHTTP) GetHTTPOnly() (value bool) {
	if p == nil {
		return
	}
	return p.HTTPOnly
}

// ProxyTypeMtproto represents TL type `proxyTypeMtproto#8ae31ffd`.
type ProxyTypeMtproto struct {
	// The proxy's secret in hexadecimal encoding
	Secret string
}

// ProxyTypeMtprotoTypeID is TL type id of ProxyTypeMtproto.
const ProxyTypeMtprotoTypeID = 0x8ae31ffd

// construct implements constructor of ProxyTypeClass.
func (p ProxyTypeMtproto) construct() ProxyTypeClass { return &p }

// Ensuring interfaces in compile-time for ProxyTypeMtproto.
var (
	_ bin.Encoder     = &ProxyTypeMtproto{}
	_ bin.Decoder     = &ProxyTypeMtproto{}
	_ bin.BareEncoder = &ProxyTypeMtproto{}
	_ bin.BareDecoder = &ProxyTypeMtproto{}

	_ ProxyTypeClass = &ProxyTypeMtproto{}
)

func (p *ProxyTypeMtproto) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Secret == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ProxyTypeMtproto) String() string {
	if p == nil {
		return "ProxyTypeMtproto(nil)"
	}
	type Alias ProxyTypeMtproto
	return fmt.Sprintf("ProxyTypeMtproto%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ProxyTypeMtproto) TypeID() uint32 {
	return ProxyTypeMtprotoTypeID
}

// TypeName returns name of type in TL schema.
func (*ProxyTypeMtproto) TypeName() string {
	return "proxyTypeMtproto"
}

// TypeInfo returns info about TL type.
func (p *ProxyTypeMtproto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "proxyTypeMtproto",
		ID:   ProxyTypeMtprotoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ProxyTypeMtproto) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeMtproto#8ae31ffd as nil")
	}
	b.PutID(ProxyTypeMtprotoTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ProxyTypeMtproto) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeMtproto#8ae31ffd as nil")
	}
	b.PutString(p.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (p *ProxyTypeMtproto) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeMtproto#8ae31ffd to nil")
	}
	if err := b.ConsumeID(ProxyTypeMtprotoTypeID); err != nil {
		return fmt.Errorf("unable to decode proxyTypeMtproto#8ae31ffd: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ProxyTypeMtproto) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeMtproto#8ae31ffd to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode proxyTypeMtproto#8ae31ffd: field secret: %w", err)
		}
		p.Secret = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ProxyTypeMtproto) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode proxyTypeMtproto#8ae31ffd as nil")
	}
	b.ObjStart()
	b.PutID("proxyTypeMtproto")
	b.Comma()
	b.FieldStart("secret")
	b.PutString(p.Secret)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *ProxyTypeMtproto) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode proxyTypeMtproto#8ae31ffd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("proxyTypeMtproto"); err != nil {
				return fmt.Errorf("unable to decode proxyTypeMtproto#8ae31ffd: %w", err)
			}
		case "secret":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode proxyTypeMtproto#8ae31ffd: field secret: %w", err)
			}
			p.Secret = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSecret returns value of Secret field.
func (p *ProxyTypeMtproto) GetSecret() (value string) {
	if p == nil {
		return
	}
	return p.Secret
}

// ProxyTypeClassName is schema name of ProxyTypeClass.
const ProxyTypeClassName = "ProxyType"

// ProxyTypeClass represents ProxyType generic type.
//
// Example:
//  g, err := tdapi.DecodeProxyType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.ProxyTypeSocks5: // proxyTypeSocks5#caf342b3
//  case *tdapi.ProxyTypeHTTP: // proxyTypeHttp#a3c7c777
//  case *tdapi.ProxyTypeMtproto: // proxyTypeMtproto#8ae31ffd
//  default: panic(v)
//  }
type ProxyTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ProxyTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeProxyType implements binary de-serialization for ProxyTypeClass.
func DecodeProxyType(buf *bin.Buffer) (ProxyTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ProxyTypeSocks5TypeID:
		// Decoding proxyTypeSocks5#caf342b3.
		v := ProxyTypeSocks5{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", err)
		}
		return &v, nil
	case ProxyTypeHTTPTypeID:
		// Decoding proxyTypeHttp#a3c7c777.
		v := ProxyTypeHTTP{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", err)
		}
		return &v, nil
	case ProxyTypeMtprotoTypeID:
		// Decoding proxyTypeMtproto#8ae31ffd.
		v := ProxyTypeMtproto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONProxyType implements binary de-serialization for ProxyTypeClass.
func DecodeTDLibJSONProxyType(buf tdjson.Decoder) (ProxyTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "proxyTypeSocks5":
		// Decoding proxyTypeSocks5#caf342b3.
		v := ProxyTypeSocks5{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", err)
		}
		return &v, nil
	case "proxyTypeHttp":
		// Decoding proxyTypeHttp#a3c7c777.
		v := ProxyTypeHTTP{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", err)
		}
		return &v, nil
	case "proxyTypeMtproto":
		// Decoding proxyTypeMtproto#8ae31ffd.
		v := ProxyTypeMtproto{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ProxyTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ProxyType boxes the ProxyTypeClass providing a helper.
type ProxyTypeBox struct {
	ProxyType ProxyTypeClass
}

// Decode implements bin.Decoder for ProxyTypeBox.
func (b *ProxyTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ProxyTypeBox to nil")
	}
	v, err := DecodeProxyType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ProxyType = v
	return nil
}

// Encode implements bin.Encode for ProxyTypeBox.
func (b *ProxyTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ProxyType == nil {
		return fmt.Errorf("unable to encode ProxyTypeClass as nil")
	}
	return b.ProxyType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ProxyTypeBox.
func (b *ProxyTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ProxyTypeBox to nil")
	}
	v, err := DecodeTDLibJSONProxyType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ProxyType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ProxyTypeBox.
func (b *ProxyTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ProxyType == nil {
		return fmt.Errorf("unable to encode ProxyTypeClass as nil")
	}
	return b.ProxyType.EncodeTDLibJSON(buf)
}
