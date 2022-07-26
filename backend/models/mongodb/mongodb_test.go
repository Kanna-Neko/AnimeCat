package mongodb

import "testing"

func TestGetSetting(t *testing.T) {
	res, err := GetSetting()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("BucketName: %s\nBucketRegion: %s\nEndPoint: %s\nFooter: %s\nGlobalCss: %s\nGolbalJs: %s\nLanguage: %s\nLogo: %s\nPageSize: %d\nSecretId: %s\nSecretKey: %s\nTheme: %s\nWallpaper: %s\nWebsiteTitle: %s\n",
			res.BucketName, res.BucketRegion, res.EndPoint, res.Footer, res.GlobalCSS, res.GlobalJS, res.Language, res.Logo, res.PageSize, res.SecretId, res.SecretKey, res.Theme, res.Wallpaper, res.WebsiteTitle)
	}
}

func TestInitSetting(t *testing.T) {
	err := InitSetting()
	if err != nil {
		t.Error(err)
	}else {
		t.Log("Init setting success!")
	}
}