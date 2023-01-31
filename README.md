# belajar-golang-validation

## Contents
- [Pengenalan Validation](#pengenalan-validation)
- [Validator Package](#Validator-Package)
- [Validate Struct](#Validate-Struct)
- [Validasi Variable](#Validasi-Variable)
- [Validasi Dua Variable](#Validasi-Dua-Variable)
- [Baked-In Validations](#Baked-In-Validation)
- [Multiple Tag Validation](#Multiple-Tag-Validation)
- [Tag Parameter](#Tag-Parameter)
- [Validasi Struct](#Validasi-Struct)
- [Validation Errors](#Validation-Errors)
- [Validasi Cross Field](#Validasi-Cross-Field)
- [Validasi Nested Struct](#Validasi-Nested-Struct)
- [Validasi Collection](#Validasi-Collection)
- [Validasi Basic Collection](#Validasi-Basic-Collection)
- [Validasi Map](#Validasi-Map)
- [Validasi Basic Map](#Validasi-Basic-Map)
- [Alias Tag](#Alias-Tag)
- [Custom Validation](#Custom-Validation)
- [Custom Validation Paramaeter](#Custom-Validation-Parameter)
- [Or Rule](#Or-Rule)
- [Custom Validation Cross Field](#Custom-Validation-Cross-Field)
- [Struct Level Validation](#Struct-Level-Validation)


# Pengenalan Validation
Validation
- Saat kita membuat aplikasi, validasi adalah salah satu hal yang selalu dibuat
- Validasi dilakukan untuk memastikan bahwa data yang diproses sudah benar
- Validasi adalah sesuatu yang wajib dilakukan saat pembuatan aplikasi, agar kesalahan pada data bisa ditemukan secepat mungkin sebelum data tersebut di proses
Tempat Melakukan Validasi
- Validasi sering dilakukan di banyak bagian dalam aplikasi, seperti
- Web, validasi request dari pengguna
- Business Logic, validasi data 
- Database, validasi constraint
- Beberapa bagian, kadang menggunakan validasi yang sama. Oleh karena itu pembuatan validasi secara manual sangat memakan waktu, dan kesalahan sedikit bisa menyebabkan validasi tidak konsisten
Kenapa Butuh Validasi?
- Sederhana, untuk memastikan request atau data yang dikirim oleh pengguna sudah sesuai dengan yang kita inginkan
- Never trust user input

# Validator Package
 Manual Validation
- Saat kita melakukan validasi, biasanya kita akan melakukan validasi secara manual
- Rata-rata, validasi manual akan menggunakan if statement
- Semakin banyak validasi yang diperlukan, semakin banyak if statement yang harus dibuat
Validation Library
- Penggunaan library untuk melakukan validasi sangat direkomendasikan
- Hal ini agar kode validasi bisa lebih mudah, rapi dan juga sama antar programmer
- Ada banyak sekali library yang bisa kita gunakan untuk mempermudah dalam pembuatan validasi di kode program kita
- Salah satunya, yang akan kita gunakan di kelas ini adalah Validator Package
Validator Package
- Salah satu library yang banyak digunakan untuk membantu melakukan validation adalah Validator Package
- Validator Package adalah opensource library untuk melakukan validation di golang
- Validator Package memiliki banyak sekali fitur yang bisa kita gunakan untuk mempermudah kita melakukan validasi
```go
https://github.com/go-playground/validator
```
# Validate Struct
Validate Struct
- Validator Package di desain agar thread safe dan digunakan sebagai object singleton (cukup buat satu object saja)
- Validator Package akan melakukan cache informasi seperti rules, tags, dan lain-lain yang berhubungan dengan validation kita
- Cache adalah menyimpan informasi di memory, sehingga bisa digunakan lagi tanpa harus melakukan eksekusi kode program lagi, hal ini akan mempercepat proses
- Jika kita selalu membuat object baru, maka keuntungan cache tidak bisa didapatkan
- Package Validator merepresentasikan object untuk validation nya dalam struct bernama Validate di package validator

# Validasi Variable
Validasi Variable
- Sekarang kita akan coba melakukan validasi terhadap variable menggunakan Validator Package
- Saat kita melakukan validasi, biasanya kita akan melakukan validasi terhadap variable yang tersedia, baik itu variable yang terdapat di Struct atau function parameter
- Untuk melakukan validasi variable, kita bisa gunakan method :
    - Validate.Var(variable, tag)
    - Validate.VarCtx(ctx, variable, tag)
Validator Tag
- Validator Package menggunakan konsep Tag untuk menyebutkan validation yang ingin digunakan
- Setiap jenis validation, memiliki nama tag masing-masing, kita akan bahas di materi khusus untuk ini
- Contoh validasi bahwa sebuah value harus ada isinya, bukan default value, misal kita bisa gunakan tag “required”, jika number pastikan bukan default number 0, dan jika string bukan default string “”, dan jika array/slice, bukan default array/slice kosong

# Validasi Dua Variable
Validasi Dua Variable
- Kadang, ada kasus dimana kita ingin melakukan validasi untuk membandingkan dua buah variable
- Misal kita mau memastikan variable password dan confirmPassword harus sama, kita bisa gunakan tag eqfield
- Untuk melakukan validasi dua variable, kita bisa gunakan method :
    - Validate.VarWithValue(first, second, tag)
    - Validate.VarWithValueCtx(ctx, first, second, tag)

# Baked-In Validations
Baked-in Validations
- Validator Package sudah menyediakan banyak sekali validation yang biasanya kita butuhkan, dari yang sederhana seperti required (string tidak boleh kosong), email (format data harus email), dan lain-lain
- Ada banyak sekali Baked-in Validation yang sudah tersedia, dan kita hanya perlu menggunakan nama tag nya saja
```go
https://pkg.go.dev/github.com/go-playground/validator/v10#readme-baked-in-validations
```

# Multiple Tag Validation
Multiple Tag Validation
- Tag pada validation bisa lebih dari satu, kita bisa tambahkan , (koma) dan diikuti dengan tag selanjutnya

# Tag Parameter
Tag Parameter
- Validator Package, mendukung penggunaan parameter ketika menggunakan tag
- Ini sangat berguna pada kasus validasi yang memang butuh data parameter, contohnya min, max, length, dan lain-lain
- Untuk menggunakan parameter, kita bisa langsung menggunakan tanda = (sama dengan) setelah tag, dan diikuti dengan nilai parameter nya

# Validasi Struct
Validasi Struct
- Selain melakukan validasi variable, Validator Package juga bisa digunakan untuk melakukan validasi terhadap struct
- Dengan begitu, kita bisa langsung melakukan validasi terhadap semua field yang terdapat di Struct tersebut
- Tag untuk validasi, bisa kita tambahkan dengan menambah reflection tag di Struct field-nya dengan tag validate

# Validation Errors
Error
- Saat kita melakukan validasi, Validator Package akan mengembalikan data error
- Jika error tersebut bernilai nil, artinya semua data valid, tidak terjadi validation error
- Namun jika tidak nill, artinya terdapat data yang error
Validation Errors
- Kita tahu bahwa error adalah kontrak interface dari golang untuk membuat error
- Validator Package sendiri sebenarnya memiliki detail struct untuk implementasi error ini, yaitu ValidationErrors
- Kita bisa melakukan konversi ke ValidationErrors ketika terjadi validation error
- Terdapat banyak sekali informasi yang bisa kita ambil dari ValidationErrors
- ValidationErrors sendiri sebenarnya adalah alias untuk []FieldError

# Validasi Cross Field
Validasi Cross Field
- Sebelumnya kita sudah tahu Validator Package memiliki validation khusus untuk validasi dua variable
- Validation tersebut juga bisa digunakan untuk validasi cross Field jika di dalam Struct
- Caranya kita bisa menggunakan validation tag yang sama, namun perlu sebutkan field kedua-nya
- Misal pada Field Password, kita bisa tambahkan validate:eqfield=ConfirmPassword
# Validasi Nested Struct
Validasi Nested Struct
- Secara default, saat kita membuat Struct yang berisikan Field Struct lainnya
- Validator Package akan melakukan validasi terhadap Field Struct tersebut secara otomatis

# Validasi Collection
Validasi Collection
- Tidak seperti tipe data Struct, jika kita memiliki field dengan tipe data Collection seperti Array, Slice atau Map, secara default Validator Package tidak akan melakukan validasi terhadap data-data yang terdapat di dalam collection tersebut
- Namun, jika kita ingin melakukan validasi semua data yang terdapat di Collection-nya, kita bisa tambahkan tag dive

# Validasi Basic Collection
Basic Collection
- Bagaimana jika data Collection nya adalah tipe data yang bukan Struct, misal []String?
- Pada kasus seperti ini, kita bisa tambahkan validation nya langsung setelah menambahkan tag dive
- Misal pada field Hobbies []String, kita bisa tambahkan tag dive,required,min=1, artinya tiap String di []String harus required dan min=1

# Validasi Map
Validasi Map
- Selain Collection Array/Slice, kita juga bisa melakukan validasi terhadap Field Map
- Karena dalam Map terdapat Key dan Value, kita bisa menggunakan dive untuk key dan dive untuk value
- Namun khusus untuk key, kita harus tandai dengan tag keys dan diakhiri dengan endkeys

# Validasi Basic Map
Basic Map
- Kadang, kita juga sering membuat tipe data Map dengan key dan value berupa tipe data bukan Struct
- Pada kasus ini, jika kita ingin menambah validasi, caranya sama seperti pada Basic Collection, namun khusus untuk key, perlu ditambahkan keys dan endkeys
- Field Wallet map[string]int dengan tag dive,keys,required,endkeys,required,gt=0, artinya key string required, dan value required dan gt=0
- Karena value bukanlah Struct, jadi kita tidak perlu menambahkan dive lagi pada value

# Alias Tag
Alias Tag
- Pada beberapa kasus, kadang kita sering menggunakan beberapa tag validation yang sama untuk Field yang berbeda
- Validator Package memiliki fitur untuk menambahkan alias, yaitu nama tag baru untuk tag lain, bisa satu atau lebih tag lain
- Kita bisa meregistrasikan alias tag baru dengan menggunakan method :
    - Validate.RegisterAlias(alias, tag)

# Custom Validation
Custom Validation
- Bagaimana jika ternyata validation yang kita butuhkan tidak tersedia di baked-in Validator Package?
- Tenang saja, kita bisa membuat Custom Validation sendiri, dengan membuat function dengan parameter validator.FieldLevel, lalu registrasikan ke Validate menggunakan :
Validate.RegisterValidation(tag, function)
Field Level
- FieldLevel merupakan parameter dari Validation Function yang kita buat
- FieldLevel berisikan informasi Reflection seperti Field Value, Name dan lain-lain

# Custom Validation Parameter
Custom Validation Parameter
- Di Baked-in Validation, beberapa validation memiliki parameter, misal min=10
- Kita juga mengambil informasi nilai parameter nya di FieldLevel.Param()
- Ini sangat cocok ketika kita membuat Validation yang memang membutuhkan parameter tambahan

# Or Rule
Or Rule
- Pada beberapa kasus, kadang kita ingin membuat kondisi OR pada validation
- Misal sebuah Field boleh email atau nomor telepon misalnya, artinya validasinya tidak bisa email,numeric
- Karena jika seperti itu, artinya Field wajib Email dan Numeric, sedangkan dua hal itu pasti berbeda
- Pada kasus ini, kita bisa gunakan | (pipe) sebagai pemisah untuk menandakan bahwa itu adalah OR
- Secara default, ketika menggunakan , (koma) artinya adalah AND

# Custom Validation Cross Field
Custom Validation Cross Field
- Saat membuat custom validation, parameter FieldLevel juga bisa digunakan untuk mendapatkan value kedua secara langsung dari Struct
- Kita bisa menggunakan method FieldLevel.GetStructFieldOK2()
- Method GetStructFieldOK2() secara otomatis akan menggunakan Param sebagai nama Field di Struct nya
- Misal ketika kita gunakan tag xxx=Yyy, maka GetStructFieldOK2() akan mengembalikan Field Yyy didalam Struct nya

# Struct Level Validation
Struct Level Validation
- Kadang ada kasus untuk melakukan validasi butuh kombinasi lebih dari dua field
- Sampai saat ini, kita hanya membuat validasi untuk single field, atau cross field
- Validator Package mendukung pembuatan validasi di level Struct, namun kita perlu membuat validation function menggunakan parameter StructLevel
- Kita bisa meregistrasikan validation nya menggunakan method Validate.RegisterStructValidation()
