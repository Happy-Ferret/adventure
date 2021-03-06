package game

import (
	"fmt"
	"sort"
	"strings"

	"github.com/abiosoft/ishell"
)

const (
	welcomeMessage     = "Adventure!"
	byeMessage         = "Good bye! Player 1"
	genericMessage     = "You don’t know how to do that."
	notCarryingMessage = "You are not carrying anything."
	carryingMessage    = "You are carrying: "
	canPerformMessage  = "You can perform the following commands:\n"
	cantGoMessage      = "You can’t go there!"
)

type Game struct {
	*ishell.Shell
	Player *Player
	Places Places
}

func NewGame() *Game {
	g := &Game{ishell.NewShell(), NewPlayer(), NewPlaces()}

	g.Setup()

	return g
}

func Start() {
	g := NewGame()

	printLogo(g)

	g.Println(welcomeMessage)
	g.Start()
}

func (g *Game) Setup() {
	g.SetPrompt("▶ ")
	g.RegisterGeneric(g.generic)

	for n, c := range g.commands() {
		g.Register(n, c)

		if n != "exit" {
			g.Register(string(n[0]), c)
		}
	}
}

func (g *Game) Place() *Place {
	if p := g.Places[g.Player.Position]; p != nil {
		return p
	}

	return &Place{Name: "void", Paths: []string{"nowhere"}}
}

func (g *Game) commands() map[string]func(...string) (string, error) {
	return map[string]func(...string) (string, error){
		"drop":     g.drop,
		"exit":     g.exit,
		"help":     g.help,
		"items":    g.items,
		"look":     g.look,
		"take":     g.take,
		"teleport": g.teleport,
		"use":      g.use,
		"walk":     g.walk,
	}
}

func (g *Game) generic(args ...string) (string, error) {
	return genericMessage, nil
}

func (g *Game) items(args ...string) (string, error) {
	items := g.Player.Items

	if len(items) == 0 {
		return notCarryingMessage, nil
	}

	return carryingMessage + items.String(), nil
}

func (g *Game) help(args ...string) (string, error) {
	commandNames := []string{}

	for n, _ := range g.commands() {
		commandNames = append(commandNames, n)
	}

	sort.Strings(commandNames)

	return canPerformMessage + strings.Join(commandNames, ", "), nil
}

func (g *Game) look(args ...string) (string, error) {
	p := g.Place()

	if p.Look != nil {
		text, err := p.Look(g)
		if err != nil {
			return err.Error(), nil
		}

		return text + "\n\n" + p.describe(), nil
	}

	return p.describe(), nil
}

func (g *Game) take(args ...string) (string, error) {
	if len(args) == 0 {
		return "You didn’t tell me what to take.", nil
	}

	l := []string{}

	for _, name := range args {
		if item, ok := g.takeItem(name); ok {
			if item.Take != nil {
				return item.Take(g), nil
			}

			l = append(l, "You took the "+name)
		}
	}

	if len(l) > 0 {
		return strings.Join(l, "\n"), nil
	}

	return "You can’t take that which doesn’t exist.", nil
}

func (g *Game) takeItem(name string) (*Item, bool) {
	p := g.Place()

	if item, ok := p.Items[name]; ok {
		delete(p.Items, name)
		g.Player.AddItem(name, item)

		return item, true
	}

	return nil, false
}

func (g *Game) drop(args ...string) (string, error) {
	if len(args) == 0 {
		return "You didn’t tell me what to drop.", nil
	}

	l := []string{}

	for _, name := range args {
		if g.dropItem(name) {
			l = append(l, "You dropped the "+name)
		}
	}

	if len(l) > 0 {
		return strings.Join(l, "\n"), nil
	}

	return "Unable to drop something you are not carrying.", nil
}

func (g *Game) dropItem(name string) bool {
	if item, ok := g.Player.Items[name]; ok {
		delete(g.Player.Items, name)
		g.Place().AddItem(name, item)

		return true
	}

	return false
}

func (g *Game) use(args ...string) (string, error) {
	if len(args) == 0 {
		return "You didn’t tell me what to use.", nil
	}

	name := args[0]

	if item, ok := g.Player.Items[name]; ok {
		if item.Use != nil {
			return item.Use(g), nil
		}

		return "You can’t use the " + name, nil
	}

	return "You are not carrying that item.", nil
}

func (g *Game) teleport(args ...string) (string, error) {
	if !g.Player.HasItem("teleporter") {
		return "You need to have the teleporter in order to teleport.", nil
	}

	if len(args) != 1 {
		return "You can only teleport to a single place.", nil
	}

	return g.teleportPlayerTo(strings.ToLower(args[0])), nil
}

func (g *Game) teleportPlayerTo(target string) string {
	p := g.Places[target]

	if p.VisitCount == 0 {
		return "You can’t teleport to places you haven’t visited before."
	}

	p.VisitCount++
	g.Player.Position = target

	return fmt.Sprintf("Teleported to %s", target)
}

func (g *Game) walk(args ...string) (string, error) {
	if len(args) == 0 {
		return "You need to specify where to go.", nil
	}

	l := []string{}

	for _, name := range args {
		p := g.Place()

		target := strings.ToLower(name)

		if p.Name == target {
			return "You are already in the " + p.Name, nil
		}

		if p.IsNextTo(target) {
			tp := g.Places[target]

			var err error

			if tp.Enter != nil {
				_, err = tp.Enter(g)
			}

			if err != nil {
				l = append(l, err.Error())
			} else {
				g.Player.Position = target
				g.Places[target].VisitCount++

				l = append(l, "You walked to the "+target)
			}
		} else {
			l = append(l, "You are not next to the "+target+", staying in the "+g.Player.Position)
		}
	}

	if len(l) > 0 {
		return strings.Join(l, "\n"), nil
	}

	return cantGoMessage, nil
}

func (g *Game) exit(args ...string) (string, error) {
	g.Stop()

	return byeMessage, nil
}

func printLogo(g *Game) {
	for idx, l := range strings.Split(logo, "\n") {
		if idx < 4 {
			g.Println("\x1b[36m" + l + "\x1b[39m")
		} else if idx < 7 {
			g.Println("\x1b[34m" + l + "\x1b[39m")
		} else {
			g.Println("\x1b[37m" + l + "\x1b[39m")
		}
	}
}

var logo = `
 ▄▄▄      ▓█████▄  ██▒   █▓▓█████ ███▄    █ ▄▄▄█████▓ █    ██  ██▀███  ▓█████
▒████▄    ▒██▀ ██▌▓██░   █▒▓█   ▀ ██ ▀█   █ ▓  ██▒ ▓▒ ██  ▓██▒▓██ ▒ ██▒▓█   ▀
▒██  ▀█▄  ░██   █▌ ▓██  █▒░▒███  ▓██  ▀█ ██▒▒ ▓██░ ▒░▓██  ▒██░▓██ ░▄█ ▒▒███
░██▄▄▄▄██ ░▓█▄   ▌  ▒██ █░░▒▓█  ▄▓██▒  ▐▌██▒░ ▓██▓ ░ ▓▓█  ░██░▒██▀▀█▄  ▒▓█  ▄
 ▓█   ▓██▒░▒████▓    ▒▀█░  ░▒████▒██░   ▓██░  ▒██▒ ░ ▒▒█████▓ ░██▓ ▒██▒░▒████▒
 ▒▒   ▓▒█░ ▒▒▓  ▒    ░ ▐░  ░░ ▒░ ░ ▒░   ▒ ▒   ▒ ░░   ░▒▓▒ ▒ ▒ ░ ▒▓ ░▒▓░░░ ▒░ ░
  ▒   ▒▒ ░ ░ ▒  ▒    ░ ░░   ░ ░  ░ ░░   ░ ▒░    ░    ░░▒░ ░ ░   ░▒ ░ ▒░ ░ ░  ░
  ░   ▒    ░ ░  ░      ░░     ░     ░   ░ ░   ░       ░░░ ░ ░   ░░   ░    ░
      ░  ░   ░          ░     ░  ░        ░             ░        ░        ░  ░
           ░           ░
`
