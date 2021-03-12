package objects

type CatalogImage struct {
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	Caption string `json:"caption,omitempty"`
}

func (*CatalogImage) isCatalogObjectType() {}
