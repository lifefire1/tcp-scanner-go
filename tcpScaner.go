package main

import(
	"net"
	"fmt"
	"sync"
)

func ScanPort(port int, wg *sync.WaitGroup){
	defer wg.Done()
	IP := "write_ip_or_domen_name"
	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	fmt.Printf("%d is open\n", port)
	connection.Close()
}

func main(){

	var wg sync.WaitGroup
	for i := 1; i < 1000; i++{
		wg.Add(1)
		go ScanPort(i, &wg);
	} 

	wg.Wait()
}
