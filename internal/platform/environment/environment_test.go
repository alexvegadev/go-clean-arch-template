package environment

import (
	"os"
	"testing"
)

func TestGetEnvOrConfig(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return empty value",
			args: args{
				"PORT",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvOrConfig(tt.args.env); got != tt.want {
				t.Errorf("GetEnvOrConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentEnvironment(t *testing.T) {
	tests := []struct {
		name string
		want Environment
	}{
		{
			name: "Should return development environment",
			want: Development,
		},
		{
			name: "Should return production environment",
			want: Production,
		},
		{
			name: "Should return stage environment",
			want: Stage,
		},
		{
			name: "Should return test environment",
			want: Testing,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != Development {
				os.Setenv("SCOPE", tt.want.String())
				loadScope()
			}
			if got := GetCurrentEnvironment(); got != tt.want {
				t.Errorf("GetCurrentEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvScope(t *testing.T) {
	type args struct {
		defaultScope string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `Should return current environment ("DEV" in this case)`,
			args: args{defaultScope: "dev"},
			want: "dev",
		},
		{
			name: `Should return current environment ("testing" in this case)`,
			args: args{defaultScope: "testing"},
			want: "testing",
		},
		{
			name: `Should return current environment ("testing" in this case)`,
			args: args{defaultScope: "testing"},
			want: "production",
		},
		{
			name: `Should return current environment ("testing" in this case)`,
			args: args{defaultScope: "testing"},
			want: "stage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("SCOPE", tt.want)
			if got := GetEnvScope(tt.args.defaultScope); got != tt.want {
				t.Errorf("GetEnvScope() = %v, want %v", got, tt.want)
			}
		})
	}
}
