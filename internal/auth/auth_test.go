package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey1(t *testing.T) {
	h := http.Header{
		"Authorization": []string{"ApiKey abc123"},
	}

	//h.Add("Authorization", "ApiKey abc123")
	got, gErr := GetAPIKey(h)

	want, wErr := "abc123", error(nil)

	if !reflect.DeepEqual(got, want) && !reflect.DeepEqual(gErr, wErr) {
		t.Fatalf("expected: %v and %v, got: %v and %v", want, wErr, got, gErr)
	}
}

func TestGetAPIKey2(t *testing.T) {
	type test struct {
		name    string
		headers http.Header
		wantVal string
		wantErr error
	}
	tests := []test{
		{
			name:    "Valid",
			headers: http.Header{"Authorization": []string{"ApiKey abc123"}},
			wantVal: "abc123",
			wantErr: nil,
		},
		{
			name:    "Invalid header",
			headers: http.Header{"Authorissssssssation": []string{"ApiKey abc123"}},
			wantVal: "",
			wantErr: errors.New("no authorization header included"),
		},
		{
			name:    "Invalid value",
			headers: http.Header{"Authorization": []string{"ApiKeys abc123"}},
			wantVal: "",
			wantErr: errors.New("malformed authorization header"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotErr := GetAPIKey(tt.headers)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("%v GetAPIKey() gotVal = %v, want %v", tt.name, gotVal, tt.wantVal)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("%v GetAPIKey() gotErr = %v, want %v", tt.name, gotErr, tt.wantErr)
			}
		})
	}
}
