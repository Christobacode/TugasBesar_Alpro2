package main

import "fmt"

// Struktur Data
type SparePart struct {
	Nama  string
	Harga int
	Stock int
}

type Transaksi struct {
	NamaPelanggan string
	Waktu         string
	SpareParts    []SparePart 
	JumlahBeli    int
	BiayaService  int
	Total         int
}

var spareParts []SparePart = []SparePart{
	{"Kampas Rem", 30000, 10},
	{"Aki", 200000, 5},
	{"Busi", 15000, 20},
	{"Oli Mesin", 45000, 10},
	{"Filter Udara", 40000, 7},
	{"Bearing Roda", 40000, 12},
	{"Shockbreaker", 200000, 6},
	{"Kampas Kopling", 125000, 4},
	{"Piston", 200000, 5},
	{"Stator", 200000, 6},
	
}

var transaksiList []Transaksi
var jumlahTransaksi int = 0

// Fungsi Login
func login() string {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	if username == "admin" && password == "admin123" {
		return "admin"
	} else if username == "pengguna" && password == "pengguna123" {
		return "pengguna"
	}
	return ""
}

// Menu Admin
func menuAdmin() {
	for {
		fmt.Println("\nMenu Admin:")
		fmt.Println("1. Tambahkan Spare-part")
		fmt.Println("2. Ubah Spare-part")
		fmt.Println("3. Hapus Spare-part")
		fmt.Println("4. Daftar Spare-part")
		fmt.Println("5. Keluar")
		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("==============================================================")
			fmt.Println("=                  Menu Tambah Spare-part                    =")
			fmt.Println("==============================================================")
			tambahkanSparePart()
		case 2:
			fmt.Println("==============================================================")
			fmt.Println("=                    Menu Ubah Spare-part                    =")
			fmt.Println("==============================================================")
			ubahSparePart()
		case 3:
			fmt.Println("==============================================================")
			fmt.Println("=                    Menu Hapus Spare-part                   =")
			fmt.Println("==============================================================")
			hapusSparePart()
		case 4:
			fmt.Println("==============================================================")
			fmt.Println("=                   Menu Daftar Spare-part                   =")
			fmt.Println("==============================================================")
			daftarSparePart()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Menu Pengguna
func menuPengguna() {
	for {
		fmt.Println("\nMenu Pengguna:")
		fmt.Println("1. Transaksi")
		fmt.Println("2. Pencarian Pelanggan")
		fmt.Println("3. Daftar Spare-part Terurut")
		fmt.Println("4. Keluar")
		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("==============================================================")
			fmt.Println("=                        Menu Transaksi                      =")
			fmt.Println("==============================================================")
			transaksi()
		case 2:
			fmt.Println("==============================================================")
			fmt.Println("=              Cari Pelanggan(sesuai waktu)                  =")
			fmt.Println("==============================================================")
			cariPelanggan()
		case 3:
			fmt.Println("==============================================================")
			fmt.Println("=          Daftar Sparepart(paling sering diganti)           =")
			fmt.Println("==============================================================")
			daftarSparePartTerurut()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Fungsi Admin
func tambahkanSparePart() {
	var nama string
	var harga, stock int
	fmt.Print("Nama Spare-part: ")
	fmt.Scan(&nama)
	fmt.Print("Harga: ")
	fmt.Scan(&harga)
	fmt.Print("Stok: ")
	fmt.Scan(&stock)

	spareParts = append(spareParts, SparePart{nama, harga, stock})

	fmt.Println("Spare-part berhasil ditambahkan.")
}

func ubahSparePart() {
	var index int
	fmt.Print("Pilih nomor Spare-part yang ingin diubah (1 - ", len(spareParts), "): ")
	fmt.Scan(&index)

	if index > 0 && index <= len(spareParts) {
		var nama string
		var harga, stock int
		fmt.Print("Nama Baru: ")
		fmt.Scan(&nama)
		fmt.Print("Harga Baru: ")
		fmt.Scan(&harga)
		fmt.Print("Stok Baru: ")
		fmt.Scan(&stock)

		spareParts[index-1] = SparePart{nama, harga, stock}
		fmt.Println("Spare-part berhasil diubah.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusSparePart() {
	var index int
	fmt.Print("Pilih nomor Spare-part yang ingin dihapus (1 - ", len(spareParts), "): ")
	fmt.Scan(&index)

	if index > 0 && index <= len(spareParts) {
		spareParts = append(spareParts[:index-1], spareParts[index:]...) // Menghapus spare-part
		fmt.Println("Spare-part berhasil dihapus.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func daftarSparePart() {
	fmt.Println("\nDaftar Spare-part:")
	for i, sp := range spareParts {
		fmt.Printf("%d. Nama: %s, Harga: Rp%d, Stok: %d\n", i+1, sp.Nama, sp.Harga, sp.Stock)
	}
}

// Fungsi untuk mencari pelanggan (belum diimplementasikan)
func cariPelanggan() {
	fmt.Println("Fungsi pencarian pelanggan belum diimplementasikan.")
}

// Fungsi untuk menampilkan daftar spare-part terurut
func daftarSparePartTerurut() {
	fmt.Println("\nDaftar Spare-part Terurut:")
	for i := 0; i < len(spareParts)-1; i++ {
		for j := 0; j < len(spareParts)-i-1; j++ {
			if spareParts[j].Harga > spareParts[j+1].Harga {
				spareParts[j], spareParts[j+1] = spareParts[j+1], spareParts[j]
			}
		}
	}
	daftarSparePart()
}

// Fungsi untuk melakukan transaksi
func transaksi() {
	if jumlahTransaksi >= 10 {
		fmt.Println("Kapasitas Transaksi penuh.")
		return
	}

	var namaPelanggan, waktu string
	var pilihan int
	var total int

	fmt.Print("Nama Pelanggan: ")
	fmt.Scan(&namaPelanggan)
	fmt.Print("Waktu Transaksi: ")
	fmt.Scan(&waktu)

	var daftarBeli []SparePart

	for {
		daftarSparePart()
		fmt.Println("==============================================================")
		fmt.Print("Pilih Spare-part untuk dibeli (0 untuk selesai): " )
		fmt.Scan(&pilihan)
		fmt.Println("==============================================================")

		if pilihan == 0 {
			break
		}

		if pilihan > 0 && pilihan <= len(spareParts) {
			if spareParts[pilihan-1].Stock > 0 {
				daftarBeli = append(daftarBeli, spareParts[pilihan-1])
				total += spareParts[pilihan-1].Harga
				spareParts[pilihan-1].Stock-- // Kurangi stok
			} else {
				fmt.Println("Stok tidak cukup untuk spare-part ini.")
			}
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}

	var biayaService int
	fmt.Print("Apakah ingin melakukan service? (1 untuk Ya, 0 untuk Tidak): ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		biayaService = 15000 * len(daftarBeli)
		total += biayaService
	}

	transaksiList = append(transaksiList, Transaksi{
		NamaPelanggan: namaPelanggan,
		Waktu:         waktu,
		SpareParts:    daftarBeli,
		JumlahBeli:    len(daftarBeli),
		BiayaService:  biayaService,
		Total:         total,
	})
	jumlahTransaksi++

	fmt.Printf("Total Pembayaran: Rp%d\n", total)
}

// Fungsi Utama
func main() {
	role := login()
	if role == "" {
		fmt.Println("Login gagal.")
		return
	}

	if role == "admin" {
		menuAdmin()
	} else if role == "pengguna" {
		menuPengguna()
	}
}