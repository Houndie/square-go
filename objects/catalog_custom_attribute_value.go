package objects

import (
	"encoding/json"
	"errors"
	"fmt"
)

type catalogCustomAttributeValue struct {
	*catalogCustomAttributeValueAlias
	BooleanValue       bool                                 `json:"boolean_value,omitempty"`
	NumberValue        string                               `json:"number_value,omitempty"`
	SelectionUIDValues []string                             `json:"selection_uid_values,omitempty"`
	StringValue        string                               `json:"string_value,omitempty"`
	Type               catalogCustomAttributeDefinitionType `json:"type,omitempty"`
}

type catalogCustomAttributeValueAlias CatalogCustomAttributeValue

type CatalogCustomAttributeValueType interface {
	isCatalogCustomAttributeValueType()
}

type CatalogCustomAttributeValueBoolean bool
type CatalogCustomAttributeValueNumber string
type CatalogCustomAttributeValueSelection []string
type CatalogCustomAttributeValueString string

func (CatalogCustomAttributeValueBoolean) isCatalogCustomAttributeValueType()   {}
func (CatalogCustomAttributeValueNumber) isCatalogCustomAttributeValueType()    {}
func (CatalogCustomAttributeValueSelection) isCatalogCustomAttributeValueType() {}
func (CatalogCustomAttributeValueString) isCatalogCustomAttributeValueType()    {}

type CatalogCustomAttributeValue struct {
	CustomAttributeDefinitionID string                          `json:"custom_attribute_definition_id,omitempty"`
	Key                         string                          `json:"key,omitempty"`
	Name                        string                          `json:"name,omitempty"`
	Type                        CatalogCustomAttributeValueType `json:"-"`
}

func (c *CatalogCustomAttributeValue) MarshalJSON() ([]byte, error) {
	cJson := catalogCustomAttributeValue{
		catalogCustomAttributeValueAlias: (*catalogCustomAttributeValueAlias)(c),
	}

	switch t := c.Type.(type) {
	case CatalogCustomAttributeValueBoolean:
		cJson.BooleanValue = bool(t)
		cJson.Type = catalogCustomAttributeDefinitionTypeBoolean
	case CatalogCustomAttributeValueNumber:
		cJson.NumberValue = string(t)
		cJson.Type = catalogCustomAttributeDefinitionTypeNumber
	case CatalogCustomAttributeValueSelection:
		cJson.SelectionUIDValues = []string(t)
		cJson.Type = catalogCustomAttributeDefinitionTypeSelection
	case CatalogCustomAttributeValueString:
		cJson.StringValue = string(t)
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

func (c *CatalogCustomAttributeValue) UnmarshalJSON(data []byte) error {
	cJson := catalogCustomAttributeValue{
		catalogCustomAttributeValueAlias: (*catalogCustomAttributeValueAlias)(c),
	}
	err := json.Unmarshal(data, &cJson)
	if err != nil {
		return fmt.Errorf("Error unmarshaling catalog object: %w", err)
	}

	switch cJson.Type {
	case catalogCustomAttributeDefinitionTypeBoolean:
		c.Type = CatalogCustomAttributeValueBoolean(cJson.BooleanValue)
	case catalogCustomAttributeDefinitionTypeNumber:
		c.Type = CatalogCustomAttributeValueNumber(cJson.NumberValue)
	case catalogCustomAttributeDefinitionTypeString:
		c.Type = CatalogCustomAttributeValueString(cJson.StringValue)
	case catalogCustomAttributeDefinitionTypeSelection:
		c.Type = CatalogCustomAttributeValueSelection(cJson.SelectionUIDValues)
	default:
		return fmt.Errorf("Found unknown custom attribute type %s", cJson.Type)
	}
	return nil
}
