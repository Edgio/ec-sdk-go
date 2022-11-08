package enums_test

import (
	"testing"

	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"
)

func TestEnums(t *testing.T) {
	cases := []struct {
		Name                  string
		Enum                  enums.Platform
		ExpectedHyphen        string
		ExpectedWithoutHyphen string
	}{
		{
			Name:                  "Happy path HTTP Large",
			Enum:                  enums.HttpLarge,
			ExpectedHyphen:        "http-large",
			ExpectedWithoutHyphen: "httplarge",
		},
		{
			Name:                  "Happy path HTTP Small",
			Enum:                  enums.HttpSmall,
			ExpectedHyphen:        "http-small",
			ExpectedWithoutHyphen: "httpsmall",
		},
		{
			Name:                  "Happy path ADN",
			Enum:                  enums.ADN,
			ExpectedHyphen:        "adn",
			ExpectedWithoutHyphen: "adn",
		},
	}

	for _, v := range cases {
		withHyphen := v.Enum.String()
		withoutHyphen := v.Enum.StringWithoutHyphen()

		if v.ExpectedHyphen != withHyphen {
			t.Fatalf("Failed for case: '%+v'. Expected: %s, Got: %s",
				v.Name, v.ExpectedHyphen, withHyphen)
		}

		if v.ExpectedWithoutHyphen != withoutHyphen {
			t.Fatalf("Failed for case: '%+v'. Expected: %s, Got: %s",
				v.Name, v.ExpectedWithoutHyphen, withoutHyphen)
		}

	}
}
