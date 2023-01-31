# belajar-golang-validation

## Contents
- [Pengenalan Validation](#pengenalan-validation)
- [Validator Package](#Validator-Package)
- [Pengenalan Validation](#pengenalan-validation)
- [Pengenalan Validation](#pengenalan-validation)
- [Pengenalan Validation](#pengenalan-validation)
- [Pengenalan Validation](#pengenalan-validation)
- [Pengenalan Validation](#pengenalan-validation)
- [Pengenalan Validation](#pengenalan-validation)
- [Pengenalan Validation](#pengenalan-validation)


# Pengenalan Validation
# Validation
- Saat kita membuat aplikasi, validasi adalah salah satu hal yang selalu dibuat
- Validasi dilakukan untuk memastikan bahwa data yang diproses sudah benar
- Validasi adalah sesuatu yang wajib dilakukan saat pembuatan aplikasi, agar kesalahan pada data bisa ditemukan secepat mungkin sebelum data tersebut di proses
# Tempat Melakukan Validasi
- Validasi sering dilakukan di banyak bagian dalam aplikasi, seperti
- Web, validasi request dari pengguna
- Business Logic, validasi data 
- Database, validasi constraint
- Beberapa bagian, kadang menggunakan validasi yang sama. Oleh karena itu pembuatan validasi secara manual sangat memakan waktu, dan kesalahan sedikit bisa menyebabkan validasi tidak konsisten
# Kenapa Butuh Validasi?
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
