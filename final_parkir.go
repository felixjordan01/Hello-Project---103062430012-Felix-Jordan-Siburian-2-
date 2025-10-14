package main

import (
	"fmt"
	"time"
)

const maxData = 100

type Officer struct {
	ID       string
	Username string
	Password string
	Role     string 
}

type Vehicle struct {
	PlateNumber string
	VehicleType string 
	EntryTime   time.Time
	ExitTime    time.Time
	Fee         float64
}

var officers [maxData]Officer
var officerCount int

var vehicles [maxData]Vehicle
var vehicleCount int

// Login
func login(username, password string) int {
	var i int
	for i = 0; i < officerCount; i++ {
		if officers[i].Username == username && officers[i].Password == password {
			return i
		}
	}
	return -1
}

// Admin functions
func addOfficer(id, username, password, role string) {
	if officerCount < maxData {
		officers[officerCount] = Officer{id, username, password, role}
		officerCount++
	}
}

func editOfficer(index int, username, password, role string) {
	if index >= 0 && index < officerCount {
		officers[index].Username = username
		officers[index].Password = password
		officers[index].Role = role
	}
}

func deleteOfficer(id string) {
	var i, found int
	found = -1
	for i = 0; i < officerCount; i++ {
		if officers[i].ID == id {
			found = i
		}
	}
	if found != -1 {
		for i = found; i < officerCount-1; i++ {
			officers[i] = officers[i+1]
		}
		officerCount--
	}
}

// Petugas functions
func addVehicle(plate, vehicleType string) {
	if vehicleCount < maxData {
		vehicles[vehicleCount] = Vehicle{
			PlateNumber: plate,
			VehicleType: vehicleType,
			EntryTime:   time.Now(),
		}
		vehicleCount++
	}
}

func calculateFee(entry, exit time.Time) float64 {
	hours := int(exit.Sub(entry).Hours())
	if exit.Sub(entry).Minutes() > float64(hours*60) {
		hours++
	}
	return float64(hours) * 5000
}

func checkoutVehicle(plate string) {
	var i int
	for i = 0; i < vehicleCount; i++ {
		if vehicles[i].PlateNumber == plate {
			vehicles[i].ExitTime = time.Now()
			vehicles[i].Fee = calculateFee(vehicles[i].EntryTime, vehicles[i].ExitTime)
		}
	}
}

func deleteVehicle(plate string) {
	var i, found int
	found = -1
	for i = 0; i < vehicleCount; i++ {
		if vehicles[i].PlateNumber == plate {
			found = i
		}
	}
	if found != -1 {
		for i = found; i < vehicleCount-1; i++ {
			vehicles[i] = vehicles[i+1]
		}
		vehicleCount--
	}
}

// Search
func sequentialSearch(plate string) int {
	var i int
	for i = 0; i < vehicleCount; i++ {
		if vehicles[i].PlateNumber == plate {
			return i
		}
	}
	return -1
}

func binarySearch(plate string) int {
	low := 0
	high := vehicleCount - 1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if vehicles[mid].PlateNumber == plate {
			return mid
		} else {
			if plate < vehicles[mid].PlateNumber {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// Sort
func selectionSortAscending() {
	var i, j, minIdx int
	for i = 0; i < vehicleCount-1; i++ {
		minIdx = i
		for j = i + 1; j < vehicleCount; j++ {
			if vehicles[j].PlateNumber < vehicles[minIdx].PlateNumber {
				minIdx = j
			}
		}
		vehicles[i], vehicles[minIdx] = vehicles[minIdx], vehicles[i]
	}
}

func insertionSortDescending() {
	var i, j int
	var key Vehicle
	for i = 1; i < vehicleCount; i++ {
		key = vehicles[i]
		j = i - 1
		for j >= 0 && vehicles[j].PlateNumber < key.PlateNumber {
			vehicles[j+1] = vehicles[j]
			j--
		}
		vehicles[j+1] = key
	}
}

func printVehicles() {
	var i int
	var total float64
	for i = 0; i < vehicleCount; i++ {
		fmt.Printf("%d. Plat: %s | Tipe: %s | Masuk: %s | Keluar: %s | Biaya: Rp. %.2f\n",
			i+1, vehicles[i].PlateNumber, vehicles[i].VehicleType,
			vehicles[i].EntryTime.Format("2006-01-02 15:04:05"),
			vehicles[i].ExitTime.Format("2006-01-02 15:04:05"),
			vehicles[i].Fee)
		total += vehicles[i].Fee
	}
	fmt.Printf("\nTotal Pendapatan: Rp. %.2f\n", total)
}

// Main menu
func main() {
	addOfficer("admin1", "admin", "admin123", "admin")
	addOfficer("ptgs1", "petugas", "petugas123", "petugas")

	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	index := login(username, password)
	if index == -1 {
		fmt.Println("Login gagal!")
		return
	}
	fmt.Println("Login berhasil!")
	if officers[index].Role == "admin" {
		adminMenu()
	} else {
		petugasMenu()
	}
}

func adminMenu() {
	var choice int
	var id, username, password, role string
	for {
		fmt.Println("\n[Admin Menu] 1.Tambah 2.Edit 3.Hapus 4.Exit")
		fmt.Print("Pilih: ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Print("ID: ")
			fmt.Scan(&id)
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)
			fmt.Print("Role: ")
			fmt.Scan(&role)
			addOfficer(id, username, password, role)
		} else if choice == 2 {
			fmt.Print("Index Edit: ")
			fmt.Scan(&choice)
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)
			fmt.Print("Role: ")
			fmt.Scan(&role)
			editOfficer(choice, username, password, role)
		} else if choice == 3 {
			fmt.Print("ID Hapus: ")
			fmt.Scan(&id)
			deleteOfficer(id)
		} else if choice == 4 {
			fmt.Println("Keluar dari admin menu.")
			return
		}
	}
}

func petugasMenu() {
	var choice int
	var plate, vehicleType string
	for {
		fmt.Println("\n[Petugas Menu] 1.Masuk 2.Keluar 3.Hapus 4.Cetak 5.Sort A-Z 6.Sort Z-A 7.Exit")
		fmt.Print("Pilih: ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Print("Plat: ")
			fmt.Scan(&plate)
			fmt.Print("Tipe [mobil/motor]: ")
			fmt.Scan(&vehicleType)
			addVehicle(plate, vehicleType)
		} else if choice == 2 {
			fmt.Print("Plat keluar: ")
			fmt.Scan(&plate)
			checkoutVehicle(plate)
		} else if choice == 3 {
			fmt.Print("Plat hapus: ")
			fmt.Scan(&plate)
			deleteVehicle(plate)
		} else if choice == 4 {
			printVehicles()
		} else if choice == 5 {
			selectionSortAscending()
		} else if choice == 6 {
			insertionSortDescending()
		} else if choice == 7 {
			fmt.Println("Keluar dari petugas menu.")
			return
		}
	}
}
