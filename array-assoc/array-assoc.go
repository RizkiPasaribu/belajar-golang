package main //nama package harus main gak tw kenapa belum dapat alasan nya

import "fmt" // import package sama lah kek di javascript gunanya untuk dapat memanggil fungsi" yang terdapat didalam package

func main() {
	/* -------------------------------------------------------------------------- */
	/*                            Pembelajaran Pertama                            */
	/* -------------------------------------------------------------------------- */
	
	// biasa nya kita map di gunakan unttuk melooping array dengan tujuan untuk mentranformasikan
	// nilai yang ada di dalam array karna fungsi map mereturn array itu kembali akan tetapi
	// di golang beda map di gunakan untuk membuat array assosiatif, yaitu array yang key adalah
	// string
	
	// membuat array asosiatif
	var chicken map[string]int
	chicken = map[string]int{}
	
	chicken["januari"] = 50
	chicken["februari"] = 40
	
	fmt.Println(chicken)
	
	
	/* -------------------------------------------------------------------------- */
	/*                             Pembelajaran Kedua                             */
	/* -------------------------------------------------------------------------- */
	
	//membuat array assosiatif langsung mendefinisikan valuenya
	var chicken1 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	
	fmt.Println(chicken1)
	
	// dan item di map ini dia bisa dihapus dengan fungsi delete
	delete(chicken1, "januari")
	fmt.Println(chicken1)
	
}