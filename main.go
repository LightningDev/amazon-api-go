package main

import (
	"net/url"
	"os"

	amazonspapi "github.com/amazinsellers/amazon-sp-api-sdk-go"
	"github.com/amazinsellers/amazon-sp-api-sdk-go/resources"
)

func setupEnv() {
	os.Setenv("SELLING_PARTNER_APP_CLIENT_ID", "")
	os.Setenv("SELLING_PARTNER_APP_CLIENT_SECRET", "")
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")
	os.Setenv("AWS_SELLING_PARTNER_ROLE", "")
}

func main() {
	setupEnv()
	spConfig := amazonspapi.NewSellingPartnerConfig()
	spConfig.Options.Debug = true
	spConfig.RefreshToken = ""
	spConfig.Region = "na"

	sp, err := amazonspapi.NewSellingPartner(spConfig)

	if err != nil {
		println(err.Error())
		return
	}

	queryValues := url.Values{}
	queryValues.Add("MarketplaceId", "ATVPDKIKX0DER")
	queryValues.Add("SellerId", "A36XKU1EALFF7S")

	params := resources.SellingPartnerParams{
		Operation: "getDefinitionsProductType",
		Query:     queryValues,
		PathParams: map[string]string{
			"productType": "GUILD_HOME",
		},
	}

	resp, err := sp.CallAPI(params)
	if err != nil {
		println(err.Error())
		return
	}

	println(*resp)
}
