package before

import (
	"testing"

	"github.com/gembaadvantage/uplift/internal/config"
	"github.com/gembaadvantage/uplift/internal/context"
	"github.com/purpleclay/gitz/gittest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	assert.Equal(t, "before hooks", Task{}.String())
}

func TestSkip(t *testing.T) {
	cmd := []string{"echo 'HELLO'"}

	assert.True(t, Task{}.Skip(&context.Context{
		Config: config.Uplift{
			Hooks: &config.Hooks{
				Before:          []string{},
				BeforeBump:      cmd,
				BeforeTag:       cmd,
				BeforeChangelog: cmd,
				After:           cmd,
				AfterBump:       cmd,
				AfterTag:        cmd,
				AfterChangelog:  cmd,
			},
		},
	}))
}

func TestRun(t *testing.T) {
	// git.MkTmpDir(t)
	gittest.InitRepository(t)

	err := Task{}.Run(&context.Context{
		Config: config.Uplift{
			Hooks: &config.Hooks{
				Before: []string{"touch a.out"},
			},
		},
	})
	require.NoError(t, err)
	assert.FileExists(t, "a.out")
}
