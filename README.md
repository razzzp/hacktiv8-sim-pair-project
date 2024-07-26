# Simulasi Pair Project - Inventory Stock Management

## Description

Anda dan tim Anda ditugaskan untuk mengembangkan aplikasi Command Line Interface (CLI) untuk manajemen toko menggunakan Golang dan database MySQL. Aplikasi ini akan memudahkan pengguna dalam mengelola aspek-aspek penting dari sebuah toko seperti stok barang, staff, dan laporan penjualan.

## DB Schema

- products: id, name, price, stock
- staff: id, name, email, position
- sales: id, product_id, quantity, sale_date

## TODO

1. [ ] Tambah Produk Memungkinkan pengguna untuk menambah produk baru ke dalam database. Pengguna akan diminta untuk memasukkan nama produk, harga, dan jumlah stok awal.
2. [ ] Ubah Stok Produk Mengizinkan pengguna untuk mengubah stok produk yang ada. Pengguna dapat menambah atau mengurangi jumlah stok berdasarkan kebutuhan.
3. [ ] Tambah Staff Memungkinkan pengguna untuk menambah staff baru. Pengguna akan diminta untuk memasukkan nama, email, dan posisi dari staff yang akan ditambahkan.
4. [ ] Rekap Penjualan Menampilkan laporan penjualan berdasarkan periode tertentu. Pengguna dapat memilih periode waktu untuk melihat total penjualan, jumlah produk terjual, dan total pendapatan.
5. [ ] Exit Keluar dari aplikasi. Mengakhiri sesi penggunaan aplikasi CLI.
6. [ ] Create DB SQL scripts
7. [ ] Create DB connection


