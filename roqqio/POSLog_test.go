package roqqio

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestCustomerXml(t *testing.T) {
	posLog := handleXmlFile(t, "testdata/sale_with_customer.xml")
	assertNotNil(t, posLog.Transaction.RetailTransaction.Customer)
	assertNotNil(t, posLog.GetCustomer())
	assertThat(t, posLog.Transaction.RetailTransaction.Customer, posLog.GetCustomer())
	assertThat(t, posLog.GetCustomer().Name.OfficialName, "Erika Mustermann")

	assertNotNil(t, posLog.GetLoyaltyAccount())
	assertThat(t, posLog.GetLoyaltyAccount().LoyaltyCard.PrimaryAccountNumber, "T220001293")
	assertThat(t, posLog.GetLoyaltyAccount().CustomerID, posLog.GetCustomer().CustomerID)

	assertNotNil(t, posLog.GetLineItems())
	assertThat(t, len(*posLog.GetLineItems()), 2)

	lineItems := *posLog.GetLineItems()
	assertThat(t, lineItems[0].GetType(), LineTypeSale)
	assertThat(t, lineItems[0].IsSale(), true)
	assertThat(t, lineItems[0].IsTender(), false)
	assertThat(t, lineItems[0].IsTenderChange(), false)

	assertThat(t, lineItems[1].GetType(), LineTypeTender)
	assertThat(t, lineItems[1].IsSale(), false)
	assertThat(t, lineItems[1].IsTender(), true)
	assertThat(t, lineItems[1].IsTenderChange(), false)
}

func TestTenderChangeXml(t *testing.T) {
	posLog := handleXmlFile(t, "testdata/sale_with_tenderChange.xml")
	assertNil(t, posLog.Transaction.RetailTransaction.Customer)
	assertNil(t, posLog.GetCustomer())
	assertNil(t, posLog.GetLoyaltyAccount())

	assertNotNil(t, posLog.GetLineItems())
	assertThat(t, len(*posLog.GetLineItems()), 3)

	lineItems := *posLog.GetLineItems()
	assertThat(t, lineItems[0].GetType(), LineTypeSale)
	assertThat(t, lineItems[0].IsSale(), true)
	assertThat(t, lineItems[0].IsTender(), false)
	assertThat(t, lineItems[0].IsTenderChange(), false)

	assertThat(t, lineItems[1].GetType(), LineTypeTender)
	assertThat(t, lineItems[1].IsSale(), false)
	assertThat(t, lineItems[1].IsTender(), true)
	assertThat(t, lineItems[1].IsTenderChange(), false)
	assertThat(t, lineItems[1].Tender.TypeCode, TenderTypeCodeSale)

	assertThat(t, lineItems[2].GetType(), LineTypeTenderChange)
	assertThat(t, lineItems[2].IsSale(), false)
	assertThat(t, lineItems[2].IsTender(), false)
	assertThat(t, lineItems[2].IsTenderChange(), true)
}

func TestRetoureXml(t *testing.T) {
	posLog := handleXmlFile(t, "testdata/retoure.xml")
	assertNotNil(t, posLog.Transaction.RetailTransaction.Customer)
	assertNotNil(t, posLog.GetCustomer())
	assertNotNil(t, posLog.GetLoyaltyAccount())

	assertNotNil(t, posLog.GetLineItems())
	assertThat(t, len(*posLog.GetLineItems()), 2)

	lineItems := *posLog.GetLineItems()
	assertThat(t, lineItems[0].GetType(), LineTypeSale)
	assertThat(t, lineItems[0].IsSale(), true)
	assertThat(t, lineItems[0].IsTender(), false)
	assertThat(t, lineItems[0].IsTenderChange(), false)
	assertThat(t, lineItems[0].Sale.Quantity.Value < 0, true)
	assertThat(t, lineItems[0].Sale.Quantity.Value, -1.0)
	assertThat(t, lineItems[0].Sale.ActualSalesUnitPrice.Value > 0, true) // Prices are Always Positive

	assertThat(t, lineItems[1].GetType(), LineTypeTender)
	assertThat(t, lineItems[1].IsSale(), false)
	assertThat(t, lineItems[1].IsTender(), true)
	assertThat(t, lineItems[1].IsTenderChange(), false)
	assertThat(t, lineItems[1].Tender.TypeCode, TenderTypeCodeRefund)

	// should not contain a TransactionLink, because retoure without Receipt
	assertNil(t, lineItems[0].Sale.TransactionLink)
}

func TestRetoureWithLinkXml(t *testing.T) {
	posLog := handleXmlFile(t, "testdata/retoure_with_link.xml")
	assertNotNil(t, posLog.Transaction.RetailTransaction.Customer)
	assertNotNil(t, posLog.GetCustomer())
	assertNotNil(t, posLog.GetLoyaltyAccount())

	assertNotNil(t, posLog.GetLineItems())
	assertThat(t, len(*posLog.GetLineItems()), 2)

	lineItems := *posLog.GetLineItems()
	assertThat(t, lineItems[0].GetType(), LineTypeSale)
	assertThat(t, lineItems[0].IsSale(), true)
	assertThat(t, lineItems[0].IsTender(), false)
	assertThat(t, lineItems[0].IsTenderChange(), false)
	assertThat(t, lineItems[0].Sale.Quantity.Value < 0, true)
	assertThat(t, lineItems[0].Sale.Quantity.Value, -1.0)
	assertThat(t, lineItems[0].Sale.ActualSalesUnitPrice.Value > 0, true) // Prices are Always Positive

	assertThat(t, lineItems[1].GetType(), LineTypeTender)
	assertThat(t, lineItems[1].IsSale(), false)
	assertThat(t, lineItems[1].IsTender(), true)
	assertThat(t, lineItems[1].IsTenderChange(), false)
	assertThat(t, lineItems[1].Tender.TypeCode, TenderTypeCodeRefund)

	// should handle the same as normal retoure but LineItem should contain a Linkage
	assertNotNil(t, lineItems[0].Sale.TransactionLink)
	assertThat(t, lineItems[0].Sale.TransactionLink.SequenceNumber, "3569")
	assertThat(t, lineItems[0].Sale.TransactionLink.BusinessDayDate, "2024-10-15")
}

func TestPosStorno(t *testing.T) {
	posLog := handleXmlFile(t, "testdata/sale_with_posStorno.xml")
	assertNil(t, posLog.Transaction.RetailTransaction.Customer)
	assertNil(t, posLog.GetCustomer())
	assertNil(t, posLog.GetLoyaltyAccount())

	assertNotNil(t, posLog.GetLineItems())
	assertThat(t, len(*posLog.GetLineItems()), 3)

	lineItems := *posLog.GetLineItems()
	assertThat(t, lineItems[0].GetType(), LineTypeSale)
	assertThat(t, lineItems[0].IsSale(), true)
	assertThat(t, lineItems[0].IsTender(), false)
	assertThat(t, lineItems[0].IsTenderChange(), false)

	assertThat(t, lineItems[1].GetType(), LineTypeSale)
	assertThat(t, lineItems[1].IsSale(), true)
	assertThat(t, lineItems[1].IsTender(), false)
	assertThat(t, lineItems[1].IsTenderChange(), false)

	assertThat(t, lineItems[2].GetType(), LineTypeTender)
	assertThat(t, lineItems[2].IsSale(), false)
	assertThat(t, lineItems[2].IsTender(), true)
	assertThat(t, lineItems[2].IsTenderChange(), false)
	assertThat(t, lineItems[2].Tender.TypeCode, TenderTypeCodeSale)
}

func assertNotNil[C any](t *testing.T, a *C) {
	t.Helper()
	if a == nil {
		t.Error("is nil")
	}
}

func assertNil[C any](t *testing.T, a *C) {
	t.Helper()
	if a != nil {
		t.Error("is not nil")
	}
}

func assertThat[C comparable](t *testing.T, a C, b C) {
	t.Helper()
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}

func handleXmlFile(t *testing.T, f string) *POSLog {
	t.Helper()
	bytes, _ := os.ReadFile(f)
	posLog := POSLog{}
	err := xml.Unmarshal(bytes, &posLog)
	if err != nil {
		t.Fatal(err)
	}

	/*
		jOut, err := json.MarshalIndent(posLog, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(jOut))
	*/
	return &posLog
}
