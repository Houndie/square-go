package objects

import (
	"fmt"
)

type ErrorCategory string

const (
	ErrorCategoryAPIError            ErrorCategory = "API_ERROR"
	ErrorCategoryAuthenticationError ErrorCategory = "AUTHENTICATION_ERROR"
	ErrorCategoryInvalidRequestError ErrorCategory = "INVALID_REQUEST_ERROR"
	ErrorCategoryRateLimitError      ErrorCategory = "RATE_LIMIT_ERROR"
	ErrorCategoryPaymentMethodError  ErrorCategory = "PAYMENT_METHOD_ERROR"
	ErrorCategoryRefundError         ErrorCategory = "REFUND_ERROR"
)

type ErrorCode string

const (
	ErrorCodeInternalServerError                           ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrorCodeUnauthorized                                  ErrorCode = "UNAUTHORIZED"
	ErrorCodeAccessTokenExpired                            ErrorCode = "ACCESS_TOKEN_EXPIRED"
	ErrorCodeAccessTokenRevoked                            ErrorCode = "ACCESS_TOKEN_REVOKED"
	ErrorCodeForbidden                                     ErrorCode = "FORBIDDEN"
	ErrorCodeInsufficientScopes                            ErrorCode = "INSUFFICIENT_SCOPES"
	ErrorCodeApplicationDisabled                           ErrorCode = "APPLICATION_DISABLED"
	ErrorCodeV1Application                                 ErrorCode = "V1_APPLICATION"
	ErrorCodeV1AccessToken                                 ErrorCode = "V1_ACCESS_TOKEN"
	ErrorCodeCardProcessingNotEnabled                      ErrorCode = "CARD_PROCESSING_NOT_ENABLED"
	ErrorCodeBadRequest                                    ErrorCode = "BAD_REQUEST"
	ErrorCodeMissingRequiredParameter                      ErrorCode = "MISSING_REQUIRED_PARAMETER"
	ErrorCodeIncorrectType                                 ErrorCode = "INCORRECT_TYPE"
	ErrorCodeInvalidTime                                   ErrorCode = "INVALID_TIME"
	ErrorCodeInvalidTimeRange                              ErrorCode = "INVALID_TIME_RANGE"
	ErrorCodeInvalidValue                                  ErrorCode = "INVALID_VALUE"
	ErrorCodeInvalidCursor                                 ErrorCode = "INVALID_CURSOR"
	ErrorCodeUnknownQueryParameter                         ErrorCode = "UNKNOWN_QUERY_PARAMETER"
	ErrorCodeConflictingParameters                         ErrorCode = "CONFLICTING_PARAMETERS"
	ErrorCodeExpectedJSONBody                              ErrorCode = "EXPECTED_JSON_BODY"
	ErrorCodeInvalidSortOrder                              ErrorCode = "INVALID_SORT_ORDER"
	ErrorCodeValueRegexMismatch                            ErrorCode = "VALUE_REGEX_MISMATCH"
	ErrorCodeValueTooShort                                 ErrorCode = "VALUE_TOO_SHORT"
	ErrorCodeValueTooLong                                  ErrorCode = "VALUE_TOO_LONG"
	ErrorCodeValueTooLow                                   ErrorCode = "VALUE_TOO_LOW"
	ErrorCodeValueTooHigh                                  ErrorCode = "VALUE_TOO_HIGH"
	ErrorCodeValueEmpty                                    ErrorCode = "VALUE_EMPTY"
	ErrorCodeArrayLengthTooLong                            ErrorCode = "ARRAY_LENGTH_TOO_LONG"
	ErrorCodeArrayLengthTooShort                           ErrorCode = "ARRAY_LENGTH_TOO_SHORT"
	ErrorCodeArrayEmpty                                    ErrorCode = "ARRAY_EMPTY"
	ErrorCodeExpectedBoolean                               ErrorCode = "EXPECTED_BOOLEAN"
	ErrorCodeExpectedInteger                               ErrorCode = "EXPECTED_INTEGER"
	ErrorCodeExpectedFloat                                 ErrorCode = "EXPECTED_FLOAT"
	ErrorCodeExpectedString                                ErrorCode = "EXPECTED_STRING"
	ErrorCodeExpectedObject                                ErrorCode = "EXPECTED_OBJECT"
	ErrorCodeExpectedArray                                 ErrorCode = "EXPECTED_ARRAY"
	ErrorCodeExpectedMap                                   ErrorCode = "EXPECTED_MAP"
	ErrorCodeExpectedBase64EncodedByteArray                ErrorCode = "EXPECTED_BASE64_ENCODED_BYTE_ARRAY"
	ErrorCodeInvalidArrayValue                             ErrorCode = "INVALID_ARRAY_VALUE"
	ErrorCodeInvalidEnumValue                              ErrorCode = "INVALID_ENUM_VALUE"
	ErrorCodeInvalidContentType                            ErrorCode = "INVALID_CONTENT_TYPE"
	ErrorCodeInvalidFormValue                              ErrorCode = "INVALID_FORM_VALUE"
	ErrorCodeOneInstrumentExpected                         ErrorCode = "ONE_INSTRUMENT_EXPECTED"
	ErrorCodeNoFieldsSet                                   ErrorCode = "NO_FIELDS_SET"
	ErrorCodeDeprecatedFieldSet                            ErrorCode = "DEPRECATED_FIELD_SET"
	ErrorCodeRetiredFieldSet                               ErrorCode = "RETIRED_FIELD_SET"
	ErrorCodeCardExpired                                   ErrorCode = "CARD_EXPIRED"
	ErrorCodeInvalidExpiration                             ErrorCode = "INVALID_EXPIRATION"
	ErrorCodeInvalidExpirationYear                         ErrorCode = "INVALID_EXPIRATION_YEAR"
	ErrorCodeInvalidExpirationDate                         ErrorCode = "INVALID_EXPIRATION_DATE"
	ErrorCodeUnsupportedCardBrand                          ErrorCode = "UNSUPPORTED_CARD_BRAND"
	ErrorCodeUnsupportedEntryMethod                        ErrorCode = "UNSUPPORTED_ENTRY_METHOD"
	ErrorCodeInvalidEncryptedCard                          ErrorCode = "INVALID_ENCRYPTED_CARD"
	ErrorCodeInvalidCard                                   ErrorCode = "INVALID_CARD"
	ErrorCodeDelayedTransactionExpired                     ErrorCode = "DELAYED_TRANSACTION_EXPIRED"
	ErrorCodeDelayedTransactionCanceled                    ErrorCode = "DELAYED_TRANSACTION_CANCELED"
	ErrorCodeDelayedTransactionCaptured                    ErrorCode = "DELAYED_TRANSACTION_CAPTURED"
	ErrorCodeDelayedTransactionFailed                      ErrorCode = "DELAYED_TRANSACTION_FAILED"
	ErrorCodeCardTokenExpired                              ErrorCode = "CARD_TOKEN_EXPIRED"
	ErrorCodeCardTokenUsed                                 ErrorCode = "CARD_TOKEN_USED"
	ErrorCodeAmountTooHigh                                 ErrorCode = "AMOUNT_TOO_HIGH"
	ErrorCodeUnsupportedInstrumentType                     ErrorCode = "UNSUPPORTED_INSTRUMENT_TYPE"
	ErrorCodeRefundAmountInvalid                           ErrorCode = "REFUND_AMOUNT_INVALID"
	ErrorCodeRefundAlreadyPending                          ErrorCode = "REFUND_ALREADY_PENDING"
	ErrorCodePaymentNotRefundable                          ErrorCode = "PAYMENT_NOT_REFUNDABLE"
	ErrorCodeInvalidCardData                               ErrorCode = "INVALID_CARD_DATA"
	ErrorCodeLocationMismatch                              ErrorCode = "LOCATION_MISMATCH"
	ErrorCodeIdempotencyKeyReused                          ErrorCode = "IDEMPOTENCY_KEY_REUSED"
	ErrorCodeUnexpectedValue                               ErrorCode = "UNEXPECTED_VALUE"
	ErrorCodeSandboxNotSupported                           ErrorCode = "SANDBOX_NOT_SUPPORTED"
	ErrorCodeInvalidEmailAddress                           ErrorCode = "INVALID_EMAIL_ADDRESS"
	ErrorCodeInvalidPhoneNumber                            ErrorCode = "INVALID_PHONE_NUMBER"
	ErrorCodeCheckoutExpired                               ErrorCode = "CHECKOUT_EXPIRED"
	ErrorCodeBadCertificate                                ErrorCode = "BAD_CERTIFICATE"
	ErrorCodeInvalidSquareVersionFormat                    ErrorCode = "INVALID_SQUARE_VERSION_FORMAT"
	ErrorCodeAPIVersionIncompatible                        ErrorCode = "API_VERSION_INCOMPATIBLE"
	ErrorCodeCardDeclined                                  ErrorCode = "CARD_DECLINED"
	ErrorCodeVerifyCvvFailure                              ErrorCode = "VERIFY_CVV_FAILURE"
	ErrorCodeVerifyAvsFailure                              ErrorCode = "VERIFY_AVS_FAILURE"
	ErrorCodeCardDeclinedCallIssuer                        ErrorCode = "CARD_DECLINED_CALL_ISSUER"
	ErrorCodeNotFound                                      ErrorCode = "NOT_FOUND"
	ErrorCodeApplePaymentProcessingCertificateHashNotFound ErrorCode = "APPLE_PAYMENT_PROCESSING_CERTIFICATE_HASH_NOT_FOUND"
	ErrorCodeMethodNotAllowed                              ErrorCode = "METHOD_NOT_ALLOWED"
	ErrorCodeNotAcceptable                                 ErrorCode = "NOT_ACCEPTABLE"
	ErrorCodeRequestTimeout                                ErrorCode = "REQUEST_TIMEOUT"
	ErrorCodeConflict                                      ErrorCode = "CONFLICT"
	ErrorCodeRequestEntityTooLarge                         ErrorCode = "REQUEST_ENTITY_TOO_LARGE"
	ErrorCodeUnsupportedMediaType                          ErrorCode = "UNSUPPORTED_MEDIA_TYPE"
	ErrorCodeRateLimited                                   ErrorCode = "RATE_LIMITED"
	ErrorCodeNotImplemented                                ErrorCode = "NOT_IMPLEMENTED"
	ErrorCodeServiceUnavailable                            ErrorCode = "SERVICE_UNAVAILABLE"
	ErrorCodeGatewayTimeout                                ErrorCode = "GATEWAY_TIMEOUT"
)

type Error struct {
	Category ErrorCategory `json:"category,omitempty"`
	Code     ErrorCode     `json:"code,omitempty"`
	Detail   string        `json:"detail,omitempty"`
	Field    string        `json:"field,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Square Error: Category: %s; Code: %s; Detail: %s; Field: %s", e.Category, e.Code, e.Detail, e.Field)
}

type ErrorList struct {
	Errors []*Error
}

func (e *ErrorList) Error() string {
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}

	retVal := "Multiple Square Errors Returned:"
	for _, err := range e.Errors {
		retVal = fmt.Sprintf(retVal+" %s", err.Error())
	}

	return retVal
}

type UnexpectedCodeError int

func (u UnexpectedCodeError) Error() string {
	return fmt.Sprintf("found unxpected http error code %v", int(u))
}
