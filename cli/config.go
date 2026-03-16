package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func resolveClaimPollConfig() (claimPollConfig, error) {
	interval, err := parseDurationConfig("claim_poll_interval", "claim poll interval")
	if err != nil {
		return claimPollConfig{}, err
	}
	timeout, err := parseDurationConfig("claim_poll_timeout", "claim timeout")
	if err != nil {
		return claimPollConfig{}, err
	}

	return claimPollConfig{interval: interval, timeout: timeout}, nil
}

func parseDurationConfig(key, label string) (time.Duration, error) {
	raw := strings.TrimSpace(viper.GetString(key))
	if raw == "" {
		return 0, fmt.Errorf("missing %s", label)
	}
	parsed, err := time.ParseDuration(raw)
	if err != nil || parsed <= 0 {
		return 0, fmt.Errorf("invalid %s: %q", label, raw)
	}
	return parsed, nil
}

func configureViperDefaults() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	viper.SetDefault("claim_poll_interval", defaultClaimPollInterval.String())
	viper.SetDefault("claim_poll_timeout", defaultClaimPollTimeout.String())
	viper.SetDefault("check_offer", "caffe-nero")
	viper.SetDefault("claim_offer", "caffe-nero")
	viper.SetDefault("watch_offer", "caffe-nero")
	viper.SetDefault("watch_interval", "30s")
	viper.SetDefault("watch_auto_claim", false)
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_format", "text")
}
