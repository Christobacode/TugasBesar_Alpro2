package main

import (
	"fmt"
	"strings"
)

// Struktur Data
type SparePart struct {
	Nama    string
	Harga   int
	Stock   int
	terjual int
}

type Transaksi struct {
	NamaPelanggan string
	Waktu         string
	SpareParts    [10]SparePart
	JumlahBeli    int
	BiayaService  int
	Total         int
}

var spareParts [11]SparePart = [11]SparePart{
	{"KampasRem", 30000, 10, 3},
	{"Aki", 200000, 5, 4},
	{"Busi", 15000, 20, 10},
	{"OliMesin", 45000, 10, 7},
	{"FilterUdara", 40000, 7, 2},
	{"BearingRoda", 40000, 12, 3},
	{"Shockbreaker", 200000, 6, 2},
	{"KampasKopling", 125000, 4, 5},
	{"Piston", 200000, 5, 2},
	{"Stator", 200000, 6, 1},
}

var transaksiList [15]Transaksi
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
		fmt.Println("==============================================================")
		fmt.Println("=                        Menu Admin                          =")
		fmt.Println("==============================================================")
		fmt.Println("1. Tambahkan Spare-part")
		fmt.Println("2. Ubah Spare-part")
		fmt.Println("3. Hapus Spare-part")
		fmt.Println("4. Daftar Spare-part")
		fmt.Println("5. Kembali ke menu login")
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
		fmt.Println("==============================================================")
		fmt.Println("=                       Menu Pengguna                        =")
		fmt.Println("==============================================================")
		fmt.Println("1. Transaksi")
		fmt.Println("2. Pencarian Pelanggan")
		fmt.Println("3. Daftar Spare-part Terurut")
		fmt.Println("4. Kembali ke menu login")
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

// tambah sparepart
func tambahkanSparePart() {
	var nama string
	var harga, stock, terjual int
	fmt.Print("Nama Spare-part: ")
	fmt.Scan(&nama)
	fmt.Print("Harga: ")
	fmt.Scan(&harga)
	fmt.Print("Stok: ")
	fmt.Scan(&stock)
	fmt.Print("terjual: ")
	fmt.Scan(&terjual)
	for i := 0; i < len(spareParts); i++ {
		if spareParts[i].Nama == "" {
			spareParts[i] = SparePart{nama, harga, stock, terjual}
			fmt.Println("Spare-part berhasil ditambahkan.")
		}
	}
	fmt.Println("Array penuh! tidak bisa menambahkan lagi.")
}

// ubah sparepart
func ubahSparePart() {
	var index int
	fmt.Print("Pilih nomor Spare-part yang ingin diubah: ")
	fmt.Scan(&index)

	if index > 0 && index <= len(spareParts) {
		var nama string
		var harga, stock, terjual int
		fmt.Print("Nama Baru: ")
		fmt.Scan(&nama)
		fmt.Print("Harga Baru: ")
		fmt.Scan(&harga)
		fmt.Print("Stok Baru: ")
		fmt.Scan(&stock)
		fmt.Print("Terjual Baru: ")
		fmt.Scan(&terjual)

		spareParts[index-1] = SparePart{nama, harga, stock, terjual}
		fmt.Println("Spare-part berhasil diubah.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

// hapus sparepart
func hapusSparePart() {
	var index int
	fmt.Print("Pilih nomor Spare-part yang ingin dihapus : ")
	fmt.Scan(&index)

	if index > 0 && index <= len(spareParts) {
		spareParts[index-1] = SparePart{} // Empty the slot
		fmt.Println("Spare-part berhasil dihapus.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

// daftar sparepart
func daftarSparePart() {
	for i, sp := range spareParts {
		if sp.Nama != "" {
			fmt.Printf("%d. Nama: %s, Harga: Rp%d, Stok: %d\n", i+1, sp.Nama, sp.Harga, sp.Stock)
		}
	}
}

// Fungsi Pencarian Pelanggan
func cariPelanggan() {
	for {
		fmt.Println("\nMenu Pencarian Pelanggan:")
		fmt.Println("1. Cari pelanggan berdasarkan waktu transaksi")
		fmt.Println("2. Cari pelanggan berdasarkan spare-part yang dibeli")
		fmt.Println("3. Kembali ke menu sebelumnya")
		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			cariBerdasarkanWaktu()
		case 2:
			cariBerdasarkanSparePart()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi Sequential Search berdasarkan waktu transaksi
func cariBerdasarkanWaktu() {
	var waktu string
	fmt.Print("Masukkan waktu transaksi (YYYY-MM-DD): ")
	fmt.Scan(&waktu)

	ketemu := false
	fmt.Printf("\nDaftar Pelanggan yang melakukan transaksi pada waktu %s:\n", waktu)
	for i := 0; i < len(transaksiList); i++ {
		if transaksiList[i].Waktu == waktu {
			fmt.Printf("- Nama: %s, Waktu: %s, Total: Rp%d\n", transaksiList[i].NamaPelanggan, transaksiList[i].Waktu, transaksiList[i].Total)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada transaksi yang ditemukan pada waktu tersebut.")
	}
}

// Fungsi Sequential Search berdasarkan spare-part
func cariBerdasarkanSparePart() {
	var namaSparePart string
	fmt.Print("Masukkan nama spare-part: ")
	fmt.Scan(&namaSparePart)

	ketemu := false
	totalSparePart := 0
	fmt.Printf("\nDaftar Pelanggan yang membeli spare-part %s:\n", namaSparePart)
	for i := 0; i < len(transaksiList); i++ {
		for j := 0; j < len(transaksiList[i].SpareParts); j++ {
			if strings.EqualFold(transaksiList[i].SpareParts[j].Nama, namaSparePart) {
				totalSparePart += transaksiList[i].SpareParts[j].Harga
				fmt.Printf("- Nama: %s, Waktu: %s, Harga Spare-Part: Rp%d, Total Transaksi: Rp%d\n", transaksiList[i].NamaPelanggan,
					transaksiList[i].Waktu, totalSparePart, transaksiList[i].Total)
				ketemu = true
				break
			}
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada transaksi yang ditemukan dengan spare-part tersebut.")
	}
}

// daftar data yang diurutkan sesuai penjualan terbanyak
func daftarSparePartTerurut() {
	// Menggunakan algoritma Selection Sort untuk mengurutkan spare-part berdasarkan jumlah terjual (terjual) secara menurun (descending)
	sparePartsTerurut := spareParts
	var i int
	var j int
	var maxIndex int
	for i = 0; i < len(sparePartsTerurut)-1; i++ {
		maxIndex = i
		for j = i + 1; j < len(sparePartsTerurut); j++ {
			if sparePartsTerurut[j].terjual > sparePartsTerurut[maxIndex].terjual {
				maxIndex = j
			}
		}
		sparePartsTerurut[i], sparePartsTerurut[maxIndex] = sparePartsTerurut[maxIndex], sparePartsTerurut[i]
	}
	for i, sp := range sparePartsTerurut {
		if sp.Nama != "" {
			fmt.Printf("%d. Nama: %s, Terjual: %d, Harga: Rp%d, Stok: %d\n", i+1, sp.Nama, sp.terjual, sp.Harga, sp.Stock)
		}
	}
}

// Fungsi untuk melakukan transaksi
func transaksi() {
	if jumlahTransaksi > 15 {
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

	var daftarBeli [10]SparePart
	var indexBeli int = 0

	for {
		daftarSparePart()
		fmt.Println("==============================================================")
		fmt.Print("Pilih Spare-part untuk dibeli (0 untuk selesai): ")
		fmt.Scan(&pilihan)
		fmt.Println("==============================================================")

		if pilihan == 0 {
			break
		}

		if pilihan > 0 && pilihan <= len(spareParts) {
			if spareParts[pilihan-1].Stock > 0 {
				spareParts[pilihan-1].Stock--   // Kurangi stok
				spareParts[pilihan-1].terjual++ // Tambah jumlah terjual
				daftarBeli[indexBeli] = spareParts[pilihan-1]
				indexBeli++
				total += spareParts[pilihan-1].Harga
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
		biayaService = 15000 * indexBeli
		total += biayaService
	}

	transaksiList[jumlahTransaksi] = Transaksi{
		NamaPelanggan: namaPelanggan,
		Waktu:         waktu,
		SpareParts:    daftarBeli,
		JumlahBeli:    indexBeli,
		BiayaService:  biayaService,
		Total:         total,
	}
	jumlahTransaksi++

	fmt.Printf("Total Pembayaran: Rp%d\n", total)
}

// Fungsi Utama
func main() {
	for {
		fmt.Println("==============================================================")
		fmt.Println("=                         Menu Login                         =")
		fmt.Println("==============================================================")
		fmt.Println("1. Admin")
		fmt.Println("2. Pengguna")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("==============================================================")
			fmt.Println("=                    Login Sebagai Admin                     =")
			fmt.Println("==============================================================")
			role := login()
			if role == "admin" {
				menuAdmin()
			} else {
				fmt.Println("Login gagal. Username atau password salah.")
			}
		case 2:
			fmt.Println("==============================================================")
			fmt.Println("=                  Login Sebagai Pengguna                    =")
			fmt.Println("==============================================================")
			role := login()
			if role == "pengguna" {
				menuPengguna()
			} else {
				fmt.Println("Login gagal. Username atau password salah.")
			}
		case 3:
			fmt.Println("\nTerima kasih telah menggunakan sistem ini. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
