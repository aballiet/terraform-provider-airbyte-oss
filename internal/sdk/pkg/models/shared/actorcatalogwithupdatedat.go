// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Catalog struct {
}

// ActorCatalogWithUpdatedAt - A source actor catalog with the timestamp it was mostly recently updated
type ActorCatalogWithUpdatedAt struct {
	UpdatedAt *int64   `json:"updatedAt,omitempty"`
	Catalog   *Catalog `json:"catalog,omitempty"`
}

func (o *ActorCatalogWithUpdatedAt) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ActorCatalogWithUpdatedAt) GetCatalog() *Catalog {
	if o == nil {
		return nil
	}
	return o.Catalog
}
