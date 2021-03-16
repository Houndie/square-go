package objects

import (
	"encoding/json"
	"fmt"
	"time"

	"errors"
)

type tender struct {
	*tenderAlias
	Type        TenderType         `json:"type,omitempty"`
	CardDetails *TenderCardDetails `json:"card_details,omitempty"`
	CashDetails *TenderCashDetails `json:"cash_details,omitempty"`
}

type tenderAlias Tender

type tenderType interface {
	isTenderType()
}

type ThirdPartyCardDetails struct{}
type SquareGiftCardDetails struct{}
type NoSaleDetails struct{}
type OtherTenderTypeDetails struct{}

func (*ThirdPartyCardDetails) isTenderType()  {}
func (*SquareGiftCardDetails) isTenderType()  {}
func (*NoSaleDetails) isTenderType()          {}
func (*OtherTenderTypeDetails) isTenderType() {}

type Tender struct {
	ID                   string                 `json:"id,omitempty"`
	LocationID           string                 `json:"location_id,omitempty"`
	TransactionID        string                 `json:"transaction_id,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	Note                 string                 `json:"note,omitempty"`
	AmountMoney          *Money                 `json:"amount_money,omitempty"`
	TipMoney             *Money                 `json:"tip_money,omitempty"`
	ProcessingFeeMoney   *Money                 `json:"processing_fee_money,omitempty"`
	CustomerID           string                 `json:"customer_id,omitempty"`
	Type                 tenderType             `json:"-"`
	AdditionalRecipients []*AdditionalRecipient `json:"additional_recipients,omitempty"`
	PaymentID            string                 `json:"payment_id,omitempty"`
}

type TenderType string

const (
	tenderTypeCard           TenderType = "CARD"
	tenderTypeCash           TenderType = "CASH"
	tenderTypeThirdPartyCard TenderType = "THIRD_PARTY_CARD"
	tenderTypeSquareGiftCard TenderType = "SQUARE_GIFT_CARD"
	tenderTypeNoSale         TenderType = "NO_SALE"
	tenderTypeOther          TenderType = "OTHER"
)

func (t *Tender) MarshalJSON() ([]byte, error) {
	tJson := tender{
		tenderAlias: (*tenderAlias)(t),
	}

	switch details := t.Type.(type) {
	case *TenderCardDetails:
		tJson.Type = tenderTypeCard
		tJson.CardDetails = details
	case *TenderCashDetails:
		tJson.Type = tenderTypeCash
		tJson.CashDetails = details
	case *ThirdPartyCardDetails:
		tJson.Type = tenderTypeThirdPartyCard
	case *SquareGiftCardDetails:
		tJson.Type = tenderTypeSquareGiftCard
	case *NoSaleDetails:
		tJson.Type = tenderTypeNoSale
	case *OtherTenderTypeDetails:
		tJson.Type = tenderTypeOther
	default:
		return nil, errors.New("Found unknown tender type")
	}

	return json.Marshal(tJson)
}

func (t *Tender) UnmarshalJSON(b []byte) error {
	tJson := tender{
		tenderAlias: (*tenderAlias)(t),
	}
	err := json.Unmarshal(b, &tJson)
	if err != nil {
		return fmt.Errorf("Error unmarshaling Tender json: %w", err)
	}

	switch tJson.Type {
	case tenderTypeCard:
		t.Type = tJson.CardDetails
	case tenderTypeCash:
		t.Type = tJson.CashDetails
	case tenderTypeThirdPartyCard:
		t.Type = &ThirdPartyCardDetails{}
	case tenderTypeSquareGiftCard:
		t.Type = &SquareGiftCardDetails{}
	case tenderTypeNoSale:
		t.Type = &NoSaleDetails{}
	case tenderTypeOther:
		t.Type = &OtherTenderTypeDetails{}
	default:
		return errors.New("Unknown tender type found")
	}
	return nil
}
