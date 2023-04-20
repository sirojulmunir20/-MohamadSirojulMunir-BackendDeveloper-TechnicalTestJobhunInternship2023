CREATE TABLE jurusan (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nama_jurusan VARCHAR(255) NOT NULL
);

CREATE TABLE mahasiswa (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    usia INT NOT NULL,
    gender INT NOT NULL,
    tanggal_reg DATETIME NOT NULL,
    id_jurusan INT NOT NULL,
    FOREIGN KEY (id_jurusan) REFERENCES jurusan(id)
);


CREATE TABLE hobi (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    nama_hobi VARCHAR(255) NOT NULL
);

CREATE TABLE mahasiswa_hobi (
    id_mahasiswa INT NOT NULL,
    id_hobi INT NOT NULL,
    FOREIGN KEY (id_mahasiswa) REFERENCES mahasiswa(id),
    FOREIGN KEY (id_hobi) REFERENCES hobi(id)
);