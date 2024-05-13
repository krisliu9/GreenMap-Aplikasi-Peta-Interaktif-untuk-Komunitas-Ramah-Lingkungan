# Green Map

## About Project

GreenMap adalah aplikasi mobile revolusioner yang lahir dari kebutuhan untuk mengatasi tantangan dalam mencari toilet umum sambil juga mendorong gaya hidup aktif dan sehat. Terinspirasi dari kesulitan dalam menemukan fasilitas umum, khususnya toilet, dan menyadari pentingnya aktivitas fisik untuk kesejahteraan secara keseluruhan, GreenMap didesain sebagai solusi yang menyeluruh. Aplikasi ini tidak hanya memberikan akses kepada pengguna untuk informasi vital seperti lokasi toilet umum, tetapi juga menggabungkan elemen-elemen permainan untuk mendorong pengguna berpartisipasi aktif dalam upaya pelestarian lingkungan. Melalui permainan interaktif yang membutuhkan gerakan fisik, GreenMap mendorong pengguna untuk mengadopsi gaya hidup lebih sehat sambil berkontribusi pada pelestarian lingkungan. Dengan menggabungkan pengalaman bermain permainan dengan mencari toilet umum dan mempromosikan aktivitas fisik, kami berusaha untuk membentuk rasa kepemilikan komunitas dan tanggung jawab kolektif terhadap lingkungan.

## Features

Dalam aplikasi ini terdapat 2 role, yaitu User dan Admin.

### User

- User bisa Login
- User bisa Register
- User bisa menambah lokasi ( map / pinpoint )
- User bisa melihat lokasi ( map / pinpoint )
- User bisa menghapus lokasi yang dibuat oleh dirinya sendiri ( map / pinpoint )
- User bisa mengupdate lokasi yang diupload olehnya ( map / pinpoint )
- User bisa mendapatkan point berdasarkan misi yang diselesaikan oleh user
- User bisa chatbot
- User bisa mereport pinpoint yang tidak sesuai (ex. tidak sesuai deskripsi, tidak benar sehingga bisa dihapus oleh admin )

### Admin

- Admin bisa melihat report yang dibuat
- Admin bisa menghapus report yang dibuat
- Admin bisa melihat dan menghapus pinpoint
- Admin bisa membuat misi
- Admin bisa mengupdate misi
- Admin bisa menghapus misi

## Tech Stacks

sebutkan daftar tools dan framework yang digunakan dalam bentuk list seperti ini:

- Golang
- MySQL
- GCP
- JWT
- ECHO
- Github
- Git
- Docker
- Postman
- Gorm
- Open AI

## API Documentation

https://documenter.getpostman.com/view/33039625/2sA3JNcgL2

## ERD

![](https://github.com/krisliu9/GreenMap-Aplikasi-Peta-Interaktif-untuk-Komunitas-Ramah-Lingkungan/blob/main/mini_project-ERD.jpg)

## Setup

1. buka aplikasi IDE
2. lakukan git clone https://github.com/krisliu9/GreenMap-Aplikasi-Peta-Interaktif-untuk-Komunitas-Ramah-Lingkungan.git
3. buat file .env di root folder dan masukkan `DB_USER=root
DB_PASSWORD=password
DB_HOST=34.87.99.166
DB_PORT=3306
DB_NAME=test_miniproject
DB_CHARSET=utf8mb4
DB_PARSE_TIME=True
DB_LOC=Local` data bisa disesuaikan dengan kebutuhan user
4. masukkan command mysql di mysql lokal berdasarkan query table file query.sql
5. jalankan go run main.go di terminal
