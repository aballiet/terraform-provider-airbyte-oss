// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type ConnectionScheduleData struct {
	BasicSchedule *ConnectionScheduleDataBasicSchedule `tfsdk:"basic_schedule"`
	Cron          *ConnectionScheduleDataCron          `tfsdk:"cron"`
}