// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UserStatus - user status
type UserStatus string

const (
	UserStatusInvited    UserStatus = "invited"
	UserStatusRegistered UserStatus = "registered"
	UserStatusDisabled   UserStatus = "disabled"
)

func (e UserStatus) ToPointer() *UserStatus {
	return &e
}

func (e *UserStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invited":
		fallthrough
	case "registered":
		fallthrough
	case "disabled":
		*e = UserStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserStatus: %v", v)
	}
}
