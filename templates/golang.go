package templates

const GoModFileTemplate = `module {{ .ProjectName }}

go {{ .GolangVersion }}
`
