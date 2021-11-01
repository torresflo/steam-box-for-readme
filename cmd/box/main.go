package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/YouEclipse/steam-box/pkg/steambox"
)

func main() {
	var err error
	steamAPIKey := os.Getenv("STEAM_API_KEY")
	steamID, _ := strconv.ParseUint(os.Getenv("STEAM_ID"), 10, 64)
	appIDs := os.Getenv("APP_ID")
	appIDList := make([]uint32, 0)

	for _, appID := range strings.Split(appIDs, ",") {
		appid, err := strconv.ParseUint(appID, 10, 32)
		if err != nil {
			continue
		}
		appIDList = append(appIDList, uint32(appid))
	}

	steamOption := "ALLTIME" // options for types of games to list: RECENT (recently played games), ALLTIME <default> (playtime of games in descending order), ALLTIME_AND_RECENT for both
	if os.Getenv("STEAM_OPTION") != "" {
		steamOption = os.Getenv("STEAM_OPTION")
	}

	multiLined := false // boolean for whether hours should have their own line - YES = true, NO = false
	if os.Getenv("MULTILINE") != "" {
		lineOption := os.Getenv("MULTILINE")
		if lineOption == "YES" {
			multiLined = true
		}
	}

	markdownFile := os.Getenv("MARKDOWN_FILE") // the markdown filename (e.g. MYFILE.md)

	updateAllTime := steamOption == "ALLTIME" || steamOption == "ALLTIME_AND_RECENT"
	updateRecent := steamOption == "RECENT" || steamOption == "ALLTIME_AND_RECENT"

	box := steambox.NewBox(steamAPIKey)
	ctx := context.Background()

	var (
		filename string
		lines    []string
	)

	if updateAllTime {
		filename = "🎮 Steam playtime leaderboard"
		lines, err = box.GetPlayTime(ctx, steamID, multiLined, appIDList...)
		if err != nil {
			panic("GetPlayTime err:" + err.Error())
		}

		if markdownFile != "" {
			content := bytes.NewBuffer(nil)
			content.WriteString(strings.Join(lines, "\n"))

			start := []byte("<!-- steam-box-playtime start -->")
			end := []byte("<!-- steam-box-playtime end -->")

			err = box.UpdateMarkdown(ctx, filename, markdownFile, content.Bytes(), start, end)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("updating markdown successfully on ", markdownFile)
		}
	}

	if updateRecent {
		filename = "🎮 Recently played Steam games"
		lines, err = box.GetRecentGames(ctx, steamID, multiLined)
		if err != nil {
			panic("GetRecentGames err:" + err.Error())
		}

		if markdownFile != "" {
			content := bytes.NewBuffer(nil)
			content.WriteString(strings.Join(lines, "\n"))

			start := []byte("<!-- steam-box-recent start -->")
			end := []byte("<!-- steam-box-recent end -->")

			err = box.UpdateMarkdown(ctx, filename, markdownFile, content.Bytes(), start, end)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("updating markdown successfully on ", markdownFile)
		}
	}
}
