// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UserWithPermissionInfoReadList struct {
	Users []UserWithPermissionInfoRead `json:"users"`
}

func (o *UserWithPermissionInfoReadList) GetUsers() []UserWithPermissionInfoRead {
	if o == nil {
		return []UserWithPermissionInfoRead{}
	}
	return o.Users
}
