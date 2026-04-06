INSERT INTO User VALUES
('U1', 'Budi', 'budi@mail.com', '123'),
('U2', 'Ani', 'ani@mail.com', '123'),
('U3', 'Joko', 'joko@mail.com', '123'),
('U4', 'Siti', 'siti@mail.com', '123'),
('U5', 'Dina', 'dina@mail.com', '123');

INSERT INTO Mahasiswa VALUES
('U1', 'NIM001'),
('U2', 'NIM002'),
('U3', 'NIM003'),
('U4', 'NIM004'),
('U5', 'NIM005');

INSERT INTO Dosen VALUES
('U1', 'D001'),
('U2', 'D002'),
('U3', 'D003'),
('U4', 'D004'),
('U5', 'D005');

INSERT INTO Ruangan (nama_ruangan, kapasitas) VALUES
('Ruang A', 30),
('Ruang B', 25),
('Ruang C', 40),
('Ruang D', 35),
('Ruang E', 20);

INSERT INTO Mata_Kuliah VALUES
('MK01', 'U1', 1, 'Basis Data', 'Senin 08:00'),
('MK02', 'U2', 2, 'Algoritma', 'Selasa 10:00'),
('MK03', 'U3', 3, 'Pemrograman', 'Rabu 13:00'),
('MK04', 'U4', 4, 'Jaringan', 'Kamis 09:00'),
('MK05', 'U5', 5, 'AI', 'Jumat 14:00');

INSERT INTO KRS VALUES
('U1', 'MK01', 1),
('U2', 'MK02', 1),
('U3', 'MK03', 1),
('U4', 'MK04', 1),
('U5', 'MK05', 1);

INSERT INTO Nilai VALUES
('U1', 'MK01', 85, 'A'),
('U2', 'MK02', 78, 'B'),
('U3', 'MK03', 90, 'A'),
('U4', 'MK04', 70, 'B'),
('U5', 'MK05', 60, 'C');

INSERT INTO Absensi (id_user, kode_mk, tanggal_abs, status_abs, lokasi_abs) VALUES
('U1', 'MK01', '2024-01-01', 'Hadir', 'Kampus'),
('U2', 'MK02', '2024-01-02', 'Izin', 'Rumah'),
('U3', 'MK03', '2024-01-03', 'Sakit', 'Rumah'),
('U4', 'MK04', '2024-01-04', 'Hadir', 'Kampus'),
('U5', 'MK05', '2024-01-05', 'Alpha', 'Tidak Ada');

id_absensiid_absensiINSERT INTO Laporan (id_user, periode) VALUES
('U1', '2024-1'),
('U2', '2024-1'),
('U3', '2024-1'),
('U4', '2024-1'),
('U5', '2024-1');