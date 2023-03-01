package bilibili

import (
	"context"
	"testing"
)

func Test_GetSubtitle(t *testing.T) {
	t.Log(GetSubtitle(context.Background(), "https://www.bilibili.com/video/BV1To4y1a7wW/?spm_id_from=333.1007.tianma.3-1-5.click"))
}
