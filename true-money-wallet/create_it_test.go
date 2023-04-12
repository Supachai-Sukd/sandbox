//go:build integration

package true_money_wallet

import (
	"errors"

	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Mock ของ
func handlerMock(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id": 1, "name": "Supachai", "Info": "gopher"}`))
}

func TestMakeHttp(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(handlerMock))
	defer server.Close()

	want := &Response{
		ID:   1,
		Name: "Supachai",
		Info: "gopher",
	}

	t.Run("Happy server response", func(t *testing.T) {
		resp, err := MakeHTTPCall(server.URL)

		// เราสามารถใช้การเทียบ Struct ได้ตามด้านล่าง
		// reflect.DeepEqual คือ ลึกๆแล้ว Struct นั้นหน้าตาเหมือนกันไหม
		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%v), got (%v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})

}
