package defaultController

import "testing"

func TestGetIndexTilte(t *testing.T) {
	golden := "博客文章"
	title := GetIndexTilte()

	if golden != title {
		t.Error(golden, title)
	}
}
