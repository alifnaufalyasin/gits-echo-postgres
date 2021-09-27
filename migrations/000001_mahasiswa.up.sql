CREATE TABLE IF NOT EXISTS mahasiswa(
  id varchar(50) PRIMARY KEY NOT NULL,
  nama varchar(255) NOT NULL,
  umur bigint NOT NULL,
  kelas varchar(50) NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL
);