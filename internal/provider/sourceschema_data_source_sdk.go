// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSchemaDataSourceModel) ToGetSDKType() *shared.SourceDiscoverSchemaRequestBody {
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

func (r *SourceSchemaDataSourceModel) RefreshFromGetResponse(resp *shared.SourceDiscoverSchemaRead) {
	if resp.BreakingChange != nil {
		r.BreakingChange = types.BoolValue(*resp.BreakingChange)
	} else {
		r.BreakingChange = types.BoolNull()
	}
	if resp.Catalog == nil {
		r.Catalog = nil
	} else {
		r.Catalog = &AirbyteCatalog{}
		if len(r.Catalog.Streams) > len(resp.Catalog.Streams) {
			r.Catalog.Streams = r.Catalog.Streams[:len(resp.Catalog.Streams)]
		}
		for streamsCount, streamsItem := range resp.Catalog.Streams {
			var streams1 AirbyteStreamAndConfiguration
			if streamsItem.Stream == nil {
				streams1.Stream = nil
			} else {
				streams1.Stream = &AirbyteStream{}
				streams1.Stream.Name = types.StringValue(streamsItem.Stream.Name)
				if streams1.Stream.JSONSchema == nil && len(streamsItem.Stream.JSONSchema) > 0 {
					streams1.Stream.JSONSchema = make(map[string]types.String)
					for key, value := range streamsItem.Stream.JSONSchema {
						result, _ := json.Marshal(value)
						streams1.Stream.JSONSchema[key] = types.StringValue(string(result))
					}
				}
				streams1.Stream.SupportedSyncModes = nil
				for _, v := range streamsItem.Stream.SupportedSyncModes {
					streams1.Stream.SupportedSyncModes = append(streams1.Stream.SupportedSyncModes, types.StringValue(string(v)))
				}
				if streamsItem.Stream.SourceDefinedCursor != nil {
					streams1.Stream.SourceDefinedCursor = types.BoolValue(*streamsItem.Stream.SourceDefinedCursor)
				} else {
					streams1.Stream.SourceDefinedCursor = types.BoolNull()
				}
				streams1.Stream.DefaultCursorField = nil
				for _, v := range streamsItem.Stream.DefaultCursorField {
					streams1.Stream.DefaultCursorField = append(streams1.Stream.DefaultCursorField, types.StringValue(v))
				}
				streams1.Stream.SourceDefinedPrimaryKey = nil
				for _, sourceDefinedPrimaryKeyItem := range streamsItem.Stream.SourceDefinedPrimaryKey {
					var sourceDefinedPrimaryKey1 []types.String
					sourceDefinedPrimaryKey1 = nil
					for _, v := range sourceDefinedPrimaryKeyItem {
						sourceDefinedPrimaryKey1 = append(sourceDefinedPrimaryKey1, types.StringValue(v))
					}
					streams1.Stream.SourceDefinedPrimaryKey = append(streams1.Stream.SourceDefinedPrimaryKey, sourceDefinedPrimaryKey1)
				}
				if streamsItem.Stream.Namespace != nil {
					streams1.Stream.Namespace = types.StringValue(*streamsItem.Stream.Namespace)
				} else {
					streams1.Stream.Namespace = types.StringNull()
				}
			}
			if streamsItem.Config == nil {
				streams1.Config = nil
			} else {
				streams1.Config = &AirbyteStreamConfiguration{}
				streams1.Config.SyncMode = types.StringValue(string(streamsItem.Config.SyncMode))
				streams1.Config.CursorField = nil
				for _, v := range streamsItem.Config.CursorField {
					streams1.Config.CursorField = append(streams1.Config.CursorField, types.StringValue(v))
				}
				streams1.Config.DestinationSyncMode = types.StringValue(string(streamsItem.Config.DestinationSyncMode))
				streams1.Config.PrimaryKey = nil
				for _, primaryKeyItem := range streamsItem.Config.PrimaryKey {
					var primaryKey1 []types.String
					primaryKey1 = nil
					for _, v := range primaryKeyItem {
						primaryKey1 = append(primaryKey1, types.StringValue(v))
					}
					streams1.Config.PrimaryKey = append(streams1.Config.PrimaryKey, primaryKey1)
				}
				if streamsItem.Config.AliasName != nil {
					streams1.Config.AliasName = types.StringValue(*streamsItem.Config.AliasName)
				} else {
					streams1.Config.AliasName = types.StringNull()
				}
				if streamsItem.Config.Selected != nil {
					streams1.Config.Selected = types.BoolValue(*streamsItem.Config.Selected)
				} else {
					streams1.Config.Selected = types.BoolNull()
				}
				if streamsItem.Config.Suggested != nil {
					streams1.Config.Suggested = types.BoolValue(*streamsItem.Config.Suggested)
				} else {
					streams1.Config.Suggested = types.BoolNull()
				}
				if streamsItem.Config.FieldSelectionEnabled != nil {
					streams1.Config.FieldSelectionEnabled = types.BoolValue(*streamsItem.Config.FieldSelectionEnabled)
				} else {
					streams1.Config.FieldSelectionEnabled = types.BoolNull()
				}
				if len(streams1.Config.SelectedFields) > len(streamsItem.Config.SelectedFields) {
					streams1.Config.SelectedFields = streams1.Config.SelectedFields[:len(streamsItem.Config.SelectedFields)]
				}
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
			}
			if streamsCount+1 > len(r.Catalog.Streams) {
				r.Catalog.Streams = append(r.Catalog.Streams, streams1)
			} else {
				r.Catalog.Streams[streamsCount].Stream = streams1.Stream
				r.Catalog.Streams[streamsCount].Config = streams1.Config
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
			transforms1.TransformType = types.StringValue(string(transformsItem.TransformType))
			transforms1.StreamDescriptor.Name = types.StringValue(transformsItem.StreamDescriptor.Name)
			if transformsItem.StreamDescriptor.Namespace != nil {
				transforms1.StreamDescriptor.Namespace = types.StringValue(*transformsItem.StreamDescriptor.Namespace)
			} else {
				transforms1.StreamDescriptor.Namespace = types.StringNull()
			}
			if len(transforms1.UpdateStream) > len(transformsItem.UpdateStream) {
				transforms1.UpdateStream = transforms1.UpdateStream[:len(transformsItem.UpdateStream)]
			}
			for updateStreamCount, updateStreamItem := range transformsItem.UpdateStream {
				var updateStream1 FieldTransform
				updateStream1.TransformType = types.StringValue(string(updateStreamItem.TransformType))
				updateStream1.FieldName = nil
				for _, v := range updateStreamItem.FieldName {
					updateStream1.FieldName = append(updateStream1.FieldName, types.StringValue(v))
				}
				updateStream1.Breaking = types.BoolValue(updateStreamItem.Breaking)
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
				if updateStreamItem.UpdateFieldSchema == nil {
					updateStream1.UpdateFieldSchema = nil
				} else {
					updateStream1.UpdateFieldSchema = &FieldSchemaUpdate{}
				}
				if updateStreamCount+1 > len(transforms1.UpdateStream) {
					transforms1.UpdateStream = append(transforms1.UpdateStream, updateStream1)
				} else {
					transforms1.UpdateStream[updateStreamCount].TransformType = updateStream1.TransformType
					transforms1.UpdateStream[updateStreamCount].FieldName = updateStream1.FieldName
					transforms1.UpdateStream[updateStreamCount].Breaking = updateStream1.Breaking
					transforms1.UpdateStream[updateStreamCount].AddField = updateStream1.AddField
					transforms1.UpdateStream[updateStreamCount].RemoveField = updateStream1.RemoveField
					transforms1.UpdateStream[updateStreamCount].UpdateFieldSchema = updateStream1.UpdateFieldSchema
				}
			}
			if transformsCount+1 > len(r.CatalogDiff.Transforms) {
				r.CatalogDiff.Transforms = append(r.CatalogDiff.Transforms, transforms1)
			} else {
				r.CatalogDiff.Transforms[transformsCount].TransformType = transforms1.TransformType
				r.CatalogDiff.Transforms[transformsCount].StreamDescriptor = transforms1.StreamDescriptor
				r.CatalogDiff.Transforms[transformsCount].UpdateStream = transforms1.UpdateStream
			}
		}
	}
	if resp.CatalogID != nil {
		r.CatalogID = types.StringValue(*resp.CatalogID)
	} else {
		r.CatalogID = types.StringNull()
	}
	if resp.ConnectionStatus != nil {
		r.ConnectionStatus = types.StringValue(string(*resp.ConnectionStatus))
	} else {
		r.ConnectionStatus = types.StringNull()
	}
	r.JobInfo.ID = types.StringValue(resp.JobInfo.ID)
	r.JobInfo.ConfigType = types.StringValue(string(resp.JobInfo.ConfigType))
	if resp.JobInfo.ConfigID != nil {
		r.JobInfo.ConfigID = types.StringValue(*resp.JobInfo.ConfigID)
	} else {
		r.JobInfo.ConfigID = types.StringNull()
	}
	r.JobInfo.CreatedAt = types.Int64Value(resp.JobInfo.CreatedAt)
	r.JobInfo.EndedAt = types.Int64Value(resp.JobInfo.EndedAt)
	r.JobInfo.Succeeded = types.BoolValue(resp.JobInfo.Succeeded)
	if resp.JobInfo.ConnectorConfigurationUpdated != nil {
		r.JobInfo.ConnectorConfigurationUpdated = types.BoolValue(*resp.JobInfo.ConnectorConfigurationUpdated)
	} else {
		r.JobInfo.ConnectorConfigurationUpdated = types.BoolNull()
	}
	if resp.JobInfo.Logs == nil {
		r.JobInfo.Logs = nil
	} else {
		r.JobInfo.Logs = &LogRead{}
		r.JobInfo.Logs.LogLines = nil
		for _, v := range resp.JobInfo.Logs.LogLines {
			r.JobInfo.Logs.LogLines = append(r.JobInfo.Logs.LogLines, types.StringValue(v))
		}
	}
	if resp.JobInfo.FailureReason == nil {
		r.JobInfo.FailureReason = nil
	} else {
		r.JobInfo.FailureReason = &FailureReason{}
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
		if resp.JobInfo.FailureReason.ExternalMessage != nil {
			r.JobInfo.FailureReason.ExternalMessage = types.StringValue(*resp.JobInfo.FailureReason.ExternalMessage)
		} else {
			r.JobInfo.FailureReason.ExternalMessage = types.StringNull()
		}
		if resp.JobInfo.FailureReason.InternalMessage != nil {
			r.JobInfo.FailureReason.InternalMessage = types.StringValue(*resp.JobInfo.FailureReason.InternalMessage)
		} else {
			r.JobInfo.FailureReason.InternalMessage = types.StringNull()
		}
		if resp.JobInfo.FailureReason.Stacktrace != nil {
			r.JobInfo.FailureReason.Stacktrace = types.StringValue(*resp.JobInfo.FailureReason.Stacktrace)
		} else {
			r.JobInfo.FailureReason.Stacktrace = types.StringNull()
		}
		if resp.JobInfo.FailureReason.Retryable != nil {
			r.JobInfo.FailureReason.Retryable = types.BoolValue(*resp.JobInfo.FailureReason.Retryable)
		} else {
			r.JobInfo.FailureReason.Retryable = types.BoolNull()
		}
		r.JobInfo.FailureReason.Timestamp = types.Int64Value(resp.JobInfo.FailureReason.Timestamp)
	}
}
