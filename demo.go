package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

func main() {
	// Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	flag.Parse()

	// Initialize SDK
	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these IDS credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, _ := waf.New(sdkConfig)

	// Test Data
	accountNumber := "4FDBB"
	accessRuleID := "fm5mFgZM"
	rateRuleID := "Eb74fbxl"

	// Access Rule
	accessRule, _ := wafService.GetAccessRuleByID(accountNumber, accessRuleID)
	fmt.Println("******* WAF Access Rule *******")
	fmt.Printf("ID: %s\nNAME: %s\n\n", accessRule.ID, accessRule.Name)

	// Rate Rule
	rateRule, _ := wafService.GetRateRuleByID(accountNumber, rateRuleID)
	fmt.Println("******* WAF Rate Rule *******")
	fmt.Printf("ID: %s\nName: %s\n\n", rateRule.ID, rateRule.Name)

	// Custom Rule Create
	customRuleRequest := createCustomRule()
	resp, _ := wafService.AddCustomRuleSet(accountNumber, customRuleRequest)
	fmt.Println("******* WAF Custom Rule Create Response *******")
	fmt.Printf("ID: %s\n\n", resp.ID)

	// Custom Rule Delete
	deleteResp, _ := wafService.DeleteCustomRuleSet(accountNumber, resp.ID)
	fmt.Println("******* WAF Custom Rule Delete Response *******")
	fmt.Printf("ID: %s\nStatus: %s\n\n", deleteResp.ID, deleteResp.Status)
}

func createCustomRule() waf.CustomRule {
	return waf.CustomRule{
		Name: "SDK Custom Rule 1",
		Directives: []waf.Directive{
			{
				SecRule: waf.SecRule{
					Name: "Block Bots",
					Action: waf.Action{
						ID:  "66000000",
						MSG: "Invalid user agent.",
						T:   []string{"NONE"},
					},
					Operator: waf.Operator{
						IsNegated: false,
						Type:      "CONTAINS",
						Value:     "bot",
					},
					Variables: []waf.Variable{
						{
							IsCount: false,
							Matches: []waf.Match{
								{
									IsNegated: false,
									IsRegex:   false,
									Value:     "User-Agent",
								},
							},
							Type: "REQUEST_HEADERS",
						},
					},
				},
			},
		},
	}
}

// func printPretty(data interface{}) {
// 	bytes, _ := json.MarshalIndent(data, "", "\t")

// 	fmt.Printf("%v\n", string(bytes))
// }

// func createAccessRule() waf.AccessRule {
// 	filePath := "/Users/fcontreras35/repos/ec-sdk-go/example/waf/access_rules/add/waf_access_rule.json"
// 	jsonFile, err := os.Open(filePath)

// 	if err != nil {
// 		fmt.Printf("Error reading json file: %+v\n", err)
// 	}

// 	defer jsonFile.Close()

// 	bytes, err := ioutil.ReadAll(jsonFile)

// 	if err != nil {
// 		fmt.Printf("Error reading json file: %+v\n", err)
// 	}

// 	var rule waf.AccessRule

// 	err = json.Unmarshal(bytes, &rule)

// 	if err != nil {
// 		fmt.Printf("Error parsing json file: %+v\n", err)
// 	}

// 	return rule
// }

// waf.CustomRule{
// 	Name: "SDK Custom Rule 1",
// 	Directives: { []waf.Directive{
// 		{
// 			waf.SecRule{
// 				Action: waf.Action{
// 					ID:  "66000000",
// 					MSG: "Invalid user agent.",
// 					T:   []string{"NONE"},
// 				},
// 				Operator: waf.Operator{
// 					IsNegated: false,
// 					Type:      "CONTAINS",
// 					Value:     "bot",
// 				},
// 				Variables: []waf.Variable{
// 					{
// 						IsCount: false,
// 						Type:    "REQUEST_HEADERS",
// 						Matches: []waf.Match{
// 							{
// 								IsNegated: false,
// 								IsRegex:   false,
// 								Value:     "User-Agent",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	},
