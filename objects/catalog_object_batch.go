package objects

type CatalogObjectBatch struct {
	Objects []*CatalogObject `json:"objects,omitempty"`
}
