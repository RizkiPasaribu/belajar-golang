package main //nama package harus main gak tw kenapa belum dapat alasan nya

import "fmt" // import package sama lah kek di javascript gunanya untuk dapat memanggil fungsi" yang terdapat didalam package

func main() {
    // ! ||--------------------------------------------------------------------------------||
    // ! ||                              Pembelajaran Pertama                              ||
    // ! ||--------------------------------------------------------------------------------||
    
    /*
     *   Cara assign variabel
     */ 
    
    var firstName string = "Rizki" //lansung assign value kedalam variabel yang tipe data nya string varibel
    var lastName string // mendeklarasikan varibel dengan tipe data string tanpa assign value kedalam nya
    lastName = "Pasaribu" // assign value Pasaribu lastname
    fmt.Printf("Halo %s %s!\n", firstName, lastName) //memanggil fungsi Printf yang terdapat dalam package fmt
    
    
    // ! ||--------------------------------------------------------------------------------||
    // ! ||                               Pembelajaran Kedua                               ||
    // ! ||--------------------------------------------------------------------------------||
    
    /*
     *   Assigment variabel tanpa tipe data
     */
    name := "wick" // deklarasikan varibel langsung tanpa tipe data, si golang bakalan tw bahwa ini tipe data nya string karna nilai yang di berikan adalah string
    fmt.Printf("Halo %s!\n", name) //memanggil fungsi Printf yang terdapat dalam package fmt
    
    // ! ||--------------------------------------------------------------------------------||
    // ! ||                               Pembelajaran Ketiga                              ||
    // ! ||--------------------------------------------------------------------------------||
    
    /*
     *   Assigment variabel underscore, variabel yang berguna untuk membuang nilai yang tidak digunakan, seperti tempat sampah
     *   inti nya apapun nilai yang dimasukkan kedalam variabel tersebut tidak akan bisa di panggil ataupun digunakan akan menghilang sepenuh nya
     */ 
    _ = "belajar Golang"
    _ = "Golang itu mudah"
}