package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input      http.Header
		want       string
		wantError  bool
		wantErrMsg string
	}{
		"valid header": {
			input: http.Header{
				"Authorization": {"ApiKey test"},
			},
			want:       "test",
			wantError:  false,
			wantErrMsg: "",
		},
		"missing header": {
			input:      http.Header{},
			want:       "",
			wantError:  true,
			wantErrMsg: "no authorization header included",
		},
		"malformed header": {
			input:      http.Header{"Authorization": {"test"}},
			want:       "",
			wantError:  true,
			wantErrMsg: "malformed authorization header",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if tc.wantError {
				if err == nil {
					t.Fatalf("expected err, got nil")
				}
				if err.Error() != tc.wantErrMsg {
					t.Errorf("expected error message %q, got %q", tc.wantErrMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect error, got %v", err)
				}
				if got != tc.want {
					t.Errorf("expected result %v, got %v", tc.want, got)
				}
			}
		})
	}
}
