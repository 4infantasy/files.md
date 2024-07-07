package insights

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"zakirullin/stuffbot/internal/fs"
)

//go:embed testdata/month_habits.md
var monthMD string

//go:embed testdata/last_month_habits.md
var lastMonthMD string

func TestRead(t *testing.T) {
	r := require.New(t)

	botFS, err := fs.NewFS("/", afero.NewMemMapFs())
	r.NoError(err)
	botFS.Put(fs.DirInsights,  "1970 Habits.md",  monthMD)

	habits, err := Read(botFS, 1970)
	r.NoError(err)

	r.Len(habits, 6)
	year, ok := habits["Went to gym"]
	r.True(ok)

	r.Len(year, 31)

	completed, ok := year[1]
	r.True(ok)
	r.Equal(false, completed)

	completed, ok = year[31]
	r.True(ok)
	r.Equal(true, completed)
}

func TestReadLastMonthHabits(t *testing.T) {
	r := require.New(t)

	botFS, err := fs.NewFS("/", afero.NewMemMapFs())
	r.NoError(err)
	botFS.Put(fs.DirInsights,  "1970 Habits.md",  lastMonthMD)

	habits, err := Read(botFS, 1970)
	r.NoError(err)

	r.Len(habits, 1)
	year, ok := habits["Habit"]
	r.True(ok)

	r.Len(year, 31)

	fmt.Printf("%v", year)
	completed, ok := year[335]
	r.True(ok)
	r.Equal(false, completed)

	completed, ok = year[365]
	r.True(ok)
	r.Equal(true, completed)
}
