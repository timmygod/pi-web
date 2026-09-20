# Larian Pintasan Papan Kekunci

Pintasan ini terpakai pada edisi model-tempatan bagi pi-web. Kelakuan masa-eksekusi khusus edisi
didokumentasikan dalam [Pembangunan edisi model-tempatan](../../docs/dev/local-llm-development.md).

## Halaman indeks (`/`)

### Skrol halaman (gaya vim)

Pintasan gaya vim yang sama berfungsi pada semua halaman apabila tumpuan **tidak** berada dalam elemen input, textarea, atau contenteditable.

| Pintasan | Tindakan |
|----------|--------|
| `j` | Skrol turun 300px |
| `k` | Skrol naik 300px |
| `g g` | Skrol ke bahagian atas halaman |
| `G` (Shift+G) | Skrol ke bahagian bawah halaman |
| `Escape` | Keluarkan tumpuan daripada input aktif supaya navigasi j/k berfungsi |

### Perintah indeks

| Pintasan | Konteks | Tindakan |
|----------|---------|--------|
| `⌘K` / `Ctrl+K` | Tahap halaman | Buka palet carian/sesi |
| `⌘⇧L` / `Ctrl+Shift+L` | Tahap halaman | Togol tema sistem (cerah/gelap) |
| `Escape` | Tahap halaman | Tutup palet, menu, atau modal |
| `Enter` | Input laluan sesi baharu | Cipta sesi baharu |

> `⌘K` / `Ctrl+K` juga merupakan pintasan "tumpun ke bar alamat" Chrome. Pelayar mungkin memintasnya kecuali tumpuan berada di dalam input teks.

## Halaman butiran sesi (`/session?id=...`)

### Skrol halaman (gaya vim)

Pintasan ini berfungsi pada kedua-dua halaman indeks dan sesi apabila tumpuan **tidak** berada dalam elemen input, textarea, atau contenteditable.

| Pintasan | Tindakan |
|----------|--------|
| `j` | Skrol turun 300px |
| `k` | Skrol naik 300px |
| `g g` | Skrol ke bahagian atas halaman |
| `G` (Shift+G) | Skrol ke bahagian bawah halaman |
| `I` (Shift+I) | Tumpun ke textarea pengarang sembang |
| `Escape` | Keluarkan tumpuan daripada input aktif supaya navigasi j/k berfungsi |

### Bar sisi & navigasi

| Pintasan | Konteks | Tindakan |
|----------|---------|--------|
| `⌘B` / `Ctrl+B` | Tahap halaman | Togol kelihatan bar sisi |
| `⌘K` / `Ctrl+K` | Tahap halaman | Buka palet senarai sesi |
| `⌘T` / `Ctrl+T` | Tahap halaman | Sesi baharu |
| `⌘⇧L` / `Ctrl+Shift+L` | Tahap halaman | Togol tema sistem (cerah/gelap) |
| `⌘⇧N` / `Ctrl+Shift+N` | Tahap halaman | Togol bar sisi skrap / nota |

> `⌘K` dan `⌘T` juga merupakan pintasan pelayar (tumpun ke bar alamat / tab baharu). Pelayar mungkin memintaskannya kecuali tumpuan berada di dalam input teks.

### Pengarang sembang

| Pintasan | Konteks | Tindakan |
|----------|---------|--------|
| `Enter` | Textarea sembang | Hantar mesej |
| `Shift+Enter` | Textarea sembang | Sisip baris baharu |
| `Shift+Tab` | Textarea sembang | Berpindah ke tahap pemikiran berikutnya (`off` → `minimal` → … → `xhigh` → `off`) |
| `Ctrl+I` / `Ctrl+L` | Textarea sembang | Buka popap pemilih model (ketik untuk tapisan, Enter untuk pilih, tumpuan kembali ke textarea) |

### Togol kelihatan entri

| Pintasan | Konteks | Tindakan |
|----------|---------|--------|
| `t` | Apabila tumpuan **tidak** berada dalam input/textarea | Togol kelihatan pemikiran |
| `o` | Apabila tumpuan **tidak** berada dalam input/textarea | Togol kelihatan alat |
| `p` | Apabila tumpuan **tidak** berada dalam input/textarea | Togol keluaran alat |

### Palet, menu & lembaran

| Pintasan | Konteks | Tindakan |
|----------|---------|--------|
| `Escape` | Tahap halaman | Tutup sebarang palet, menu, atau lembaran yang terbuka |
| `⌘K` / `Ctrl+K` | Tahap halaman | Buka palet senarai sesi |
| `ArrowUp` / `ArrowDown` | Palet senarai sesi | Navigasi hasil sesi |
| `Enter` | Palet senarai sesi | Buka sesi yang dipilih (atau pertama) |
| `ArrowUp` / `ArrowDown` | Popap pemilih model | Navigasi senarai model |
| `Enter` | Popap pemilih model | Pilih model yang disorot |
| `ArrowUp` / `ArrowDown` | Modal garahan | Navigasi mesej |
| `Enter` | Modal garahan | Garahan daripada mesej yang disorot |
| `Tab` | Lembaran skrin penuh | Berpindah tumpuan dalam lembaran |
| `Escape` | Lembaran skrin penuh | Tutup lembaran |
