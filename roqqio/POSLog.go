package roqqio

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	LineTypeSale         = "Sale"
	LineTypeTender       = "Tender"
	LineTypeTenderChange = "TenderChange"
	LineTypeUnknown      = "Unknown"
)

const (
	SaleLineTypeVoucherSale = "VOUCHER_SALE"
	SaleLineTypeItemSale    = "ITEM_SALE"
)

const (
	TenderTypeCodeSale   = "Sale"
	TenderTypeCodeRefund = "Refund"
)

type POSLog struct {
	Transaction Transaction `xml:"Transaction"`
}

type Transaction struct {
	BeginDateTime     time.Time         `xml:"BeginDateTime"`
	EndDateTime       time.Time         `xml:"EndDateTime"`
	BusinessDayDate   string            `xml:"BusinessDayDate"`
	Hashcode          string            `xml:"Hashcode"`
	POSLogDateTime    string            `xml:"POSLogDateTime"`
	ReceiptDateTime   string            `xml:"ReceiptDateTime"`
	ReceiptNumber     string            `xml:"ReceiptNumber"`
	SequenceNumber    string            `xml:"SequenceNumber"`
	BusinessUnit      BusinessUnit      `xml:"BusinessUnit"`
	OperatorID        OperatorID        `xml:"OperatorID"`
	WorkstationID     WorkstationID     `xml:"WorkstationID"`
	RetailTransaction RetailTransaction `xml:"RetailTransaction"`
}

type BusinessUnit struct {
	UnitID UnitID `xml:"UnitID"`
}

type UnitID struct {
	Name     string `xml:"Name,attr"`
	CharData string `xml:",chardata"`
}

type OperatorID struct {
	OperatorName string `xml:"OperatorName,attr"`
	OperatorType string `xml:"OperatorType,attr"`
	CharData     string `xml:",chardata"`
}

type WorkstationID struct {
	TypeCode string `xml:"TypeCode,attr"`
	CharData string `xml:",chardata"`
}

type RetailTransaction struct {
	Customer       *Customer       `xml:"Customer"`
	LineItems      []LineItem      `xml:"LineItem"`
	LoyaltyAccount *LoyaltyAccount `xml:"LoyaltyAccount"`
	Total          []struct {
		CurrencyCode string `xml:"CurrencyCode,attr"`
		TotalType    string `xml:"TotalType,attr"`
		CharData     string `xml:",chardata"`
	} `xml:"Total"`
}

type Customer struct {
	Gender        string  `xml:"Gender,attr"`
	Address       Address `xml:"Address"`
	BirthDayMonth string  `xml:"BirthDayMonth"`
	BirthYear     string  `xml:"BirthYear"`
	CustomerID    string  `xml:"CustomerID"`
	Name          struct {
		OfficialName string `xml:"OfficialName"`
	} `xml:"Name"`
}

type Address struct {
	AddressType string `xml:"AddressType,attr"`
	AddressLine struct {
		TypeCode string `xml:"TypeCode,attr"`
		CharData string `xml:",chardata"`
	} `xml:"AddressLine"`
	City    string `xml:"City"`
	Country struct {
		Code string `xml:"Code,attr"`
	} `xml:"Country"`
	PostalCode string `xml:"PostalCode"`
	Territory  struct {
		TypeCode string `xml:"TypeCode,attr"`
	} `xml:"Territory"`
}

type LoyaltyAccount struct {
	CustomerID  string `xml:"CustomerID"`
	LoyaltyCard *struct {
		PrimaryAccountNumber string `xml:"PrimaryAccountNumber"`
	} `xml:"LoyaltyCard"`
}

type LineItem struct {
	SequenceNumber string        `xml:"SequenceNumber"`
	Sale           *Sale         `xml:"Sale"`
	Tender         *Tender       `xml:"Tender"`
	TenderChange   *TenderChange `xml:"TenderChange"`
}

type Sale struct {
	ActualSalesUnitPrice   CurrencyAmount `xml:"ActualSalesUnitPrice"`
	Dimension1             string         `xml:"Dimension1"`
	Dimension2             string         `xml:"Dimension2"`
	Dimension3             string         `xml:"Dimension3"`
	ExtendedAmount         CurrencyAmount `xml:"ExtendedAmount"`
	ExtendedDiscountAmount CurrencyAmount `xml:"ExtendedDiscountAmount"`
	ItemID                 struct {
		Name     string `xml:"Name,attr"`
		CharData string `xml:",chardata"`
	} `xml:"ItemID"`
	LineType              string           `xml:"LineType"`
	ProductGroup          string           `xml:"ProductGroup"`
	Quantity              Quantity         `xml:"Quantity"`
	RegularSalesUnitPrice CurrencyAmount   `xml:"RegularSalesUnitPrice"`
	Rounding              CurrencyAmount   `xml:"Rounding"`
	SalesMode             string           `xml:"SalesMode"`
	TransactionLink       *TransactionLink `xml:"TransactionLink"`
}

type Quantity struct {
	UnitOfMeasureCode string `xml:"UnitOfMeasureCode,attr"`
	Value             string `xml:",chardata"`
}

func (q *Quantity) AsFloat() float64 {
	f, _ := strconv.ParseFloat(q.Value, 64)
	return f
}

func (q *Quantity) AsInt() int64 {
	val := strings.Split(q.Value, ".")[0]
	i, _ := strconv.ParseInt(val, 10, 64)
	return i
}

type TransactionLink struct {
	TransactionID          string        `xml:"TransactionID"`
	BusinessUnit           UnitID        `xml:"BusinessUnit"`
	WorkstationID          WorkstationID `xml:"WorkstationID"`
	BusinessDayDate        string        `xml:"BusinessDayDate>Date"`
	SequenceNumber         string        `xml:"SequenceNumber"`
	LineItemSequenceNumber string        `xml:"LineItemSequenceNumber"`
}

type Tender struct {
	TenderType             string         `xml:"TenderType,attr"`
	TypeCode               string         `xml:"TypeCode,attr"`
	Amount                 CurrencyAmount `xml:"Amount"`
	PaymentTypeDescription string         `xml:"PaymentTypeDescription"`
	PaymentTypeExternalID  string         `xml:"PaymentTypeExternalId"`
	Rounding               float64        `xml:"Rounding"`
}
type TenderChange struct {
	TypeCode     string  `xml:"TypeCode,attr"`
	Rounding     float64 `xml:"Rounding"`
	TenderChange struct {
		TenderType string         `xml:"TenderType,attr"`
		Amount     CurrencyAmount `xml:"Amount"`
	} `xml:"TenderChange"`
}

type CurrencyAmount struct {
	Currency string `xml:"Currency,attr"`
	Value    string `xml:",chardata"`
}

func (c *CurrencyAmount) AsFloat() float64 {
	f, _ := strconv.ParseFloat(c.Value, 64)
	return f
}

func (c *CurrencyAmount) AsCents() int64 {
	return c.AsCurrencyDecimals(2)
}

func (c *CurrencyAmount) AsCurrencyDecimals(decimals int) int64 {
	factor := int64(math.Pow10(decimals))
	split := strings.Split(c.Value, ".")
	major, _ := strconv.ParseInt(split[0], 10, 64)
	for {
		if decimals > len(split[1]) {
			split[1] = split[1] + "0"
		} else {
			break
		}
	}
	minor, _ := strconv.ParseInt(split[1][0:decimals], 10, 64)
	return major*factor + minor
}

func (l *LineItem) GetType() string {
	if l.Sale != nil {
		return LineTypeSale
	}
	if l.Tender != nil {
		return LineTypeTender
	}
	if l.TenderChange != nil {
		return LineTypeTenderChange
	}

	return LineTypeUnknown
}

func (p *POSLog) GetCustomer() *Customer {
	return p.Transaction.RetailTransaction.Customer
}

func (p *POSLog) GetLoyaltyAccount() *LoyaltyAccount {
	return p.Transaction.RetailTransaction.LoyaltyAccount
}

func (p *POSLog) GetLineItems() *[]LineItem {
	return &p.Transaction.RetailTransaction.LineItems
}

func (l *LineItem) IsSale() bool {
	return l.GetType() == LineTypeSale
}

func (l *LineItem) IsTender() bool {
	return l.GetType() == LineTypeTender
}

func (l *LineItem) IsTenderChange() bool {
	return l.GetType() == LineTypeTenderChange
}
