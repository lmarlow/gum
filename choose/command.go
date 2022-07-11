package choose

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

// Run provides a shell script interface for choosing between different through
// options.
func (o Options) Run() {
	items := []list.Item{}
	for _, option := range o.Options {
		if option == "" {
			continue
		}
		items = append(items, item(option))
	}

	const defaultWidth = 20

	id := itemDelegate{
		indicator:         o.Indicator,
		indicatorStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(o.IndicatorColor)),
		itemStyle:         lipgloss.NewStyle().Padding(0, runewidth.StringWidth(o.Indicator)).Foreground(lipgloss.Color(o.UnselectedColor)),
		selectedItemStyle: lipgloss.NewStyle().Foreground(lipgloss.Color(o.SelectedColor)),
	}

	l := list.New(items, id, defaultWidth, o.Height)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowPagination(!o.HidePagination)

	m, err := tea.NewProgram(model{list: l}, tea.WithOutput(os.Stderr)).StartReturningModel()

	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	fmt.Println(m.(model).choice)
}
