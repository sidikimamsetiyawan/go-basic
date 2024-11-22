# Soal

Given two arrays of strings `a1` and `a2` return a sorted array `r` in lexicographical order of the strings of `a1` which are substrings of strings of `a2`.

#### Example 1:

`a1 = ["arp", "live", "strong"]`

`a2 = ["lively", "alive", "harp", "sharp", "armstrong"]`

returns `["arp", "live", "strong"]`

#### Example 2:

`a1 = ["tarp", "mice", "bull"]`

`a2 = ["lively", "alive", "harp", "sharp", "armstrong"]`

returns `[]`

#### Notes:

* Arrays are written in "general" notation. See "Your Test Cases" for examples in your language.
* In Shell bash `a1` and `a2` are strings. The return is a string where words are separated by commas.
* Beware: In some languages `r` must be without duplicates.

---

# Penyelesaian


Untuk menyelesaikan masalah ini, mari kita ikuti langkah-langkah berikut dalam bahasa yang lebih sederhana:

1. **Pahami Input dan Output**:

   - Kita memiliki dua array string sebagai input: `a1` dan `a2`.
   - Tujuan kita adalah mencari string dari `a1` yang muncul sebagai bagian dari string di `a2`.
   - Hasil yang diinginkan adalah array `r` yang berisi string dari `a1` yang ditemukan sebagai substring di `a2`, disusun dalam urutan leksikografis (alfabetis).
2. **Identifikasi Tugas Utama**:

   - Untuk setiap string di `a1`, kita perlu memeriksa apakah string tersebut muncul sebagai substring di salah satu string dalam `a2`.
   - Substring berarti seluruh string dari `a1` muncul berturut-turut di dalam sebuah string di `a2`, setidaknya satu kali.
3. **Logika Penyelesaian**:

   - **Inisialisasi Array Hasil Kosong**: Mulai dengan array kosong `r` untuk menyimpan string dari `a1` yang cocok.
   - **Iterasi pada `a1`**: Untuk setiap string dalam `a1`, lakukan langkah-langkah berikut:
     - **Periksa Setiap String di `a2`**: Untuk setiap string di `a2`, cek apakah string saat ini dari `a1` adalah substring di dalamnya.
     - **Simpan Hasil yang Cocok**: Jika string dari `a1` ditemukan sebagai substring di salah satu string `a2`, tambahkan ke array hasil `r`.
     - **Hindari Duplikasi**: Pastikan setiap string yang cocok hanya ditambahkan sekali ke `r` (jika ada duplikasi di `a1`, gunakan struktur data yang hanya menyimpan elemen unik).
   - **Urutkan Hasil**: Setelah semua string yang cocok terkumpul di `r`, urutkan array tersebut secara leksikografis untuk memenuhi syarat soal.
4. **Perhatikan Kasus Khusus**:

   - **Array Kosong**: Jika `a1` atau `a2` kosong, hasilnya harus berupa array kosong karena tidak ada string yang bisa dicocokkan.
   - **Tidak Ada Substring yang Cocok**: Jika tidak ada string dari `a1` yang ditemukan sebagai substring di `a2`, hasilnya juga harus berupa array kosong.
   - **Duplikasi di `a1`**: Jika `a1` memiliki duplikasi, pastikan setiap hasil yang cocok hanya muncul sekali di array hasil.

Dengan mengikuti pendekatan ini, kita bisa membuat solusi yang efisien tanpa langsung masuk ke kode, yang membantu memahami logika dan mengidentifikasi kasus khusus yang mungkin muncul.


# Links

* [Codewars detail code](https://www.codewars.com/kata/reviews/5c18fb45e4a709a26a001952/groups/672eda979c0da96bae9feb7e)
