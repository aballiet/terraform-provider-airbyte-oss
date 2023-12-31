// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PermissionDeleteUserFromWorkspaceRequestBody struct {
	// The user ID for which to remove workspace permissions
	UserIDToRemove string `json:"userIdToRemove"`
	// The workspace ID from which to remove all workspace-level permissions for the indicated user
	WorkspaceID string `json:"workspaceId"`
}

func (o *PermissionDeleteUserFromWorkspaceRequestBody) GetUserIDToRemove() string {
	if o == nil {
		return ""
	}
	return o.UserIDToRemove
}

func (o *PermissionDeleteUserFromWorkspaceRequestBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
