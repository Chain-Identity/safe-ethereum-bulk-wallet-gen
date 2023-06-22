package progressBar

import (
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

const DefaultStandardModeTemplate = `{{counters . }} | {{bar . "" "█" "█" "" "" | rndcolor}} | {{percent . }} | {{speed . }} | {{string . "resolved"}}`

// ProgressBar progress bar interface
type ProgressBar interface {
	Increment() error
	Finish() error
}

// progressBar logs data
type progressBar struct {
	Bar *progressbar.ProgressBar
}

// NewCompatibleProgressBar returns a new progress bar set to compatible mode.
func NewCompatibleProgressBar(number int) ProgressBar {
	bar := progressbar.NewOptions(number,
		progressbar.OptionSetItsString("w"),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionShowIts(),
		progressbar.OptionShowCount(),
		progressbar.OptionFullWidth(),
	)
	_ = bar.RenderBlank()
	return &progressBar{
		Bar: bar,
	}
}

// Increment increment progress
func (bar *progressBar) Increment() error {
	if bar.Bar != nil {
		return errors.WithStack(bar.Bar.Add(1))
	}
	return nil
}

// Finish close progress bar
func (bar *progressBar) Finish() error {
	if bar.Bar != nil {
		return errors.WithStack(bar.Bar.Close())
	}
	return errors.WithStack(bar.Bar.Finish())
}
