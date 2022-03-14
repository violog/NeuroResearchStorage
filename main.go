package main

func main() {
	// 2,5 GiB was too much, 2,0 GiB is fine (2147479552 bytes)
	// When I picked file 1,8 GiB, the program used 4,2 GiB of RAM;
	// There was a major freeze when I tried debugging with 2,5 GiB:
	// with this formulae it should have used 5,83 GiB
	// Total RAM usage was 6,2 GiB and 7,83 GiB respectively
	// That's why file size limit was created (currently 1 GiB)
	title := "Тестовое название содержит разрешённые символы: Fd'їґ \", -. (8%)"
	t := createTransaction(title, "test_attachment.txt", 0)
	t.print()
}
