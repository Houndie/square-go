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
	tJSON := tender{
		tenderAlias: (*tenderAlias)(t),
	}

	switch details := t.Type.(type) {
	case *TenderCardDetails:
		tJSON.Type = tenderTypeCard
		tJSON.CardDetails = details
	case *TenderCashDetails:
		tJSON.Type = tenderTypeCash
		tJSON.CashDetails = details
	case *ThirdPartyCardDetails:
		tJSON.Type = tenderTypeThirdPartyCard
	case *SquareGiftCardDetails:
		tJSON.Type = tenderTypeSquareGiftCard
	case *NoSaleDetails:
		tJSON.Type = tenderTypeNoSale
	case *OtherTenderTypeDetails:
		tJSON.Type = tenderTypeOther
	default:
		return nil, errors.New("found unknown tender type")
	}

	j, err := json.Marshal(tJSON)
	if err != nil {
		return nil, fmt.Errorf("error marshaling tender: %w", err)
	}

	return j, nil
}

func (t *Tender) UnmarshalJSON(b []byte) error {
	tJSON := tender{
		tenderAlias: (*tenderAlias)(t),
	}
	if err := json.Unmarshal(b, &tJSON); err != nil {
		return fmt.Errorf("error unmarshaling Tender json: %w", err)
	}

	switch tJSON.Type {
	case tenderTypeCard:
		t.Type = tJSON.CardDetails
	case tenderTypeCash:
		t.Type = tJSON.CashDetails
	case tenderTypeThirdPartyCard:
		t.Type = &ThirdPartyCardDetails{}
	case tenderTypeSquareGiftCard:
		t.Type = &SquareGiftCardDetails{}
	case tenderTypeNoSale:
		t.Type = &NoSaleDetails{}
	case tenderTypeOther:
		t.Type = &OtherTenderTypeDetails{}
	default:
		return errors.New("unknown tender type found")
	}

	return nil
}
