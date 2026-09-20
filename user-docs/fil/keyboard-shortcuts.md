# Mga Pindutan sa Keyboard

Ang mga pindutan na ito ay angkop sa local-model edition ng pi-web. Ang isahan ng runtime na nakabase sa edisyong ito ay nakadokumento sa [Local-model edition development](../../docs/dev/local-llm-development.md).

## Index page (`/`)

### Pag-scroll sa pahina (vim-style)

Ang mga vim-style na pindutang parehong ito ay gumagana sa lahat ng pahina kapag **hindi** ang focus nasa input, textarea, o contenteditable na elemento.

| Pindutan | Gawa |
|----------|--------|
| `j` | Mag-scroll pababa ng 300px |
| `k` | Mag-scroll pataas ng 300px |
| `g g` | Pumunta sa itaas ng pahina |
| `G` (Shift+G) | Pumunta sa ibaba ng pahina |
| `Escape` | I-blur ang aktibong input upang gumana ang j/k na pag-navigate |

### Mga command sa index

| Pindutan | Konteksto | Gawa |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Antas ng pahina | Buksan ang search/sessions palette |
| `⌘⇧L` / `Ctrl+Shift+L` | Antas ng pahina | I-toggle ang system theme (light/dark) |
| `Escape` | Antas ng pahina | Isara ang palette, menu, o modal |
| `Enter` | Input ng bago-sesyon | Lumikha ng bagong sesyon |

> Ang `⌘K` / `Ctrl+K` ay pangalagang shortcut ng Chrome na "focus address bar". Maaaring i-intercept ng browser ang pindutan na ito maliban kapag nasa loob ng text input ang focus.

## Sesyon-detail na pahina (`/session?id=...`)

### Pag-scroll sa pahina (vim-style)

Gumagana ang mga ito sa parehong index at sesyon na pahina kapag **hindi** ang focus nasa input, textarea, o contenteditable na elemento.

| Pindutan | Gawa |
|----------|--------|
| `j` | Mag-scroll pababa ng 300px |
| `k` | Mag-scroll pataas ng 300px |
| `g g` | Pumunta sa itaas ng pahina |
| `G` (Shift+G) | Pumunta sa ibaba ng pahina |
| `I` (Shift+I) | Itaguyod ang focus sa chat composer na textarea |
| `Escape` | I-blur ang aktibong input upang gumana ang j/k na pag-navigate |

### Sidebar at pag-navigate

| Pindutan | Konteksto | Gawa |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Antas ng pahina | I-toggle ang pagkakita ng sidebar |
| `⌘K` / `Ctrl+K` | Antas ng pahina | Buksan ang session list palette |
| `⌘T` / `Ctrl+T` | Antas ng pahina | Bagong sesyon |
| `⌘⇧L` / `Ctrl+Shift+L` | Antas ng pahina | I-toggle ang system theme (light/dark) |
| `⌘⇧N` / `Ctrl+Shift+N` | Antas ng pahina | I-toggle ang scratchpad / notes na sidebar |

> Ang `⌘K` at `⌘T` ay pangalagang shortcut ng browser (focus address bar / bagong tab). Maaaring i-intercept ng browser ang mga ito maliban kapag nasa loob ng text input ang focus.

### Chat composer

| Pindutan | Konteksto | Gawa |
|----------|---------|--------|
| `Enter` | Chat na textarea | I-submit ang mensahe |
| `Shift+Enter` | Chat na textarea | Ilagay ang newline |
| `Shift+Tab` | Chat na textarea | Lumipat sa susunod na antas ng pag-iisip (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Chat na textarea | Buksan ang model selector na popup (tangiin ang pagsulat, piliin sa pamamagitan ng Enter, lilitaw ang focus sa textarea) |

### Mga toggle sa pagkakita ng entry

| Pindutan | Konteksto | Gawa |
|----------|---------|--------|
| `t` | Kapag **hindi** nasa input/textarea ang focus | I-toggle ang pagkakita ng pag-iisip |
| `o` | Kapag **hindi** nasa input/textarea ang focus | I-toggle ang pagkakita ng mga tool |
| `p` | Kapag **hindi** nasa input/textarea ang focus | I-toggle ang mga output ng tool |

### Palettes, mga menu at sheet

| Pindutan | Konteksto | Gawa |
|----------|---------|--------|
| `Escape` | Antas ng pahina | Isara ang anumang bukas na palette, menu, o sheet |
| `⌘K` / `Ctrl+K` | Antas ng pahina | Buksan ang session list palette |
| `ArrowUp` / `ArrowDown` | Session list palette | Lumipat sa mga resulta ng sesyon |
| `Enter` | Session list palette | Buksan ang napiling (o unang) sesyon |
| `ArrowUp` / `ArrowDown` | Model selector na popup | Lumipat sa listahan ng model |
| `Enter` | Model selector na popup | Pumili ng naka-highlight na model |
| `ArrowUp` / `ArrowDown` | Fork na modal | Lumipat sa mga mensahe |
| `Enter` | Fork na modal | Gumawa ng fork mula sa naka-highlight na mensahe |
| `Tab` | Full-screen na sheet | Lumipat ang focus sa loob ng sheet |
| `Escape` | Full-screen na sheet | Isara ang sheet |
