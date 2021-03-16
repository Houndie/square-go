package objects

import (
	"encoding/json"
	"errors"
	"fmt"
)

type catalogCustomAttributeValue struct {
	BooleanValue                bool                                 `json:"boolean_value,omitempty"`
	CustomAttributeDefinitionID string                               `json:"custom_attribute_definition_id,omitempty"`
	Key                         string                               `json:"key,omitempty"`
	Name                        string                               `json:"name,omitempty"`
	NumberValue                 string                               `json:"number_value,omitempty`
	SelectionUIDValues          []string                             `json:"selection_uid_values,omitempty`
	StringValue                 string                               `json:"string_value,omitempty"`
	Type                        catalogCustomAttributeDefinitionType `json:"type,omitempty"`
}

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
	CustomAttributeDefinitionID string
	Key                         string
	Name                        string
	Type                        CatalogCustomAttributeValueType
}

func (c *CatalogCustomAttributeValue) MarshalJSON() ([]byte, error) {
	cJson := catalogCustomAttributeValue{
		CustomAttributeDefinitionID: c.CustomAttributeDefinitionID,
		Key:                         c.Key,
		Name:                        c.Name,
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
	cJson := &catalogCustomAttributeValue{}
	err := json.Unmarshal(data, &cJson)
	if err != nil {
		return fmt.Errorf("Error unmarshaling catalog object: %w", err)
	}
	c.CustomAttributeDefinitionID = cJson.CustomAttributeDefinitionID
	c.Key = cJson.Key
	c.Name = cJson.Name

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
