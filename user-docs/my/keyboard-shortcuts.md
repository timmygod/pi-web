# Keyboard Shortcuts

pi-web ၏ local-model edition အတွက် shortcut များသည် အောက်ပါအတိုင်း ဖြစ်သည်။ Edition အလိုက် ကွဲပြားသော runtime ဝန်ဆောင်မှုများအကြောင်း [Local-model edition development](../../docs/dev/local-llm-development.md) တွင် ဖော်ပြထားသည်။

## Index page (`/`)

### Page scrolling (vim-style)

Focus သည် input, textarea, သို့မဟုတ် contenteditable element ထဲတွင် **မ**ရှိပါက j/k နှင့်တူညီသော vim-style shortcut များသည် pages များအားလုံးတွင် အလုပ်လုပ်သည်။

| Shortcut | Action |
|----------|--------|
| `j` | Scroll down 300px |
| `k` | Scroll up 300px |
| `g g` | Scroll to top of page |
| `G` (Shift+G) | Scroll to bottom of page |
| `Escape` | j/k navigation အလုပ်လုပ်ရန် active input ကို blur သွင်းစေသည် |

### Index commands

| Shortcut | Context | Action |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Page-level | Open search/sessions palette |
| `⌘⇧L` / `Ctrl+Shift+L` | Page-level | Toggle system theme (light/dark) |
| `Escape` | Page-level | Close palette, menu, or modal |
| `Enter` | New-session path input | Create new session |

> `⌘K` / `Ctrl+K` သည် Chrome ၏ "focus address bar" shortcut နှင့်လည်း တူသည်။ Focus သည် text input ထဲတွင် မပါဝင်ပါက browser က အကွက်ပြုလုပ်နိုင်သည်။

## Session detail page (`/session?id=...`)

### Page scrolling (vim-style)

Focus သည် input, textarea, သို့မဟုတ် contenteditable element ထဲတွင် **မ**ရှိပါက index နှင့် session pages နှစ်ခုလုံးတွင် အလုပ်လုပ်သည်။

| Shortcut | Action |
|----------|--------|
| `j` | Scroll down 300px |
| `k` | Scroll up 300px |
| `g g` | Scroll to top of page |
| `G` (Shift+G) | Scroll to bottom of page |
| `I` (Shift+I) | Focus the chat composer textarea |
| `Escape` | j/k navigation အလုပ်လုပ်ရန် active input ကို blur သွင်းစေသည် |

### Sidebar & navigation

| Shortcut | Context | Action |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Page-level | Toggle sidebar visibility |
| `⌘K` / `Ctrl+K` | Page-level | Open session list palette |
| `⌘T` / `Ctrl+T` | Page-level | New session |
| `⌘⇧L` / `Ctrl+Shift+L` | Page-level | Toggle system theme (light/dark) |
| `⌘⇧N` / `Ctrl+Shift+N` | Page-level | Toggle scratchpad / notes sidebar |

> `⌘K` နှင့် `⌘T` သည် browser shortcuts နှင့်လည်း တူသည် (focus address bar / new tab)။ Focus သည် text input ထဲတွင် မပါဝင်ပါက browser က အကွက်ပြုလုပ်နိုင်သည်။

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
| `t` | Focus သည် input/textarea ထဲတွင် **မ**ရှိချိန် | Toggle thinking visibility |
| `o` | Focus သည် input/textarea ထဲတွင် **မ**ရှိချိန် | Toggle tools visibility |
| `p` | Focus သည် input/textarea ထဲတွင် **မ**ရှိချိန် | Toggle tool outputs |

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
