package zspace

import "testing"

func Test_Sync(t *testing.T) {
	if err := Sync("/mnt/c/Users/14028/Desktop/Wallpaper", "Wallpaper"); err != nil {
		t.Fatal(err)
	}
}
