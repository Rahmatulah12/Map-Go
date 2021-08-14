package main

import "fmt"

func main() {

	/*
		Map
		Map adalah tipe data asosiatif yang ada di Golang.
		Bentuknya key-value, artinya setiap data (atau value) yang disimpan,
		disiapkan juga key-nya. Key tersebut harus unik, karena digunakan sebagai
		penanda (atau identifier) untuk pengaksesan data atau item yang tersimpan.
		Penggunaan Map Cara menggunakan map adalah dengan menuliskan keyword map
		diikuti tipe data key dan value-nya.
	*/
	var chickens map[string]int
	chickens = map[string]int{}

	chickens["januari"] = 12
	chickens["februari"] = 27
	chickens["maret"] = 13

	fmt.Println(chickens)
	fmt.Println(chickens["april"]) // jika key nya beluum tersedia maka nilai defaultnya bawaan tipe data

	// vertikal
	var students = map[string]string{"name": "Rahmatulah Sidik", "gender": "Pria"}

	// horizontal
	var profile = map[string]int{
		"usia": 27,
		"maks": 60,
	}

	fmt.Println(students)
	fmt.Println(profile)

	var coba = [...]map[string]int{{"usia": 10}, {"usia": 20}}
	fmt.Println(coba)

	/*
		Variabel map bisa diinisialisasi dengan tanpa nilai awal,
		caranya cukup menggunakan tanda kurung kurawal, contoh: map[string]int{} .
		Atau bisa juga dengan menggunakan keyword make dan new .
	*/
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	/*
		Khusus inisialisasi data menggunakan keyword new , yang dihasilkan adalah data pointer.
		Untuk mengambil nilai aslinya bisa dengan menggunakan tanda asterisk ( * ).
	*/

	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

	/*
		Item variabel map bisa di iterasi menggunakan for - range .
		Cara penerapannya masih sama seperti pada slice,
		pembedanya data yang dikembalikan di tiap perulangan
		adalah key dan value, bukan indeks dan elemen.
	*/

	var chicken = map[string]int{
		"januari":  100,
		"februari": 50,
		"maret":    150,
		"april":    200,
		"mei":      125,
	}

	for key, value := range chicken {
		fmt.Println(key, ":", value)
	}

	/*
		menghapus item map
		delete()
		digunakan untuk menghapus item dengan key tertentu pada variabel map.
	*/

	delete(chicken, "mei")
	fmt.Println(chicken)

	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exists")
	}

}
