package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	// TODO
	fmt.Println("Hello World!")
	var x int // değişken tanımlama
	x = 5     // değişkene değer atama

	fmt.Print("x = ", x)

	y := 10 // değişken tanımlama ve değer atama. Go bunun int olduğunu anlıyor.
	fmt.Println("y = ", y)

	var a, b int = 1, 2 // birden fazla değişkene değer atama
	fmt.Println("a = ", a, "b = ", b)

	fmt.Println(reflect.TypeOf(x)) // değişkenin tipini öğrenme
	fmt.Println(reflect.TypeOf(y))

	name := "Golang" // string değişken tanımlama
	fmt.Println("name = ", name)

	var ch rune = 'a' // rune değişken tanımlama-ascii karakter. rune int32'ye eşdeğer yani unicode karşılığğını yazdı
	fmt.Println("ch = ", ch, " type = ", reflect.TypeOf(ch))

	// byte değişken tanımlama-ascii karakter
	var ch2 byte = 'a'
	fmt.Println("ch2 = ", ch2, " type = ", reflect.TypeOf(ch2))

	var f float32 = 3.14 // float değişken tanımlama
	fmt.Println("f = ", f, " type = ", reflect.TypeOf(f))

	var b2 bool = true // bool değişken tanımlama
	fmt.Println("b2 = ", b2, " type = ", reflect.TypeOf(b2))

	var resp []byte = []byte("Hello") // byte dizisi tanımlama. Go dilinde stringler aslında byte dizileridir.
	fmt.Println("resp = ", resp, " type = ", reflect.TypeOf(resp))
	// string'e dönüştürme

	fmt.Println("string(resp) = ", string(resp), " type = ", reflect.TypeOf(string(resp)))
	// byte'dan string'e string'den byte'a dönüşüm yaparken dikkat edilmesi gereken şeyler var.
	// string'den byte'a dönüşüm yaparken string'in içindeki karakterlerin ascii karşılıkları alınır.
	// byte'dan string'e dönüşüm yaparken de byte'ın ascii karşılığına göre karakterler oluşturulur.

	var num complex64 = 3 + 2i // complex değişken tanımlama
	fmt.Println("num = ", num, " type = ", reflect.TypeOf(num))
	fmt.Println(real(num)) // complex sayının reel kısmını verir
	fmt.Println(imag(num)) // complex sayının sanal kısmını verir

	// Go'da değişken tanımlarken var keyword'ü kullanılmaz. var keyword'ü sadece fonksiyon dışında kullanılır.
	// Go'da değişken tanımlarken := kullanılır. := değişken tanımlarken ve değer atarken kullanılır.

	var arr [3]int = [3]int{1, 3} // array tanımlama. array'ler sabit boyutludur.
	fmt.Println("arr = ", arr, " type = ", reflect.TypeOf(arr))
	fmt.Printf("len(arr) = %d cap(slice) = %d\n", len(arr), cap(arr)) // len(arr) array'ın uzunluğunu verir.
	var slice []int = []int{1, 2, 3, 4, 5}                            // slice tanımlama.Belli bir kapasite yok. slice'lar dinamik boyutludur.
	// UI dan veri alırken ya da veri tabanından veri çekerken slice kullanılır.
	fmt.Println("slice = ", slice, " type = ", reflect.TypeOf(slice))

	fmt.Printf("len(slice) = %d cap(slice) = %d\n", len(slice), cap(slice)) // len(slice) slice'ın uzunluğunu verir.
	//make fonksiyonu ile slice oluşturma
	slice2 := make([]int, 5, 10) // 5 elemanlı 10 kapasiteli slice oluşturur.
	fmt.Println("slice2 = ", slice2, " type = ", reflect.TypeOf(slice2))
	fmt.Printf("len(slice2) = %d cap(slice2) = %d\n", len(slice2), cap(slice2))
	slice2[0] = 1 // slice'a değer atama
	//len(slice2) slice'ın uzunluğunu verir.
	//cap(slice2) slice'ın kapasitesini verir.
	//slice'ın kapasitesi dolunca yeni bir dizi oluşturur ve eski değerleri yeni diziye kopyalar.
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7, 8, 9) // slice'a eleman ekleme
	fmt.Println("slice2 = ", slice2, " type = ", reflect.TypeOf(slice2))

	// slice'ın uzunluğu ve kapasitesi dolunca yeni bir dizi oluşturur ve eski değerleri yeni diziye kopyalar.
	sort.Ints(slice2) // slice'ı sıralama

	var m map[string]int = map[string]int{
		"one": 1,
		"two": 2, // map tanımlama
	} // map tanımlama
	fmt.Println("m = ", m, " type = ", reflect.TypeOf(m))
	n := make(map[string]int)
	n["one"] = 1
	fmt.Println("n = ", n["one"], " type = ", reflect.TypeOf(n))
	// var l map[string]int // neden tehlikeli olduğunu anlat. Initialize etmeden kullanırsan nil pointer exception verir.
	// map'ler key-value şeklinde çalışır. Key'ler unique olmalıdır.

	for k, v := range m {
		//fmt.Println(k, v)
		fmt.Printf("key = %s value = %d\n", k, v)
	}
	// map'lerde key yoksa default değer döner.

	for i := 0; i < len(slice)-1; i++ {
		fmt.Println(slice[i])
	}

	for _, v := range slice {
		fmt.Println(v)
	}

	// go da while döngüsü yoktur.
	// go da tanımlanan değikenlerin hepsi kullanılmalıdır. Kullanılmayan değişkenler hata verir.
	// go da switch case de break kullanmaya gerek yoktur. Otomatik olarak break yapar.
	// go da switch case de default case olmak zorundadır.

	// map den bir değer çıkarmak için delete fonksiyonu kullanılır.

	delete(m, "one")
	fmt.Println("m = ", m, " type = ", reflect.TypeOf(m))

	// Kendi tipler yazılabilir

	users := []Person{
		{name: "Ahmet", age: 20},
		{name: "Mehmet", age: 30},
		{name: "Ayşe", age: 40},
	}
	fmt.Println("users = ", users, " type = ", reflect.TypeOf(users))
	fmt.Println("len(users) = ", len(users), " type = ", reflect.TypeOf(len(users)))
	for _, v := range users {
		//fmt.Println(v.name, v.age)
		if v.age > 20 {
			fmt.Println("20 den buyuk kisiler ", v.name, " ", v.age, " yasinda")
		}
	}
	// slice'ı sıralama
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].age < users[j].age
	})
	fmt.Println("users = ", users, " type = ", reflect.TypeOf(users))

	// condition  tanımlama

	if len(users) > 0 {
		// TODO
		fmt.Println("users is not empty")
	} else if len(users) == 0 {
		// TODO
		fmt.Println("users is empty")
	} else {
		// TODO
		fmt.Println("users is nil")
	}

	// switch case tanımlama
	switch len(users) {
	case 0:
		// TODO
		fmt.Println("users is empty")
	case 1:
		// TODO
		fmt.Println("users has one element")
	default:
		// TODO
		fmt.Println("users has more than one element")
	}

	// uintpr ve uintptr kullanımı anlatılacak.
	// initialize go modular project : go init github.com/username/projectname
	// how we can cleanup downloaded module dependenceies packages : go mod tidy
	// Which data structures dynamic and which data structures static : array static, slice dynamic

}
// go da eş zamanlılık için goroutine kullanılır. Goroutine'ler thread'lerden daha hafiftir.
type Person struct {
	name string
	age  int
}
