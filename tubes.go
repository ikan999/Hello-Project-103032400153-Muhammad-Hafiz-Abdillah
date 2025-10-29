package main

import "fmt"

const NMAX = 5

type Produk struct {
	Nama  string
	Harga int
}

type ItemKeranjang struct {
	Produk Produk
	Jumlah int
}

type DaftarProduk [NMAX]Produk
type KeranjangBelanja [NMAX]ItemKeranjang

func main() {
	daftarProduk := DaftarProduk{
		{"Susu", 15000},
		{"Roti", 12000},
		{"Teh Botol", 8000},
		{"Indomie", 3500},
		{"Gula", 10000},
	}

	var keranjang KeranjangBelanja
	var jumlahItem int

	for {
		fmt.Println("\n=== APLIKASI KASIR MINIMARKET ===")
		fmt.Println("1. Tampilkan Menu & Harga (urutkan dan beli langsung)")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih menu (1-2): ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			jumlahItem = urutTampilTambahBayar(&daftarProduk, &keranjang, jumlahItem)
		case 2:
			fmt.Println("Terima kasih telah berbelanja!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func urutTampilTambahBayar(produk *DaftarProduk, keranjang *KeranjangBelanja, jumlahItem int) int {
	for {
		fmt.Println("\nPilih urutan tampilan harga:")
		fmt.Println("1. Ascending (Harga Terendah ke Tertinggi)")
		fmt.Println("2. Descending (Harga Tertinggi ke Terendah)")
		fmt.Print("Pilih (1/2): ")

		var urut int
		fmt.Scan(&urut)

		if urut == 1 {
			urutkanAscending(produk)
			break
		} else if urut == 2 {
			urutkanDescending(produk)
			break
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}

	tampilkanMenu(*produk)

	for {
		fmt.Println("\nMasukkan nama produk yang ingin dibeli (ketik '0' untuk selesai):")
		fmt.Print("Nama produk: ")
		var nama string
		fmt.Scan(&nama)

		if nama == "0" {
			break
		}

		var produkDitemukan Produk
		var ditemukan bool
		cariProdukByNama(*produk, nama, &produkDitemukan, &ditemukan)
		if !ditemukan {
			fmt.Println("Produk tidak ditemukan, coba lagi.")
			continue
		}

		var jumlah int
		fmt.Printf("Masukkan jumlah %s: ", produkDitemukan.Nama)
		fmt.Scan(&jumlah)
		if jumlah <= 0 {
			fmt.Println("Jumlah harus lebih dari 0, coba lagi.")
			continue
		}

		keranjang[jumlahItem] = ItemKeranjang{
			Produk: produkDitemukan,
			Jumlah: jumlah,
		}
		jumlahItem++
		fmt.Printf("%s sebanyak %d telah ditambahkan ke keranjang.\n", produkDitemukan.Nama, jumlah)
	}

	prosesPembayaran(*keranjang, jumlahItem)

	// Kosongkan keranjang setelah bayar
	for i := 0; i < jumlahItem; i++ {
		keranjang[i] = ItemKeranjang{}
	}
	jumlahItem = 0

	return jumlahItem
}

func urutkanAscending(produk *DaftarProduk) {
	for i := 0; i < NMAX-1; i++ {
		for j := 0; j < NMAX-1-i; j++ {
			if produk[j].Harga > produk[j+1].Harga {
				temp := produk[j]
				produk[j] = produk[j+1]
				produk[j+1] = temp
			}
		}
	}
}

func urutkanDescending(produk *DaftarProduk) {
	for i := 0; i < NMAX-1; i++ {
		for j := 0; j < NMAX-1-i; j++ {
			if produk[j].Harga < produk[j+1].Harga {
				temp := produk[j]
				produk[j] = produk[j+1]
				produk[j+1] = temp
			}
		}
	}
}

func tampilkanMenu(produk DaftarProduk) {
	fmt.Println("\n=== DAFTAR MENU & HARGA ===")
	fmt.Printf("%-15s | %-10s\n", "Produk", "Harga")
	fmt.Println("-------------------------------")
	for _, p := range produk {
		fmt.Printf("%-15s | Rp%8d\n", p.Nama, p.Harga)
	}
}

func cariProdukByNama(produk DaftarProduk, nama string, hasil *Produk, ditemukan *bool) {
	*ditemukan = false
	for i := 0; i < NMAX; i++ {
		if produk[i].Nama == nama {
			*hasil = produk[i]
			*ditemukan = true
			break
		}
	}
}

func prosesPembayaran(keranjang KeranjangBelanja, jumlahItem int) {
	if jumlahItem == 0 {
		fmt.Println("\nKeranjang kosong. Tidak ada transaksi.")
		return
	}

	total := 0
	fmt.Println("\n================= RINCIAN BELANJA ===================")
	fmt.Printf("%-15s | %-8s | %-6s | %s\n", "Produk", "Harga", "Jumlah", "Subtotal")
	fmt.Println("-------------------------------------------------------")

	for i := 0; i < jumlahItem; i++ {
		item := keranjang[i]
		subtotal := item.Produk.Harga * item.Jumlah
		fmt.Printf("%-15s | Rp%-7d | %-6d | Rp%d\n",
			item.Produk.Nama, item.Produk.Harga, item.Jumlah, subtotal)
		total += subtotal
	}

	diskon := 0
	if total > 150000 {
		diskon = total * 15 / 100
		fmt.Println("\nSelamat! Anda mendapatkan diskon 15%")
	} else if total > 100000 {
		diskon = total * 10 / 100
		fmt.Println("\nSelamat! Anda mendapatkan diskon 10%")
	} else if total > 50000 {
		diskon = total * 5 / 100
		fmt.Println("\nSelamat! Anda mendapatkan diskon 5%")
	} else {
		fmt.Println("\nBelanja lagi untuk mendapatkan diskon!")
	}

	totalSetelahDiskon := total - diskon

	fmt.Printf("\nTOTAL HARGA     : Rp%d\n", total)
	fmt.Printf("DISKON          : Rp%d\n", diskon)
	fmt.Printf("TOTAL BAYAR     : Rp%d\n", totalSetelahDiskon)

	for {
		var bayar int
		fmt.Print("Masukkan uang: Rp")
		fmt.Scan(&bayar)

		if bayar < totalSetelahDiskon {
			kekurangan := totalSetelahDiskon - bayar
			fmt.Printf("Maaf, uang Anda kurang Rp%d. Coba isi ulang.\n", kekurangan)
		} else {
			kembalian := bayar - totalSetelahDiskon
			fmt.Printf("Pembayaran berhasil. Kembalian Anda: Rp%d\n", kembalian)
			break
		}
	}

	fmt.Println("Terima kasih telah berbelanja!")
}
