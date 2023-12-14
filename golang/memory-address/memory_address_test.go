package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

var globalVariable int

func TestSimpleMemoryAddress(t *testing.T) {
	var localVariable0 int
	var (
		localVariable1 int
		localVariable2 int
	)
	var localArray [4]int
	var localSlice []int
	var localSlice2 []int = []int{1, 2, 3, 4}
	var localSlice3 []float64 = []float64{1, 2, 3, 4}

	fmt.Printf("globalVariable:\t%p\tType:\t%T\tSize:\t%d\n", &globalVariable, globalVariable, unsafe.Sizeof(globalVariable))
	fmt.Printf("localVariable0:\t%p\tType:\t%T\tSize:\t%d\n", &localVariable0, localVariable0, unsafe.Sizeof(localVariable0))
	fmt.Printf("localVariable1:\t%p\tType:\t%T\tSize:\t%d\n", &localVariable1, localVariable1, unsafe.Sizeof(localVariable1))
	fmt.Printf("localVariable2:\t%p\tType:\t%T\tSize:\t%d\n", &localVariable2, localVariable2, unsafe.Sizeof(localVariable2))
	fmt.Printf("localArray:\t%p\tType:\t%T\tSize:\t%d\n", &localArray, localArray, unsafe.Sizeof(localArray))
	fmt.Printf("localSlice:\t%p\tType:\t%T\tSize:\t%d\n", &localSlice, localSlice, unsafe.Sizeof(localSlice))
	fmt.Printf("localSlice2:\t%p\tType:\t%T\tSize:\t%d\n", &localSlice2, localSlice2, unsafe.Sizeof(localSlice2))
	fmt.Printf("localSlice3:\t%p\tType:\t%T\tSize:\t%d\n", &localSlice3, localSlice3, unsafe.Sizeof(localSlice3))

	/*
		globalVariable:	0x6a2b48		Type:	int			Size:	8
		localVariable0:	0xc0001ca1f0	Type:	int			Size:	8
		localVariable1:	0xc0001ca1f8	Type:	int			Size:	8
		localVariable2:	0xc0001ca200	Type:	int			Size:	8
		localArray:		0xc0001e2060	Type:	[4]int		Size:	32
		localSlice:		0xc0001bc090	Type:	[]int		Size:	24
		localSlice2:	0xc0001bc0a8	Type:	[]int		Size:	24
		localSlice3:	0xc0001bc0c0	Type:	[]float64	Size:	24
	*/
}

func TestStructMemoryAddress(t *testing.T) {
	type StructCase1 struct {
		a bool
		b string
		c bool
	}
	type StructCase2 struct {
		a  bool
		b  bool
		c1 bool
		c  string
	}

	var case1 StructCase1
	var case2 StructCase2

	fmt.Printf("case1  :\t%p\tType:\t%T\tSize:\t%d\n", &case1, case1, unsafe.Sizeof(case1))
	fmt.Printf("case1.a:\t%p\tType:\t%T\tSize:\t%d\n", &case1.a, case1.a, unsafe.Sizeof(case1.a))
	fmt.Printf("case1.b:\t%p\tType:\t%T\tSize:\t%d\n", &case1.b, case1.b, unsafe.Sizeof(case1.b))
	fmt.Printf("case1.c:\t%p\tType:\t%T\tSize:\t%d\n", &case1.c, case1.c, unsafe.Sizeof(case1.c))
	fmt.Printf("case2  :\t%p\tType:\t%T\tSize:\t%d\n", &case2, case2, unsafe.Sizeof(case2))
	fmt.Printf("case2.a:\t%p\tType:\t%T\tSize:\t%d\n", &case2.a, case2.a, unsafe.Sizeof(case2.a))
	fmt.Printf("case2.b:\t%p\tType:\t%T\tSize:\t%d\n", &case2.b, case2.b, unsafe.Sizeof(case2.b))
	fmt.Printf("case2.b:\t%p\tType:\t%T\tSize:\t%d\n", &case2.c1, case2.c1, unsafe.Sizeof(case2.c1))
	fmt.Printf("case2.c:\t%p\tType:\t%T\tSize:\t%d\n", &case2.c, case2.c, unsafe.Sizeof(case2.c))

	/*
		case1  :	0xc000098060	Type:	main.StructCase1	Size:	32	// bool(1) + padding(7) + string(16) + bool(1) + padding(7)
		case1.a:	0xc000098060	Type:	bool				Size:	1
		case1.b:	0xc000098068	Type:	string				Size:	16 	// case1.a address +1 +7(padding)
		case1.c:	0xc000098078	Type:	bool				Size:	1	// case1.b address +10

		case2  :	0xc0000100a8	Type:	main.StructCase2	Size:	24  // bool(1) + bool(1) + padding(6) + string(16)
		case2.a:	0xc0000100a8	Type:	bool				Size:	1
		case2.b:	0xc0000100a9	Type:	bool				Size:	1	// case2.a address +1
		case2.c:	0xc0000100b0	Type:	string				Size:	16	// case2.b address +7
	*/
}

func TestArrayAndSliceMemoryAddress(t *testing.T) {
	var localArray [4]int
	var localSlice []int = []int{1, 2, 3, 4}
	var localArray2 [4]int

	fmt.Printf("localArray:\t%p\tType:\t%T\tSize:\t%d\n", &localArray, localArray, unsafe.Sizeof(localArray))
	for i := 0; i < 4; i++ {
		fmt.Printf("localArray[i]:\t%p\tType:\t%T\tSize:\t%d\n", &localArray[i], localArray[i], unsafe.Sizeof(localArray[i]))
	}
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("localSlice:\t%p\tType:\t%T\tSize:\t%d\n", &localSlice, localSlice, unsafe.Sizeof(localSlice))
	for i := 0; i < 4; i++ {
		fmt.Printf("localSlice[i]:\t%p\tType:\t%T\tSize:\t%d\n", &localSlice[i], localSlice[i], unsafe.Sizeof(localSlice[i]))
	}
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("localArray2:\t%p\tType:\t%T\tSize:\t%d\n", &localArray2, localArray2, unsafe.Sizeof(localArray2))
	for i := 0; i < 4; i++ {
		fmt.Printf("localArray2[i]:\t%p\tType:\t%T\tSize:\t%d\n", &localArray2[i], localArray2[i], unsafe.Sizeof(localArray2[i]))
	}

	/*
		localArray:		0xc000018120	Type:	[4]int	Size:	32
		localArray[i]:	0xc000018120	Type:	int		Size:	8
		localArray[i]:	0xc000018128	Type:	int		Size:	8
		localArray[i]:	0xc000018130	Type:	int		Size:	8
		localArray[i]:	0xc000018138	Type:	int		Size:	8
		---------------------------------------------------------------------------
		localSlice:		0xc0000100a8	Type:	[]int	Size:	24  <- slice descriptor
		localSlice[i]:	0xc000018140	Type:	int		Size:	8
		localSlice[i]:	0xc000018148	Type:	int		Size:	8
		localSlice[i]:	0xc000018150	Type:	int		Size:	8
		localSlice[i]:	0xc000018158	Type:	int		Size:	8
		---------------------------------------------------------------------------
		localArray2:	0xc000018160	Type:	[4]int	Size:	32
		localArray2[i]:	0xc000018160	Type:	int		Size:	8
		localArray2[i]:	0xc000018168	Type:	int		Size:	8
		localArray2[i]:	0xc000018170	Type:	int		Size:	8
		localArray2[i]:	0xc000018178	Type:	int		Size:	8
	*/
}

func TestSliceDescriptorMemoryAddress(t *testing.T) {
	data := make([]int, 5, 10)

	// Slice Header에 직접접근
	sliceHeaderAddress := (*uintptr)(unsafe.Pointer(&data))
	sliceHeaderLength := (*int)(unsafe.Add(unsafe.Pointer(&data), unsafe.Sizeof(*sliceHeaderAddress)))
	sliceHeaderCapacity := (*int)(unsafe.Add(unsafe.Pointer(&data), unsafe.Sizeof(*sliceHeaderAddress)+unsafe.Sizeof(*sliceHeaderLength)))
	fmt.Println("sliceHeaderAddress:", sliceHeaderAddress)
	fmt.Println("sliceHeaderLength:", *sliceHeaderLength)
	fmt.Println("sliceHeaderCapacity:", *sliceHeaderCapacity)
	fmt.Println("---------------------------------------------------------------------------")

	// reflect.SliceHeader 사용
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	fmt.Printf("SliceDataAddress:%x \n", sliceHeader.Data)
	fmt.Println("SliceLength:", sliceHeader.Len)
	fmt.Println("SliceCapacity:", sliceHeader.Cap)
	fmt.Println("---------------------------------------------------------------------------")

	fmt.Printf("sliceHeader:%p\n", &data)
	fmt.Printf("sliceData:%p\n", &data[0])
}
