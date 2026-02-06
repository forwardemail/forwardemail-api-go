schema_version = 1

project {
  license        = "MIT"
  copyright_year = 2026
  copyright_holder = "Forward Email"

  header_ignore = [
    # GitHub issue template configuration
    ".github/ISSUE_TEMPLATE/*.yml",

    # golangci-lint tooling configuration
    ".golangci.yml",
  ]
}

