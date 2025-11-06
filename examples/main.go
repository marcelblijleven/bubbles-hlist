package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelblijleven/bubbles-hlist/hlist"
)

var (
	frame      = lipgloss.NewStyle().Padding(1)
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(2).
			MarginBottom(2).
			Foreground(lipgloss.Color("80"))
)

type Model struct {
	items hlist.Model
}

func New() Model {
	delegate := hlist.NewDefaultDelegate()
	items := hlist.New(newItems(), delegate, 0, 0)

	return Model{
		items: items,
	}
}

// Init satisfies tea.Model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update satisfies tea.Model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		frameWidth, _ := frame.GetFrameSize()
		m.items.SetHeight(5)
		m.items.SetWidth(msg.Width - frameWidth)
	}

	// Make sure to send msgs to the nested hlist model
	m.items, cmd = m.items.Update(msg)
	return m, cmd
}

// View satisfies tea.Model
func (m Model) View() string {
	content := lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("hlist example"),
		m.items.View(),
	)
	return frame.Render(content)
}

type item struct {
	title       string
	description string
}

// FilterValue satisfies hlist.Item
func (i item) FilterValue() string {
	return i.title
}

// Title satisfies hlist.DefaultItem
func (i item) Title() string {
	return i.title
}

// Title satisfies hlist.DefaultItem
func (i item) Description() string {
	return i.description
}

func newItems() []hlist.Item {
	return []hlist.Item{
		item{title: "Plastic Centipede", description: "Cold crawl"},
		item{title: "Digital Sorcerer", description: "Code and dust"},
		item{title: "Nonagon Loop", description: "Never ends"},
		item{title: "Acid Canyon", description: "Glass river"},
		item{title: "Gamma Ritual", description: "Burning rite"},
		item{title: "Mind Melt Motel", description: "No exit"},
		item{title: "Robot Lagoon", description: "Still metal"},
		item{title: "Cosmic Biscuit", description: "Tiny portal"},
		item{title: "Lava Lamp God", description: "Soft chaos"},
		item{title: "Interstellar Tadpole", description: "Young star"},
		item{title: "Crystalline Lizard", description: "Sharp shine"},
		item{title: "Boogieman Nebula", description: "Dark bloom"},
		item{title: "Electric Druid", description: "Wired roots"},
		item{title: "Honeycomb Reactor", description: "Sweet core"},
		item{title: "The Eternal Flute", description: "Endless tone"},
		item{title: "Warped Orchard", description: "Bent fruit"},
		item{title: "Magnetic Wombat", description: "Pulling chaos"},
		item{title: "Sun Drip City", description: "Melted noon"},
		item{title: "Dream Slug", description: "Slow light"},
		item{title: "Turbo Lich", description: "Fast death"},
		item{title: "Neon Bog Witch", description: "Glows mad"},
		item{title: "Volcano Disco", description: "Hot floor"},
		item{title: "Quantum Alpaca", description: "Split fluff"},
		item{title: "Glass Crocodile", description: "Clear bite"},
		item{title: "Cosmo Grub", description: "Star eater"},
		item{title: "Vapor Serpent", description: "Mist coil"},
		item{title: "Microtonal Ghost", description: "Half seen"},
		item{title: "Toad of Time", description: "Old leap"},
	}
}

func main() {
	m := New()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
