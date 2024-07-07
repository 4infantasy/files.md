package insights

import (
	_ "embed"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"zakirullin/stuffbot/internal/fs"
)

//go:embed testdata/month_habits.md
var monthMD string

func TestRead(t *testing.T) {
	r := require.New(t)

	botFS, err := fs.NewFS("/", afero.NewMemMapFs())
	r.NoError(err)
	botFS.Put(fs.DirInsights,  "1970 Habits.md",  monthMD)

	habits, err := Read(botFS, 1970)
	r.NoError(err)

	r.Len(habits, 7)
	year, ok := habits["Went to gym"]
	r.True(ok)

	r.Len(year, 31)
}
