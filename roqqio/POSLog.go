package roqqio

import "time"

type POSLog struct {
	Transaction struct {
		BeginDateTime   time.Time `xml:"BeginDateTime"`
		BusinessDayDate string    `xml:"BusinessDayDate"`
		BusinessUnit    struct {
			UnitID struct {
				Name     string `xml:"Name,attr"`
				CharData string `xml:",chardata"`
			} `xml:"UnitID"`
		} `xml:"BusinessUnit"`
		EndDateTime time.Time `xml:"EndDateTime"`
		Hashcode    string    `xml:"Hashcode"`
		OperatorID  struct {
			OperatorName string `xml:"OperatorName,attr"`
			OperatorType string `xml:"OperatorType,attr"`
			CharData     string `xml:",chardata"`
		} `xml:"OperatorID"`
		POSLogDateTime    string `xml:"POSLogDateTime"`
		ReceiptDateTime   string `xml:"ReceiptDateTime"`
		ReceiptNumber     string `xml:"ReceiptNumber"`
		RetailTransaction struct {
			Customer *struct {
				Gender  string `xml:"Gender,attr"`
				Address struct {
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
				} `xml:"Address"`
				BirthDayMonth string `xml:"BirthDayMonth"`
				BirthYear     string `xml:"BirthYear"`
				CustomerID    string `xml:"CustomerID"`
				Name          struct {
					OfficialName string `xml:"OfficialName"`
				} `xml:"Name"`
			} `xml:"Customer"`
			LineItem []struct {
				Sale *struct {
					ActualSalesUnitPrice struct {
						Currency string `xml:"Currency,attr"`
						CharData string `xml:",chardata"`
					} `xml:"ActualSalesUnitPrice"`
					Dimension1     string `xml:"Dimension1"`
					Dimension2     string `xml:"Dimension2"`
					Dimension3     string `xml:"Dimension3"`
					ExtendedAmount struct {
						Currency string `xml:"Currency,attr"`
						CharData string `xml:",chardata"`
					} `xml:"ExtendedAmount"`
					ExtendedDiscountAmount struct {
						Currency string `xml:"Currency,attr"`
						CharData string `xml:",chardata"`
					} `xml:"ExtendedDiscountAmount"`
					ItemID struct {
						Name     string `xml:"Name,attr"`
						CharData string `xml:",chardata"`
					} `xml:"ItemID"`
					LineType     string `xml:"LineType"`
					ProductGroup string `xml:"ProductGroup"`
					Quantity     struct {
						UnitOfMeasureCode string `xml:"UnitOfMeasureCode,attr"`
						CharData          string `xml:",chardata"`
					} `xml:"Quantity"`
					RegularSalesUnitPrice struct {
						Currency string `xml:"Currency,attr"`
						CharData string `xml:",chardata"`
					} `xml:"RegularSalesUnitPrice"`
					Rounding struct {
						Currency string `xml:"Currency,attr"`
						CharData string `xml:",chardata"`
					} `xml:"Rounding"`
					SalesMode string `xml:"SalesMode"`
				} `xml:"Sale"`
				SequenceNumber string `xml:"SequenceNumber"`
				Tender         *struct {
					TenderType string `xml:"TenderType,attr"`
					TypeCode   string `xml:"TypeCode,attr"`
					Amount     struct {
						Currency string `xml:"Currency,attr"`
						CharData string `xml:",chardata"`
					} `xml:"Amount"`
					PaymentTypeDescription string  `xml:"PaymentTypeDescription"`
					PaymentTypeExternalID  string  `xml:"PaymentTypeExternalId"`
					Rounding               float64 `xml:"Rounding"`
				} `xml:"Tender"`
				TenderChange *struct {
					TypeCode     string  `xml:"TypeCode,attr"`
					Rounding     float64 `xml:"Rounding"`
					TenderChange struct {
						TenderType string `xml:"TenderType,attr"`
						Amount     struct {
							Currency string `xml:"Currency,attr"`
							CharData string `xml:",chardata"`
						} `xml:"Amount"`
					} `xml:"TenderChange"`
				} `xml:"TenderChange"`
			} `xml:"LineItem"`
			LoyaltyAccount *struct {
				CustomerID  string `xml:"CustomerID"`
				LoyaltyCard struct {
					PrimaryAccountNumber string `xml:"PrimaryAccountNumber"`
				} `xml:"LoyaltyCard"`
			} `xml:"LoyaltyAccount"`
			Total []struct {
				CurrencyCode string `xml:"CurrencyCode,attr"`
				TotalType    string `xml:"TotalType,attr"`
				CharData     string `xml:",chardata"`
			} `xml:"Total"`
		} `xml:"RetailTransaction"`
		SequenceNumber string `xml:"SequenceNumber"`
		WorkstationID  struct {
			TypeCode string `xml:"TypeCode,attr"`
			CharData string `xml:",chardata"`
		} `xml:"WorkstationID"`
	} `xml:"Transaction"`
}
