# Technical Requirements Document (TRD) - Project Pantau Gizi V2.1

## 1. Ringkasan Eksekutif
Platform **Pantau Gizi** adalah sistem crowdsourcing independen untuk memantau kualitas program Makan Bergizi Gratis (MBG). Fokus pengembangan: efisiensi biaya, skalabilitas nasional, dan integritas data publik melalui transparansi kode sumber (Open Source).

## 2. Arsitektur Infrastruktur (Hybrid Model)
Sistem menggunakan pendekatan hybrid untuk optimasi performa vs biaya:
- **Backend API:** Go (Golang) 1.22+ (Dockerized on VPS).
- **Frontend:** Next.js (PWA enabled) on Vercel/Cloudflare Pages.
- **Database Utama:** PostgreSQL (Managed/Serverless - Neon.tech/Supabase).
- **Cache & Rate Limiting:** Redis (Upstash/Self-hosted).
- **Object Storage:** Cloudflare R2 (S3-compatible) untuk foto laporan.

## 3. Spesifikasi Teknis Backend (Go)
- **Framework:** Gin atau Echo.
- **Database Handler:** GORM atau sqlx dengan connection pooling.
- **Logging:** Structured logging (zerolog/zap).
- **Middleware:** CORS, Recovery, dan Redis-based Rate Limiting.

## 4. Strategi Pengelolaan Data & Media
### A. Mekanisme Presigned URL
Server Go **TIDAK DIPERBOLEHKAN** memproses upload file biner secara langsung demi efisiensi resource server.
1. Client `POST` ke `/v1/upload/request`.
2. Backend generate **S3 Presigned URL** (Cloudflare R2, expired 5-10 mnt).
3. Client upload biner langsung ke R2 via method `PUT`.
4. Client mengirim metadata (URL, rating, koordinat) ke API utama setelah upload sukses.

### B. Kompresi Sisi Client (Frontend)
Wajib dilakukan sebelum upload untuk efisiensi bandwidth:
- **Max File Size:** 500 KB.
- **Max Resolution:** 1280px (long side).
- **Format:** WebP/JPEG (Quality 0.7 - 0.8).

### C. Offline Persistence & PWA Support
- **Storage:** Menggunakan `IndexedDB` atau `LocalStorage` untuk menyimpan draft laporan jika koneksi internet terputus.
- **Sync:** Implementasi *Background Sync* agar data terkirim otomatis saat perangkat kembali online.
- **Service Worker:** Caching aset statis agar aplikasi tetap dapat diakses tanpa internet.

## 5. Open Data & Export Engine
- **Mechanism:** Generator CSV/JSON berbasis *stream* untuk efisiensi penggunaan RAM pada volume data besar.
- **Anonymization:** Logic wajib menghapus kolom sensitif (`user_id`, `email`) sebelum data dipublikasikan ke publik.
- **Caching:** File hasil ekspor disimpan di Redis/R2 selama 24 jam untuk menghindari regenerasi berulang yang membebani database.

## 6. Skema Database (Ringkasan)
| Tabel | Kolom Utama | Indeks Kritis |
| :--- | :--- | :--- |
| **users** | id, email, name, role | email |
| **schools** | id, npsn, name, city_code, lat, long | npsn, city_code |
| **reports** | id, user_id, school_id, photo_url, rating, is_verified | school_id, created_at |

## 7. Keamanan
- **Captcha:** Cloudflare Turnstile pada setiap submit laporan.
- **Validation:** Backend wajib memvalidasi domain `photo_url` (harus berasal dari bucket R2 resmi).
- **Export Protection:** Rate limiting khusus pada endpoint Open Data untuk mencegah *resource exhaustion*.

## 8. Panduan Deployment
Wajib menyertakan `Dockerfile` dan `docker-compose.yml` yang mencakup:
- Go API container.
- Redis container.
- Nginx Reverse Proxy (SSL/HTTPS).