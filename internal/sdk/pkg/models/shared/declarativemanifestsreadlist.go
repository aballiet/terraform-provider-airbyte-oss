// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeclarativeManifestsReadList struct {
	ManifestVersions []DeclarativeManifestVersionRead `json:"manifestVersions"`
}

func (o *DeclarativeManifestsReadList) GetManifestVersions() []DeclarativeManifestVersionRead {
	if o == nil {
		return []DeclarativeManifestVersionRead{}
	}
	return o.ManifestVersions
}
