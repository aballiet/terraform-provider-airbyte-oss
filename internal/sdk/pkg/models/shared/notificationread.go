// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Status string

const (
	StatusSucceeded Status = "succeeded"
	StatusFailed    Status = "failed"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "succeeded":
		fallthrough
	case "failed":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type NotificationRead struct {
	Status  Status  `json:"status"`
	Message *string `json:"message,omitempty"`
}

func (o *NotificationRead) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}

func (o *NotificationRead) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
