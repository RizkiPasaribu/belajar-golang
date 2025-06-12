package main //nama package harus main gak tw kenapa belum dapat alasan nya

import "fmt" // import package sama lah kek di javascript gunanya untuk dapat memanggil fungsi" yang terdapat didalam package

func main() {
    /* -------------------------------------------------------------------------- */
    /*                            Pembelajaran Pertama                            */
    /* -------------------------------------------------------------------------- */
	
	//Membuat array yang data nya bisa terisi sebanyak 4 data
	var names [4]string
	
	//Assign satu satu nilai ke 4 data yang telah di deklarasikan sebelum nya
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	
	fmt.Println(names[0])
	
	/* -------------------------------------------------------------------------- */
	/*                             Pembelajaran Kedua                             */
	/* -------------------------------------------------------------------------- */
	
	//Langsung init value kedalam data array yang panjang data nya 4 begitu array di deklarasikan
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])
	
	/* -------------------------------------------------------------------------- */
	/*                             Pembelajaran Ketiga                            */
	/* -------------------------------------------------------------------------- */
	
	//Deklarasi array yang panjang data nya otomatis menyesuaikan panjang nilai yang di assign kedalam array nya
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println(numbers[0])	
	
	/* -------------------------------------------------------------------------- */
	/*                            Pembelajaran Keempat                            */
	/* -------------------------------------------------------------------------- */
	
	// Array looping
	// Disini bisa pakai for loop tapi itu sudah terlalu biasa disini saya menggunakan range
	
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	
	// ketika tidak ingin menggunakan index (i) bisa diganti dengan variabel kosong (_) karna
	// digolang ketika variabel di deklarasikan variabel tersebut harus digunakan kalau tidak kode akan eror
	
	for _, fruit := range fruits {
		fmt.Printf("fruit : %s\n", fruit)
	}
	
	/* -------------------------------------------------------------------------- */
	/*                                Fungsi array                                */
	/* -------------------------------------------------------------------------- */
	// Slice	= Memotong array
	// Len 		= Length of array
	// Cap		= Capacity of array
	// append	= Append
	// Copy		= coppy array take 2 argument src(source) and dst (destination)
}