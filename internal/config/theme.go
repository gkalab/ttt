package config

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

type StyleDef struct {
	Fg     string `json:"fg,omitempty"`
	Bg     string `json:"bg,omitempty"`
	Bold   bool   `json:"bold,omitempty"`
	Italic bool   `json:"italic,omitempty"`
}

type BorderChars struct {
	Horizontal  string `json:"horizontal"`
	Vertical    string `json:"vertical"`
	TopLeft     string `json:"topLeft"`
	TopRight    string `json:"topRight"`
	BottomLeft  string `json:"bottomLeft"`
	BottomRight string `json:"bottomRight"`
	TopTee      string `json:"topTee"`
	BottomTee   string `json:"bottomTee"`
	LeftTee     string `json:"leftTee"`
	RightTee    string `json:"rightTee"`
}

type TabStyles struct {
	Active   StyleDef `json:"active"`
	Inactive StyleDef `json:"inactive"`
	Selected StyleDef `json:"selected"`
}

type SidebarStyles struct {
	Header   StyleDef `json:"header"`
	Item     StyleDef `json:"item"`
	Selected StyleDef `json:"selected"`
}

type DialogStyles struct {
	Input    StyleDef `json:"input"`
	Item     StyleDef `json:"item"`
	Selected StyleDef `json:"selected"`
	Muted    StyleDef `json:"muted"`
}

type InputStyles struct {
	Item        StyleDef `json:"item"`
	Placeholder StyleDef `json:"placeholder"`
	Action      StyleDef `json:"action"`
}

type ButtonStyles struct {
	Item    StyleDef `json:"item"`
	Focused StyleDef `json:"focused"`
}

type MenuStyles struct {
	Item   StyleDef `json:"item"`
	Active StyleDef `json:"active"`
}

type DiagnosticStyles struct {
	Error   StyleDef `json:"error"`
	Warning StyleDef `json:"warning"`
	Info    StyleDef `json:"info"`
	Hint    StyleDef `json:"hint"`
}

type EditorStyles struct {
	LineNumber    StyleDef         `json:"lineNumber"`
	ActiveLine    StyleDef         `json:"activeLine"`
	Selection     StyleDef         `json:"selection"`
	SearchMatch   StyleDef         `json:"searchMatch"`
	SearchActive  StyleDef         `json:"searchActive"`
	BracketMatch  StyleDef         `json:"bracketMatch"`
	BracketColors []string         `json:"bracketColors,omitempty"`
	Diagnostics   DiagnosticStyles `json:"diagnostics"`
}

type DiffStyles struct {
	Added             StyleDef `json:"added"`
	Deleted           StyleDef `json:"deleted"`
	Modified          StyleDef `json:"modified"`
	CollapsedEmphasis StyleDef `json:"collapsedEmphasis,omitempty"`
	CollapsedHover    StyleDef `json:"collapsedHover,omitempty"`
	// Collapsed is retained for themes created against the original collapsed-gap contract.
	// Deprecated: use CollapsedHover.
	Collapsed      StyleDef `json:"collapsed,omitempty"`
	GutterAdded    StyleDef `json:"gutterAdded,omitempty"`
	GutterDeleted  StyleDef `json:"gutterDeleted,omitempty"`
	GutterModified StyleDef `json:"gutterModified,omitempty"`
}

type SyntaxStyles struct {
	Comment     StyleDef `json:"comment"`
	String      StyleDef `json:"string"`
	Keyword     StyleDef `json:"keyword"`
	Number      StyleDef `json:"number"`
	Operator    StyleDef `json:"operator"`
	Function    StyleDef `json:"function"`
	Type        StyleDef `json:"type"`
	Builtin     StyleDef `json:"builtin"`
	Variable    StyleDef `json:"variable"`
	Punctuation StyleDef `json:"punctuation"`
	Tag         StyleDef `json:"tag"`
	Attribute   StyleDef `json:"attribute"`
}

type TerminalColors struct {
	Foreground    string `json:"foreground,omitempty"`
	Background    string `json:"background,omitempty"`
	Black         string `json:"black,omitempty"`
	Red           string `json:"red,omitempty"`
	Green         string `json:"green,omitempty"`
	Yellow        string `json:"yellow,omitempty"`
	Blue          string `json:"blue,omitempty"`
	Magenta       string `json:"magenta,omitempty"`
	Cyan          string `json:"cyan,omitempty"`
	White         string `json:"white,omitempty"`
	BrightBlack   string `json:"brightBlack,omitempty"`
	BrightRed     string `json:"brightRed,omitempty"`
	BrightGreen   string `json:"brightGreen,omitempty"`
	BrightYellow  string `json:"brightYellow,omitempty"`
	BrightBlue    string `json:"brightBlue,omitempty"`
	BrightMagenta string `json:"brightMagenta,omitempty"`
	BrightCyan    string `json:"brightCyan,omitempty"`
	BrightWhite   string `json:"brightWhite,omitempty"`
}

func DefaultTerminalColors() TerminalColors {
	return TerminalColors{
		Black:         "#1e1e1e",
		Red:           "#f44747",
		Green:         "#6a9955",
		Yellow:        "#d7ba7d",
		Blue:          "#569cd6",
		Magenta:       "#c586c0",
		Cyan:          "#4ec9b0",
		White:         "#d4d4d4",
		BrightBlack:   "#808080",
		BrightRed:     "#f14c4c",
		BrightGreen:   "#73c991",
		BrightYellow:  "#e2c08d",
		BrightBlue:    "#6cb6ff",
		BrightMagenta: "#d670d6",
		BrightCyan:    "#58d1c9",
		BrightWhite:   "#e5e5e5",
	}
}

// ANSIPalette returns the 16 ANSI colors as an ordered array [0..15].
func (tc TerminalColors) ANSIPalette() [16]string {
	return [16]string{
		tc.Black, tc.Red, tc.Green, tc.Yellow,
		tc.Blue, tc.Magenta, tc.Cyan, tc.White,
		tc.BrightBlack, tc.BrightRed, tc.BrightGreen, tc.BrightYellow,
		tc.BrightBlue, tc.BrightMagenta, tc.BrightCyan, tc.BrightWhite,
	}
}

func (tc TerminalColors) ColorByName(name string) string {
	switch name {
	case "black":
		return tc.Black
	case "red":
		return tc.Red
	case "green":
		return tc.Green
	case "yellow":
		return tc.Yellow
	case "blue":
		return tc.Blue
	case "magenta":
		return tc.Magenta
	case "cyan":
		return tc.Cyan
	case "white":
		return tc.White
	case "brightBlack":
		return tc.BrightBlack
	case "brightRed":
		return tc.BrightRed
	case "brightGreen":
		return tc.BrightGreen
	case "brightYellow":
		return tc.BrightYellow
	case "brightBlue":
		return tc.BrightBlue
	case "brightMagenta":
		return tc.BrightMagenta
	case "brightCyan":
		return tc.BrightCyan
	case "brightWhite":
		return tc.BrightWhite
	}
	return ""
}

type HoverStyles struct {
	Bold   StyleDef `json:"bold"`
	Italic StyleDef `json:"italic"`
	Code   StyleDef `json:"code"`
}

type ThemeConfig struct {
	Default      StyleDef       `json:"default"`
	Muted        StyleDef       `json:"muted"`
	Success      StyleDef       `json:"success"`
	Danger       StyleDef       `json:"danger"`
	Warning      StyleDef       `json:"warning"`
	StatusBar    StyleDef       `json:"statusBar"`
	CommitHeader StyleDef       `json:"commitHeader"`
	Tabs         TabStyles      `json:"tabs"`
	Sidebar      SidebarStyles  `json:"sidebar"`
	Dialog       DialogStyles   `json:"dialog"`
	Editor       EditorStyles   `json:"editor"`
	Menu         MenuStyles     `json:"menu"`
	Input        InputStyles    `json:"input"`
	Button       ButtonStyles   `json:"button"`
	Hover        HoverStyles    `json:"hover"`
	Border       StyleDef       `json:"border"`
	BorderActive StyleDef       `json:"borderActive"`
	Diff         DiffStyles     `json:"diff"`
	Scrollbar    StyleDef       `json:"scrollbar"`
	Syntax       SyntaxStyles   `json:"syntax"`
	Borders      BorderChars    `json:"borders"`
	Terminal     TerminalColors `json:"terminal,omitempty"`
}

func DefaultTheme() ThemeConfig {
	t := ThemeConfig{
		Terminal: DefaultTerminalColors(),
		Default:  StyleDef{Fg: "#d8dee9", Bg: "#303841"},
		Muted:    StyleDef{Fg: "#888888"},

		Menu: MenuStyles{
			Active: StyleDef{Fg: "#d8dee9", Bg: "#4f5b66", Bold: true},
		},
		StatusBar: StyleDef{},

		Tabs: TabStyles{
			Active:   StyleDef{Fg: "#d8dee9", Bold: true},
			Inactive: StyleDef{Fg: "#65737e"},
		},

		Sidebar: SidebarStyles{
			Header:   StyleDef{Fg: "#d8dee9", Bold: true},
			Selected: StyleDef{Fg: "#d8dee9", Bg: "#4f5b66"},
		},

		Dialog: DialogStyles{
			Selected: StyleDef{Fg: "#d8dee9", Bg: "#4f5b66"},
		},

		Border: StyleDef{Fg: "#4f5b66"},

		Editor: EditorStyles{
			ActiveLine:    StyleDef{Bg: "#4c5863"},
			Selection:     StyleDef{Bg: "#4f5b66"},
			LineNumber:    StyleDef{Fg: "#848b95"},
			SearchMatch:   StyleDef{Fg: "#333333", Bg: "#fac761"},
			SearchActive:  StyleDef{Fg: "#333333", Bg: "#f97b58"},
			BracketMatch:  StyleDef{Bg: "#3a3a3a"}, // TODO
			BracketColors: []string{"yellow", "magenta", "blue"}, // TODO
		},
		Scrollbar: StyleDef{Fg: "#697076", Bg: "#444c54"},
		Diff: DiffStyles{
			CollapsedEmphasis: StyleDef{Bold: true},
		},

		Syntax: SyntaxStyles{
			Comment:     StyleDef{Fg: "#a6acb9"},
			String:      StyleDef{Fg: "#99c794"},
			Keyword:     StyleDef{Fg: "#c695c6"},
			Number:      StyleDef{Fg: "#f9ae58"},
			Operator:    StyleDef{Fg: "#f97b58"},
			Function:    StyleDef{Fg: "#5fb4b4"},
			Type:        StyleDef{Fg: "#6699cc"},
			Builtin:     StyleDef{Fg: "#ec5f66"},
			Variable:    StyleDef{Fg: "#d8dee9"},
			Punctuation: StyleDef{Fg: "#ffffff"},
			Tag:         StyleDef{Fg: "#c695c6"},
			Attribute:   StyleDef{Fg: "#6699cc"},
		},

		Borders: BorderChars{
			Horizontal:  "─",
			Vertical:    "│",
			TopLeft:     "╭",
			TopRight:    "╮",
			BottomLeft:  "╰",
			BottomRight: "╯",
			TopTee:      "┬",
			BottomTee:   "┴",
			LeftTee:     "├",
			RightTee:    "┤",
		},
	}
	return t
}

func (t *ThemeConfig) ResolveColors() {
	fillFg(&t.CommitHeader, t.Default.Fg)
	fillBg(&t.CommitHeader, t.Default.Bg)
	fillFg(&t.Dialog.Muted, t.Muted.Fg)
	fillBg(&t.Input.Item, t.Default.Bg)
	fillFg(&t.Input.Item, t.Default.Fg)
	fillFg(&t.Input.Placeholder, t.Muted.Fg)
	fillFg(&t.Input.Action, t.Muted.Fg)
	fillBg(&t.Button.Item, t.Default.Bg)
	fillFg(&t.Button.Item, t.Default.Fg)
	fillBg(&t.Button.Focused, t.Sidebar.Selected.Bg)
	fillFg(&t.Button.Focused, t.Sidebar.Selected.Fg)
	fillFg(&t.BorderActive, t.Default.Fg)
	fillBg(&t.Diff.Added, "#3d686e")
	fillBg(&t.Diff.Deleted, "#62434b")
	fillBg(&t.Diff.Modified, "#394d55")
	fillFg(&t.Diff.CollapsedEmphasis, t.Default.Fg)
	emphasisBg := t.Diff.CollapsedEmphasis.Bg
	if emphasisBg == "" {
		emphasisBg = t.Default.Bg
	}
	t.Diff.CollapsedEmphasis.Fg = contrastSafeForeground(t.Diff.CollapsedEmphasis.Fg, emphasisBg, t.Default.Fg)
	if t.Diff.CollapsedHover == (StyleDef{}) && t.Diff.Collapsed != (StyleDef{}) {
		t.Diff.CollapsedHover = t.Diff.Collapsed
	}
	fillBg(&t.Diff.CollapsedHover, t.Editor.ActiveLine.Bg)
	fillFg(&t.Diff.GutterAdded, "#99c794")
	fillFg(&t.Diff.GutterDeleted, "#ec5f66")
	fillFg(&t.Diff.GutterModified, "#f9ae58")
	t.Diff.GutterAdded.Fg = contrastSafeForeground(t.Diff.GutterAdded.Fg, t.Diff.Added.Bg, t.Default.Fg)
	t.Diff.GutterDeleted.Fg = contrastSafeForeground(t.Diff.GutterDeleted.Fg, t.Diff.Deleted.Bg, t.Default.Fg)
	fillFg(&t.Success, "#99c794")
	fillFg(&t.Danger, "#ec5f66")
	fillFg(&t.Warning, "#fac761")
	fillFg(&t.Editor.Diagnostics.Error, t.Danger.Fg)
	fillFg(&t.Editor.Diagnostics.Warning, t.Warning.Fg)
	fillFg(&t.Editor.Diagnostics.Info, t.Default.Fg)
	fillFg(&t.Editor.Diagnostics.Hint, t.Default.Fg)
	if !t.Hover.Bold.Bold {
		t.Hover.Bold.Bold = true
	}
	if !t.Hover.Italic.Italic {
		t.Hover.Italic.Italic = true
	}
	fillFg(&t.Hover.Bold, t.Default.Fg)
	fillFg(&t.Hover.Italic, t.Default.Fg)
	fillFg(&t.Hover.Code, t.Syntax.String.Fg)
}

func fillFg(s *StyleDef, color string) {
	if s.Fg == "" {
		s.Fg = color
	}
}

func fillBg(s *StyleDef, color string) {
	if s.Bg == "" {
		s.Bg = color
	}
}

type themeRGB struct {
	r float64
	g float64
	b float64
}

func parseThemeRGB(color string) (themeRGB, bool) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(color, "#"))
	if err != nil || len(decoded) != 3 {
		return themeRGB{}, false
	}
	return themeRGB{r: float64(decoded[0]), g: float64(decoded[1]), b: float64(decoded[2])}, true
}

func themeRelativeLuminance(color themeRGB) float64 {
	channel := func(value float64) float64 {
		value /= 255
		if value <= 0.04045 {
			return value / 12.92
		}
		return math.Pow((value+0.055)/1.055, 2.4)
	}
	return 0.2126*channel(color.r) + 0.7152*channel(color.g) + 0.0722*channel(color.b)
}

func themeContrast(foreground, background themeRGB) float64 {
	foregroundLuminance := themeRelativeLuminance(foreground)
	backgroundLuminance := themeRelativeLuminance(background)
	if foregroundLuminance < backgroundLuminance {
		foregroundLuminance, backgroundLuminance = backgroundLuminance, foregroundLuminance
	}
	return (foregroundLuminance + 0.05) / (backgroundLuminance + 0.05)
}

func formatThemeRGB(color themeRGB) string {
	return fmt.Sprintf("#%02x%02x%02x", int(math.Round(color.r)), int(math.Round(color.g)), int(math.Round(color.b)))
}

func contrastSafeForeground(foreground, background, fallback string) string {
	bg, bgOK := parseThemeRGB(background)
	if !bgOK {
		return foreground
	}
	fg, fgOK := parseThemeRGB(foreground)
	if !fgOK {
		fg, fgOK = parseThemeRGB(fallback)
		if !fgOK {
			return foreground
		}
		foreground = fallback
	}
	if themeContrast(fg, bg) >= 4.5 {
		return foreground
	}

	black := themeRGB{}
	white := themeRGB{r: 255, g: 255, b: 255}
	target := black
	if themeContrast(white, bg) > themeContrast(black, bg) {
		target = white
	}
	for step := 1; step <= 255; step++ {
		amount := float64(step) / 255
		candidate := themeRGB{
			r: fg.r + (target.r-fg.r)*amount,
			g: fg.g + (target.g-fg.g)*amount,
			b: fg.b + (target.b-fg.b)*amount,
		}
		resolved, _ := parseThemeRGB(formatThemeRGB(candidate))
		if themeContrast(resolved, bg) >= 4.5 {
			return formatThemeRGB(resolved)
		}
	}
	return formatThemeRGB(target)
}
