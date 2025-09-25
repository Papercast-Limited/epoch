package epoch

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEpochHandlerFormats(t *testing.T) {
	tests := []struct {
		format     string
		expectKeys []string
	}{
		{"epoch", []string{"epoch"}},
		{"epoch_s", []string{"epoch"}},
		{"epoch_ns", []string{"epoch"}},
		{"rfc3339", []string{"epoch"}},
		{"all", []string{"epoch", "epoch_s", "epoch_ns", "rfc3339"}},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			cfg := CreateConfig()
			cfg.Format = tt.format

			handler, err := New(context.Background(), http.NotFoundHandler(), cfg, "test-epoch")
			if err != nil {
				t.Fatalf("failed to create handler: %v", err)
			}

			req := httptest.NewRequest(http.MethodGet, "/epoch", nil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Fatalf("expected status 200, got %d", rr.Code)
			}

			var resp map[string]any
			if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			// Check required keys exist
			for _, key := range tt.expectKeys {
				if _, ok := resp[key]; !ok {
					t.Errorf("expected key %q in response, got %v", key, resp)
				}
			}
		})
	}
}

func TestEpochHandlerCustomKeyName(t *testing.T) {
	cfg := CreateConfig()
	cfg.KeyName = "custom_time"
	cfg.Format = "epoch"

	handler, err := New(context.Background(), http.NotFoundHandler(), cfg, "test-epoch")
	if err != nil {
		t.Fatalf("failed to create handler: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/epoch", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if _, ok := resp["custom_time"]; !ok {
		t.Errorf("expected key %q in response, got %v", "custom_time", resp)
	}
}
