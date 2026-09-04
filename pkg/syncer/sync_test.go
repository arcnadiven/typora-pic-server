package syncer

import "testing"

func Test_doSync(t *testing.T) {
	if err := doSync(); err != nil {
		t.Errorf("%+v", err)
	}
}
