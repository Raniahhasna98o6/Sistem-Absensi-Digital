-- Insert User Mahasiswa
INSERT INTO User (email, password, role) VALUES 
('ahmad.fauzan@student.telkomuniversity.ac.id', 'fauzan123', 'mahasiswa'),
('rizky.pratama@student.telkomuniversity.ac.id', 'rizky123', 'mahasiswa'),
('dimas.saputra@student.telkomuniversity.ac.id', 'dimas123', 'mahasiswa'),
('andika.putra@student.telkomuniversity.ac.id', 'andika123', 'mahasiswa'),
('fajar.hidayat@student.telkomuniversity.ac.id', 'fajar123', 'mahasiswa'),
('nurul.aisyah@student.telkomuniversity.ac.id', 'aisyah123', 'mahasiswa'),
('siti.khairunnisa@student.telkomuniversity.ac.id', 'khairun123', 'mahasiswa'),
('annisa.putri@student.telkomuniversity.ac.id', 'annisa123', 'mahasiswa'),
('dewi.lestari@student.telkomuniversity.ac.id', 'dewi123', 'mahasiswa'),
('tiara.maharani@student.telkomuniversity.ac.id', 'tiara123', 'mahasiswa');

-- Insert Mahasiswa
INSERT INTO mahasiswa (nim, id_user, id_kelas, nama) VALUES 
('1301221001', 3, 1, 'Ahmad Fauzan'),
('1301221002', 4, 1, 'Rizky Pratama'),
('1301221003', 5, 1, 'Dimas Saputra'),
('1301221004', 6, 1, 'Andika Putra'),
('1301221005', 7, 1, 'Fajar Hidayat'),
('1301221006', 8, 2, 'Nurul Aisyah'),
('1301221007', 9, 2, 'Siti Khairunnisa'),
('1301221008', 10, 2, 'Annisa Putri'),
('1301221009', 11, 2, 'Dewi Lestari'),
('1301221010', 12, 2, 'Tiara Maharani');

-- Insert User Dosen
INSERT INTO User (email, password, role) VALUES 
('budi.santoso@telkomuniversity.ac.id', 'budi123', 'dosen'),
('agus.setiawan@telkomuniversity.ac.id', 'agus123', 'dosen'),
('hendro.wibowo@telkomuniversity.ac.id', 'hendro123', 'dosen'),
('ratna.dewi@telkomuniversity.ac.id', 'ratna123', 'dosen'),
('sri.mulyani@telkomuniversity.ac.id', 'sri123', 'dosen');

-- Insert Dosen
INSERT INTO dosen (nidn, id_user, nama) VALUES 
('0412345684', 13, 'Budi Santoso'),
('0412345685', 14, 'Agus Setiawan'),
('0412345686', 15, 'Hendro Wibowo'),
('0412345687', 16, 'Ratna Dewi'),
('0412345688', 17, 'Sri Mulyani');
