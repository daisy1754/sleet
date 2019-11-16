package test

import (
	"github.com/BoltApp/sleet"
	"github.com/BoltApp/sleet/gateways/adyen"
	"github.com/Pallinder/go-randomdata"
	"testing"
)

func TestAdyenAuthorize(t *testing.T) {
	client := adyen.NewClient(getEnv("ADYEN_KEY"), getEnv("ADYEN_ACCOUNT"))
	amount := sleet.Amount{
		Amount:   100,
		Currency: "USD",
	}
	postalCode := "94103"
	address := sleet.BillingAddress{PostalCode: &postalCode}
	card := sleet.CreditCard{
		FirstName:       "Bolt",
		LastName:        "Checkout",
		Number:          "4111111111111111",
		ExpirationMonth: 10,
		ExpirationYear:  2020,
		CVV:             "737",
	}
	options := make(map[string]interface{})
	options2 := make(map[string]interface{})
	options["reference"] = randomdata.Letters(10) // so we don't collide with adyen
	options2["reference"] = randomdata.Letters(10)
	auth, _ := client.Authorize(&sleet.AuthorizationRequest{Amount: &amount, CreditCard: &card, BillingAddress: &address, Options: options})
	client.Capture(&sleet.CaptureRequest{Amount: &amount, TransactionReference: auth.TransactionReference})
	client.Refund(&sleet.RefundRequest{Amount: &amount, TransactionReference: auth.TransactionReference})
	auth2, _ := client.Authorize(&sleet.AuthorizationRequest{Amount: &amount, CreditCard: &card, BillingAddress: &address, Options: options2})
	client.Void(&sleet.VoidRequest{TransactionReference: auth2.TransactionReference})
}