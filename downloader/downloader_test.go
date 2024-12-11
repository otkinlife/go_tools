package downloader

import (
	"testing"
)

func TestDownloadImage(t *testing.T) {
	ret := DownloadImage("https://cdn-dashboard-2.aihelp.net/FileService/UserFile/12767/202408/20240809033833643b22718d945_lite.png", "/Users/admin/Developer/go_tools/output")
	if ret.Err != nil {
		t.Errorf("DownloadImage() failed, err: %s", ret.Err)
		return
	}
	t.Logf("DownloadImage() succeeded, ret: %v", ret)
	ret = DownloadImage("https://bi-material-center.oss-cn-beijing.aliyuncs.com/rg_cs/discord_feedback_img/1595a17698487f525b0db39e43130a24.jpg?x-oss-process=image/resize,w_300", "/Users/admin/Developer/go_tools/output")
	if ret.Err != nil {
		t.Errorf("DownloadImage() failed, err: %s", ret.Err)
		return
	}
	t.Logf("DownloadImage() succeeded, ret: %v", ret)
	return
}

func TestDownloadFile(t *testing.T) {
	ret := DownloadFile("https://cdn-dashboard-2.aihelp.net/FileService/UserFile/12767/202408/20240809033833643b22718d945_lite.png", "/Users/admin/Developer/go_tools/output/test.png", 0)
	if ret != nil {
		t.Errorf("DownloadFile() failed, err: %s", ret)
		return
	}
	t.Logf("DownloadFile() succeeded")
	ret = DownloadFile("https://bi-material-center.oss-cn-beijing.aliyuncs.com/rg_cs/discord_feedback_img/1595a17698487f525b0db39e43130a24.jpg?x-oss-process=image/resize,w_300", "/Users/admin/Developer/go_tools/output/test2.png", 12)
	if ret != nil {
		t.Errorf("DownloadFile() failed, err: %s", ret)
		return
	}
	t.Logf("DownloadFile() succeeded")
	return
}
