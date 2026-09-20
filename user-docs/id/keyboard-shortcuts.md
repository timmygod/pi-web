# Pintasan Keyboard

Pintasan-pintasan ini berlaku untuk edisi local-model dari pi-web. Perilaku runtime yang spesifik per edisi didokumentasikan di [Local-model edition development](../../docs/dev/local-llm-development.md).

## Halaman indeks (`/`)

### Gulir halaman (gaya vim)

Pintasan gaya vim yang sama berfungsi di semua halaman saat fokus **tidak** berada di elemen input, textarea, atau contenteditable.

| Pintasan | Aksi |
|----------|--------|
| `j` | Gulir ke bawah 300px |
| `k` | Gulir ke atas 300px |
| `g g` | Gulir ke atas halaman |
| `G` (Shift+G) | Gulir ke bawah halaman |
| `Escape` | Lepaskan fokus dari input aktif agar navigasi j/k berfungsi |

### Perintah indeks

| Pintasan | Konteks | Aksi |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Level halaman | Buka palet pencarian/sesi |
| `⌘⇧L` / `Ctrl+Shift+L` | Level halaman | Ganti tema sistem (terang/gelap) |
| `Escape` | Level halaman | Tutup palet, menu, atau modal |
| `Enter` | Input jalur sesi baru | Buat sesi baru |

> `⌘K` / `Ctrl+K` juga merupakan pintasan "fokus bilah alamat" di Chrome. Browser mungkin menginterupsinya kecuali fokus berada di dalam input teks.

## Halaman detail sesi (`/session?id=...`)

### Gulir halaman (gaya vim)

Pintasan-pintasan ini berfungsi di halaman indeks maupun halaman sesi saat fokus **tidak** berada di elemen input, textarea, atau contenteditable.

| Pintasan | Aksi |
|----------|--------|
| `j` | Gulir ke bawah 300px |
| `k` | Gulir ke atas 300px |
| `g g` | Gulir ke atas halaman |
| `G` (Shift+G) | Gulir ke bawah halaman |
| `I` (Shift+I) | Fokus ke textarea chat composer |
| `Escape` | Lepaskan fokus dari input aktif agar navigasi j/k berfungsi |

### Sidebar & navigasi

| Pintasan | Konteks | Aksi |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Level halaman | Ganti visibilitas sidebar |
| `⌘K` / `Ctrl+K` | Level halaman | Buka palet daftar sesi |
| `⌘T` / `Ctrl+T` | Level halaman | Sesi baru |
| `⌘⇧L` / `Ctrl+Shift+L` | Level halaman | Ganti tema sistem (terang/gelap) |
| `⌘⇧N` / `Ctrl+Shift+N` | Level halaman | Ganti sidebar scratchpad / catatan |

> `⌘K` dan `⌘T` juga merupakan pintasan browser (fokus bilah alamat / tab baru). Browser mungkin menginterupsinya kecuali fokus berada di dalam input teks.

### Chat composer

| Pintasan | Konteks | Aksi |
|----------|---------|--------|
| `Enter` | Textarea chat | Kirim pesan |
| `Shift+Enter` | Textarea chat | Sisipkan baris baru |
| `Shift+Tab` | Textarea chat | Beralih ke level berpikir berikutnya (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Textarea chat | Buka popup pemilih model (ketik untuk memfilter, Enter untuk memilih, fokus kembali ke textarea) |

### Pengalih visibilitas entri

| Pintasan | Konteks | Aksi |
|----------|---------|--------|
| `t` | Saat fokus **tidak** berada di input/textarea | Ganti visibilitas thinking |
| `o` | Saat fokus **tidak** berada di input/textarea | Ganti visibilitas tools |
| `p` | Saat fokus **tidak** berada di input/textarea | Ganti output tools |

### Palet, menu & sheet

| Pintasan | Konteks | Aksi |
|----------|---------|--------|
| `Escape` | Level halaman | Tutup palet, menu, atau sheet yang terbuka |
| `⌘K` / `Ctrl+K` | Level halaman | Buka palet daftar sesi |
| `ArrowUp` / `ArrowDown` | Palet daftar sesi | Navigasi hasil sesi |
| `Enter` | Palet daftar sesi | Buka sesi yang dipilih (atau pertama) |
| `ArrowUp` / `ArrowDown` | Popup pemilih model | Navigasi daftar model |
| `Enter` | Popup pemilih model | Pilih model yang disorot |
| `ArrowUp` / `ArrowDown` | Modal fork | Navigasi pesan |
| `Enter` | Modal fork | Fork dari pesan yang disorot |
| `Tab` | Sheet layar penuh | Beralih fokus di dalam sheet |
| `Escape` | Sheet layar penuh | Tutup sheet |
