package bilibili

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gtoxlili/give-advice/common/ht"
)

var (
	bvidRegexp = regexp.MustCompile(`BV[0-9a-zA-Z]+`)
	pageRegexp = regexp.MustCompile(`p=[0-9]+`)
)

func GetSubtitle(ctx context.Context, url string) (string, error) {
	// https://www.bilibili.com/video/BV1qD4y1G7YK/?spm_id_from=333.1007.tianma.1-1-1.click
	// 获取 bv 号
	bvid := bvidRegexp.FindString(url)
	// 获取页码
	pIStr := strings.TrimPrefix(pageRegexp.FindString(url), "p=")
	if pIStr == "" {
		pIStr = "1"
	}
	pageIndex, _ := strconv.Atoi(pIStr)
	if bvid == "" {
		// 返回错误
		return "", errors.New("unrecognizable url")
	}
	return getVideoInfo(ctx, bvid, pageIndex-1)
}

type videoInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Aid   int `json:"aid"`
		Pages []struct {
			Cid int `json:"cid"`
		}
	} `json:"data"`
}

func getVideoInfo(ctx context.Context, bvid string, pageIndex int) (string, error) {
	// https://api.bilibili.com/x/web-interface/view?bvid="+bvid
	res, err := ht.Get[videoInfo](ctx, "https://api.bilibili.com/x/web-interface/view?bvid="+bvid, nil)
	if err != nil {
		return "", err
	}
	info := res.Data
	if info.Code != 0 {
		return "", errors.New(info.Message)
	}
	return getPlayerInfo(ctx, info.Data.Aid, info.Data.Pages[pageIndex].Cid)
}

type playerInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Subtitle struct {
			Subtitles []struct {
				SubtitleUrl string `json:"subtitle_url"`
			} `json:"subtitles"`
		} `json:"subtitle"`
	} `json:"data"`
}

func getPlayerInfo(ctx context.Context, aid, cid int) (string, error) {
	// https://api.bilibili.com/x/player/v2?aid=821145051&cid=989862555
	res, err := ht.Get[playerInfo](ctx, fmt.Sprintf("https://api.bilibili.com/x/player/v2?aid=%d&cid=%d", aid, cid), nil)
	if err != nil {
		return "", err
	}
	info := res.Data
	if info.Code != 0 {
		return "", errors.New(info.Message)
	}
	if len(info.Data.Subtitle.Subtitles) == 0 {
		return "", errors.New("no subtitle")
	}
	return getSubtitle(ctx, info.Data.Subtitle.Subtitles[0].SubtitleUrl)
}

type subtitle struct {
	Body []struct {
		Content string `json:"content"`
	} `json:"body"`
}

func getSubtitle(ctx context.Context, url string) (string, error) {
	res, err := ht.Get[subtitle](ctx, "https:"+url, nil)
	if err != nil {
		return "", err
	}
	body := res.Data
	sb := strings.Builder{}
	for _, v := range body.Body {
		sb.WriteString(v.Content)
		sb.WriteString("\n")
	}
	return sb.String(), nil
}
