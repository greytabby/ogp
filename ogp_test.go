package ogp_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/greytabby/ogp"
)

func TestParse(t *testing.T) {
	document, err := ioutil.ReadFile(filepath.Join(".", "testdata", "ogp.html"))
	if err != nil {
		t.Fatal("Cannot find test html document")
	}
	og, err := ogp.Parse(document)
	wantTitle := "Open Graph protocol"
	if og.Title != wantTitle {
		t.Errorf("og:title Want: %v; Got: %v;", wantTitle, og.Title)
	}
	t.Logf("%+v", og)
}
