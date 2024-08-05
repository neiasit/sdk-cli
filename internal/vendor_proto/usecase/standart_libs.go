package usecase

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/labstack/gommon/color"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	vendorProtobufFolder = "vendor.protobuf"
)

func VendorStandardLibraries() error {
	steps := []struct {
		name string
		fn   func() error
	}{
		{"Reset vendor directory", resetVendorDirectory},
		{"Vendor google-protobuf", vendorGoogleProtobuf},
		{"Vendor googleapis", vendorGoogleAPIs},
		{"Vendor protoc-gen-openapiv2", vendorProtocGenOpenapiv2},
		{"Vendor protovalidate", vendorProtovalidate},
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Используем анимацию из CharSets
	err := s.Color("yellow")
	if err != nil {
		return err
	}

	for _, step := range steps {
		s.Suffix = " " + step.name
		s.Start()
		err := step.fn()
		s.Stop()
		if err != nil {
			return fmt.Errorf("failed to %s: %w", step.name, err)
		}
		fmt.Println(color.Green(fmt.Sprintf("✔ %s", step.name)))
	}

	return nil
}

func resetVendorDirectory() error {
	err := os.RemoveAll(vendorProtobufFolder)
	if err != nil {
		return err
	}
	return os.Mkdir(vendorProtobufFolder, os.ModePerm)
}

func vendorGoogleProtobuf() error {
	return runGitSparseCheckout(
		"https://github.com/protocolbuffers/protobuf",
		"main",
		"src/google/protobuf",
		"protobuf/src/google/protobuf",
	)
}

func vendorGoogleAPIs() error {
	return runGitSparseCheckout(
		"https://github.com/googleapis/googleapis",
		"master",
		"google",
		"googleapis/google",
	)
}

func vendorProtocGenOpenapiv2() error {
	return runGitSparseCheckout(
		"https://github.com/grpc-ecosystem/grpc-gateway",
		"main",
		"protoc-gen-openapiv2/options",
		"grpc-gateway/protoc-gen-openapiv2/options",
	)
}

func vendorProtovalidate() error {
	return runGitClone(
		"https://github.com/bufbuild/protovalidate",
		"main",
		"protovalidate/proto/protovalidate/buf",
	)
}

func runGitSparseCheckout(repo, branch, sparsePath, destPath string) error {
	dir := filepath.Join(vendorProtobufFolder, filepath.Dir(destPath))
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "clone", "-b", branch, "--single-branch", "--depth=1", "--filter=tree:0", repo, filepath.Join(vendorProtobufFolder, destPath))
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "-C", filepath.Join(vendorProtobufFolder, destPath), "sparse-checkout", "set", "--no-cone", sparsePath)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	err = cmd.Run()
	if err != nil {
		return err
	}
	return os.RemoveAll(filepath.Join(vendorProtobufFolder, destPath, ".git"))
}

func runGitClone(repo, branch, destPath string) error {
	cmd := exec.Command("git", "clone", "-b", branch, "--single-branch", "--depth=1", repo, filepath.Join(vendorProtobufFolder, destPath))
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run()
}
