-- Membuat Database
CREATE DATABASE IF NOT EXISTS db_absensi;
USE db_absensi;

-- 1. Tabel USER (Untuk Autentikasi Terpusat)
CREATE TABLE User (
    id_user INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role ENUM('mahasiswa', 'dosen') NOT NULL
);

-- 2. Tabel KELAS (Dibuat lebih dulu karena Mahasiswa butuh id_kelas)
CREATE TABLE kelas (
    id_kelas INT AUTO_INCREMENT PRIMARY KEY,
    nama_kelas VARCHAR(50) NOT NULL,
    kode_mk VARCHAR(20) NOT NULL
);

-- 3. Tabel MAHASISWA
CREATE TABLE mahasiswa (
    nim VARCHAR(20) PRIMARY KEY,
    id_user INT NOT NULL,
    id_kelas INT NOT NULL,
    nama VARCHAR(100) NOT NULL,
    FOREIGN KEY (id_user) REFERENCES User(id_user) ON DELETE CASCADE,
    FOREIGN KEY (id_kelas) REFERENCES kelas(id_kelas) ON DELETE CASCADE
);

-- 4. Tabel DOSEN
CREATE TABLE dosen (
    nidn VARCHAR(20) PRIMARY KEY,
    id_user INT NOT NULL,
    nama VARCHAR(100) NOT NULL,
    FOREIGN KEY (id_user) REFERENCES User(id_user) ON DELETE CASCADE
);

-- 5. Tabel ABSENSI (foto_abs menggunakan LONGTEXT untuk menampung Base64)
CREATE TABLE absensi (
    id_absensi INT AUTO_INCREMENT PRIMARY KEY,
    nim VARCHAR(20) NOT NULL,
    tanggal_abs DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_abs VARCHAR(20) NOT NULL,
    lokasi_abs TEXT NOT NULL,
    foto_abs LONGTEXT NOT NULL,
    latitude DOUBLE NOT NULL,
    longitude DOUBLE NOT NULL,
    FOREIGN KEY (nim) REFERENCES mahasiswa(nim) ON DELETE CASCADE
);
