package command

import (
	"io"
	"os"
	"strconv"

	"github.com/docker/cli/cli/streams"
	"github.com/docker/docker/client"
	"github.com/moby/term"
)

// CLIOption applies a modification on a DockerCli.
type CLIOption func(cli *DockerCli) error

// DockerCliOption applies a modification on a DockerCli.
//
// Deprecated: use [CLIOption] instead.
type DockerCliOption = CLIOption

// InitializeOpt is the type of the functional options passed to DockerCli.Initialize
//
// Deprecated: use [CLIOption] instead.
type InitializeOpt = CLIOption

// WithStandardStreams sets a cli in, out and err streams with the standard streams.
func WithStandardStreams() CLIOption {
	return func(cli *DockerCli) error {
		// Set terminal emulation based on platform as required.
		stdin, stdout, stderr := term.StdStreams()
		cli.in = streams.NewIn(stdin)
		cli.out = streams.NewOut(stdout)
		cli.err = stderr
		return nil
	}
}

// WithCombinedStreams uses the same stream for the output and error streams.
func WithCombinedStreams(combined io.Writer) CLIOption {
	return func(cli *DockerCli) error {
		cli.out = streams.NewOut(combined)
		cli.err = combined
		return nil
	}
}

// WithInputStream sets a cli input stream.
func WithInputStream(in io.ReadCloser) CLIOption {
	return func(cli *DockerCli) error {
		cli.in = streams.NewIn(in)
		return nil
	}
}

// WithOutputStream sets a cli output stream.
func WithOutputStream(out io.Writer) CLIOption {
	return func(cli *DockerCli) error {
		cli.out = streams.NewOut(out)
		return nil
	}
}

// WithErrorStream sets a cli error stream.
func WithErrorStream(err io.Writer) CLIOption {
	return func(cli *DockerCli) error {
		cli.err = err
		return nil
	}
}

// WithContentTrustFromEnv enables content trust on a cli from environment variable DOCKER_CONTENT_TRUST value.
func WithContentTrustFromEnv() CLIOption {
	return func(cli *DockerCli) error {
		cli.contentTrust = false
		if e := os.Getenv("DOCKER_CONTENT_TRUST"); e != "" {
			if t, err := strconv.ParseBool(e); t || err != nil {
				// treat any other value as true
				cli.contentTrust = true
			}
		}
		return nil
	}
}

// WithContentTrust enables content trust on a cli.
func WithContentTrust(enabled bool) CLIOption {
	return func(cli *DockerCli) error {
		cli.contentTrust = enabled
		return nil
	}
}

// WithDefaultContextStoreConfig configures the cli to use the default context store configuration.
func WithDefaultContextStoreConfig() CLIOption {
	return func(cli *DockerCli) error {
		cli.contextStoreConfig = DefaultContextStoreConfig()
		return nil
	}
}

// WithAPIClient configures the cli to use the given API client.
func WithAPIClient(c client.APIClient) CLIOption {
	return func(cli *DockerCli) error {
		cli.client = c
		return nil
	}
}
