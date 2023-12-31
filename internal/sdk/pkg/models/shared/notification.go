// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Notification struct {
	NotificationType        NotificationType                     `json:"notificationType"`
	SendOnSuccess           *bool                                `default:"false" json:"sendOnSuccess"`
	SendOnFailure           *bool                                `default:"true" json:"sendOnFailure"`
	SlackConfiguration      *SlackNotificationConfiguration      `json:"slackConfiguration,omitempty"`
	CustomerioConfiguration *CustomerioNotificationConfiguration `json:"customerioConfiguration,omitempty"`
}

func (n Notification) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *Notification) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Notification) GetNotificationType() NotificationType {
	if o == nil {
		return NotificationType("")
	}
	return o.NotificationType
}

func (o *Notification) GetSendOnSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.SendOnSuccess
}

func (o *Notification) GetSendOnFailure() *bool {
	if o == nil {
		return nil
	}
	return o.SendOnFailure
}

func (o *Notification) GetSlackConfiguration() *SlackNotificationConfiguration {
	if o == nil {
		return nil
	}
	return o.SlackConfiguration
}

func (o *Notification) GetCustomerioConfiguration() *CustomerioNotificationConfiguration {
	if o == nil {
		return nil
	}
	return o.CustomerioConfiguration
}
