package templates

const GithubCiCdWorkflowTemplate = `# This workflow will build a golang project
# For more information see: https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go

name: Go

on:
  push:
    branches: [ "root" ]
  pull_request:
    branches: [ "root" ]

jobs:

  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '{{ .GolangVersion }}'

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...
`
