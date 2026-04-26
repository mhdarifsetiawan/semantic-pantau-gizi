# 🍱 Semantic Pantau Gizi - Open Source Crowdsourcing Platform

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Next.js](https://img.shields.io/badge/Next.js-14-black?style=flat&logo=next.js)](https://nextjs.org/)

**Semantic Pantau Gizi** adalah platform independen berbasis masyarakat untuk memantau, melaporkan, dan mengevaluasi kualitas program Makan Bergizi Gratis (MBG) di Indonesia. Kami menggabungkan kekuatan **komunitas open-source** dan **kecerdasan buatan (AI)** untuk menciptakan transparansi radikal demi gizi anak bangsa yang lebih baik.

---

## 🌟 Mengapa Bergabung dengan Proyek Ini?

Kami percaya bahwa transparansi sejati hanya bisa dicapai jika kodenya dapat diaudit oleh publik. Bergabung di sini berarti:

1. **Elite Tech-Stack Portfolio:** Implementasi nyata arsitektur modern (*Go, Redis, Cloudflare R2, Next.js*). Ini adalah aset berharga untuk portofolio di level perusahaan Big Tech.
2. **AI-Driven Impact:** Terlibat dalam pengembangan Phase 3—otomasi validasi gizi menggunakan *Computer Vision*.
3. **Radical Transparency:** Berbeda dengan platform tertutup, di sini setiap baris kode berkontribusi pada akuntabilitas publik yang bisa dipertanggungjawabkan.
4. **Professional Networking:** Berkolaborasi dengan sesama engineer yang visioner dan peduli pada isu sosial-nasional.
5. **Future Priorities:** Kontributor inti (Core Contributors) akan menjadi prioritas utama jika proyek ini berkembang menjadi entitas formal atau menerima dukungan hibah/investasi.

---

## 🚀 Key Engineering Features

Didesain untuk **High Performance** dengan **Low Cost Operations**:

- **Hybrid Infrastructure:** Optimasi biaya menggunakan VPS untuk *Compute* dan Serverless untuk *Data/Storage*.
- **Zero-Load File Upload:** Menggunakan **S3 Presigned URLs** (Cloudflare R2). Backend tidak memproses file biner, menghemat RAM dan CPU secara signifikan.
- **Atomic Caching:** Redis digunakan untuk menyajikan data agregat secara real-time tanpa membebani PostgreSQL.
- **Client-Side Edge Optimization:** Kompresi gambar otomatis di sisi client sebelum upload untuk efisiensi bandwidth.

## 🛠 Tech Stack

- **Backend:** [Go (Golang)](https://go.dev/)
- **Frontend:** [Next.js](https://nextjs.org/) (PWA Ready)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Cache:** [Redis](https://redis.io/)
- **Storage:** [Cloudflare R2](https://www.cloudflare.com/developer-platform/products/r2/)
- **Infrastructure:** Docker & Docker Compose

## 🏗 System Architecture & Spec

Cetak biru sistem, alur data, dan standar koding dapat diakses di [Technical Requirements Document (TRD)](/docs/TRD.md).

## 📂 Project Structure

```text
semantic-pantau-gizi/
├── cmd/                # Entry point aplikasi (Main apps)
│   └── api/
│       └── main.go     # HTTP Server entry point
├── docs/               # Dokumentasi teknis & TRD
├── internal/           # Private library & Business Logic (tidak bisa diimport proyek lain)
├── pkg/                # Public library (bisa diimport proyek lain)
├── web/                # Frontend side
│   └── next-app/       # Next.js (App Router)
├── go.mod              # Go module definition
└── README.md           # Pintu utama informasi proyek
```

## 📈 Roadmap

- [x] **Phase 0:** Inisialisasi Boilerplate & Dokumen Arsitektur.
- [ ] **Phase 1 (MVP):** Sistem Laporan, Rating Sekolah, & Dashboard Dasar.
- [ ] **Phase 2 (Analytics):** Peta Geospasial Nasional & Analitik Wilayah Terpadu.
- [ ] **Phase 3 (AI Integration):** Implementasi **Computer Vision** untuk deteksi otomatis komponen makanan (nasi, protein, sayur) guna validasi standar gizi secara otomatis dan masif.

## 🤝 Cara Berkontribusi

Kami menyukai kolaborasi! Silakan mulai dengan:
1. Fork repository ini.
2. Cari tugas yang tersedia di [Issues](https://github.com/mhdarifsetiawan/semantic-pantau-gizi/issues).
3. Buat branch baru, koding, dan kirim Pull Request (PR).

## ⚖️ Legal Disclaimer

Proyek ini bersifat independen dan tidak berafiliasi dengan lembaga pemerintah mana pun. Tujuan platform ini adalah murni untuk transparansi publik, partisipasi masyarakat, dan edukasi gizi.

1. Tanggung Jawab Konten: Isi laporan, foto, dan data yang diunggah adalah tanggung jawab sepenuhnya dari pelapor (User-Generated Content).
2. Akurasi Data: Pengelola platform tidak menjamin akurasi 100% dari data crowdsourcing sebelum melalui proses verifikasi teknis/AI.
3. Penyalahgunaan: Penggunaan platform untuk penyebaran hoax, ujaran kebencian, atau pencemaran nama baik dilarang keras dan akan dimoderasi.
4. Lisensi Kode: Seluruh kode sumber disediakan "as-is" di bawah lisensi MIT tanpa jaminan dalam bentuk apa pun.

## 💰 Dukungan & Kemitraan

Kami menjaga independensi dengan model pendanaan berbasis komunitas. Dukungan Anda membantu biaya operasional server dan infrastruktur.

- **Donasi via Saweria:** [saweria.co/mhdarifsetiawan](https://saweria.co/mhdarifsetiawan)
- **Email Kontak:** [mhdarifsetiawan01@gmail.com](mailto:mhdarifsetiawan01@gmail.com)

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah **MIT License**.

---
*Membangun teknologi untuk Indonesia yang lebih sehat, transparan, dan cerdas.*

