---
--- INSERT DATA DUMMY
---

-- Insert User
INSERT INTO User (id_user, nama, email, password) VALUES
('U01', 'Budi Santoso', 'budi@mail.com', 'pass123'),
('U02', 'Ani Wijaya', 'ani@mail.com', 'pass123'),
('U03', 'Joko Susilo', 'joko@mail.com', 'pass123'),
('U04', 'Siti Aminah', 'siti@mail.com', 'pass123'),
('U05', 'Dina Lestari', 'dina@mail.com', 'pass123'),
('U06', 'Eko Prasetyo', 'eko@mail.com', 'pass123'),
('U07', 'Rina Permata', 'rina@mail.com', 'pass123'),
('U08', 'Andi Hakim', 'andi@mail.com', 'pass123'),
('U09', 'Sari Utami', 'sari@mail.com', 'pass123'),
('U10', 'Rizky Fauzi', 'rizky@mail.com', 'pass123'),
('D01', 'Dr. Aris', 'aris@staff.com', 'dosen123'),
('D02', 'Dr. Bella', 'bella@staff.com', 'dosen123'),
('D03', 'Prof. Citra', 'citra@staff.com', 'dosen123'),
('D04', 'Deni, M.T.', 'deni@staff.com', 'dosen123'),
('D05', 'Euis, M.Kom', 'euis@staff.com', 'dosen123'),
('D06', 'Fajar, Ph.D', 'fajar@staff.com', 'dosen123'),
('D07', 'Gita, M.Sc', 'gita@staff.com', 'dosen123'),
('D08', 'Hadi, M.Eng', 'hadi@staff.com', 'dosen123'),
('D09', 'Indah, M.Si', 'indah@staff.com', 'dosen123'),
('D10', 'Jaka, M.T.', 'jaka@staff.com', 'dosen123');

-- Insert Mahasiswa
INSERT INTO Mahasiswa (id_user, nim) VALUES
('U01', 'NIM2024001'), ('U02', 'NIM2024002'), ('U03', 'NIM2024003'),
('U04', 'NIM2024004'), ('U05', 'NIM2024005'), ('U06', 'NIM2024006'),
('U07', 'NIM2024007'), ('U08', 'NIM2024008'), ('U09', 'NIM2024009'),
('U10', 'NIM2024010');

-- Insert Dosen (Kolom NIDN disesuaikan)
INSERT INTO Dosen (id_user, NIDN) VALUES
('D01', 'NIDN1001'), ('D02', 'NIDN1002'), ('D03', 'NIDN1003'),
('D04', 'NIDN1004'), ('D05', 'NIDN1005'), ('D06', 'NIDN1006'),
('D07', 'NIDN1007'), ('D08', 'NIDN1008'), ('D09', 'NIDN1009'),
('D10', 'NIDN1010');

-- Insert Ruangan
INSERT INTO Ruangan (nama_ruangan, kapasitas) VALUES
('Lab Komputer 1', 40), ('Lab Komputer 2', 40), ('Ruang Teori A', 50),
('Ruang Teori B', 50), ('Aula Utama', 200), ('Ruang Seminar', 30),
('Lab AI', 25), ('Lab Jaringan', 30), ('Ruang 301', 45), ('Ruang 302', 45);

-- Insert Mata_Kuliah (Kolom id_user sesuai CREATE TABLE)
INSERT INTO Mata_Kuliah (kode_mk, id_user, id_ruangan, nama_mk, jadwal) VALUES
('MK01', 'D01', 1, 'Basis Data', 'Senin 08:00'),
('MK02', 'D02', 2, 'Algoritma', 'Selasa 10:00'),
('MK03', 'D03', 3, 'Pemrograman Web', 'Rabu 13:00'),
('MK04', 'D04', 4, 'Jaringan Komputer', 'Kamis 09:00'),
('MK05', 'D05', 7, 'Kecerdasan Buatan', 'Jumat 14:00'),
('MK06', 'D06', 8, 'Keamanan Siber', 'Senin 13:00'),
('MK07', 'D07', 9, 'Sistem Operasi', 'Selasa 08:00'),
('MK08', 'D08', 10, 'Arsitektur Komputer', 'Rabu 10:00'),
('MK09', 'D09', 3, 'Etika Profesi', 'Kamis 13:00'),
('MK10', 'D10', 4, 'Bahasa Inggris', 'Jumat 08:00');

-- Insert KRS
INSERT INTO KRS (id_user, kode_mk, semester) VALUES
('U01', 'MK01', 1), ('U02', 'MK01', 1), ('U03', 'MK01', 1),
('U04', 'MK02', 1), ('U05', 'MK02', 1), ('U06', 'MK03', 1),
('U07', 'MK03', 1), ('U08', 'MK04', 1), ('U09', 'MK04', 1),
('U10', 'MK05', 1);

-- Insert Nilai (Nama kolom disesuaikan: nilai_angka, nilai_index)
INSERT INTO Nilai (id_user, kode_mk, nilai_angka, nilai_index) VALUES
('U01', 'MK01', 88, 'A'), ('U02', 'MK01', 75, 'B'),
('U03', 'MK01', 0, 'E'), ('U04', 'MK02', 90, 'A'),
('U05', 'MK02', 82, 'A'), ('U06', 'MK03', 70, 'B'),
('U07', 'MK03', 85, 'A'), ('U08', 'MK04', 65, 'C'),
('U09', 'MK04', 78, 'B'), ('U10', 'MK05', 92, 'A');

-- Insert Absensi
INSERT INTO Absensi (id_user, kode_mk, tanggal_abs, status_abs, lokasi_abs) VALUES
('U01', 'MK01', '2024-05-20', 'Hadir', 'Gedung A - Lab 1'),
('U02', 'MK01', '2024-05-20', 'Hadir', 'Gedung A - Lab 1'),
('U03', 'MK01', '2024-05-20', 'Tidak Hadir', 'Rumah (Surat Terlampir)'),
('U04', 'MK02', '2024-05-21', 'Hadir', 'Gedung B - Lab 2'),
('U05', 'MK02', '2024-05-21', 'Hadir', 'Rumah'),
('U06', 'MK03', '2024-05-22', 'Hadir', 'Gedung C - Ruang 3'),
('U07', 'MK03', '2024-05-22', 'Hadir', 'Gedung C - Ruang 3'),
('U08', 'MK04', '2024-05-23', 'Tidak Hadir', 'Tidak Terdeteksi'),
('U09', 'MK04', '2024-05-23', 'Hadir', 'Gedung D - Ruang 4'),
('U10', 'MK05', '2024-05-24', 'Hadir', 'Gedung A - Lab AI');

-- Insert laporan
INSERT INTO Laporan (id_user, periode) VALUES
('D01', 'Ganjil 2023/2024'),
('D02', 'Ganjil 2023/2024'),
('D03', 'Genap 2023/2024'),
('D04', 'Genap 2023/2024'),
('U01', 'Mei 2024'),
('U02', 'Mei 2024'),
('U03', 'Mei 2024'),
('D05', 'Juni 2024'),
('D06', 'Juni 2024'),
('U04', 'Juli 2024');

-- update
-- 1. Mahasiswa U08 awalnya Tidak Hadir, Dosen memverifikasi karena masalah jaringan/GPS.
UPDATE Absensi 
SET status_abs = 'Hadir', lokasi_abs = 'Terverifikasi Dosen (Kendala GPS)' 
WHERE id_user = 'U08' AND kode_mk = 'MK04' AND tanggal_abs = '2024-05-23';

-- 2. Mahasiswa U05 sudah mengirim surat dokter, status diubah dari Hadir ke Tidak Hadir (prosedur kampus).
UPDATE Absensi 
SET status_abs = 'Tidak Hadir', lokasi_abs = 'Rumah (Surat Dokter Terlampir)' 
WHERE id_user = 'U05' AND kode_mk = 'MK02' AND tanggal_abs = '2024-05-21';

-- 3. Mahasiswa U03 yang tadinya Tidak Hadir, ternyata tetap masuk kelas (Hadir).
UPDATE Absensi 
SET status_abs = 'Hadir', lokasi_abs = 'Gedung A - Lab 1' 
WHERE id_user = 'U03' AND kode_mk = 'MK01' AND tanggal_abs = '2024-05-20';

-- 4. Koreksi lokasi absen Mahasiswa U01 agar lebih spesifik.
UPDATE Absensi 
SET lokasi_abs = 'Gedung A - Lantai 2, Ruang 201' 
WHERE id_user = 'U01' AND kode_mk = 'MK01' AND tanggal_abs = '2024-05-20';

-- 5. Mahasiswa U09 absen Hadir tapi lupa mengisi lokasi, sistem mengupdate lokasinya.
UPDATE Absensi 
SET lokasi_abs = 'Gedung D - Ruang 4 (Auto-Detect)' 
WHERE id_user = 'U09' AND kode_mk = 'MK04' AND tanggal_abs = '2024-05-23';

-- 6. Mahasiswa U02 awalnya Hadir, tapi ketahuan titip absen (diubah ke Tidak Hadir).
UPDATE Absensi 
SET status_abs = 'Tidak Hadir', lokasi_abs = 'Diblokir (Pelanggaran Absensi)' 
WHERE id_user = 'U02' AND kode_mk = 'MK01' AND tanggal_abs = '2024-05-20';

-- 7. Update status mahasiswa U04 dari Hadir ke Tidak Hadir karena pulang di tengah perkuliahan.
UPDATE Absensi 
SET status_abs = 'Tidak Hadir', lokasi_abs = 'UKS Kampus' 
WHERE id_user = 'U04' AND kode_mk = 'MK02' AND tanggal_abs = '2024-05-21';


-- 8. Mahasiswa U07 melakukan update lokasi karena berpindah ruangan saat praktikum.
UPDATE Absensi 
SET lokasi_abs = 'Gedung C - Ruang 5 (Pindah Ruang)' 
WHERE id_user = 'U07' AND kode_mk = 'MK03' AND tanggal_abs = '2024-05-22';

-- 9. Mengubah status Hadir menjadi Tidak Hadir karena mahasiswa memberikan alasan yang logis setelah kelas.
UPDATE Absensi 
SET status_abs = 'Tidak Hadir', lokasi_abs = 'Keluarga Berduka (Verifikasi Susulan)' 
WHERE status_abs = 'Hadir' AND kode_mk = 'MK04';

-- 10. Menandai lokasi absen Tidak Hadir agar diverifikasi oleh admin
UPDATE Absensi 
SET lokasi_abs = CONCAT(IFNULL(lokasi_abs, ''), ' - Menunggu Verifikasi Admin') 
WHERE status_abs = 'Tidak Hadir' AND tanggal_abs = '2024-05-21';
