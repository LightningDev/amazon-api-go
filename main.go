package main

import (
	"net/url"
	"os"

	amazonspapi "github.com/amazinsellers/amazon-sp-api-sdk-go"
	"github.com/amazinsellers/amazon-sp-api-sdk-go/resources"
)

func setupEnv() {
	os.Setenv("SELLING_PARTNER_APP_CLIENT_ID", "amzn1.application-oa2-client.eac76cca8e604d4d9a93fa7f1a79d94d")
	os.Setenv("SELLING_PARTNER_APP_CLIENT_SECRET", "c57341408893ffd93d911a447102bd7e005b147ca83f4845c7f9da928112323e")
	os.Setenv("AWS_ACCESS_KEY_ID", "AKIAU3AGNTHQJH5CHRBO")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "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")
	os.Setenv("AWS_SELLING_PARTNER_ROLE", "arn:aws:iam::332873439712:role/SellingAPIRole")
}

func main() {
	setupEnv()
	spConfig := amazonspapi.NewSellingPartnerConfig()
	spConfig.Options.Debug = true
	spConfig.RefreshToken = "Atzr|IwEBIK7qjeYN0hSHlRAWELQs-b2HJpBN6dcbX0_7pPNJMeoUakOHR7cFacuGufDr4F1Eq9EzRpkXX410L2cjVV6eorZpmVNPLtNjLLJGrkfSmxWS3cIxvHWl81oTEQCtfxgP2l1djQIGNvX6hodrAQ4Ioor3cMxCJ-y8oaTrmCyzkO4UsaKTwN9okPKdU0HRYAVb25YBwxPVemQqPEp_lvWGRMtsg3nCTVZ5dgNpcF7TY24rM3lCeBbTIryNidwnEFR5KEC7y01VB5IHmitR_yeA88JjkdbV8j__IMm7H4iD21sWPjeXttN6F_Yvzs0yRBz6FuM"
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
