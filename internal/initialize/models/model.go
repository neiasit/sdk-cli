package models

type ProjectData struct {
	ProjectFolder string
	ProjectName   string
	GolangVersion string
	Libraries     []string
	OtherOptions  []string
}

var PlatformLibraries = map[string]string{
	"auth":         "github.com/neiasit/auth-library",
	"logging":      "github.com/neiasit/logging-library",
	"redis":        "github.com/neiasit/redis-library",
	"grpc":         "github.com/neiasit/grpc-library",
	"http-support": "github.com/neiasit/http-support-library",
}

var GolangVersions = []string{
	"1.22",
	"1.21",
	"1.20",
}

var AdditionalOptions = []string{
	GitInitializationOption,
	GithubCiCdOption,
	DockerOption,
}

const (
	GitInitializationOption = "Git Initialization"
	GithubCiCdOption        = "Github CI&CD"
	DockerOption            = "Dockerfile"
)
