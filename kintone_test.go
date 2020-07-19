package cuikinton

import (
	"fmt"
	"os"
	"testing"
)

func TestGetHeadline(t *testing.T) {
	config := KintoneConfig{
		Domain:       os.Getenv("KINTONE_DOMAIN"),
		AppId:        97,
		GuestSpaceId: 0,
		ApiToken:     os.Getenv("KINTONE_APITOKEN"),
	}
	app := getKintoneApp(config)

	if records, err := getRecords(app); err != nil {
		Die(err)
	} else {
		for _, r := range records {
			line := getHeadline(r)
			fmt.Println(line)
		}
	}

}
