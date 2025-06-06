// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSchemaDataSourceModel) ToSharedSourceDiscoverSchemaRequestBody() *shared.SourceDiscoverSchemaRequestBody {
	sourceID := r.SourceID.ValueString()
	connectionID := new(string)
	if !r.ConnectionID.IsUnknown() && !r.ConnectionID.IsNull() {
		*connectionID = r.ConnectionID.ValueString()
	} else {
		connectionID = nil
	}
	disableCache := new(bool)
	if !r.DisableCache.IsUnknown() && !r.DisableCache.IsNull() {
		*disableCache = r.DisableCache.ValueBool()
	} else {
		disableCache = nil
	}
	notifySchemaChange := new(bool)
	if !r.NotifySchemaChange.IsUnknown() && !r.NotifySchemaChange.IsNull() {
		*notifySchemaChange = r.NotifySchemaChange.ValueBool()
	} else {
		notifySchemaChange = nil
	}
	out := shared.SourceDiscoverSchemaRequestBody{
		SourceID:           sourceID,
		ConnectionID:       connectionID,
		DisableCache:       disableCache,
		NotifySchemaChange: notifySchemaChange,
	}
	return &out
}

func (r *SourceSchemaDataSourceModel) RefreshFromSharedSourceDiscoverSchemaRead(resp *shared.SourceDiscoverSchemaRead) {
	r.BreakingChange = types.BoolPointerValue(resp.BreakingChange)
	if resp.Catalog == nil {
		r.Catalog = nil
	} else {
		r.Catalog = &AirbyteCatalog{}
		if len(r.Catalog.Streams) > len(resp.Catalog.Streams) {
			r.Catalog.Streams = r.Catalog.Streams[:len(resp.Catalog.Streams)]
		}
		for streamsCount, streamsItem := range resp.Catalog.Streams {
			var streams1 AirbyteStreamAndConfiguration
			if streamsItem.Config == nil {
				streams1.Config = nil
			} else {
				streams1.Config = &AirbyteStreamConfiguration{}
				streams1.Config.AliasName = types.StringPointerValue(streamsItem.Config.AliasName)
				streams1.Config.CursorField = nil
				for _, v := range streamsItem.Config.CursorField {
					streams1.Config.CursorField = append(streams1.Config.CursorField, types.StringValue(v))
				}
				streams1.Config.DestinationSyncMode = types.StringValue(string(streamsItem.Config.DestinationSyncMode))
				streams1.Config.FieldSelectionEnabled = types.BoolPointerValue(streamsItem.Config.FieldSelectionEnabled)
				streams1.Config.PrimaryKey = nil
				for _, primaryKeyItem := range streamsItem.Config.PrimaryKey {
					var primaryKey1 []types.String
					primaryKey1 = nil
					for _, v := range primaryKeyItem {
						primaryKey1 = append(primaryKey1, types.StringValue(v))
					}
					streams1.Config.PrimaryKey = append(streams1.Config.PrimaryKey, primaryKey1)
				}
				streams1.Config.Selected = types.BoolPointerValue(streamsItem.Config.Selected)
				for selectedFieldsCount, selectedFieldsItem := range streamsItem.Config.SelectedFields {
					var selectedFields1 SelectedFieldInfo
					selectedFields1.FieldPath = nil
					for _, v := range selectedFieldsItem.FieldPath {
						selectedFields1.FieldPath = append(selectedFields1.FieldPath, types.StringValue(v))
					}
					if selectedFieldsCount+1 > len(streams1.Config.SelectedFields) {
						streams1.Config.SelectedFields = append(streams1.Config.SelectedFields, selectedFields1)
					} else {
						streams1.Config.SelectedFields[selectedFieldsCount].FieldPath = selectedFields1.FieldPath
					}
				}
				streams1.Config.Suggested = types.BoolPointerValue(streamsItem.Config.Suggested)
				streams1.Config.SyncMode = types.StringValue(string(streamsItem.Config.SyncMode))
			}
			if streamsItem.Stream == nil {
				streams1.Stream = nil
			} else {
				streams1.Stream = &AirbyteStream{}
				streams1.Stream.DefaultCursorField = nil
				for _, v := range streamsItem.Stream.DefaultCursorField {
					streams1.Stream.DefaultCursorField = append(streams1.Stream.DefaultCursorField, types.StringValue(v))
				}
				if len(streamsItem.Stream.JSONSchema) > 0 {
					streams1.Stream.JSONSchema = make(map[string]types.String)
					for key, value := range streamsItem.Stream.JSONSchema {
						result, _ := json.Marshal(value)
						streams1.Stream.JSONSchema[key] = types.StringValue(string(result))
					}
				}
				streams1.Stream.Name = types.StringValue(streamsItem.Stream.Name)
				streams1.Stream.Namespace = types.StringPointerValue(streamsItem.Stream.Namespace)
				streams1.Stream.SourceDefinedCursor = types.BoolPointerValue(streamsItem.Stream.SourceDefinedCursor)
				streams1.Stream.SourceDefinedPrimaryKey = nil
				for _, sourceDefinedPrimaryKeyItem := range streamsItem.Stream.SourceDefinedPrimaryKey {
					var sourceDefinedPrimaryKey1 []types.String
					sourceDefinedPrimaryKey1 = nil
					for _, v := range sourceDefinedPrimaryKeyItem {
						sourceDefinedPrimaryKey1 = append(sourceDefinedPrimaryKey1, types.StringValue(v))
					}
					streams1.Stream.SourceDefinedPrimaryKey = append(streams1.Stream.SourceDefinedPrimaryKey, sourceDefinedPrimaryKey1)
				}
				streams1.Stream.SupportedSyncModes = nil
				for _, v := range streamsItem.Stream.SupportedSyncModes {
					streams1.Stream.SupportedSyncModes = append(streams1.Stream.SupportedSyncModes, types.StringValue(string(v)))
				}
			}
			if streamsCount+1 > len(r.Catalog.Streams) {
				r.Catalog.Streams = append(r.Catalog.Streams, streams1)
			} else {
				r.Catalog.Streams[streamsCount].Config = streams1.Config
				r.Catalog.Streams[streamsCount].Stream = streams1.Stream
			}
		}
	}
	if resp.CatalogDiff == nil {
		r.CatalogDiff = nil
	} else {
		r.CatalogDiff = &CatalogDiff{}
		if len(r.CatalogDiff.Transforms) > len(resp.CatalogDiff.Transforms) {
			r.CatalogDiff.Transforms = r.CatalogDiff.Transforms[:len(resp.CatalogDiff.Transforms)]
		}
		for transformsCount, transformsItem := range resp.CatalogDiff.Transforms {
			var transforms1 StreamTransform
			transforms1.StreamDescriptor.Name = types.StringValue(transformsItem.StreamDescriptor.Name)
			transforms1.StreamDescriptor.Namespace = types.StringPointerValue(transformsItem.StreamDescriptor.Namespace)
			transforms1.TransformType = types.StringValue(string(transformsItem.TransformType))
			for updateStreamCount, updateStreamItem := range transformsItem.UpdateStream {
				var updateStream1 FieldTransform
				if updateStreamItem.AddField == nil {
					updateStream1.AddField = nil
				} else {
					updateStream1.AddField = &FieldAdd{}
					if updateStreamItem.AddField.Schema == nil {
						updateStream1.AddField.Schema = nil
					} else {
						updateStream1.AddField.Schema = &DeclarativeManifest{}
					}
				}
				updateStream1.Breaking = types.BoolValue(updateStreamItem.Breaking)
				updateStream1.FieldName = nil
				for _, v := range updateStreamItem.FieldName {
					updateStream1.FieldName = append(updateStream1.FieldName, types.StringValue(v))
				}
				if updateStreamItem.RemoveField == nil {
					updateStream1.RemoveField = nil
				} else {
					updateStream1.RemoveField = &FieldAdd{}
					if updateStreamItem.RemoveField.Schema == nil {
						updateStream1.RemoveField.Schema = nil
					} else {
						updateStream1.RemoveField.Schema = &DeclarativeManifest{}
					}
				}
				updateStream1.TransformType = types.StringValue(string(updateStreamItem.TransformType))
				if updateStreamItem.UpdateFieldSchema == nil {
					updateStream1.UpdateFieldSchema = nil
				} else {
					updateStream1.UpdateFieldSchema = &FieldSchemaUpdate{}
				}
				if updateStreamCount+1 > len(transforms1.UpdateStream) {
					transforms1.UpdateStream = append(transforms1.UpdateStream, updateStream1)
				} else {
					transforms1.UpdateStream[updateStreamCount].AddField = updateStream1.AddField
					transforms1.UpdateStream[updateStreamCount].Breaking = updateStream1.Breaking
					transforms1.UpdateStream[updateStreamCount].FieldName = updateStream1.FieldName
					transforms1.UpdateStream[updateStreamCount].RemoveField = updateStream1.RemoveField
					transforms1.UpdateStream[updateStreamCount].TransformType = updateStream1.TransformType
					transforms1.UpdateStream[updateStreamCount].UpdateFieldSchema = updateStream1.UpdateFieldSchema
				}
			}
			if transformsCount+1 > len(r.CatalogDiff.Transforms) {
				r.CatalogDiff.Transforms = append(r.CatalogDiff.Transforms, transforms1)
			} else {
				r.CatalogDiff.Transforms[transformsCount].StreamDescriptor = transforms1.StreamDescriptor
				r.CatalogDiff.Transforms[transformsCount].TransformType = transforms1.TransformType
				r.CatalogDiff.Transforms[transformsCount].UpdateStream = transforms1.UpdateStream
			}
		}
	}
	r.CatalogID = types.StringPointerValue(resp.CatalogID)
	if resp.ConnectionStatus != nil {
		r.ConnectionStatus = types.StringValue(string(*resp.ConnectionStatus))
	} else {
		r.ConnectionStatus = types.StringNull()
	}
	r.JobInfo.ConfigID = types.StringPointerValue(resp.JobInfo.ConfigID)
	r.JobInfo.ConfigType = types.StringValue(string(resp.JobInfo.ConfigType))
	r.JobInfo.ConnectorConfigurationUpdated = types.BoolPointerValue(resp.JobInfo.ConnectorConfigurationUpdated)
	r.JobInfo.CreatedAt = types.Int64Value(resp.JobInfo.CreatedAt)
	r.JobInfo.EndedAt = types.Int64Value(resp.JobInfo.EndedAt)
	if resp.JobInfo.FailureReason == nil {
		r.JobInfo.FailureReason = nil
	} else {
		r.JobInfo.FailureReason = &FailureReason{}
		r.JobInfo.FailureReason.ExternalMessage = types.StringPointerValue(resp.JobInfo.FailureReason.ExternalMessage)
		if resp.JobInfo.FailureReason.FailureOrigin != nil {
			r.JobInfo.FailureReason.FailureOrigin = types.StringValue(string(*resp.JobInfo.FailureReason.FailureOrigin))
		} else {
			r.JobInfo.FailureReason.FailureOrigin = types.StringNull()
		}
		if resp.JobInfo.FailureReason.FailureType != nil {
			r.JobInfo.FailureReason.FailureType = types.StringValue(string(*resp.JobInfo.FailureReason.FailureType))
		} else {
			r.JobInfo.FailureReason.FailureType = types.StringNull()
		}
		r.JobInfo.FailureReason.InternalMessage = types.StringPointerValue(resp.JobInfo.FailureReason.InternalMessage)
		r.JobInfo.FailureReason.Retryable = types.BoolPointerValue(resp.JobInfo.FailureReason.Retryable)
		r.JobInfo.FailureReason.Stacktrace = types.StringPointerValue(resp.JobInfo.FailureReason.Stacktrace)
		r.JobInfo.FailureReason.Timestamp = types.Int64Value(resp.JobInfo.FailureReason.Timestamp)
	}
	r.JobInfo.ID = types.StringValue(resp.JobInfo.ID)
	if resp.JobInfo.Logs == nil {
		r.JobInfo.Logs = nil
	} else {
		r.JobInfo.Logs = &LogRead{}
		r.JobInfo.Logs.LogLines = nil
		for _, v := range resp.JobInfo.Logs.LogLines {
			r.JobInfo.Logs.LogLines = append(r.JobInfo.Logs.LogLines, types.StringValue(v))
		}
	}
	r.JobInfo.Succeeded = types.BoolValue(resp.JobInfo.Succeeded)
}
