package main
import "fmt"

func main(){
	var nik int
	var tanggal, bulan, tahun int
	fmt.Scan(&nik)
	cariTTL(nik, &tanggal, &bulan, &tahun)
}
func cariTTL(nik int, tanggal, bulan, tahun *int){
	*tanggal = (nik/100000000)%100
	*bulan = (nik/1000000)%100
	*tahun = (nik/10000)%100
	fmt.Println(*tanggal, *bulan, *tahun)
}