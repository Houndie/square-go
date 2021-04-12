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
	Type            catalogCustomAttributeDefinitionType             `json:"type,omitempty"`
	NumberConfig    *CatalogCustomAttributeDefinitionNumberConfig    `json:"number_config,omitempty"`
	SelectionConfig *CatalogCustomAttributeDefinitionSelectionConfig `json:"selection_config,omitempty"`
	StringConfig    *CatalogCustomAttributeDefinitionStringConfig    `json:"string_config,omitempty"`
	*catalogCustomAttributeDefinitionAlias
}

type catalogCustomAttributeDefinitionAlias CatalogCustomAttributeDefinition

type CatalogCustomAttributeDefinitionNumberConfig struct {
	Precision *int `json:"precision,omitempty"`
}

type CatalogCustomAttributeDefinitionSelectionConfig struct {
	AllowedSelections    []*CatalogCustomAttributeDefinitionSelectionConfigCustomAttributeSelection `json:"allowed_selections,omitempty"`
	MaxAllowedSelections *int                                                                       `json:"max_allowed_selections,omitempty"`
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
	AllowedObjectTypes        CatalogObjectEnumType                            `json:"allowed_object_types,omitempty"`
	Name                      string                                           `json:"name,omitempty"`
	Type                      CatalogCustomAttributeDefinitionType             `json:"-"`
	AppVisibility             CatalogCustomAttributeDefinitionAppVisibility    `json:"app_visibility,omitempty"`
	CustomAttributeUsageCount int                                              `json:"custom_attribute_usage_count,omitempty"`
	Description               string                                           `json:"description,string"`
	Key                       string                                           `json:"key,omitempty"`
	SellerVisibility          CatalogCustomAttributeDefinitionSellerVisibility `json:"seller_visibility,omitempty"`
	SourceApplication         *SourceApplication                               `json:"source_application,omitempty"`
}

func (*CatalogCustomAttributeDefinition) isCatalogObjectType() {}

func (c *CatalogCustomAttributeDefinition) MarshalJSON() ([]byte, error) {
	cJSON := catalogCustomAttributeDefinition{
		catalogCustomAttributeDefinitionAlias: (*catalogCustomAttributeDefinitionAlias)(c),
	}

	switch t := c.Type.(type) {
	case *CatalogCustomAttributeDefinitionTypeBoolean:
		cJSON.Type = catalogCustomAttributeDefinitionTypeBoolean
	case *CatalogCustomAttributeDefinitionTypeNumber:
		cJSON.NumberConfig = t.Config
		cJSON.Type = catalogCustomAttributeDefinitionTypeNumber
	case *CatalogCustomAttributeDefinitionTypeSelection:
		cJSON.SelectionConfig = t.Config
		cJSON.Type = catalogCustomAttributeDefinitionTypeSelection
	case *CatalogCustomAttributeDefinitionTypeString:
		cJSON.StringConfig = t.Config
		cJSON.Type = catalogCustomAttributeDefinitionTypeString
	default:
		return nil, errors.New("found unknown custom attribute data type")
	}

	j, err := json.Marshal(&cJSON)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json catalog object: %w", err)
	}

	return j, nil
}

func (c *CatalogCustomAttributeDefinition) UnmarshalJSON(data []byte) error {
	cJSON := catalogCustomAttributeDefinition{
		catalogCustomAttributeDefinitionAlias: (*catalogCustomAttributeDefinitionAlias)(c),
	}

	err := json.Unmarshal(data, &cJSON)
	if err != nil {
		return fmt.Errorf("error unmarshaling catalog object: %w", err)
	}

	switch cJSON.Type {
	case catalogCustomAttributeDefinitionTypeBoolean:
		c.Type = &CatalogCustomAttributeDefinitionTypeBoolean{}
	case catalogCustomAttributeDefinitionTypeNumber:
		c.Type = &CatalogCustomAttributeDefinitionTypeNumber{Config: cJSON.NumberConfig}
	case catalogCustomAttributeDefinitionTypeString:
		c.Type = &CatalogCustomAttributeDefinitionTypeString{Config: cJSON.StringConfig}
	case catalogCustomAttributeDefinitionTypeSelection:
		c.Type = &CatalogCustomAttributeDefinitionTypeSelection{Config: cJSON.SelectionConfig}
	default:
		return fmt.Errorf("found unknown custom attribute type %s", cJSON.Type)
	}

	return nil
}
