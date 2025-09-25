package epoch

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	Passthrough bool   `json:"passthrough,omitempty"`
	MatchPath   string `json:"matchPath,omitempty"`
	KeyName     string `json:"keyName,omitempty"`
	Format      string `json:"format,omitempty"` // epoch | epoch_s | epoch_ns | rfc3339 | all
}

func CreateConfig() *Config {
	return &Config{
		Passthrough: false,
		MatchPath:   "/epoch",
		KeyName:     "epoch",
		Format:      "epoch", // default = milliseconds
	}
}

type Epoch struct {
	next        http.Handler
	name        string
	passthrough bool
	matchPath   string
	keyName     string
	format      string
}

func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Epoch{
		next:        next,
		name:        name,
		passthrough: config.Passthrough,
		matchPath:   config.MatchPath,
		keyName:     config.KeyName,
		format:      strings.ToLower(config.Format),
	}, nil
}

func (e *Epoch) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if e.passthrough {
		if !strings.HasPrefix(req.URL.Path, e.matchPath) {
			e.next.ServeHTTP(rw, req)
			return
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	now := time.Now()

	var resp map[string]any

	switch e.format {
	case "epoch":
		resp = map[string]any{e.keyName: now.UnixMilli()}
	case "epoch_s":
		resp = map[string]any{e.keyName: now.Unix()}
	case "epoch_ns":
		resp = map[string]any{e.keyName: now.UnixNano()}
	case "rfc3339":
		resp = map[string]any{e.keyName: now.UTC().Format(time.RFC3339)}
	case "all":
		resp = map[string]any{
			"epoch":    now.UnixMilli(),
			"epoch_s":  now.Unix(),
			"epoch_ns": now.UnixNano(),
			"rfc3339":  now.UTC().Format(time.RFC3339),
		}
	default:
		resp = map[string]any{e.keyName: now.UnixMilli()}
	}

	_ = json.NewEncoder(rw).Encode(resp)
}
