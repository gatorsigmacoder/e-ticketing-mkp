# Baca aku

## database

- Untuk database saya menggunakan postgresql 17
- Sudah disediakan docker compose jika menggunakan docker
- Pada project ini juga disediakan command untuk create database dan migrate database

1. Untuk create DB

```bash
go run cmd/createDatabase/createDatabase.go
```

2. Untuk migrate DB

```bash
go run cmd/migrateDatabase/migrateDatabase.go
```

- Disediakan juga backupan database postgresql pada project ini jika ingin backup manual

- Design database tersedia di https://dbdiagram.io/d/Tabel-e-ticketing-MKP-68dcffa7d2b621e422c37c7e

## How to run

- bisa menggunakan air (hot reload) atau manual dengan go
- User yang tersedia :

1. Admin  
   Email : admintest@mail.com  
   password : password
2. User biasa  
   Email : usertest@mail.com  
   password : password
