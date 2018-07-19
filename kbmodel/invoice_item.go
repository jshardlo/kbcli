// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceItem invoice item
// swagger:model InvoiceItem
type InvoiceItem struct {

	// account Id
	// Required: true
	AccountID *strfmt.UUID `json:"accountId"`

	// amount
	Amount float64 `json:"amount,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// bundle Id
	BundleID strfmt.UUID `json:"bundleId,omitempty"`

	// child account Id
	ChildAccountID strfmt.UUID `json:"childAccountId,omitempty"`

	// child items
	ChildItems []*InvoiceItem `json:"childItems"`

	// currency
	Currency InvoiceItemCurrencyEnum `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end date
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// invoice Id
	InvoiceID strfmt.UUID `json:"invoiceId,omitempty"`

	// invoice item Id
	// Required: true
	InvoiceItemID *strfmt.UUID `json:"invoiceItemId"`

	// item details
	ItemDetails string `json:"itemDetails,omitempty"`

	// item type
	ItemType InvoiceItemItemTypeEnum `json:"itemType,omitempty"`

	// linked invoice item Id
	LinkedInvoiceItemID strfmt.UUID `json:"linkedInvoiceItemId,omitempty"`

	// phase name
	PhaseName string `json:"phaseName,omitempty"`

	// plan name
	PlanName string `json:"planName,omitempty"`

	// pretty phase name
	PrettyPhaseName string `json:"prettyPhaseName,omitempty"`

	// pretty plan name
	PrettyPlanName string `json:"prettyPlanName,omitempty"`

	// pretty product name
	PrettyProductName string `json:"prettyProductName,omitempty"`

	// pretty usage name
	PrettyUsageName string `json:"prettyUsageName,omitempty"`

	// product name
	ProductName string `json:"productName,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// rate
	Rate float64 `json:"rate,omitempty"`

	// start date
	StartDate strfmt.Date `json:"startDate,omitempty"`

	// subscription Id
	SubscriptionID strfmt.UUID `json:"subscriptionId,omitempty"`

	// usage name
	UsageName string `json:"usageName,omitempty"`
}

// Validate validates this invoice item
func (m *InvoiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBundleID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChildAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChildItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceItemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinkedInvoiceItemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceItem) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateAuditLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {

		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {

			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *InvoiceItem) validateBundleID(formats strfmt.Registry) error {

	if swag.IsZero(m.BundleID) { // not required
		return nil
	}

	if err := validate.FormatOf("bundleId", "body", "uuid", m.BundleID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateChildAccountID(formats strfmt.Registry) error {

	if swag.IsZero(m.ChildAccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("childAccountId", "body", "uuid", m.ChildAccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateChildItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ChildItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ChildItems); i++ {

		if swag.IsZero(m.ChildItems[i]) { // not required
			continue
		}

		if m.ChildItems[i] != nil {

			if err := m.ChildItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childItems" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

var invoiceItemTypeCurrencyPropEnum []interface{}

func init() {
	var res []InvoiceItemCurrencyEnum
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceItemTypeCurrencyPropEnum = append(invoiceItemTypeCurrencyPropEnum, v)
	}
}

type InvoiceItemCurrencyEnum string

const (

	// InvoiceItemCurrencyAED captures enum value "AED"
	InvoiceItemCurrencyAED InvoiceItemCurrencyEnum = "AED"

	// InvoiceItemCurrencyAFN captures enum value "AFN"
	InvoiceItemCurrencyAFN InvoiceItemCurrencyEnum = "AFN"

	// InvoiceItemCurrencyALL captures enum value "ALL"
	InvoiceItemCurrencyALL InvoiceItemCurrencyEnum = "ALL"

	// InvoiceItemCurrencyAMD captures enum value "AMD"
	InvoiceItemCurrencyAMD InvoiceItemCurrencyEnum = "AMD"

	// InvoiceItemCurrencyANG captures enum value "ANG"
	InvoiceItemCurrencyANG InvoiceItemCurrencyEnum = "ANG"

	// InvoiceItemCurrencyAOA captures enum value "AOA"
	InvoiceItemCurrencyAOA InvoiceItemCurrencyEnum = "AOA"

	// InvoiceItemCurrencyARS captures enum value "ARS"
	InvoiceItemCurrencyARS InvoiceItemCurrencyEnum = "ARS"

	// InvoiceItemCurrencyAUD captures enum value "AUD"
	InvoiceItemCurrencyAUD InvoiceItemCurrencyEnum = "AUD"

	// InvoiceItemCurrencyAWG captures enum value "AWG"
	InvoiceItemCurrencyAWG InvoiceItemCurrencyEnum = "AWG"

	// InvoiceItemCurrencyAZN captures enum value "AZN"
	InvoiceItemCurrencyAZN InvoiceItemCurrencyEnum = "AZN"

	// InvoiceItemCurrencyBAM captures enum value "BAM"
	InvoiceItemCurrencyBAM InvoiceItemCurrencyEnum = "BAM"

	// InvoiceItemCurrencyBBD captures enum value "BBD"
	InvoiceItemCurrencyBBD InvoiceItemCurrencyEnum = "BBD"

	// InvoiceItemCurrencyBDT captures enum value "BDT"
	InvoiceItemCurrencyBDT InvoiceItemCurrencyEnum = "BDT"

	// InvoiceItemCurrencyBGN captures enum value "BGN"
	InvoiceItemCurrencyBGN InvoiceItemCurrencyEnum = "BGN"

	// InvoiceItemCurrencyBHD captures enum value "BHD"
	InvoiceItemCurrencyBHD InvoiceItemCurrencyEnum = "BHD"

	// InvoiceItemCurrencyBIF captures enum value "BIF"
	InvoiceItemCurrencyBIF InvoiceItemCurrencyEnum = "BIF"

	// InvoiceItemCurrencyBMD captures enum value "BMD"
	InvoiceItemCurrencyBMD InvoiceItemCurrencyEnum = "BMD"

	// InvoiceItemCurrencyBND captures enum value "BND"
	InvoiceItemCurrencyBND InvoiceItemCurrencyEnum = "BND"

	// InvoiceItemCurrencyBOB captures enum value "BOB"
	InvoiceItemCurrencyBOB InvoiceItemCurrencyEnum = "BOB"

	// InvoiceItemCurrencyBRL captures enum value "BRL"
	InvoiceItemCurrencyBRL InvoiceItemCurrencyEnum = "BRL"

	// InvoiceItemCurrencyBSD captures enum value "BSD"
	InvoiceItemCurrencyBSD InvoiceItemCurrencyEnum = "BSD"

	// InvoiceItemCurrencyBTN captures enum value "BTN"
	InvoiceItemCurrencyBTN InvoiceItemCurrencyEnum = "BTN"

	// InvoiceItemCurrencyBWP captures enum value "BWP"
	InvoiceItemCurrencyBWP InvoiceItemCurrencyEnum = "BWP"

	// InvoiceItemCurrencyBYR captures enum value "BYR"
	InvoiceItemCurrencyBYR InvoiceItemCurrencyEnum = "BYR"

	// InvoiceItemCurrencyBZD captures enum value "BZD"
	InvoiceItemCurrencyBZD InvoiceItemCurrencyEnum = "BZD"

	// InvoiceItemCurrencyCAD captures enum value "CAD"
	InvoiceItemCurrencyCAD InvoiceItemCurrencyEnum = "CAD"

	// InvoiceItemCurrencyCDF captures enum value "CDF"
	InvoiceItemCurrencyCDF InvoiceItemCurrencyEnum = "CDF"

	// InvoiceItemCurrencyCHF captures enum value "CHF"
	InvoiceItemCurrencyCHF InvoiceItemCurrencyEnum = "CHF"

	// InvoiceItemCurrencyCLP captures enum value "CLP"
	InvoiceItemCurrencyCLP InvoiceItemCurrencyEnum = "CLP"

	// InvoiceItemCurrencyCNY captures enum value "CNY"
	InvoiceItemCurrencyCNY InvoiceItemCurrencyEnum = "CNY"

	// InvoiceItemCurrencyCOP captures enum value "COP"
	InvoiceItemCurrencyCOP InvoiceItemCurrencyEnum = "COP"

	// InvoiceItemCurrencyCRC captures enum value "CRC"
	InvoiceItemCurrencyCRC InvoiceItemCurrencyEnum = "CRC"

	// InvoiceItemCurrencyCUC captures enum value "CUC"
	InvoiceItemCurrencyCUC InvoiceItemCurrencyEnum = "CUC"

	// InvoiceItemCurrencyCUP captures enum value "CUP"
	InvoiceItemCurrencyCUP InvoiceItemCurrencyEnum = "CUP"

	// InvoiceItemCurrencyCVE captures enum value "CVE"
	InvoiceItemCurrencyCVE InvoiceItemCurrencyEnum = "CVE"

	// InvoiceItemCurrencyCZK captures enum value "CZK"
	InvoiceItemCurrencyCZK InvoiceItemCurrencyEnum = "CZK"

	// InvoiceItemCurrencyDJF captures enum value "DJF"
	InvoiceItemCurrencyDJF InvoiceItemCurrencyEnum = "DJF"

	// InvoiceItemCurrencyDKK captures enum value "DKK"
	InvoiceItemCurrencyDKK InvoiceItemCurrencyEnum = "DKK"

	// InvoiceItemCurrencyDOP captures enum value "DOP"
	InvoiceItemCurrencyDOP InvoiceItemCurrencyEnum = "DOP"

	// InvoiceItemCurrencyDZD captures enum value "DZD"
	InvoiceItemCurrencyDZD InvoiceItemCurrencyEnum = "DZD"

	// InvoiceItemCurrencyEGP captures enum value "EGP"
	InvoiceItemCurrencyEGP InvoiceItemCurrencyEnum = "EGP"

	// InvoiceItemCurrencyERN captures enum value "ERN"
	InvoiceItemCurrencyERN InvoiceItemCurrencyEnum = "ERN"

	// InvoiceItemCurrencyETB captures enum value "ETB"
	InvoiceItemCurrencyETB InvoiceItemCurrencyEnum = "ETB"

	// InvoiceItemCurrencyEUR captures enum value "EUR"
	InvoiceItemCurrencyEUR InvoiceItemCurrencyEnum = "EUR"

	// InvoiceItemCurrencyFJD captures enum value "FJD"
	InvoiceItemCurrencyFJD InvoiceItemCurrencyEnum = "FJD"

	// InvoiceItemCurrencyFKP captures enum value "FKP"
	InvoiceItemCurrencyFKP InvoiceItemCurrencyEnum = "FKP"

	// InvoiceItemCurrencyGBP captures enum value "GBP"
	InvoiceItemCurrencyGBP InvoiceItemCurrencyEnum = "GBP"

	// InvoiceItemCurrencyGEL captures enum value "GEL"
	InvoiceItemCurrencyGEL InvoiceItemCurrencyEnum = "GEL"

	// InvoiceItemCurrencyGGP captures enum value "GGP"
	InvoiceItemCurrencyGGP InvoiceItemCurrencyEnum = "GGP"

	// InvoiceItemCurrencyGHS captures enum value "GHS"
	InvoiceItemCurrencyGHS InvoiceItemCurrencyEnum = "GHS"

	// InvoiceItemCurrencyGIP captures enum value "GIP"
	InvoiceItemCurrencyGIP InvoiceItemCurrencyEnum = "GIP"

	// InvoiceItemCurrencyGMD captures enum value "GMD"
	InvoiceItemCurrencyGMD InvoiceItemCurrencyEnum = "GMD"

	// InvoiceItemCurrencyGNF captures enum value "GNF"
	InvoiceItemCurrencyGNF InvoiceItemCurrencyEnum = "GNF"

	// InvoiceItemCurrencyGTQ captures enum value "GTQ"
	InvoiceItemCurrencyGTQ InvoiceItemCurrencyEnum = "GTQ"

	// InvoiceItemCurrencyGYD captures enum value "GYD"
	InvoiceItemCurrencyGYD InvoiceItemCurrencyEnum = "GYD"

	// InvoiceItemCurrencyHKD captures enum value "HKD"
	InvoiceItemCurrencyHKD InvoiceItemCurrencyEnum = "HKD"

	// InvoiceItemCurrencyHNL captures enum value "HNL"
	InvoiceItemCurrencyHNL InvoiceItemCurrencyEnum = "HNL"

	// InvoiceItemCurrencyHRK captures enum value "HRK"
	InvoiceItemCurrencyHRK InvoiceItemCurrencyEnum = "HRK"

	// InvoiceItemCurrencyHTG captures enum value "HTG"
	InvoiceItemCurrencyHTG InvoiceItemCurrencyEnum = "HTG"

	// InvoiceItemCurrencyHUF captures enum value "HUF"
	InvoiceItemCurrencyHUF InvoiceItemCurrencyEnum = "HUF"

	// InvoiceItemCurrencyIDR captures enum value "IDR"
	InvoiceItemCurrencyIDR InvoiceItemCurrencyEnum = "IDR"

	// InvoiceItemCurrencyILS captures enum value "ILS"
	InvoiceItemCurrencyILS InvoiceItemCurrencyEnum = "ILS"

	// InvoiceItemCurrencyIMP captures enum value "IMP"
	InvoiceItemCurrencyIMP InvoiceItemCurrencyEnum = "IMP"

	// InvoiceItemCurrencyINR captures enum value "INR"
	InvoiceItemCurrencyINR InvoiceItemCurrencyEnum = "INR"

	// InvoiceItemCurrencyIQD captures enum value "IQD"
	InvoiceItemCurrencyIQD InvoiceItemCurrencyEnum = "IQD"

	// InvoiceItemCurrencyIRR captures enum value "IRR"
	InvoiceItemCurrencyIRR InvoiceItemCurrencyEnum = "IRR"

	// InvoiceItemCurrencyISK captures enum value "ISK"
	InvoiceItemCurrencyISK InvoiceItemCurrencyEnum = "ISK"

	// InvoiceItemCurrencyJEP captures enum value "JEP"
	InvoiceItemCurrencyJEP InvoiceItemCurrencyEnum = "JEP"

	// InvoiceItemCurrencyJMD captures enum value "JMD"
	InvoiceItemCurrencyJMD InvoiceItemCurrencyEnum = "JMD"

	// InvoiceItemCurrencyJOD captures enum value "JOD"
	InvoiceItemCurrencyJOD InvoiceItemCurrencyEnum = "JOD"

	// InvoiceItemCurrencyJPY captures enum value "JPY"
	InvoiceItemCurrencyJPY InvoiceItemCurrencyEnum = "JPY"

	// InvoiceItemCurrencyKES captures enum value "KES"
	InvoiceItemCurrencyKES InvoiceItemCurrencyEnum = "KES"

	// InvoiceItemCurrencyKGS captures enum value "KGS"
	InvoiceItemCurrencyKGS InvoiceItemCurrencyEnum = "KGS"

	// InvoiceItemCurrencyKHR captures enum value "KHR"
	InvoiceItemCurrencyKHR InvoiceItemCurrencyEnum = "KHR"

	// InvoiceItemCurrencyKMF captures enum value "KMF"
	InvoiceItemCurrencyKMF InvoiceItemCurrencyEnum = "KMF"

	// InvoiceItemCurrencyKPW captures enum value "KPW"
	InvoiceItemCurrencyKPW InvoiceItemCurrencyEnum = "KPW"

	// InvoiceItemCurrencyKRW captures enum value "KRW"
	InvoiceItemCurrencyKRW InvoiceItemCurrencyEnum = "KRW"

	// InvoiceItemCurrencyKWD captures enum value "KWD"
	InvoiceItemCurrencyKWD InvoiceItemCurrencyEnum = "KWD"

	// InvoiceItemCurrencyKYD captures enum value "KYD"
	InvoiceItemCurrencyKYD InvoiceItemCurrencyEnum = "KYD"

	// InvoiceItemCurrencyKZT captures enum value "KZT"
	InvoiceItemCurrencyKZT InvoiceItemCurrencyEnum = "KZT"

	// InvoiceItemCurrencyLAK captures enum value "LAK"
	InvoiceItemCurrencyLAK InvoiceItemCurrencyEnum = "LAK"

	// InvoiceItemCurrencyLBP captures enum value "LBP"
	InvoiceItemCurrencyLBP InvoiceItemCurrencyEnum = "LBP"

	// InvoiceItemCurrencyLKR captures enum value "LKR"
	InvoiceItemCurrencyLKR InvoiceItemCurrencyEnum = "LKR"

	// InvoiceItemCurrencyLRD captures enum value "LRD"
	InvoiceItemCurrencyLRD InvoiceItemCurrencyEnum = "LRD"

	// InvoiceItemCurrencyLSL captures enum value "LSL"
	InvoiceItemCurrencyLSL InvoiceItemCurrencyEnum = "LSL"

	// InvoiceItemCurrencyLTL captures enum value "LTL"
	InvoiceItemCurrencyLTL InvoiceItemCurrencyEnum = "LTL"

	// InvoiceItemCurrencyLVL captures enum value "LVL"
	InvoiceItemCurrencyLVL InvoiceItemCurrencyEnum = "LVL"

	// InvoiceItemCurrencyLYD captures enum value "LYD"
	InvoiceItemCurrencyLYD InvoiceItemCurrencyEnum = "LYD"

	// InvoiceItemCurrencyMAD captures enum value "MAD"
	InvoiceItemCurrencyMAD InvoiceItemCurrencyEnum = "MAD"

	// InvoiceItemCurrencyMDL captures enum value "MDL"
	InvoiceItemCurrencyMDL InvoiceItemCurrencyEnum = "MDL"

	// InvoiceItemCurrencyMGA captures enum value "MGA"
	InvoiceItemCurrencyMGA InvoiceItemCurrencyEnum = "MGA"

	// InvoiceItemCurrencyMKD captures enum value "MKD"
	InvoiceItemCurrencyMKD InvoiceItemCurrencyEnum = "MKD"

	// InvoiceItemCurrencyMMK captures enum value "MMK"
	InvoiceItemCurrencyMMK InvoiceItemCurrencyEnum = "MMK"

	// InvoiceItemCurrencyMNT captures enum value "MNT"
	InvoiceItemCurrencyMNT InvoiceItemCurrencyEnum = "MNT"

	// InvoiceItemCurrencyMOP captures enum value "MOP"
	InvoiceItemCurrencyMOP InvoiceItemCurrencyEnum = "MOP"

	// InvoiceItemCurrencyMRO captures enum value "MRO"
	InvoiceItemCurrencyMRO InvoiceItemCurrencyEnum = "MRO"

	// InvoiceItemCurrencyMUR captures enum value "MUR"
	InvoiceItemCurrencyMUR InvoiceItemCurrencyEnum = "MUR"

	// InvoiceItemCurrencyMVR captures enum value "MVR"
	InvoiceItemCurrencyMVR InvoiceItemCurrencyEnum = "MVR"

	// InvoiceItemCurrencyMWK captures enum value "MWK"
	InvoiceItemCurrencyMWK InvoiceItemCurrencyEnum = "MWK"

	// InvoiceItemCurrencyMXN captures enum value "MXN"
	InvoiceItemCurrencyMXN InvoiceItemCurrencyEnum = "MXN"

	// InvoiceItemCurrencyMYR captures enum value "MYR"
	InvoiceItemCurrencyMYR InvoiceItemCurrencyEnum = "MYR"

	// InvoiceItemCurrencyMZN captures enum value "MZN"
	InvoiceItemCurrencyMZN InvoiceItemCurrencyEnum = "MZN"

	// InvoiceItemCurrencyNAD captures enum value "NAD"
	InvoiceItemCurrencyNAD InvoiceItemCurrencyEnum = "NAD"

	// InvoiceItemCurrencyNGN captures enum value "NGN"
	InvoiceItemCurrencyNGN InvoiceItemCurrencyEnum = "NGN"

	// InvoiceItemCurrencyNIO captures enum value "NIO"
	InvoiceItemCurrencyNIO InvoiceItemCurrencyEnum = "NIO"

	// InvoiceItemCurrencyNOK captures enum value "NOK"
	InvoiceItemCurrencyNOK InvoiceItemCurrencyEnum = "NOK"

	// InvoiceItemCurrencyNPR captures enum value "NPR"
	InvoiceItemCurrencyNPR InvoiceItemCurrencyEnum = "NPR"

	// InvoiceItemCurrencyNZD captures enum value "NZD"
	InvoiceItemCurrencyNZD InvoiceItemCurrencyEnum = "NZD"

	// InvoiceItemCurrencyOMR captures enum value "OMR"
	InvoiceItemCurrencyOMR InvoiceItemCurrencyEnum = "OMR"

	// InvoiceItemCurrencyPAB captures enum value "PAB"
	InvoiceItemCurrencyPAB InvoiceItemCurrencyEnum = "PAB"

	// InvoiceItemCurrencyPEN captures enum value "PEN"
	InvoiceItemCurrencyPEN InvoiceItemCurrencyEnum = "PEN"

	// InvoiceItemCurrencyPGK captures enum value "PGK"
	InvoiceItemCurrencyPGK InvoiceItemCurrencyEnum = "PGK"

	// InvoiceItemCurrencyPHP captures enum value "PHP"
	InvoiceItemCurrencyPHP InvoiceItemCurrencyEnum = "PHP"

	// InvoiceItemCurrencyPKR captures enum value "PKR"
	InvoiceItemCurrencyPKR InvoiceItemCurrencyEnum = "PKR"

	// InvoiceItemCurrencyPLN captures enum value "PLN"
	InvoiceItemCurrencyPLN InvoiceItemCurrencyEnum = "PLN"

	// InvoiceItemCurrencyPYG captures enum value "PYG"
	InvoiceItemCurrencyPYG InvoiceItemCurrencyEnum = "PYG"

	// InvoiceItemCurrencyQAR captures enum value "QAR"
	InvoiceItemCurrencyQAR InvoiceItemCurrencyEnum = "QAR"

	// InvoiceItemCurrencyRON captures enum value "RON"
	InvoiceItemCurrencyRON InvoiceItemCurrencyEnum = "RON"

	// InvoiceItemCurrencyRSD captures enum value "RSD"
	InvoiceItemCurrencyRSD InvoiceItemCurrencyEnum = "RSD"

	// InvoiceItemCurrencyRUB captures enum value "RUB"
	InvoiceItemCurrencyRUB InvoiceItemCurrencyEnum = "RUB"

	// InvoiceItemCurrencyRWF captures enum value "RWF"
	InvoiceItemCurrencyRWF InvoiceItemCurrencyEnum = "RWF"

	// InvoiceItemCurrencySAR captures enum value "SAR"
	InvoiceItemCurrencySAR InvoiceItemCurrencyEnum = "SAR"

	// InvoiceItemCurrencySBD captures enum value "SBD"
	InvoiceItemCurrencySBD InvoiceItemCurrencyEnum = "SBD"

	// InvoiceItemCurrencySCR captures enum value "SCR"
	InvoiceItemCurrencySCR InvoiceItemCurrencyEnum = "SCR"

	// InvoiceItemCurrencySDG captures enum value "SDG"
	InvoiceItemCurrencySDG InvoiceItemCurrencyEnum = "SDG"

	// InvoiceItemCurrencySEK captures enum value "SEK"
	InvoiceItemCurrencySEK InvoiceItemCurrencyEnum = "SEK"

	// InvoiceItemCurrencySGD captures enum value "SGD"
	InvoiceItemCurrencySGD InvoiceItemCurrencyEnum = "SGD"

	// InvoiceItemCurrencySHP captures enum value "SHP"
	InvoiceItemCurrencySHP InvoiceItemCurrencyEnum = "SHP"

	// InvoiceItemCurrencySLL captures enum value "SLL"
	InvoiceItemCurrencySLL InvoiceItemCurrencyEnum = "SLL"

	// InvoiceItemCurrencySOS captures enum value "SOS"
	InvoiceItemCurrencySOS InvoiceItemCurrencyEnum = "SOS"

	// InvoiceItemCurrencySPL captures enum value "SPL"
	InvoiceItemCurrencySPL InvoiceItemCurrencyEnum = "SPL"

	// InvoiceItemCurrencySRD captures enum value "SRD"
	InvoiceItemCurrencySRD InvoiceItemCurrencyEnum = "SRD"

	// InvoiceItemCurrencySTD captures enum value "STD"
	InvoiceItemCurrencySTD InvoiceItemCurrencyEnum = "STD"

	// InvoiceItemCurrencySVC captures enum value "SVC"
	InvoiceItemCurrencySVC InvoiceItemCurrencyEnum = "SVC"

	// InvoiceItemCurrencySYP captures enum value "SYP"
	InvoiceItemCurrencySYP InvoiceItemCurrencyEnum = "SYP"

	// InvoiceItemCurrencySZL captures enum value "SZL"
	InvoiceItemCurrencySZL InvoiceItemCurrencyEnum = "SZL"

	// InvoiceItemCurrencyTHB captures enum value "THB"
	InvoiceItemCurrencyTHB InvoiceItemCurrencyEnum = "THB"

	// InvoiceItemCurrencyTJS captures enum value "TJS"
	InvoiceItemCurrencyTJS InvoiceItemCurrencyEnum = "TJS"

	// InvoiceItemCurrencyTMT captures enum value "TMT"
	InvoiceItemCurrencyTMT InvoiceItemCurrencyEnum = "TMT"

	// InvoiceItemCurrencyTND captures enum value "TND"
	InvoiceItemCurrencyTND InvoiceItemCurrencyEnum = "TND"

	// InvoiceItemCurrencyTOP captures enum value "TOP"
	InvoiceItemCurrencyTOP InvoiceItemCurrencyEnum = "TOP"

	// InvoiceItemCurrencyTRY captures enum value "TRY"
	InvoiceItemCurrencyTRY InvoiceItemCurrencyEnum = "TRY"

	// InvoiceItemCurrencyTTD captures enum value "TTD"
	InvoiceItemCurrencyTTD InvoiceItemCurrencyEnum = "TTD"

	// InvoiceItemCurrencyTVD captures enum value "TVD"
	InvoiceItemCurrencyTVD InvoiceItemCurrencyEnum = "TVD"

	// InvoiceItemCurrencyTWD captures enum value "TWD"
	InvoiceItemCurrencyTWD InvoiceItemCurrencyEnum = "TWD"

	// InvoiceItemCurrencyTZS captures enum value "TZS"
	InvoiceItemCurrencyTZS InvoiceItemCurrencyEnum = "TZS"

	// InvoiceItemCurrencyUAH captures enum value "UAH"
	InvoiceItemCurrencyUAH InvoiceItemCurrencyEnum = "UAH"

	// InvoiceItemCurrencyUGX captures enum value "UGX"
	InvoiceItemCurrencyUGX InvoiceItemCurrencyEnum = "UGX"

	// InvoiceItemCurrencyUSD captures enum value "USD"
	InvoiceItemCurrencyUSD InvoiceItemCurrencyEnum = "USD"

	// InvoiceItemCurrencyUYU captures enum value "UYU"
	InvoiceItemCurrencyUYU InvoiceItemCurrencyEnum = "UYU"

	// InvoiceItemCurrencyUZS captures enum value "UZS"
	InvoiceItemCurrencyUZS InvoiceItemCurrencyEnum = "UZS"

	// InvoiceItemCurrencyVEF captures enum value "VEF"
	InvoiceItemCurrencyVEF InvoiceItemCurrencyEnum = "VEF"

	// InvoiceItemCurrencyVND captures enum value "VND"
	InvoiceItemCurrencyVND InvoiceItemCurrencyEnum = "VND"

	// InvoiceItemCurrencyVUV captures enum value "VUV"
	InvoiceItemCurrencyVUV InvoiceItemCurrencyEnum = "VUV"

	// InvoiceItemCurrencyWST captures enum value "WST"
	InvoiceItemCurrencyWST InvoiceItemCurrencyEnum = "WST"

	// InvoiceItemCurrencyXAF captures enum value "XAF"
	InvoiceItemCurrencyXAF InvoiceItemCurrencyEnum = "XAF"

	// InvoiceItemCurrencyXCD captures enum value "XCD"
	InvoiceItemCurrencyXCD InvoiceItemCurrencyEnum = "XCD"

	// InvoiceItemCurrencyXDR captures enum value "XDR"
	InvoiceItemCurrencyXDR InvoiceItemCurrencyEnum = "XDR"

	// InvoiceItemCurrencyXOF captures enum value "XOF"
	InvoiceItemCurrencyXOF InvoiceItemCurrencyEnum = "XOF"

	// InvoiceItemCurrencyXPF captures enum value "XPF"
	InvoiceItemCurrencyXPF InvoiceItemCurrencyEnum = "XPF"

	// InvoiceItemCurrencyYER captures enum value "YER"
	InvoiceItemCurrencyYER InvoiceItemCurrencyEnum = "YER"

	// InvoiceItemCurrencyZAR captures enum value "ZAR"
	InvoiceItemCurrencyZAR InvoiceItemCurrencyEnum = "ZAR"

	// InvoiceItemCurrencyZMW captures enum value "ZMW"
	InvoiceItemCurrencyZMW InvoiceItemCurrencyEnum = "ZMW"

	// InvoiceItemCurrencyZWD captures enum value "ZWD"
	InvoiceItemCurrencyZWD InvoiceItemCurrencyEnum = "ZWD"

	// InvoiceItemCurrencyBTC captures enum value "BTC"
	InvoiceItemCurrencyBTC InvoiceItemCurrencyEnum = "BTC"
)

var InvoiceItemCurrencyEnumValues = []string{
	"AED",
	"AFN",
	"ALL",
	"AMD",
	"ANG",
	"AOA",
	"ARS",
	"AUD",
	"AWG",
	"AZN",
	"BAM",
	"BBD",
	"BDT",
	"BGN",
	"BHD",
	"BIF",
	"BMD",
	"BND",
	"BOB",
	"BRL",
	"BSD",
	"BTN",
	"BWP",
	"BYR",
	"BZD",
	"CAD",
	"CDF",
	"CHF",
	"CLP",
	"CNY",
	"COP",
	"CRC",
	"CUC",
	"CUP",
	"CVE",
	"CZK",
	"DJF",
	"DKK",
	"DOP",
	"DZD",
	"EGP",
	"ERN",
	"ETB",
	"EUR",
	"FJD",
	"FKP",
	"GBP",
	"GEL",
	"GGP",
	"GHS",
	"GIP",
	"GMD",
	"GNF",
	"GTQ",
	"GYD",
	"HKD",
	"HNL",
	"HRK",
	"HTG",
	"HUF",
	"IDR",
	"ILS",
	"IMP",
	"INR",
	"IQD",
	"IRR",
	"ISK",
	"JEP",
	"JMD",
	"JOD",
	"JPY",
	"KES",
	"KGS",
	"KHR",
	"KMF",
	"KPW",
	"KRW",
	"KWD",
	"KYD",
	"KZT",
	"LAK",
	"LBP",
	"LKR",
	"LRD",
	"LSL",
	"LTL",
	"LVL",
	"LYD",
	"MAD",
	"MDL",
	"MGA",
	"MKD",
	"MMK",
	"MNT",
	"MOP",
	"MRO",
	"MUR",
	"MVR",
	"MWK",
	"MXN",
	"MYR",
	"MZN",
	"NAD",
	"NGN",
	"NIO",
	"NOK",
	"NPR",
	"NZD",
	"OMR",
	"PAB",
	"PEN",
	"PGK",
	"PHP",
	"PKR",
	"PLN",
	"PYG",
	"QAR",
	"RON",
	"RSD",
	"RUB",
	"RWF",
	"SAR",
	"SBD",
	"SCR",
	"SDG",
	"SEK",
	"SGD",
	"SHP",
	"SLL",
	"SOS",
	"SPL",
	"SRD",
	"STD",
	"SVC",
	"SYP",
	"SZL",
	"THB",
	"TJS",
	"TMT",
	"TND",
	"TOP",
	"TRY",
	"TTD",
	"TVD",
	"TWD",
	"TZS",
	"UAH",
	"UGX",
	"USD",
	"UYU",
	"UZS",
	"VEF",
	"VND",
	"VUV",
	"WST",
	"XAF",
	"XCD",
	"XDR",
	"XOF",
	"XPF",
	"YER",
	"ZAR",
	"ZMW",
	"ZWD",
	"BTC",
}

func (e InvoiceItemCurrencyEnum) IsValid() bool {
	for _, v := range InvoiceItemCurrencyEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *InvoiceItem) validateCurrencyEnum(path, location string, value InvoiceItemCurrencyEnum) error {
	if err := validate.Enum(path, location, value, invoiceItemTypeCurrencyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceItem) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateInvoiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("invoiceId", "body", "uuid", m.InvoiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateInvoiceItemID(formats strfmt.Registry) error {

	if err := validate.Required("invoiceItemId", "body", m.InvoiceItemID); err != nil {
		return err
	}

	if err := validate.FormatOf("invoiceItemId", "body", "uuid", m.InvoiceItemID.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceItemTypeItemTypePropEnum []interface{}

func init() {
	var res []InvoiceItemItemTypeEnum
	if err := json.Unmarshal([]byte(`["EXTERNAL_CHARGE","FIXED","RECURRING","REPAIR_ADJ","CBA_ADJ","CREDIT_ADJ","ITEM_ADJ","USAGE","TAX","PARENT_SUMMARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceItemTypeItemTypePropEnum = append(invoiceItemTypeItemTypePropEnum, v)
	}
}

type InvoiceItemItemTypeEnum string

const (

	// InvoiceItemItemTypeEXTERNALCHARGE captures enum value "EXTERNAL_CHARGE"
	InvoiceItemItemTypeEXTERNALCHARGE InvoiceItemItemTypeEnum = "EXTERNAL_CHARGE"

	// InvoiceItemItemTypeFIXED captures enum value "FIXED"
	InvoiceItemItemTypeFIXED InvoiceItemItemTypeEnum = "FIXED"

	// InvoiceItemItemTypeRECURRING captures enum value "RECURRING"
	InvoiceItemItemTypeRECURRING InvoiceItemItemTypeEnum = "RECURRING"

	// InvoiceItemItemTypeREPAIRADJ captures enum value "REPAIR_ADJ"
	InvoiceItemItemTypeREPAIRADJ InvoiceItemItemTypeEnum = "REPAIR_ADJ"

	// InvoiceItemItemTypeCBAADJ captures enum value "CBA_ADJ"
	InvoiceItemItemTypeCBAADJ InvoiceItemItemTypeEnum = "CBA_ADJ"

	// InvoiceItemItemTypeCREDITADJ captures enum value "CREDIT_ADJ"
	InvoiceItemItemTypeCREDITADJ InvoiceItemItemTypeEnum = "CREDIT_ADJ"

	// InvoiceItemItemTypeITEMADJ captures enum value "ITEM_ADJ"
	InvoiceItemItemTypeITEMADJ InvoiceItemItemTypeEnum = "ITEM_ADJ"

	// InvoiceItemItemTypeUSAGE captures enum value "USAGE"
	InvoiceItemItemTypeUSAGE InvoiceItemItemTypeEnum = "USAGE"

	// InvoiceItemItemTypeTAX captures enum value "TAX"
	InvoiceItemItemTypeTAX InvoiceItemItemTypeEnum = "TAX"

	// InvoiceItemItemTypePARENTSUMMARY captures enum value "PARENT_SUMMARY"
	InvoiceItemItemTypePARENTSUMMARY InvoiceItemItemTypeEnum = "PARENT_SUMMARY"
)

var InvoiceItemItemTypeEnumValues = []string{
	"EXTERNAL_CHARGE",
	"FIXED",
	"RECURRING",
	"REPAIR_ADJ",
	"CBA_ADJ",
	"CREDIT_ADJ",
	"ITEM_ADJ",
	"USAGE",
	"TAX",
	"PARENT_SUMMARY",
}

func (e InvoiceItemItemTypeEnum) IsValid() bool {
	for _, v := range InvoiceItemItemTypeEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *InvoiceItem) validateItemTypeEnum(path, location string, value InvoiceItemItemTypeEnum) error {
	if err := validate.Enum(path, location, value, invoiceItemTypeItemTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceItem) validateItemType(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateLinkedInvoiceItemID(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedInvoiceItemID) { // not required
		return nil
	}

	if err := validate.FormatOf("linkedInvoiceItemId", "body", "uuid", m.LinkedInvoiceItemID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateSubscriptionID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionID) { // not required
		return nil
	}

	if err := validate.FormatOf("subscriptionId", "body", "uuid", m.SubscriptionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceItem) UnmarshalBinary(b []byte) error {
	var res InvoiceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}