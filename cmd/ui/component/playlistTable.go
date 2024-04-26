package component

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type PlaylistTable struct {
	Table table.Model
}

func (p PlaylistTable) Init() tea.Cmd {
	return nil
}

func (p PlaylistTable) View() string {
	return p.Table.View()
}

func (p PlaylistTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, defaultKeys.ESC):
			if p.Table.Focused() {
				p.Table.Blur()
			} else {
				p.Table.Focus()
			}

		case key.Matches(msg, defaultKeys.Up):
			fmt.Println("going up")
			p.Table.MoveUp(1)
		case key.Matches(msg, defaultKeys.Down):
			fmt.Println("going down")
			p.Table.MoveDown(1)

		}
	}
	return p, cmd
}

func InitPlaylist() PlaylistTable {
	colums := []table.Column{
		{Title: "Playlist", Width: 20},
		{Title: "Total Songs", Width: 15},
	}

	rows := []table.Row{
		{"Chill Coding Playlist", "312"},
		{"Old Vibes", "123"},
		{"Liked Songs", "4532"},
	}

	modelTable := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithWidth(50),
	)
	style := table.DefaultStyles()

	modelTable.SetStyles(style)

	m := PlaylistTable{
		Table: modelTable,
	}
	return m

}
