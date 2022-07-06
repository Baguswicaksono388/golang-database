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

# Database Pooling
- sql.DB di Go-Lang sebenarnya bukanlah sebuah koneksi ke database
- Melainkan sebuah pool ke database, atau dikenal dgn konsep Database Pooling
- Di dalam sql.DB, Golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
- Dgn kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yg dibuat di Go-Lang, sehingga tdk membanjiri koneksi ke database, krn biasanya ada batas maksimal koneksi yg bisa ditangani oleh database yg kita gunakan.
- Manajemen Koneksi bisa disebut Pooling

# Pengaturan Database Pooling
- (DB) SetMaxIdleConns(number) => Pengaturan beberapa jumlah koneksi minimal yg dibuat
- (DB) SetMaxOpenConns(number) => Pengaturan beberapa jumlah koneksi maksimal yg dibuat
- (DB) SetConnMaxIdleTime(duration) => Pengaturan beberapa lama koneksi yg sdh tdk digunakan akan dihapus
- (DB) SetConnMaxLifetime(duration) => Pengaturan beberapa lama koneksi boleh digunakan

# Eksekusi Perintah SQL
- Saat membuat aplikasi menggunakan database, sdh pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL
- Di Golang menyediakan func yg bisa kita gunakan unk mengirim perintah SQL ke database menggunakan func (DB) ExecContext(context, sql, params)
- Ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yg sdh pernah kita pelajari di course Golang Context, dgn context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya