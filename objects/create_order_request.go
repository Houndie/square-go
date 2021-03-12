package objects

type CreateOrderRequest struct {
	Order          *Order `json:"order,omitempty"`
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
