// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PermissionReadList struct {
	Permissions []PermissionRead `json:"permissions"`
}

func (o *PermissionReadList) GetPermissions() []PermissionRead {
	if o == nil {
		return []PermissionRead{}
	}
	return o.Permissions
}
