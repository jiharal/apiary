# APIARY

## Kebutuhan perangkat lunak.

- Visual Studio Code (Aplikasi editor).
  > Paduan instalasi [disini](https://code.visualstudio.com/docs/setup/setup-overview).
- Go (Bahasa pemrograman).
  > Paduan instalasi [disini](https://golang.org/doc/install).
- PostgresQL (database).
  > Paduan instalasi [disini](https://www.tutorialspoint.com/postgresql/postgresql_environment.htm)
- Redis (Caching).

  > Paduan instalasi [disini](https://www.tutorialspoint.com/redis/redis_environment.htm)

## Struktur projek

```
.
├── api
│   ├── api.product.go
│   ├── api.user.go
│   ├── error.go
│   ├── init.go
│   ├── response.go
│   └── session.go
├── auth
│   ├── init.go
│   └── middleware.go
├── cmd
│   └── cmd.go
├── files
│   └── db.sql
├── main.go
├── model
│   ├── model.product.go
│   └── model.user.go
├── router
│   ├── handler.go
│   ├── handler.product.go
│   ├── handler.user.go
│   └── init.go
└── system
    └── id.go
```

- `api` direktori untuk menyimpan file \*.go yang berhubungan dengan bisnis logic.
- `auth` direktori untuk menyimpan file yang berhubungan dengan authentikasi.
- `cmd` direktori untuk penyimpan file inti yang berhubungan dengan configurasi, dll.
- `model` direktori untuk menyimpan file query database.
- `files` direktori untuk menyimpan file assets laianya seperi file `.sql, .png, etc...`
- `router` direktori untuk menyimpan handler.
- `sistem` direktori untuk menyimpan file didalamnya mengandung fungsi umum yang bisa dipakai berkali-kali.

## Cara menjalankan

- buatlah table (file terdapat pada direktori `files/db.sql`) pada database `apiary_db`.
- jalankan redis.
- pada sub direktori projek, jalankan sistem menggunakan perintah `go run main.go`
- selesai.

## Testing Response API

### login

![image](https://user-images.githubusercontent.com/21150538/67977344-e0c47500-fc52-11e9-85de-fce3e269bb7f.png)

### register

![image](https://user-images.githubusercontent.com/21150538/67977386-f5087200-fc52-11e9-9203-4e4f56e17106.png)

### Daftar produk

![image](https://user-images.githubusercontent.com/21150538/67977537-41ec4880-fc53-11e9-9bd1-1cc398e256e2.png)

### Rincian produk berdasarkan ID

![image](https://user-images.githubusercontent.com/21150538/67977586-592b3600-fc53-11e9-80c0-cfa85687eb82.png)

### Membuat produk baru

![image](https://user-images.githubusercontent.com/21150538/67977485-284b0100-fc53-11e9-88a4-f1e8ab95b0c2.png)

### Update produk berdasarkan ID

![image](https://user-images.githubusercontent.com/21150538/67977629-6f38f680-fc53-11e9-9e32-b9838a82bc88.png)

### Menghapus produk berdasarkan ID

![image](https://user-images.githubusercontent.com/21150538/67977685-85df4d80-fc53-11e9-8675-aba6c56a2fc6.png)

## Dokuementasi lengkap disini.

- [here](https://documenter.getpostman.com/view/1919076/SW12zHdM?version=latest)
