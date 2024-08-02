package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gosnmp/gosnmp"
)

func main() {
	// Настройка SNMP-клиента
	snmp := gosnmp.GoSNMP{
		Target:    "10.50.5.5", // IP-адрес камеры
		Port:      161,             // Порт SNMP
		Community: "public",        //community string
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	// Подключение к SNMP-серверу
	err := snmp.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to SNMP server: %v", err)
	} else {
		// ping := snmp.Ping()
		// fmt.Printf("Ping: %v\n", ping)
		fmt.Println("Connected to SNMP server")
	}
	defer snmp.Conn.Close()

	// Запрос состояния камеры
	oidState := "1.3.6.1.2.1.2.2.1.7.1" // OID для состояния камеры
	resultState, err := snmp.Get([]string{oidState})
	if err != nil {
		log.Fatalf("Failed to get camera state: %v", err)
	}
	if len(resultState.Variables) == 0 {
		log.Fatal("No camera state variable found")
	}
	cameraState := resultState.Variables[0].Value.(int)
	fmt.Printf("Camera State: %d\n", cameraState)

	// Запрос имени камеры
	oidName := "1.3.6.1.2.1.1.5.0" // OID для имени камеры
	resultName, err := snmp.Get([]string{oidName})
	if err != nil {
		log.Fatalf("Failed to get camera name: %v", err)
	}
	if len(resultName.Variables) == 0 {
		log.Fatal("No camera name variable found")
	}
	cameraName := resultName.Variables[0].Value.(string)
	fmt.Printf("Camera Name: %s\n", cameraName)
}
