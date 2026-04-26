# 🍱 Semantic Pantau Gizi - Open Source Crowdsourcing Platform

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Next.js](https://img.shields.io/badge/Next.js-14-black?style=flat&logo=next.js)](https://nextjs.org/)

**Pantau Gizi** adalah platform independen berbasis masyarakat untuk memantau, melaporkan, dan mengevaluasi kualitas program Makan Bergizi Gratis (MBG) di Indonesia. Kami menggunakan teknologi untuk menciptakan transparansi dan memastikan hak gizi anak-anak terpenuhi secara adil dan berkualitas.

---

## 🌟 Kenapa Berkontribusi di Sini? (Benefits)

Bergabung dengan proyek open-source bukan hanya soal menulis kode, tapi juga investasi untuk karier Anda:

1. **Elite Tech-Stack Portfolio:** Tunjukkan kemampuan Anda dalam menangani arsitektur modern (*Go, Redis, Cloudflare R2, Next.js*). Ini adalah nilai jual tinggi di hadapan rekruter perusahaan teknologi besar.
2. **Real-World Experience:** Hadapi tantangan nyata seperti *high-concurrency*, *presigned-URL implementation*, dan optimasi infrastruktur biaya rendah.
3. **Social Impact:** Kontribusi Anda berdampak langsung pada pengawasan gizi jutaan anak sekolah di Indonesia. Kode Anda adalah solusi sosial.
4. **Networking:** Terhubung dengan sesama engineer yang memiliki visi yang sama.
5. **Future Opportunities:** Proyek ini dikelola secara profesional. Kontributor utama akan menjadi prioritas jika proyek ini mendapatkan pendanaan, hibah, atau bertransformasi menjadi entitas yang lebih besar.

---

## 🚀 Key Engineering Features

Proyek ini dibangun dengan prinsip **High Performance** dan **Low Cost Operations**.

- **Hybrid Infrastructure:** Memadukan VPS untuk *Compute* dan Serverless untuk *Data/Storage*.
- **Zero-Load File Upload:** Implementasi **S3 Presigned URLs** via Cloudflare R2. Server tidak pernah memproses data biner gambar, menjaga kestabilan RAM VPS.
- **Atomic Caching:** Menggunakan Redis untuk dashboard publik agar tetap responsif meskipun diakses ribuan pengguna secara bersamaan.
- **Efficient Deployment:** Full Dockerized environment untuk kemudahan *scaling* dan *maintenance*.

## 🛠 Tech Stack

- **Backend:** [Go (Golang)](https://go.dev/)
- **Frontend:** [Next.js](https://nextjs.org/) (PWA)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Cache:** [Redis](https://redis.io/)
- **Storage:** [Cloudflare R2](https://www.cloudflare.com/lp/pg-storage-r2/)
- **Infrastructure:** Docker & Docker Compose

## 🏗 System Architecture & Spec

Detail spesifik mengenai alur data dan kebutuhan teknis dapat dilihat pada [Technical Requirements Document (TRD)](/docs/TRD.md).

## 📈 Roadmap

- [ ] **Phase 1:** MVP (Laporan, Foto, Rating, Dashboard Dasar).
- [ ] **Phase 2:** Peta Geospasial Nasional & Analisis Wilayah.
- [ ] **Phase 3:** AI-Powered Image Verification (Deteksi komponen makanan otomatis).

## 🤝 Cara Berkontribusi

1. Fork repository ini.
2. Pilih task pada menu [Issues](https://github.com/mhdarifsetiawan/semantic-pantau-gizi/issues).
3. Buat branch baru dan kirim Pull Request (PR).

## 💰 Dukung Kami (Donasi & Kemitraan)

Kami berkomitmen menjaga proyek ini tetap independen. Dukungan finansial Anda akan digunakan sepenuhnya untuk biaya infrastruktur server dan pengembangan fitur.

- **Donasi via Saweria:** [saweria.co/mhdarifsetiawan](https://saweria.co/mhdarifsetiawan)
- **Email Kontak:** [mhdarifsetiawan01@gmail.com](mailto:mhdarifsetiawan01@gmail.com)

---

## 📄 Lisensi

Proyek ini berada di bawah lisensi **MIT**.

---
*Dibuat dengan ❤️ untuk masa depan Indonesia yang lebih sehat dan transparan.*
