// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/07-28-test/v2/pkg/utils"
)

type LocType string

const (
	LocTypeStr     LocType = "str"
	LocTypeInteger LocType = "integer"
)

type Loc struct {
	Str     *string
	Integer *int64

	Type LocType
}

func CreateLocStr(str string) Loc {
	typ := LocTypeStr

	return Loc{
		Str:  &str,
		Type: typ,
	}
}

func CreateLocInteger(integer int64) Loc {
	typ := LocTypeInteger

	return Loc{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *Loc) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = LocTypeStr
		return nil
	}

	integer := int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = LocTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Loc) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ValidationError struct {
	Loc  []Loc  `json:"loc"`
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

func (o *ValidationError) GetLoc() []Loc {
	if o == nil {
		return []Loc{}
	}
	return o.Loc
}

func (o *ValidationError) GetMsg() string {
	if o == nil {
		return ""
	}
	return o.Msg
}

func (o *ValidationError) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}