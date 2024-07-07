package insights

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/uniseg"

	"zakirullin/stuffbot/internal/fs"
	"zakirullin/stuffbot/pkg/text"
)


// [1 => false, <year day> => false, ...] 
type Year map[int]bool

const (
    habitSkipped = "⚪️"
	habitCompleted = "🟢"
    habitCompletedAtWeekend = "🟡"
)

// getLastWeekHabits
// getLastMonthHabits

func Read(botFS *fs.FS, year int) (map[string]Year, error) {
	filename := "%d Habits.md"
	habitsStr, err := botFS.Content(fs.DirInsights, fmt.Sprintf(filename, year))
	if err != nil {
		return nil, fmt.Errorf("read %s error: %w", filename, err)
	}

	habits := make(map[string]Year)
	month := time.January
	lines := strings.Split(text.NormNewLines(habitsStr), "\n")
	for _, line := range(lines) {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		isMonthLine := strings.HasPrefix(line, "###")
		if isMonthLine {
			parts := strings.Split(line, " ")
			if len(parts) < 2 {
				return nil, nil
				// return "bad month line: %s"
			}

			date, err := time.Parse("January", parts[1])
			if err != nil {
				return nil, nil
			}
			month = date.Month()

			continue
		}

		// At this moment we have a habits line, which is
		// [⚪️🟢 Habit name] for the above found month

		parts := strings.SplitN(line, " ", 2)
		if len(parts) < 2 {
			return nil, nil
			// return "bad month line: %s"
		}
		habitName := strings.TrimSpace(parts[1])
		if _, ok := habits[habitName]; !ok {
			habits[habitName] = make(Year)
		}

		firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		dayOfTheYear := firstDayOfMonth.YearDay()

		days := parts[0]
		gr := uniseg.NewGraphemes(days)

		dayOffset := 0
    	for gr.Next() {
			habits[habitName][dayOfTheYear + dayOffset] = gr.Str() != habitSkipped
			dayOfTheYear++
    	}
	}

	return habits, nil
}

// func Write(botFS *fs.FS, habits []Habit) error {
// 	return nil
// }