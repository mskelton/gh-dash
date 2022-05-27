package tabs

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/dlvhdr/gh-dash/ui/styles"
)

var (
	tabsBorderHeight  = 1
	tabsContentHeight = 2
	TabsHeight        = tabsBorderHeight + tabsContentHeight

	tab = lipgloss.NewStyle().
		Faint(true).
		Padding(0, 1)

	activeTab = tab.
			Copy().
			Faint(false).
			Bold(true).
			Background(styles.DefaultTheme.SelectedBackground).
			Foreground(styles.DefaultTheme.MainText)

	tabsRow = lipgloss.NewStyle().
		Height(tabsContentHeight).
		PaddingTop(1).
		PaddingBottom(0).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottomForeground(styles.DefaultTheme.Border)
)
