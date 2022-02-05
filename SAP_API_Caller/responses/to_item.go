package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			DebitMemoRequest             string `json:"DebitMemoRequest"`
			DebitMemoRequestItem         string `json:"DebitMemoRequestItem"`
			HigherLevelItem              string `json:"HigherLevelItem"`
			DebitMemoRequestItemCategory string `json:"DebitMemoRequestItemCategory"`
			DebitMemoRequestItemText     string `json:"DebitMemoRequestItemText"`
			PurchaseOrderByCustomer      string `json:"PurchaseOrderByCustomer"`
			Material                     string `json:"Material"`
			MaterialByCustomer           string `json:"MaterialByCustomer"`
			PricingDate                  string `json:"PricingDate"`
			RequestedQuantity            string `json:"RequestedQuantity"`
			RequestedQuantityUnit        string `json:"RequestedQuantityUnit"`
			ItemGrossWeight              string `json:"ItemGrossWeight"`
			ItemNetWeight                string `json:"ItemNetWeight"`
			ItemWeightUnit               string `json:"ItemWeightUnit"`
			ItemVolume                   string `json:"ItemVolume"`
			ItemVolumeUnit               string `json:"ItemVolumeUnit"`
			TransactionCurrency          string `json:"TransactionCurrency"`
			NetAmount                    string `json:"NetAmount"`
			MaterialGroup                string `json:"MaterialGroup"`
			MaterialPricingGroup         string `json:"MaterialPricingGroup"`
			ProductTaxClassification1    string `json:"ProductTaxClassification1"`
			Batch                        string `json:"Batch"`
			Plant                        string `json:"Plant"`
			IncotermsClassification      string `json:"IncotermsClassification"`
			CustomerPaymentTerms         string `json:"CustomerPaymentTerms"`
			ItemBillingBlockReason       string `json:"ItemBillingBlockReason"`
			SalesDocumentRjcnReason      string `json:"SalesDocumentRjcnReason"`
			WBSElement                   string `json:"WBSElement"`
			ProfitCenter                 string `json:"ProfitCenter"`
			ReferenceSDDocument          string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem      string `json:"ReferenceSDDocumentItem"`
			SDProcessStatus              string `json:"SDProcessStatus"`
			OrderRelatedBillingStatus    string `json:"OrderRelatedBillingStatus"`
			ToItemPricingElement struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PricingElement"`
		} `json:"results"`
	} `json:"d"`
}
