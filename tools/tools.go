//go:build generate

package tools

import (
	_ "github.com/hashicorp/copywrite"
	_ "mvdan.cc/gofumpt"
	_ "golang.org/x/tools/cmd/goimports"
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
)

// Generate copyright headers
//go:generate go run github.com/hashicorp/copywrite headers -d .. --config ../.copywrite.hcl

// Run linters
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint run --config ../.golangci.yml ../...
