# Soal Test

## Soal Nomor 1 
Ada di struktur file golang project

## Soal Nomor 2 
Saya langsung mengimplementasikan di model dengan path direktorinya di /models

## Soal Nomor 3 
Ada di file no_3_ bank_account_system.go

## Soal Nomor 4
### Optimasi SQL Queri:
SELECT customer_id, SUM(amount) AS total_spent
FROM orders
WHERE order_date >= CURRENT_DATE - INTERVAL '1 month'
GROUP BY customer_id
ORDER BY total_spent DESC
LIMIT 5;

### Performa improvement:
#### 1. menambakan index di colomb order_date:
CREATE INDEX idx_orders_order_date ON orders (order_date);
#### 2. gunakan indeks penutup komposit:
CREATE INDEX idx_orders_date_customer_amount ON orders (order_date, customer_id, amount);
##### penjelasan:
Indeks ini memungkinkan basis data untuk:
- memfilter dengan cepat berdasarkan order_date.
- mengelompokkan berdasarkan customer_id (memanfaatkan urutan indeks).
- mengambil jumlah langsung dari indeks (menghindari akses tabel).

## Soal Nomor 5
### menguraikan monolit menjadi Layanan
- menganalisis dan memetakan sistem yang ada (mengidentifikasi tanggung jawab, menilai ketergantungan, dan memprioritaskan ekstraksi)
- menetapkan batasan layanan (desain berbasis domain (DDD) dan pemisahan basis data)
- ekstraksi inkremental (mengganti modul)
- migrasi data (penulisan ganda/CDC)
- komunikasi (API dan pesan asinkron)
### memastikan kompatibilitas
- memelihara API yang Ada (API gateway dan versioning)
- beralih fitur 
- konsistensi data (Proxy Layers dan Shadow Testing)
- koordinasi klien(Deprecation Timeline dan SDKs/Client Libraries)
- testing
### langkah-langkah pasca migrasi
- menghentikan layanan monolit
- kemampuan observasi
- mengoptimalkan