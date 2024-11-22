
---

# CRUD Album with Go and MySQL

## Deskripsi
Proyek ini adalah aplikasi sederhana yang menunjukkan bagaimana cara membuat operasi CRUD (Create, Read, Update, Delete) untuk tabel **album** menggunakan bahasa pemrograman Go dan database MySQL.

---

## Fitur
- **Create**: Menambahkan album baru ke database.
- **Read**:
  - Menampilkan semua album.
  - Menampilkan detail album berdasarkan ID.
- **Update**: Memperbarui informasi album berdasarkan ID.
- **Delete**: Menghapus album berdasarkan ID.

---

## Prasyarat
Pastikan Anda memiliki perangkat lunak berikut:
1. **Go** (minimal versi 1.18)
2. **MySQL** (atau server MySQL lainnya)
3. **Git** 

---

## Instalasi

### 1. Clone Repository
```bash
git clone https://github.com/mochavin/go-database-access.git
cd go-database-access
```

### 2. Buat Database
Jalankan skrip SQL berikut di MySQL:
```sql
CREATE DATABASE albumdb;

USE albumdb;

CREATE TABLE album (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL
);
```

### 3. Konfigurasi Koneksi Database
Edit file `db/db.go` dan sesuaikan `dsn` dengan kredensial MySQL Anda:
```go
dsn := "root:@tcp(127.0.0.1:3306)/albumdb"
```

### 4. Instal Dependency
Proyek ini menggunakan driver MySQL untuk Go. Instal dependensi dengan perintah berikut:
```bash
go get -u github.com/go-sql-driver/mysql
```

### 5. Jalankan Aplikasi
```bash
go run main.go
```

---

## Struktur Direktori
```plaintext
go-database-access/
├── main.go            # Entry point aplikasi
├── db/
│   └── db.go          # Koneksi database
├── models/
│   └── album.go       # Model data album
└── handlers/
    └── album.go       # Logika CRUD album
```

---

## API Endpoint

### **1. Menampilkan Semua Album**
- **URL**: `/albums`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "id": 1,
      "title": "Album Title",
      "artist": "Artist Name",
      "price": 19.99
    }
  ]
  ```

### **2. Menampilkan Album Berdasarkan ID**
- **URL**: `/album?id={id}`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "id": 1,
    "title": "Album Title",
    "artist": "Artist Name",
    "price": 19.99
  }
  ```

### **3. Menambahkan Album**
- **URL**: `/albums`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "title": "New Album",
    "artist": "New Artist",
    "price": 15.99
  }
  ```
- **Response**:
  ```json
  {
    "id": 2,
    "title": "New Album",
    "artist": "New Artist",
    "price": 15.99
  }
  ```

### **4. Memperbarui Album**
- **URL**: `/album?id={id}`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "title": "Updated Album",
    "artist": "Updated Artist",
    "price": 25.00
  }
  ```
- **Response**:
  ```
  Album with ID {id} updated successfully
  ```

### **5. Menghapus Album**
- **URL**: `/album?id={id}`
- **Method**: `DELETE`
- **Response**:
  ```
  Album with ID {id} deleted successfully
  ```

---

## Pengujian
Gunakan alat seperti **Postman** atau **curl** untuk menguji endpoint. Contoh penggunaan `curl`:
```bash
# Tambah album
curl -X POST -H "Content-Type: application/json" \
-d '{"title":"Test Album", "artist":"Test Artist", "price":12.99}' \
http://localhost:8080/albums

# Lihat semua album
curl http://localhost:8080/albums

# Lihat album berdasarkan ID
curl http://localhost:8080/album?id=1

# Perbarui album
curl -X PUT -H "Content-Type: application/json" \
-d '{"title":"Updated Album", "artist":"Updated Artist", "price":20.00}' \
http://localhost:8080/album?id=1

# Hapus album
curl -X DELETE http://localhost:8080/album?id=1
```

---