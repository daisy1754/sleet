package authorizenet

type TransactionType string

const (
	TransactionTypeAuthOnly         TransactionType = "authOnlyTransaction"
	TransactionTypeVoid             TransactionType = "voidTransaction"
	TransactionTypePriorAuthCapture TransactionType = "priorAuthCaptureTransaction"
	TransactionTypeRefund           TransactionType = "refundTransaction"
)

type ResponseCode string

const (
	ResponseCodeApproved ResponseCode = "1"
	ResponseCodeDeclined ResponseCode = "2"
	ResponseCodeError    ResponseCode = "3"
	ResponseCodeHeld     ResponseCode = "4"
)

type ResultCode string

const (
	ResultCodeOK    ResultCode = "Ok"
	ResultCodeError ResultCode = "Error"
)

type AVSResultCode string

const (
	AVSResultNotPresent                AVSResultCode = "B"
	AVSResultError                     AVSResultCode = "E"
	AVSResultNotApplicable             AVSResultCode = "P" // AVS is not applicable for this transaction.
	AVSResultRetry                     AVSResultCode = "R" // AVS was unavailable or timed out.
	AVSResultNotSupportedIssuer        AVSResultCode = "S"
	AVSResultInfoUnavailable           AVSResultCode = "U" // Address information is unavailable.
	AVSResultPostMatchAddressMatch     AVSResultCode = "Y" // The street address and postal code matched.
	AVSResultNoMatch                   AVSResultCode = "N"
	AVSResultPostNoMatchAddressMatch   AVSResultCode = "A"
	AVSResultZipMatchAddressNoMatch    AVSResultCode = "W"
	AVSResultZipMatchAddressMatch      AVSResultCode = "X" // Both the street address and the US ZIP+4 code matched.
	AVSResultPostMatchAddressNoMatch   AVSResultCode = "Z"
	AVSResultNotSupportedInternational AVSResultCode = "G" // The card was issued by a bank outside the U.S. and does not support AVS.
)

type CVVResultCode string

const (
	CVVResultMatched         CVVResultCode = "M"
	CVVResultNoMatch         CVVResultCode = "N"
	CVVResultNotProcessed    CVVResultCode = "P"
	CVVResultNotPresent      CVVResultCode = "S"
	CVVResultUnableToProcess CVVResultCode = "U"
)

type CAVVResultCode string

const (
	CAVVResultBadRequest        CAVVResultCode = "0"
	CAVVResultFailed            CAVVResultCode = "1"
	CAVVResultPassed            CAVVResultCode = "2"
	CAVVResultAttemptIncomplete CAVVResultCode = "3"
	CAVVResultSytemError        CAVVResultCode = "4"

// omitted
// 5 -- N/A
// 6 -- N/A
// 7 -- CAVV failed validation, but the issuer is available. Valid for U.S.-issued card submitted to non-U.S acquirer.
// 8 -- CAVV passed validation and the issuer is available. Valid for U.S.-issued card submitted to non-U.S. acquirer.
// 9 -- CAVV failed validation and the issuer is unavailable. Valid for U.S.-issued card submitted to non-U.S acquirer.
// A -- CAVV passed validation but the issuer unavailable. Valid for U.S.-issued card submitted to non-U.S acquirer.
// B -- CAVV passed validation, information only, no liability shift.
)

const expirationDateXXXX = "XXXX"

// Request contains a createTransactionRequest for authorizations
type Request struct {
	CreateTransactionRequest CreateTransactionRequest `json:"createTransactionRequest"`
}

// CreateTransactionRequest specifies the merchant authentication to be used for request as well as transaction
// details specified in transactionRequest
type CreateTransactionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication"`
	RefID                  *string                `json:"refId,omitempty"`
	TransactionRequest     TransactionRequest     `json:"transactionRequest"`
}

// MerchantAuthentication is the name/key pair to authenticate Auth.net calls
type MerchantAuthentication struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

// TransactionRequest has the raw credit card info as Payment and amount to authorize
type TransactionRequest struct {
	TransactionType  TransactionType `json:"transactionType"`
	Amount           *string         `json:"amount,omitempty"`
	Payment          *Payment        `json:"payment,omitempty"`
	BillingAddress   *BillingAddress `json:"billTo,omitempty"`
	RefTransactionID *string         `json:"refTransId,omitempty"`
	// Ignoring Line items, Shipping, Tax, Duty, etc.
}

// Payment specifies the credit card to be authorized (only payment option for now)
type Payment struct {
	CreditCard CreditCard `json:"creditCard"`
}

// CreditCard is raw cc info
type CreditCard struct {
	CardNumber     string  `json:"cardNumber"`
	ExpirationDate string  `json:"expirationDate"`
	CardCode       *string `json:"cardCode,omitempty"`
}

// BillingAddress is used in TransactionRequest for making an auth call
type BillingAddress struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Company   string  `json:"company"`
	Address   *string `json:"address"`
	City      *string `json:"city"`
	State     *string `json:"state"`
	Zip       *string `json:"zip"`
	Country   *string `json:"country"`
}

// Response is a generic Auth.net response
type Response struct {
	TransactionResponse TransactionResponse `json:"transactionResponse"`
	RefID               string              `json:"refId"`
	Messsages           Messages            `json:"messages"`
}

// TransactionResponse contains the information from issuer about AVS, CVV and whether or not authorization was successful
type TransactionResponse struct {
	ResponseCode   ResponseCode                 `json:"responseCode"`
	AuthCode       string                       `json:"authCode"`
	AVSResultCode  AVSResultCode                `json:"avsResultCode"`
	CVVResultCode  CVVResultCode                `json:"cvvResultCode"`
	CAVVResultCode CAVVResultCode               `json:"cavvResultCode"`
	TransID        string                       `json:"transId"`
	RefTransID     string                       `json:"refTransID"`
	TransHash      string                       `json:"transHash"`
	AccountNumber  string                       `json:"accountNumber"`
	AccountType    string                       `json:"accountType"`
	Messages       []TransactionResponseMessage `json:"messages"`
	Errors         []Error                      `json:"errors"`
}

// TransactionResponseMessage contains additional information about transaction result from processor
type TransactionResponseMessage struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Error specifies a code and text explaining what happened
type Error struct {
	ErrorCode string `json:"errorCode"`
	ErrorText string `json:"errorText"`
}

// Messages is used to augment responses with codes and readable messages
type Messages struct {
	ResultCode ResultCode `json:"resultCode"`
	Message    []Message  `json:"message"`
}

// Message is similar to Error with code that maps to Auth.net internals and text for human readability
type Message struct {
	Code string `json:"code"`
	Text string `json:"text"`
}
