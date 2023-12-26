package main

import "conceptos/tipos_genericos"

func main() {
	//funciones.Greet("Jennyfer", "Salcedo")
	//name := "diego"
	//funciones.ToUpperCase(&name)
	//fmt.Println(name)
	//
	//result := funciones.Suma(2, 3)
	//fmt.Println(result)
	//
	//lower, upper := funciones.Convert("PrUeBA")
	//fmt.Println("Lower:", lower, "Upper:", upper)

	//nums := []int{2, 12, 23, 98, 21, 79}
	//result := funciones.Filter(nums, lessThanTwenty)
	//fmt.Println(result)

	//result := funciones.Sum(2)(3)
	//fmt.Println(result)
	//plus10 := funciones.Sum(10)
	//fmt.Println(plus10(2))
	//fmt.Println(plus10(4))
	//fmt.Println(plus10(1))
	//fmt.Println(plus10(5))

	//fmt.Println(funciones.SumWithManyParameters(2))
	//fmt.Println(funciones.SumWithManyParameters(2, 3))
	//fmt.Println(funciones.SumWithManyParameters(2, 3, 12))
	//fmt.Println(funciones.SumWithManyParameters(2, 3, 12, 1, 24))

	// funciones anónimas
	//func(name string) {
	//	fmt.Println("✋ Hello", name)
	//}("gophers")

	//manejo_errores.ErrorsFundamentals()
	//found, err := manejo_errores.Search("34")
	//
	//if errors.Is(err, manejo_errores.ErrNotFound) {
	//	fmt.Println("Pudimos controlar el error!")
	//	return
	//}
	//
	//if err != nil {
	//	fmt.Println("Search()", err)
	//}
	//fmt.Println(found)

	//manejo_errores.DeferBasic()
	//manejo_errores.Division(100, 10)
	//manejo_errores.Division(200, 25)
	//manejo_errores.Division(34, 0)
	//manejo_errores.Division(124, 8)

	//fn_no_genericas_any.NoGenerics()
	tipos_genericos.GenericTypesFundamentals()

}

//func greaterToTen(num int) bool {
//	return num > 10
//}
//
//func lessThanTwenty(num int) bool {
//	return num < 20
//}
