package router

import (
	"github.com/gorilla/mux"
	"reflect"
	"testing"
)

func TestHttpRouter_CreateRouter(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		{"Test Router Creation", mux.NewRouter().StrictSlash(true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := testRouter{}
			if got := r.CreateRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testRouter struct{}

func (r testRouter) CreateRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}
