// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PermissionCreate struct {
	// This is a temporary and optional field just for dual write purpose during the data migration.
	PermissionID *string `json:"permissionId,omitempty"`
	// Describes what actions/endpoints the permission entitles to
	PermissionType PermissionType `json:"permissionType"`
	// Internal Airbyte user ID
	UserID         string  `json:"userId"`
	WorkspaceID    *string `json:"workspaceId,omitempty"`
	OrganizationID *string `json:"organizationId,omitempty"`
}

func (o *PermissionCreate) GetPermissionID() *string {
	if o == nil {
		return nil
	}
	return o.PermissionID
}

func (o *PermissionCreate) GetPermissionType() PermissionType {
	if o == nil {
		return PermissionType("")
	}
	return o.PermissionType
}

func (o *PermissionCreate) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *PermissionCreate) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

func (o *PermissionCreate) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}
