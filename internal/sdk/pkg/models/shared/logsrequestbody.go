// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LogsRequestBody struct {
	// type/source of logs produced
	LogType LogType `json:"logType"`
}

func (o *LogsRequestBody) GetLogType() LogType {
	if o == nil {
		return LogType("")
	}
	return o.LogType
}
