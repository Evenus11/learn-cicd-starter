package auth

import (
	  "errors"
	  "github.com/google/go-cmp/cmp"
	  "net/http"
	  "testing"
)

func TestAuth(t *testing.T) {
	  tests := map[string]struct {
			input   http.Header
			want    string
			wantErr error
	  }{
			"basic":           {input: http.Header{"Authorization": []string{"ApiKey Test-Key"}}, want: "Test-Key", wantErr: nil},
			"no auth hedder":  {input: http.Header{" ": []string{"ApiKey Test-Key"}}, want: "", wantErr: ErrNoAuthHeaderIncluded},
			"no key included": {input: http.Header{"Authorization": []string{"ApiKey"}}, want: "", wantErr: errors.New("malformed authorization_ eader")},
	  }

	  for name, tc := range tests {
			t.Run(name, func(t *testing.T) {
				  got, err := GetAPIKey(tc.input)
				  if tc.wantErr != nil && err == nil {
						t.Fatalf("GetAPIKey should have returned an error, but didn't")
				  } else if tc.wantErr == nil && err != nil {
						t.Fatalf("GetAPIKey should not have returned an error, but did: %v", err)
				  }

				  diff := cmp.Diff(tc.want, got)
				  if diff != "" {
						t.Fatalf(diff)
				  }
			})
	  }
}
