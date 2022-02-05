package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-debit-memo-request-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	DebitMemoRequest:              data.DebitMemoRequest,
	DebitMemoRequestType:          data.DebitMemoRequestType,
	SalesOrganization:             data.SalesOrganization,
	DistributionChannel:           data.DistributionChannel,
	OrganizationDivision:          data.OrganizationDivision,
	SalesGroup:                    data.SalesGroup,
	SalesOffice:                   data.SalesOffice,
	SalesDistrict:                 data.SalesDistrict,
	SoldToParty:                   data.SoldToParty,
	CreationDate:                  data.CreationDate,
	LastChangeDate:                data.LastChangeDate,
	LastChangeDateTime:            data.LastChangeDateTime,
	PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
	CustomerPurchaseOrderType:     data.CustomerPurchaseOrderType,
	CustomerPurchaseOrderDate:     data.CustomerPurchaseOrderDate,
	DebitMemoRequestDate:          data.DebitMemoRequestDate,
	TotalNetAmount:                data.TotalNetAmount,
	TransactionCurrency:           data.TransactionCurrency,
	SDDocumentReason:              data.SDDocumentReason,
	PricingDate:                   data.PricingDate,
	CustomerTaxClassification1:    data.CustomerTaxClassification1,
	HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
	IncotermsClassification:       data.IncotermsClassification,
	CustomerPaymentTerms:          data.CustomerPaymentTerms,
	PaymentMethod:                 data.PaymentMethod,
	BillingDocumentDate:           data.BillingDocumentDate,
	ReferenceSDDocument:           data.ReferenceSDDocument,
	ReferenceSDDocumentCategory:   data.ReferenceSDDocumentCategory,
	OverallSDProcessStatus:        data.OverallSDProcessStatus,
	TotalCreditCheckStatus:        data.TotalCreditCheckStatus,
	OverallSDDocumentRejectionSts: data.OverallSDDocumentRejectionSts,
	OverallOrdReltdBillgStatus:    data.OverallOrdReltdBillgStatus,
	ToHeaderPartner:               data.ToHeaderPartner.Deferred.URI,
	ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
	DebitMemoRequest:             data.DebitMemoRequest,
	DebitMemoRequestItem:         data.DebitMemoRequestItem,
	HigherLevelItem:              data.HigherLevelItem,
	DebitMemoRequestItemCategory: data.DebitMemoRequestItemCategory,
	DebitMemoRequestItemText:     data.DebitMemoRequestItemText,
	PurchaseOrderByCustomer:      data.PurchaseOrderByCustomer,
	Material:                     data.Material,
	MaterialByCustomer:           data.MaterialByCustomer,
	PricingDate:                  data.PricingDate,
	RequestedQuantity:            data.RequestedQuantity,
	RequestedQuantityUnit:        data.RequestedQuantityUnit,
	ItemGrossWeight:              data.ItemGrossWeight,
	ItemNetWeight:                data.ItemNetWeight,
	ItemWeightUnit:               data.ItemWeightUnit,
	ItemVolume:                   data.ItemVolume,
	ItemVolumeUnit:               data.ItemVolumeUnit,
	TransactionCurrency:          data.TransactionCurrency,
	NetAmount:                    data.NetAmount,
	MaterialGroup:                data.MaterialGroup,
	MaterialPricingGroup:         data.MaterialPricingGroup,
	ProductTaxClassification1:    data.ProductTaxClassification1,
	Batch:                        data.Batch,
	Plant:                        data.Plant,
	IncotermsClassification:      data.IncotermsClassification,
	CustomerPaymentTerms:         data.CustomerPaymentTerms,
	ItemBillingBlockReason:       data.ItemBillingBlockReason,
	SalesDocumentRjcnReason:      data.SalesDocumentRjcnReason,
	WBSElement:                   data.WBSElement,
	ProfitCenter:                 data.ProfitCenter,
	ReferenceSDDocument:          data.ReferenceSDDocument,
	ReferenceSDDocumentItem:      data.ReferenceSDDocumentItem,
	SDProcessStatus:              data.SDProcessStatus,
	OrderRelatedBillingStatus:    data.OrderRelatedBillingStatus,
	ToItemPricingElement:         data.ToItemPricingElement.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
		DebitMemoRequest:       data.DebitMemoRequest,
		PartnerFunction:        data.PartnerFunction,
		Customer:               data.Customer,
		Supplier:               data.Supplier,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
	DebitMemoRequest:             data.DebitMemoRequest,
	DebitMemoRequestItem:         data.DebitMemoRequestItem,
	HigherLevelItem:              data.HigherLevelItem,
	DebitMemoRequestItemCategory: data.DebitMemoRequestItemCategory,
	DebitMemoRequestItemText:     data.DebitMemoRequestItemText,
	PurchaseOrderByCustomer:      data.PurchaseOrderByCustomer,
	Material:                     data.Material,
	MaterialByCustomer:           data.MaterialByCustomer,
	PricingDate:                  data.PricingDate,
	RequestedQuantity:            data.RequestedQuantity,
	RequestedQuantityUnit:        data.RequestedQuantityUnit,
	ItemGrossWeight:              data.ItemGrossWeight,
	ItemNetWeight:                data.ItemNetWeight,
	ItemWeightUnit:               data.ItemWeightUnit,
	ItemVolume:                   data.ItemVolume,
	ItemVolumeUnit:               data.ItemVolumeUnit,
	TransactionCurrency:          data.TransactionCurrency,
	NetAmount:                    data.NetAmount,
	MaterialGroup:                data.MaterialGroup,
	MaterialPricingGroup:         data.MaterialPricingGroup,
	ProductTaxClassification1:    data.ProductTaxClassification1,
	Batch:                        data.Batch,
	Plant:                        data.Plant,
	IncotermsClassification:      data.IncotermsClassification,
	CustomerPaymentTerms:         data.CustomerPaymentTerms,
	ItemBillingBlockReason:       data.ItemBillingBlockReason,
	SalesDocumentRjcnReason:      data.SalesDocumentRjcnReason,
	WBSElement:                   data.WBSElement,
	ProfitCenter:                 data.ProfitCenter,
	ReferenceSDDocument:          data.ReferenceSDDocument,
	ReferenceSDDocumentItem:      data.ReferenceSDDocumentItem,
	SDProcessStatus:              data.SDProcessStatus,
	OrderRelatedBillingStatus:    data.OrderRelatedBillingStatus,
	ToItemPricingElement:         data.ToItemPricingElement.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
	DebitMemoRequest:               data.DebitMemoRequest,
	DebitMemoRequestItem:           data.DebitMemoRequestItem,
	PricingProcedureStep:           data.PricingProcedureStep,
	PricingProcedureCounter:        data.PricingProcedureCounter,
	ConditionApplication:           data.ConditionApplication,
	ConditionType:                  data.ConditionType,
	PricingDateTime:                data.PricingDateTime,
	PriceConditionDeterminationDte: data.PriceConditionDeterminationDte,
	ConditionCalculationType:       data.ConditionCalculationType,
	ConditionBaseValue:             data.ConditionBaseValue,
	ConditionRateValue:             data.ConditionRateValue,
	ConditionCurrency:              data.ConditionCurrency,
	ConditionQuantity:              data.ConditionQuantity,
	ConditionQuantityUnit:          data.ConditionQuantityUnit,
	ConditionToBaseQtyNmrtr:        data.ConditionToBaseQtyNmrtr,
	ConditionToBaseQtyDnmntr:       data.ConditionToBaseQtyDnmntr,
	ConditionCategory:              data.ConditionCategory,
	PricingScaleType:               data.PricingScaleType,
	ConditionRecord:                data.ConditionRecord,
	ConditionSequentialNumber:      data.ConditionSequentialNumber,
	TaxCode:                        data.TaxCode,
	ConditionAmount:                data.ConditionAmount,
	TransactionCurrency:            data.TransactionCurrency,
	PricingScaleBasis:              data.PricingScaleBasis,
	ConditionScaleBasisValue:       data.ConditionScaleBasisValue,
	ConditionScaleBasisUnit:        data.ConditionScaleBasisUnit,
	ConditionScaleBasisCurrency:    data.ConditionScaleBasisCurrency,
	ConditionIsManuallyChanged:     data.ConditionIsManuallyChanged,
		})
	}

	return toItemPricingElement, nil
}
