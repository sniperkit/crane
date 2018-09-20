package runner_test

import (
	"testing"

	// external
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// internal
	"github.com/sniperkit/snk.fork.michaelsauter-crane/pkg/repack/runner"
)

func TestNewCmdWithRunner(t *testing.T) {
	for i, currCase := range []struct {
		input       string
		expected    string
		expectedErr string
	}{
		{input: "foo", expected: "foo"},
		{input: "Bar", expected: "Bar"},
		{input: "punc_is-okay", expected: "punc_is-okay"},
		{input: "", expectedErr: "Cmd cannot be blank"},
		{input: "no whitespace", expectedErr: `Cmd cannot contain whitespace: "no whitespace"`},
		{input: "	noleadingwhitespace", expectedErr: `Cmd cannot contain whitespace: "	noleadingwhitespace"`},
		{input: "notrailingwhitespace ", expectedErr: `Cmd cannot contain whitespace: "notrailingwhitespace "`},
	} {
		actual, err := runner.NewCmdWithRunner(currCase.input, nil)

		if currCase.expectedErr != "" {
			assert.Error(t, err, currCase.expectedErr, "Case %d", i)
		} else {
			assert.Equal(t, currCase.expected, actual.Name(), "Case %d", i)
		}
	}
}

func TestCmdLibraryNewCmd(t *testing.T) {
	runner, err := runner.NewCmdWithRunner("foo", nil)
	require.NoError(t, err)
	cmdSet, err := runner.NewStringCmdSetForRunners(runner)
	require.NoError(t, err)
	cmdLibrary := runner.NewCmdLibrary(cmdSet)

	for i, currCase := range []struct {
		cmdName         string
		expectedCmdName string
		expectedErr     string
	}{
		// NewCmd succeeds for command that exists
		{cmdName: "foo", expectedCmdName: "foo"},
		// NewCmd fails for command that does not exist
		{cmdName: "bar", expectedErr: "invalid command \"bar\" (valid values: [foo])"},
	} {
		cmd, err := cmdLibrary.NewCmd(currCase.cmdName)

		if currCase.expectedErr != "" {
			assert.EqualError(t, err, currCase.expectedErr, "Case %d", i)
		} else {
			require.NoError(t, err, "Case %d", i)
		}

		if err == nil {
			assert.Equal(t, currCase.expectedCmdName, cmd.Name(), "Case %d", i)
		}
	}
}

func TestNewStringCmdSetForRunnersErrorsOnDuplicates(t *testing.T) {
	for i, currCase := range []struct {
		names         []string
		expectedError string
	}{
		// fail on duplicate names
		{
			names: []string{
				"foo",
				"foo",
			},
			expectedError: `multiple runners provided for commands: \[foo\]`,
		},
		// only report duplicate name once even if it is duplicated multiple times
		{
			names: []string{
				"foo",
				"bar",
				"foo",
				"foo",
			},
			expectedError: `multiple runners provided for commands: \[foo\]`,
		},
		// report based on order in which duplicate elements occur
		{
			names: []string{
				"foo",
				"zoo",
				"zoo",
				"bar",
				"foo",
				"bar",
			},
			expectedError: `multiple runners provided for commands: \[zoo foo bar\]`,
		},
	} {
		var runners []*runner.CmdWithRunner
		for _, name := range currCase.names {
			runner, err := runner.NewCmdWithRunner(name, nil)
			require.NoError(t, err, "Case %d", i)
			runners = append(runners, runner)
		}

		_, err := runner.NewStringCmdSetForRunners(runners...)
		require.Error(t, err)

		assert.Regexp(t, currCase.expectedError, err.Error())
	}
}
