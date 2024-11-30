package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/suzuito/linter1/linter2"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("linter2", New)
}

type Settings struct {
	One string `json:"one"`
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &linterPlugin{settings: s}, nil
}

type linterPlugin struct {
	settings Settings
}

func (t *linterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		linter2.Analyzer,
	}, nil
}

func (t *linterPlugin) GetLoadMode() string {
	// return register.LoadModeSyntax
	return register.LoadModeTypesInfo
}
