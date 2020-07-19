package cuikinton

import (
	"fmt"
	"log"
	"os"

	"github.com/k-nishijima/cuikinton/ui"
)

func Run(config KintoneConfig) {
	u := ui.NewUI()

	app := getKintoneApp(config)

	initKintoneView := func() {
		u.Console.SetText("kintoneに接続中です...")

		u.App.QueueUpdateDraw(func() {
			if records, err := getRecords(app); err != nil {
				Die(err)
			} else {
				u.Console.Clear()
				if len(records) == 0 {
					u.Console.SetText("データがありません。")
				} else {
					u.Console.SetText(fmt.Sprintf("%dレコード取得しました。", len(records)))
				}
				for _, r := range records {
					u.Side.AddItem(getHeadline(r), "", 0, nil)
					u.Side.SetSelectedFunc(func(i int, mainstr, substr string, ru rune) {
						u.Main.Clear()
						u.Main.SetText(prettyPrintKintoneRecord(records[i]))
						u.Main.ScrollToBeginning()
						u.App.SetFocus(u.Main)
					})
				}
				u.App.SetFocus(u.Side)
			}
		})
	}

	go initKintoneView()

	if err := u.Run(); err != nil {
		Die(err)
	}
}

func Die(err error) {
	log.Panicln(err)
	os.Exit(-1)
}
