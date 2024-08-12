package main //nama package harus main gak tw kenapa belum dapat alasan nya

import "fmt" // import package sama lah kek di javascript gunanya untuk dapat memanggil fungsi" yang terdapat didalam package

func main() {
    // ! ||--------------------------------------------------------------------------------||
    // ! ||                              Pembelajaran Pertama                              ||
    // ! ||--------------------------------------------------------------------------------||
    
    /*
     *   If else dengan variabel temporary
     */ 
	
	// membuat variabel dengan nama point dan value integer
	var point = 8840.0 
	
	// if kondisi dimana menggunakan variabel temporary, dimana variabel temporary langsung 
	// dipassing/diteruskan ke pengecekan if else ini lah salah satu pembeda dengan bahasa pemograman lain
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
	
	// ! ||--------------------------------------------------------------------------------||
	// ! ||                               Pembelajaran Kedua                               ||
	// ! ||--------------------------------------------------------------------------------||
	
	/*
     *  Switch case
	 *
	 *	Switch case di golang agak beda dengan bahasa pemograman lain dimana dia bisa di lakukan kondisi layak nya if bahkan bisa multi kondisi
	 *  seperti operator > < ==  && dan juga ada kita tidak perlu melakukan break untuk memberhentikan program seperti di javascript
	 *  akan tetapi ada sintaks fallthrough dimana sintaks ini berfungsi untuk tidak memberhentikan kode ketika kondisi terpenuhi
	 *  jadi program akan tetap jalan seperti swithcase di javascript tanpa sintaks break; didalam nya
     */ 
	
	var point1 = 6
	switch {
	case point1 == 8:
		fmt.Println("perfect")
	case (point1 < 8) && (point1 > 3):
		fmt.Println("awesome")
		
		// sintaks fallthrough bertujuan agar kode tidak di berhentikan (tidak di 'break') kondisi terpenuhi
		fallthrough 
	case point1 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("testing")
		}
	}
}