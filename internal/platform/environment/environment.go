package environment

import (
	"go-clean-arch/internal/platform/log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

type Environment int

var cfg *Config

const (
	Development Environment = iota
	Stage
	Production
	Testing
)

func (d Environment) String() string {
	return [...]string{"development", "stage", "production", "testing"}[d]
}

type Config struct {
	env        Environment
	configName string
}

func init() {
	scope := GetEnvScope("dev")

	cfg = getScope(scope, ".yaml")

	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName(cfg.configName)

	log.Info("Current environment:", cfg.env.String())

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Error trying to parse config %s", cfg.configName)
	}
}

func getScope(s string, suffix string) *Config {
	var match string
	re := regexp.MustCompile("(^prod[a-z]*|stag[a-z]*|test[a-z]*)-?(.+)?")

	if !re.MatchString(s) {
		return &Config{Development, Development.String() + suffix}
	}

	match = re.FindStringSubmatch(s)[1]

	for _, e := range [...]Environment{Production, Stage, Testing} {
		env := strings.ToLower(e.String())
		if strings.Contains(env, match) {
			return &Config{e, e.String() + suffix}
		}
	}

	return &Config{Development, Development.String() + suffix}
}

func GetEnvScope(defaultScope string) string {
	scope := os.Getenv("SCOPE")

	if scope == "" {
		scope = defaultScope
	}

	return scope
}

func GetEnvOrConfig(env string) string {
	value := os.Getenv(env)

	if value == "" {
		value = viper.GetString(env)
	}

	return value
}

func GetCurrentEnvironment() Environment {
	return cfg.env
}
