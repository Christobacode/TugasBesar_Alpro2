package main

import "fmt"

type Akun struct {
	user     string
	password string
	role     string
}

var jenisAkun = []Akun{
	{"admin", "admin123", "admin"},
	{"pengguna", "pengguna123", "pengguna"},
}

func login() (Akun, bool) {
	var username, password string
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)

	for _, akun := range jenisAkun {
		if akun.user == username && akun.password == password {
			return akun, true
		}
	}

	return Akun{}, false
}
func menuAdmin() {
	fmt.Println("Menu Admin")
	fmt.Println("1. Menambahkan Spare-part")
	fmt.Println("2. Mengubah Spare-part")
	fmt.Println("3. Menghapus Spare-part")
	fmt.Println("4. Tampilkan Daftar Spare-part")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih Menu: ")
}

func menuPengguna() {
	fmt.Println("Menu Pengguna")
	fmt.Println("1. Transaksi")
	fmt.Println("2. Pencarian Spare-part")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih Menu: ")
}

func tambahkanSparePart() {
	var nama string
	var harga int
	fmt.Print("Masukkan nama Spare-part: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan harga Spare-part: ")
	fmt.Scan(&harga)

}

func mengubahSparePart() {

}
func menghapusSparePart() {

}
func transaksi() {

}
func cariSparePart() {

}
func daftarSparePart() {
}

func main() {
	fmt.Println("Selamat datang! Silakan login terlebih dahulu.")

	akun, berhasil := login()
	if !berhasil {
		fmt.Println("Login gagal. Username atau password salah.")
		return
	}

	fmt.Println("Login berhasil! Selamat datang", "", akun.role)

	var pilih int
	if akun.role == "admin" {
		for {
			menuAdmin()
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				tambahkanSparePart()
			case 2:
				mengubahSparePart()
			case 3:
				menghapusSparePart()
			case 4:
				daftarSparePart()
			case 5:
				fmt.Println("Keluar dari sistem. Terima kasih!")
				return
			default:
				fmt.Println("Menu yang Anda pilih tidak tersedia.")
			}
			fmt.Println()
		}
	} else if akun.role == "pengguna" {
		for {
			menuPengguna()
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				transaksi()
			case 2:
				cariSparePart()
			case 3:
				fmt.Println("Keluar dari sistem. Terima kasih!")
				return
			default:
				fmt.Println("Menu yang Anda pilih tidak tersedia.")
			}
			fmt.Println()
		}
	}
}
