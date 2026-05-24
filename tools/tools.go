//go:build generate

package tools

import (
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
	_ "github.com/hashicorp/copywrite"
	_ "golang.org/x/tools/cmd/goimports"
	_ "mvdan.cc/gofumpt"
)

// Generate copyright headers
//go:generate go run github.com/hashicorp/copywrite headers -d .. --config ../.copywrite.hcl

// Run linters
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint run --config ../.golangci.yml ../...
