// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NotificationItem struct {
	NotificationType        []NotificationType                   `json:"notificationType,omitempty"`
	SlackConfiguration      *SlackNotificationConfiguration      `json:"slackConfiguration,omitempty"`
	CustomerioConfiguration *CustomerioNotificationConfiguration `json:"customerioConfiguration,omitempty"`
}

func (o *NotificationItem) GetNotificationType() []NotificationType {
	if o == nil {
		return nil
	}
	return o.NotificationType
}

func (o *NotificationItem) GetSlackConfiguration() *SlackNotificationConfiguration {
	if o == nil {
		return nil
	}
	return o.SlackConfiguration
}

func (o *NotificationItem) GetCustomerioConfiguration() *CustomerioNotificationConfiguration {
	if o == nil {
		return nil
	}
	return o.CustomerioConfiguration
}
