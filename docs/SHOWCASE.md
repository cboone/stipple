# Drawille & Braille Terminal Graphics: Use Cases & Showcase

A curated collection of real-world applications, tools, and creative projects that use Unicode Braille characters for terminal graphics.

---

## Table of Contents

1. [System Monitoring & Dashboards](#1-system-monitoring--dashboards)
2. [Data Visualization & Plotting](#2-data-visualization--plotting)
3. [Financial & Cryptocurrency](#3-financial--cryptocurrency)
4. [Maps & Geospatial](#4-maps--geospatial)
5. [Image & Video Viewers](#5-image--video-viewers)
6. [Audio Visualization](#6-audio-visualization)
7. [Games & Emulators](#7-games--emulators)
8. [Developer Tools & DevOps](#8-developer-tools--devops)
9. [Creative & Artistic](#9-creative--artistic)
10. [Libraries & Frameworks](#10-libraries--frameworks)

---

## 1. System Monitoring & Dashboards

### gotop
**Terminal-based graphical activity monitor**

![gotop screenshot](https://raw.githubusercontent.com/xxxserxxx/gotop/master/assets/screenshots/demo.gif)

- **GitHub**: https://github.com/cjbassi/gotop (original), https://github.com/xxxserxxx/gotop (maintained fork)
- **Language**: Go
- **Uses**: `drawille-go` for braille-based CPU/memory graphs
- **Features**: CPU, memory, disk, network, temperature monitoring with real-time braille sparklines
- **Note**: Requires font with braille characters (U+2800-U+28FF). Recommended: Terminess Powerline, Hack, JetBrains Mono

### btop / bpytop / bashtop
**Resource monitor with braille graph options**

- **GitHub**: https://github.com/aristocratos/btop
- **Language**: C++ (btop), Python (bpytop), Bash (bashtop)
- **Graph modes**:
  - `braille` - Highest resolution (2×4 dots per cell)
  - `block` - Half resolution, more common font support
  - `tty` - 3 symbols, works with most fonts/TTYs
- **Requirements**: Unicode Block "Braille Patterns" U+2800-U+28FF

### vtop
**The original Node.js terminal activity monitor**

- **GitHub**: https://github.com/MrRio/vtop
- **Language**: Node.js
- **Uses**: `node-drawille` for charts
- **Tagline**: "A graphical activity monitor for the command line"

### lazydocker
**Docker management TUI with container metrics**

- **GitHub**: https://github.com/jesseduffield/lazydocker
- **Language**: Go
- **Features**: Real-time CPU/memory graphs for containers using ASCII/braille graphics
- **Uses**: `termui` library with braille support

---

## 2. Data Visualization & Plotting

### plotille (Python)
**Plots, scatter plots, histograms in terminal using braille**

```
   ▗▄▄▖
  ▟▀  ▀▙
 ▐▌    ▐▌
 ▐▌    ▐▌
  ▜▄  ▄▞
   ▝▀▀▘
```

- **GitHub**: https://github.com/tammoippen/plotille
- **Features**:
  - Scatter plots, line charts, histograms, heatmaps
  - Foreground and background colors
  - Timeseries support (v3.2+)
  - Zero dependencies
- **Install**: `pip install plotille`

### uniplot (Python)
**Lightweight plotting with 4x/8x resolution via Unicode**

- **GitHub**: https://github.com/olavolav/uniplot
- **Modes**:
  - `block` - 4x resolution using Block Elements
  - `braille` - 8x resolution, lighter look
  - `ascii` - ASCII-only fallback
- **Use case**: Data science CI/CD pipelines without graphics dependencies

### TextPlots.jl (Julia)
**Elegant terminal plotting**

- **GitHub**: https://github.com/sunetos/TextPlots.jl
- **Features**:
  - Function plotting: `plot(sin, -π, π)`
  - Multi-series support
  - Auto-generated titles from function names
  - Up to 8× visual resolution via braille

### termui (Go)
**Cross-platform terminal dashboard library**

- **GitHub**: https://github.com/gizak/termui
- **Widgets**: BarChart, Canvas (braille dots), Gauge, Plot, Sparkline, PieChart
- **Inspired by**: blessed-contrib (Node.js)
- **Used by**: gotop, cointop, dockdash, and many others

### dashing (Python)
**Terminal dashboards with braille-filled charts**

- **GitHub**: https://github.com/FedericoCeratto/dashing
- **Features**:
  - Sub-character resolution (partially filled blocks)
  - Braille-filled charts
  - 24-bit RGB and HSV color theming
  - Automatic scaling

---

## 3. Financial & Cryptocurrency

### cointop
**Interactive cryptocurrency tracker**

```
┌─ Rank ─┬─ Name ──────┬─ Price ─────┬─ 24H ────┬─ Sparkline ──────────┐
│      1 │ Bitcoin     │ $43,521.00  │ +2.34%   │ ⣀⣀⣠⣤⣴⣶⣾⣿⣿⣿⣿       │
│      2 │ Ethereum    │ $2,891.00   │ +1.87%   │ ⣀⣠⣴⣶⣿⣿⣿⣷⣦⣤⣀       │
└────────┴─────────────┴─────────────┴──────────┴──────────────────────┘
```

- **GitHub**: https://github.com/cointop-sh/cointop
- **Language**: Go
- **Features**:
  - Real-time price tracking
  - Portfolio management with profit/loss
  - Price charts with braille sparklines
  - Vim-style keybindings
  - Desktop notifications for price alerts
- **Memory**: Only ~15MB (runs on Raspberry Pi Zero)
- **Try it**: `ssh cointop.sh`

### Sparklines in general
**Mini inline charts for financial data**

- Braille sparklines show 2 data points per cell
- Used in: stock tickers, crypto trackers, system monitors
- Example libraries: `termui`, `blessed-contrib`, `plotille`

---

## 4. Maps & Geospatial

### MapSCII
**The entire world map in your terminal**

```
      . _..::__:  ,-"-"._       |7       ,     _,.__
_.___ _ _<_>`!(._`.`-.    /        _._     `_ ,_/  '  '-._.---.-.__
>.{     " " `-==,googI googI=googI=googI=googI=_googI==-:googI=-    -googI=`googI=googleI==
.googI=- ` `  googI=(googI=-.  `.googI=-`googI=-googleI=-googI=-  `  `=googI=-(googI=-  `
(googI= \ _googI-googI=-  _googI=-  `=googI=-  )=googI=-  `=googI=-  `=googI=-  `=googI=-
```

- **GitHub**: https://github.com/rastapasta/mapscii
- **Try it NOW**: `telnet mapscii.me`
- **Online**: https://labs.maptiler.com/mapscii/
- **Features**:
  - OpenStreetMap data rendered in braille/ASCII
  - Mouse support for pan and zoom
  - Point-of-interest discovery
  - Offline mode with local vector tiles
  - Customizable styling (Mapbox Styles)
- **Controls**:
  - Arrow keys: scroll
  - `a`/`z`: zoom in/out
  - `c`: toggle braille/ASCII mode
  - Mouse drag/scroll supported

---

## 5. Image & Video Viewers

### chafa
**The gold standard for terminal image conversion**

- **GitHub**: https://github.com/hpjansson/chafa
- **Output modes**:
  - Sixel graphics
  - Kitty graphics protocol
  - iTerm2 graphics protocol
  - Unicode braille (`--symbols braille`)
  - Block characters
  - ASCII
- **Braille usage**: `chafa --symbols braille -c full -s 80x40 image.jpg`
- **Features**: Animated GIFs, dithering (Bayer, Floyd-Steinberg), SVG support

### timg
**Terminal image and video viewer**

- **GitHub**: https://github.com/hzeller/timg
- **Pixel modes**: half blocks, quarter blocks, sixel, kitty, iTerm2
- **Video support**: Plays video files in terminal
- **Grid view**: Display multiple images at once

### ascii-image-converter
**Cross-platform image to ASCII/braille converter**

- **GitHub**: https://github.com/TheZoraiz/ascii-image-converter
- **Install**: `go install github.com/TheZoraiz/ascii-image-converter@latest`
- **Braille mode**: `-b` or `--braille` flag
- **Color support**: `-C` for original colors (24-bit or 8-bit fallback)
- **Dithering**: Makes braille art more visible for binary on/off dots

### TerminalImageViewer (tiv)
**C++ image viewer using RGB ANSI and block graphics**

- **GitHub**: https://github.com/stefanhaustein/TerminalImageViewer
- **Resolution**: Uses block characters for 2 pixels per character height

---

## 6. Audio Visualization

### CAVA
**Console Audio Visualizer for ALSA/PulseAudio/PipeWire**

```
  ▁▂▃▄▅▆▇█▇▆▅▄▃▂▁
```

- **GitHub**: https://github.com/karlstav/cava
- **Features**:
  - Bar spectrum visualizer
  - Supports ALSA, PulseAudio, PipeWire, JACK, sndio
  - Custom characters: uses U+2581-2587 (partial block characters)
  - SDL mode for desktop window
- **Install**: Available in most distro repos

### cli-visualizer
**MPD/ALSA/PulseAudio visualizer**

- **GitHub**: https://github.com/PosixAlchemist/cli-visualizer
- **Visualizers**: spectrum, lorenz, ellipse
- **Real-time**: Synced to music playback

### boscillate
**Real-time audio waveform in terminal**

- Built with `node-drawille`
- Shows live microphone/audio input as braille waveform

---

## 7. Games & Emulators

### Jester-Emu
**Retro console emulator rendered entirely in braille**

- **Website**: https://jester-tui.github.io/
- **Features**:
  - GameBoy through PS1 support
  - Cycle-accurate GameBoy core
  - 60 FPS smooth playback
  - Works over SSH, in TTY, or any terminal
  - 4×4 sub-pixel resolution per character cell
- **Tech**: C++ with heavy multithreading

### Gambaterm
**Game Boy Color in your terminal**

- **Website**: https://vxgmichel.github.io/gambatte-terminal/
- **Resolution**: 200×200 monochrome canvas (100 cols × 50 lines)
- **Note**: Original GB has 4 grayscale shades; braille is monochrome (loses half the info)
- **SSH support**: Play over network

### Terminal-based Game of Life
**Conway's Game of Life with braille rendering**

- Implementations in: `rsille` (Rust), `lua-drawille`, `Term-Graille` (Perl)
- RLE pattern file support in `rsille`
- Higher resolution than block character versions

### Rotating 3D Cube (various)
**Classic graphics demo**

- Example in `drawille-go/examples/rotating_cube`
- Wireframe rendering with braille
- Demonstrates 3D projection math

---

## 8. Developer Tools & DevOps

### lazygit
**Simple terminal UI for git commands**

- **GitHub**: https://github.com/jesseduffield/lazygit
- **Uses**: `gocui` / terminal UI libraries
- **Features**: Interactive staging, rebasing, commit graphs

### lazydocker
**Docker management in terminal**

- **GitHub**: https://github.com/jesseduffield/lazydocker
- **Metrics**: Container CPU/memory with ASCII/braille graphs
- **Cross-platform**: Linux, macOS, Windows

### k9s
**Kubernetes CLI management**

- **GitHub**: https://github.com/derailed/k9s
- **Features**: Resource graphs, pod metrics visualization

### Dockdash / expvarmon
**Docker and Go monitoring dashboards**

- Built with `termui`
- Real-time metrics with braille sparklines

---

## 9. Creative & Artistic

### Image to Braille Art (Web tools)

| Tool | URL |
|------|-----|
| Braille ASCII Art | https://lachlanarthur.github.io/Braille-ASCII-Art/ |
| Image to Braille | https://505e06b2.github.io/Image-to-Braille/ |
| A.Tools Generator | https://www.a.tools/Tool.php?Id=347 |
| LAZE Software | https://lazesoftware.com/en/tool/brailleaagen/ |

### L-Systems & Fractals
**Procedural graphics in terminal**

- `lua-drawille`: L-system support for fractal generation
- Turtle graphics for drawing trees, Koch curves, Sierpinski triangles

### brailleframe
**Persistence of vision color display**

- **GitHub**: https://github.com/ashfn/brailleframe
- Uses rapid braille character updates to simulate color
- Turns terminal into crude "display"

### SparkBraille
**Accessible charts for blind users**

- **Website**: https://fizzstudio.github.io/sparkbraille/
- Shows 2 data points per braille cell on refreshable braille displays
- Enables quick trend comprehension through touch

---

## 10. Libraries & Frameworks

### By Language

| Language | Library | GitHub | Notes |
|----------|---------|--------|-------|
| **Python** | drawille | [asciimoo/drawille](https://github.com/asciimoo/drawille) | Original, Canvas + Turtle |
| **Python** | plotille | [tammoippen/plotille](https://github.com/tammoippen/plotille) | Plotting focused |
| **Python** | termgraphics | [dheera/python-termgraphics](https://github.com/dheera/python-termgraphics) | NumPy, ASCII fallback |
| **Go** | drawille-go | [exrook/drawille-go](https://github.com/exrook/drawille-go) | Basic port |
| **Go** | drawille-go | [Kerrigan29a/drawille-go](https://github.com/Kerrigan29a/drawille-go) | +Bezier, inverse Y |
| **Go** | termui | [gizak/termui](https://github.com/gizak/termui) | Dashboard widgets |
| **Rust** | drawille-rs | [ftxqxd/drawille-rs](https://github.com/ftxqxd/drawille-rs) | Turtle, colors |
| **Rust** | rsille | [nidhoggfgg/rsille](https://github.com/nidhoggfgg/rsille) | Animation, 3D, image |
| **Node.js** | node-drawille | [madbence/node-drawille](https://github.com/madbence/node-drawille) | Powers vtop |
| **Node.js** | drawille-canvas | [madbence/node-drawille-canvas](https://github.com/madbence/node-drawille-canvas) | HTML5 Canvas API |
| **Node.js** | blessed-contrib | [yaronn/blessed-contrib](https://github.com/yaronn/blessed-contrib) | Dashboard widgets |
| **C** | libdrawille | [Huulivoide/libdrawille](https://github.com/Huulivoide/libdrawille) | Low-level, fast |
| **C** | chafa | [hpjansson/chafa](https://github.com/hpjansson/chafa) | Comprehensive |
| **Perl** | Term-Graille | [saiftynet/Term-Graille](https://github.com/saiftynet/Term-Graille) | Sprites, menus, audio |
| **Lua** | lua-drawille | [asciimoo/lua-drawille](https://github.com/asciimoo/lua-drawille) | 3D, L-systems |
| **Julia** | TextPlots.jl | [sunetos/TextPlots.jl](https://github.com/sunetos/TextPlots.jl) | Function plotting |

---

## Technical Notes

### Why Braille Characters?

1. **8× resolution**: Each 2×4 cell gives 8 pixels vs 1 for regular characters
2. **256 patterns**: Full binary coverage (U+2800-U+28FF)
3. **Consistent width**: Even in variable-width fonts
4. **Wide font support**: Most modern fonts include braille
5. **Special blank**: U+2800 (BRAILLE PATTERN BLANK) for consistent spacing

### Font Requirements

Fonts with good braille support:
- **Recommended**: Hack, JetBrains Mono, Fira Code, DejaVu Sans Mono
- **Powerline fonts**: Terminess Powerline, etc.
- **Nerd Fonts**: Any Nerd Font variant

### Terminal Compatibility

| Feature | xterm | iTerm2 | Kitty | Windows Terminal | Alacritty |
|---------|-------|--------|-------|------------------|-----------|
| Braille | ✓ | ✓ | ✓ | ✓ | ✓ |
| 256 color | ✓ | ✓ | ✓ | ✓ | ✓ |
| True color | ✓ | ✓ | ✓ | ✓ | ✓ |
| Sixel | ✓ | ✗ | ✗ | ✗ | ✗ |
| Kitty graphics | ✗ | ✗ | ✓ | ✗ | ✗ |

---

## Try It Now!

```bash
# World map in your terminal
telnet mapscii.me

# Cryptocurrency tracker (if installed)
ssh cointop.sh

# Quick braille image (install ascii-image-converter first)
ascii-image-converter -b your-image.png
```

---

## Contributing

Know of a project using braille terminal graphics that's not listed? The ecosystem is always growing!

---

*Last updated: 2026-02-01*
