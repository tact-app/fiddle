// Code generated by github.com/kamilsk/egg. DO NOT EDIT.

//go:build tools

package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/godownloader"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/MontFerret/cli/ferret"
	_ "github.com/suntong/html2md"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/vuln/cmd/govulncheck"
)

//go:generate go install github.com/golang/mock/mockgen
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install github.com/goreleaser/godownloader
//go:generate go install github.com/goreleaser/goreleaser
//go:generate go install github.com/MontFerret/cli/ferret
//go:generate go install github.com/suntong/html2md
//go:generate go install golang.org/x/tools/cmd/goimports
//go:generate go install golang.org/x/vuln/cmd/govulncheck
