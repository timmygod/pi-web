# ក្លាវៀស្ទីដ

ក្លាវៀស្ទីដទាំងនេះ អនុវត្តទៅតាម edition របស់ local-model របស់ pi-web។ Runtime behavior ដែលពាក់ព័ន្ធនឹង edition ត្រូវបានសរសេរឆ្លងកាត់ [Local-model edition development](../../docs/dev/local-llm-development.md)។

## Index page (`/`)

### Page scrolling (vim-style)

ក្លាវៀស្ទីដ vim-style ដូចគ្នា ដំណើរការនៅទាំងអស់នៃទំព័រ នៅពេល focus មិនក្នុង input, textarea, ឬ contenteditable element ទេ។

| Shortcut | Action |
|----------|--------|
| `j` | Scroll down 300px |
| `k` | Scroll up 300px |
| `g g` | Scroll to top of page |
| `G` (Shift+G) | Scroll to bottom of page |
| `Escape` | Blur the active input so j/k navigation works |

### Index commands

| Shortcut | Context | Action |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Page-level | Open search/sessions palette |
| `⌘⇧L` / `Ctrl+Shift+L` | Page-level | Toggle system theme (light/dark) |
| `Escape` | Page-level | Close palette, menu, or modal |
| `Enter` | New-session path input | Create new session |

> `⌘K` / `Ctrl+K` ក៏ជា shortcut របស់ Chrome ដែរ សម្រាប់ "focus address bar"។ Browser អាច intercept វា លើកលែងតែ focus ក្នុង text input។

## Session detail page (`/session?id=...`)

### Page scrolling (vim-style)

ក្លាវៀស្ទីដទាំងនេះ ដំណើរការទាំង index និង session pages នៅពេល focus មិនក្នុង input, textarea, ឬ contenteditable element ទេ។

| Shortcut | Action |
|----------|--------|
| `j` | Scroll down 300px |
| `k` | Scroll up 300px |
| `g g` | Scroll to top of page |
| `G` (Shift+G) | Scroll to bottom of page |
| `I` (Shift+I) | Focus the chat composer textarea |
| `Escape` | Blur the active input so j/k navigation works |

### Sidebar & navigation

| Shortcut | Context | Action |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Page-level | Toggle sidebar visibility |
| `⌘K` / `Ctrl+K` | Page-level | Open session list palette |
| `⌘T` / `Ctrl+T` | Page-level | New session |
| `⌘⇧L` / `Ctrl+Shift+L` | Page-level | Toggle system theme (light/dark) |
| `⌘⇧N` / `Ctrl+Shift+N` | Page-level | Toggle scratchpad / notes sidebar |

> `⌘K` និង `⌘T` ក៏ជា shortcut របស់ browser ដែរ (focus address bar / new tab)។ Browser អាច intercept វា លើកលែងតែ focus ក្នុង text input។

### Chat composer

| Shortcut | Context | Action |
|----------|---------|--------|
| `Enter` | Chat textarea | Submit message |
| `Shift+Enter` | Chat textarea | Insert newline |
| `Shift+Tab` | Chat textarea | Cycle to next thinking level (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Chat textarea | Open model selector popup (type to filter, Enter to select, focus returns to textarea) |

### Entry visibility toggles

| Shortcut | Context | Action |
|----------|---------|--------|
| `t` | When focus is **not** in an input/textarea | Toggle thinking visibility |
| `o` | When focus is **not** in an input/textarea | Toggle tools visibility |
| `p` | When focus is **not** in an input/textarea | Toggle tool outputs |

### Palettes, menus & sheets

| Shortcut | Context | Action |
|----------|---------|--------|
| `Escape` | Page-level | Close any open palette, menu, or sheet |
| `⌘K` / `Ctrl+K` | Page-level | Open session list palette |
| `ArrowUp` / `ArrowDown` | Session list palette | Navigate session results |
| `Enter` | Session list palette | Open the selected (or first) session |
| `ArrowUp` / `ArrowDown` | Model selector popup | Navigate model list |
| `Enter` | Model selector popup | Select highlighted model |
| `ArrowUp` / `ArrowDown` | Fork modal | Navigate messages |
| `Enter` | Fork modal | Fork from highlighted message |
| `Tab` | Full-screen sheet | Cycle focus within the sheet |
| `Escape` | Full-screen sheet | Close the sheet |
