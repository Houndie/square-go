package objects

import (
	"encoding/json"
	"errors"
	"fmt"
)

type catalogCustomAttributeDefinitionType string

const (
	catalogCustomAttributeDefinitionTypeString    catalogCustomAttributeDefinitionType = "STRING"
	catalogCustomAttributeDefinitionTypeBoolean   catalogCustomAttributeDefinitionType = "BOOLEAN"
	catalogCustomAttributeDefinitionTypeNumber    catalogCustomAttributeDefinitionType = "NUMBER"
	catalogCustomAttributeDefinitionTypeSelection catalogCustomAttributeDefinitionType = "SELECTION"
)

type CatalogCustomAttributeDefinitionAppVisibility string

const (
	CatalogCustomAttributeDefinitionAppVisibilityHidden      CatalogCustomAttributeDefinitionAppVisibility = "APP_VISIBILITY_HIDDEN"
	CatalogCustomAttributeDefinitionAppVisibilityReadOnly    CatalogCustomAttributeDefinitionAppVisibility = "APP_VISIBILITY_READ_ONLY"
	CatalogCustomAttributeDefinitionAppVisibilityWriteValues CatalogCustomAttributeDefinitionAppVisibility = "APP_VISIBILITY_WRITE_VALUES"
)

type CatalogCustomAttributeDefinitionSellerVisibility string

const (
	CatalogCustomAttributeDefinitionSellerVisibilityHidden          CatalogCustomAttributeDefinitionSellerVisibility = "APP_VISIBILITY_HIDDEN"
	CatalogCustomAttributeDefinitionSellerVisibilityReadWriteValues CatalogCustomAttributeDefinitionSellerVisibility = "APP_VISIBILITY_READ_WRITE_VALUES"
)

type catalogCustomAttributeDefinition struct {
	AllowedObjectTypes        CatalogObjectType                                `json:"allowed_object_types,omitempty"`
	Name                      string                                           `json:"name,omitempty"`
	Type                      catalogCustomAttributeDefinitionType             `json:"type,omitempty"`
	AppVisibility             CatalogCustomAttributeDefinitionAppVisibility    `json:"app_visibility,omitempty"`
	CustomAttributeUsageCount int                                              `json:"custom_attribute_usage_count,omitempty"`
	Description               string                                           `json:"description,string"`
	Key                       string                                           `json:"key,omitempty"`
	NumberConfig              *CatalogCustomAttributeDefinitionNumberConfig    `json:"number_config,omitempty"`
	SelectionConfig           *CatalogCustomAttributeDefinitionSelectionConfig `json:"selection_config,omitempty"`
	SellerVisibility          CatalogCustomAttributeDefinitionSellerVisibility `json:"seller_visibility,omitempty"`
	SourceApplication         *SourceApplication                               `json:"source_application,omitempty"`
	StringConfig              *CatalogCustomAttributeDefinitionStringConfig    `json:"string_config,omitempty"`
}

type CatalogCustomAttributeDefinitionNumberConfig struct {
	Precision int `json:"precision,omitempty"`
}

type CatalogCustomAttributeDefinitionSelectionConfig struct {
	AllowedSelections    []*CatalogCustomAttributeDefinitionSelectionConfigCustomAttributeSelection `json:allowed_selections,omitempty"`
	MaxAllowedSelections int                                                                        `json:"max_allowed_selections,omitempty"`
}

type CatalogCustomAttributeDefinitionSelectionConfigCustomAttributeSelection struct {
	Name string `json:"name,omitempty"`
	UID  string `json:"uid,omitempty"`
}

type CatalogCustomAttributeDefinitionStringConfig struct {
	EnforceUniqueness bool `json:"enforce_uniqueness,omitempty"`
}

type CatalogCustomAttributeDefinitionType interface {
	isCatalogCustomAttributeDefinitionType()
}

type CatalogCustomAttributeDefinitionTypeBoolean struct{}
type CatalogCustomAttributeDefinitionTypeNumber struct {
	Config *CatalogCustomAttributeDefinitionNumberConfig
}
type CatalogCustomAttributeDefinitionTypeString struct {
	Config *CatalogCustomAttributeDefinitionStringConfig
}
type CatalogCustomAttributeDefinitionTypeSelection struct {
	Config *CatalogCustomAttributeDefinitionSelectionConfig
}

func (*CatalogCustomAttributeDefinitionTypeBoolean) isCatalogCustomAttributeDefinitionType()   {}
func (*CatalogCustomAttributeDefinitionTypeString) isCatalogCustomAttributeDefinitionType()    {}
func (*CatalogCustomAttributeDefinitionTypeNumber) isCatalogCustomAttributeDefinitionType()    {}
func (*CatalogCustomAttributeDefinitionTypeSelection) isCatalogCustomAttributeDefinitionType() {}

type CatalogCustomAttributeDefinition struct {
	AllowedObjectTypes        CatalogObjectType
	Name                      string
	Type                      CatalogCustomAttributeDefinitionType
	AppVisibility             CatalogCustomAttributeDefinitionAppVisibility
	CustomAttributeUsageCount int
	Description               string
	Key                       string
	SellerVisibility          CatalogCustomAttributeDefinitionSellerVisibility
	SourceApplication         *SourceApplication
}

func (*CatalogCustomAttributeDefinition) isCatalogObjectType() {}

func (c *CatalogCustomAttributeDefinition) MarshalJSON() ([]byte, error) {
	cJson := catalogCustomAttributeDefinition{
		AllowedObjectTypes:        c.AllowedObjectTypes,
		Name:                      c.Name,
		AppVisibility:             c.AppVisibility,
		CustomAttributeUsageCount: c.CustomAttributeUsageCount,
		Description:               c.Description,
		Key:                       c.Key,
		SellerVisibility:          c.SellerVisibility,
		SourceApplication:         c.SourceApplication,
	}

	switch t := c.Type.(type) {
	case *CatalogCustomAttributeDefinitionTypeBoolean:
		cJson.Type = catalogCustomAttributeDefinitionTypeBoolean
	case *CatalogCustomAttributeDefinitionTypeNumber:
		cJson.NumberConfig = t.Config
		cJson.Type = catalogCustomAttributeDefinitionTypeNumber
	case *CatalogCustomAttributeDefinitionTypeSelection:
		cJson.SelectionConfig = t.Config
		cJson.Type = catalogCustomAttributeDefinitionTypeSelection
	case *CatalogCustomAttributeDefinitionTypeString:
		cJson.StringConfig = t.Config
		cJson.Type = catalogCustomAttributeDefinitionTypeString
	default:
		return nil, errors.New("Found unknown custom attribute data type")
	}
	json, err := json.Marshal(&cJson)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling json catalog object: %w", err)
	}
	return json, nil
}

func (c *CatalogCustomAttributeDefinition) UnmarshalJSON(data []byte) error {
	cJson := &catalogCustomAttributeDefinition{}
	err := json.Unmarshal(data, &cJson)
	if err != nil {
		return fmt.Errorf("Error unmarshaling catalog object: %w", err)
	}
	c.AllowedObjectTypes = cJson.AllowedObjectTypes
	c.Name = cJson.Name
	c.AppVisibility = cJson.AppVisibility
	c.CustomAttributeUsageCount = cJson.CustomAttributeUsageCount
	c.Description = cJson.Description
	c.Key = cJson.Key
	c.SellerVisibility = cJson.SellerVisibility
	c.SourceApplication = cJson.SourceApplication

	switch cJson.Type {
	case catalogCustomAttributeDefinitionTypeBoolean:
		c.Type = &CatalogCustomAttributeDefinitionTypeBoolean{}
	case catalogCustomAttributeDefinitionTypeNumber:
		c.Type = &CatalogCustomAttributeDefinitionTypeNumber{Config: cJson.NumberConfig}
	case catalogCustomAttributeDefinitionTypeString:
		c.Type = &CatalogCustomAttributeDefinitionTypeString{Config: cJson.StringConfig}
	case catalogCustomAttributeDefinitionTypeSelection:
		c.Type = &CatalogCustomAttributeDefinitionTypeSelection{Config: cJson.SelectionConfig}
	default:
		return fmt.Errorf("Found unknown custom attribute type %s", cJson.Type)
	}
	return nil
}
