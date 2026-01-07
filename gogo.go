Mulai
    Deklarasi a, b sebagai real
    Deklarasi op sebagai string
    Deklarasi hasil sebagai real
    Tampilkan "Masukkan angka pertama: "
    Baca a
    Tampilkan "Masukkan operator (+, -, *, /): "
    Baca op
    Tampilkan "Masukkan angka kedua: "
    Baca b
    Jika op = "+" maka
        hasil ← a + b
    Jika op = "-" maka
        hasil ← a - b
    Jika op = "*" maka
        hasil ← a * b
    Jika op = "/" maka
        Jika b = 0 maka
            Tampilkan "Error: pembagian dengan nol tidak diperbolehkan."
            Selesai
        Endif
        hasil ← a / b
    Jika tidak maka
        Tampilkan "Operator tidak valid."
        Selesai
    Endif
    Tampilkan "Hasil: ", hasil
Selesai
