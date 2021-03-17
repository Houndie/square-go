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
	cJSON := catalogCustomAttributeValue{
		catalogCustomAttributeValueAlias: (*catalogCustomAttributeValueAlias)(c),
	}

	switch t := c.Type.(type) {
	case CatalogCustomAttributeValueBoolean:
		cJSON.BooleanValue = bool(t)
		cJSON.Type = catalogCustomAttributeDefinitionTypeBoolean
	case CatalogCustomAttributeValueNumber:
		cJSON.NumberValue = string(t)
		cJSON.Type = catalogCustomAttributeDefinitionTypeNumber
	case CatalogCustomAttributeValueSelection:
		cJSON.SelectionUIDValues = []string(t)
		cJSON.Type = catalogCustomAttributeDefinitionTypeSelection
	case CatalogCustomAttributeValueString:
		cJSON.StringValue = string(t)
		cJSON.Type = catalogCustomAttributeDefinitionTypeString
	default:
		return nil, errors.New("found unknown custom attribute data type")
	}

	json, err := json.Marshal(&cJSON)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json catalog object: %w", err)
	}

	return json, nil
}

func (c *CatalogCustomAttributeValue) UnmarshalJSON(data []byte) error {
	cJSON := catalogCustomAttributeValue{
		catalogCustomAttributeValueAlias: (*catalogCustomAttributeValueAlias)(c),
	}

	err := json.Unmarshal(data, &cJSON)
	if err != nil {
		return fmt.Errorf("error unmarshaling catalog object: %w", err)
	}

	switch cJSON.Type {
	case catalogCustomAttributeDefinitionTypeBoolean:
		c.Type = CatalogCustomAttributeValueBoolean(cJSON.BooleanValue)
	case catalogCustomAttributeDefinitionTypeNumber:
		c.Type = CatalogCustomAttributeValueNumber(cJSON.NumberValue)
	case catalogCustomAttributeDefinitionTypeString:
		c.Type = CatalogCustomAttributeValueString(cJSON.StringValue)
	case catalogCustomAttributeDefinitionTypeSelection:
		c.Type = CatalogCustomAttributeValueSelection(cJSON.SelectionUIDValues)
	default:
		return fmt.Errorf("found unknown custom attribute type %s", cJSON.Type)
	}

	return nil
}
