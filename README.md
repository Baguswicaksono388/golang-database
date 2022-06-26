# Pengenalan Package Database
- Bahasa pemograman Go-Lang secara default memiliki sebuah package bernama database
- Package Database adlah package yg berisikan kumpulan standard interface yg menjadi standart unk berkomunikasi ke database
- Hal ini menjadikan kode program yg kita buat unk mengakses jenis database apapun bisa mennggunakan kode yg sama
- Yg berbeda hanya kode SQL yg perlu kita gunakan sesuai dgn database yg kita gunakan

# Database Driver
- Seblm kita membuat kode program menggunakan database di Go-Lang, terlebih dahulu kita wajib menambahkan driver database nya
- Tanpa Driver Database, maka package database di Go-Lang  tidak mengerti apapun, karena hanya berisi kontrak interface saja.
- Kalau bisa pilih yg paling populer digunakan, agara ketika error banyak referensinyago
- Sumber (https://golang.org/s/sqldriver, https://github.com/golang/go/wiki/SQLDrivers)

# Membuat Koneksi ke Database
- Hal pertama yg kita lakukan ketika membuat aplikasi yg menggunakan database adlh melakukan koneksi ke databasenya.
- Unk melakukan koneksi ke database di Go-Lang, kita membuat object sql.DB menggunakan func sql.Open(driver, dataSourceName)
- Unk menggunakan database MySQl, kita bisa menggunakan driver "mysql"
- Sedangkan unk dataSourceName, tiap database punya cara penulisan masing - masing, misal di MySQL, kita bisa menggunakan dataSourceName seperti : (username:password@tcp(host:port)/database_name)
- Jika object sql.DB sudah tdk digunakan lagi, disarankan unk menutupnya menggunakan function Close()