/*
Copyright (c) 2021 Gemba Advantage

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gembaadvantage/uplift/internal/config"
	"github.com/gembaadvantage/uplift/internal/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBump(t *testing.T) {
	taggedRepo(t)
	data := testFileWithConfig(t, "test.txt", ".uplift.yml")

	cfg, _ := config.Load(".uplift.yml")
	cmd := newBumpCmd(&context.Context{Config: cfg})

	err := cmd.Execute()
	require.NoError(t, err)

	actual, err := ioutil.ReadFile("test.txt")
	require.NoError(t, err)
	assert.NotEqual(t, string(data), string(actual))
}

func testFileWithConfig(t *testing.T, f string, cfg string) []byte {
	t.Helper()

	c := []byte("version: 0.0.0")
	err := ioutil.WriteFile(f, c, 0644)
	require.NoError(t, err)

	yml := fmt.Sprintf(`
bumps:
  - file: %s
    regex: "version: $VERSION"`, f)

	err = ioutil.WriteFile(cfg, []byte(yml), 0644)
	require.NoError(t, err)

	return c
}

func TestBump_PrereleaseFlag(t *testing.T) {
	untaggedRepo(t)
	testFileWithConfig(t, "test.txt", ".uplift.yml")

	cfg, _ := config.Load(".uplift.yml")
	cmd := newBumpCmd(&context.Context{Config: cfg})
	cmd.SetArgs([]string{"--prerelease", "-beta.1+12345"})

	err := cmd.Execute()
	require.NoError(t, err)

	actual, err := ioutil.ReadFile("test.txt")
	require.NoError(t, err)
	assert.Equal(t, "version: 0.1.0-beta.1+12345", string(actual))
}