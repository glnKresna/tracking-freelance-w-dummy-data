package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Proyek struct {
	ID            string
	Judul         string
	Klien         string
	Status        string
	TanggalTerima string
	Deadline      string
	Catatan       string
}

// Fungsi generate ID proyek
func generateID(proyek []Proyek) string {
	newID := len(proyek) + 1

	return fmt.Sprintf("%03d", newID)
}

// Fungsi pemotong string
func truncateString(str string, num int) string {
	if len(str) <= num {
		return str
	}
	return str[0:num-3] + "..."
}

// Fungsi format tanggal
func formatTanggal(input string) string {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) != 8 {
		return ""
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return ""
		}
	}

	hari := input[0:2]
	bulan := input[2:4]
	tahun := input[4:8]

	hariInt := int(hari[0]-'0')*10 + int(hari[1]-'0')
	bulanInt := int(bulan[0]-'0')*10 + int(bulan[1]-'0')

	if hariInt < 1 || hariInt > 31 {
		return ""
	}

	if bulanInt < 1 || bulanInt > 12 {
		return ""
	}

	return hari + "/" + bulan + "/" + tahun
}

// Fungsi input tanggal
func inputTanggal(reader *bufio.Reader, inputTanggal string) string {
	for {
		fmt.Print(inputTanggal)
		tanggal, _ := reader.ReadString('\n')
		tanggal = strings.TrimSpace(tanggal)

		tanggalTerformat := formatTanggal(tanggal)
		if tanggalTerformat != "" {
			return tanggalTerformat
		}
		fmt.Println("Format tanggal tidak valid! Gunakan format DD MM YYYY")
	}
}

// Fungsi nambah proyek baru
func tambahProyek(proyek *[]Proyek) {
	var judul, klien, status, catatan string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Judul Proyek: ")
	judul, _ = reader.ReadString('\n')
	judul = strings.TrimSpace(judul)

	fmt.Print("Nama Klien: ")
	klien, _ = reader.ReadString('\n')
	klien = strings.TrimSpace(klien)

	tanggalTerima := inputTanggal(reader, "Tanggal Terima (DD MM YYYY): ")
	deadline := inputTanggal(reader, "Deadline (DD MM YYYY): ")

	fmt.Print("Status Pengerjaan (1: Pending, 2: Ongoing, 3: Selesai): ")
	status, _ = reader.ReadString('\n')
	status = strings.TrimSpace(status)

	switch status {
	case "1":
		status = "Pending"
	case "2":
		status = "Ongoing"
	case "3":
		status = "Selesai"
	}

	fmt.Print("Tambah Catatan: ")
	catatan, _ = reader.ReadString('\n')
	catatan = strings.TrimSpace(catatan)

	proyekBaru := Proyek{
		ID:            generateID(*proyek),
		Judul:         judul,
		Klien:         klien,
		Status:        status,
		TanggalTerima: tanggalTerima,
		Deadline:      deadline,
		Catatan:       catatan,
	}

	*proyek = append(*proyek, proyekBaru)
	fmt.Println("Proyek baru berhasil ditambahkan.")
}

// Fungsi lihat proyek
func lihatProyek(proyekList []Proyek) {
	if len(proyekList) == 0 {
		fmt.Println("Belum ada proyek, tambahkan untuk melihat.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPilih urutan tampilan:")
	fmt.Println("1. Berdasarkan Judul Proyek")
	fmt.Println("2. Berdasarkan ID Proyek")
	fmt.Println("3. Berdasarkan Status ")
	fmt.Print("Pilihan (1-3): ")

	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.TrimSpace(pilihan)

	proyekTerurut := make([]Proyek, len(proyekList))
	copy(proyekTerurut, proyekList)

	switch pilihan {
	case "1":
		bubbleSortByJudul(&proyekTerurut)
		fmt.Println("\nDaftar Proyek (Diurutkan berdasarkan Judul menggunakan Bubble Sort)")
	case "2":
		insertionSortByID(&proyekTerurut)
		fmt.Println("\nDaftar Proyek (Diurutkan berdasarkan ID menggunakan Insertion Sort)")
	case "3":
		selectionSortByStatus(&proyekTerurut)
		fmt.Println("\nDaftar Proyek (Diurutkan berdasarkan Status menggunakan Selection Sort)")
	default:
		fmt.Println("Pilihan tidak valid! Menampilkan data tanpa pengurutan.")
		fmt.Println("\nDaftar Proyek")
		proyekTerurut = proyekList
	}

	fmt.Println("|======================================================================================================================|")
	fmt.Println("|   ID   |   Judul Proyek   |     Klien     |     Status     |   Tanggal Terima   |    Deadline    |      Catatan      |")
	fmt.Println("|======================================================================================================================|")
	for _, p := range proyekTerurut {
		fmt.Printf("| %-6s | %-16s | %-13s | %-14s | %-18s | %-14s | %-17s |\n",
			p.ID,
			truncateString(p.Judul, 16),
			truncateString(p.Klien, 13),
			truncateString(p.Status, 14),
			truncateString(p.TanggalTerima, 18),
			truncateString(p.Deadline, 14),
			truncateString(p.Catatan, 17))
	}
	fmt.Println("|======================================================================================================================|")
}

// Fungsi edit data proyek
func editProyek(proyekList *[]Proyek) {
	if len(*proyekList) == 0 {
		fmt.Println("Belum ada proyek yang bisa diedit.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var targetIndex int = -1
	var searchInput string

	fmt.Print("Masukkan ID proyek yang ingin diedit: ")
	searchInput, _ = reader.ReadString('\n')
	searchInput = strings.TrimSpace(searchInput)

	for i, p := range *proyekList {
		if p.ID == searchInput {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		fmt.Println("Proyek tidak ditemukan.")
		return
	}

	target := &(*proyekList)[targetIndex]
	fmt.Printf("\nData proyek yang akan diedit:")
	fmt.Printf("\nID: %s", target.ID)
	fmt.Printf("\nJudul: %s", target.Judul)
	fmt.Printf("\nKlien: %s", target.Klien)
	fmt.Printf("\nStatus: %s", target.Status)
	fmt.Printf("\nTanggal Terima: %s", target.TanggalTerima)
	fmt.Printf("\nDeadline: %s", target.Deadline)
	fmt.Printf("\nCatatan: %s\n", target.Catatan)

	fmt.Printf("\nJudul Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		target.Judul = input
	}

	fmt.Printf("Klien Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		target.Klien = input
	}

	fmt.Printf("Status Baru (1: Pending, 2: Ongoing, 3: Selesai) (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		switch input {
		case "1":
			target.Status = "Pending"
		case "2":
			target.Status = "Ongoing"
		case "3":
			target.Status = "Selesai"
		}
	}

	fmt.Printf("Tanggal Terima Baru (DD MM YYYY) (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		if tanggalTerformat := formatTanggal(input); tanggalTerformat != "" {
			target.TanggalTerima = tanggalTerformat
		} else {
			fmt.Println("Format tanggal tidak valid! Tanggal tidak diubah.")
		}
	}

	fmt.Printf("Deadline Baru (DD MM YYYY) (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		if tanggalTerformat := formatTanggal(input); tanggalTerformat != "" {
			target.Deadline = tanggalTerformat
		} else {
			fmt.Println("Format tanggal tidak valid! Deadline tidak diubah.")
		}
	}

	fmt.Printf("Catatan Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		target.Catatan = input
	}

	fmt.Println("Proyek berhasil diperbarui.")
}

// Fungsi hapus proyek
func hapusProyek(proyekList *[]Proyek) {
	if len(*proyekList) == 0 {
		fmt.Println("Belum ada proyek yang bisa dihapus.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var targetIndex int = -1
	var searchInput string

	fmt.Print("Masukkan ID proyek yang ingin dihapus: ")
	searchInput, _ = reader.ReadString('\n')
	searchInput = strings.TrimSpace(searchInput)

	for i, p := range *proyekList {
		if p.ID == searchInput {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		fmt.Println("Proyek tidak ditemukan.")
		return
	}

	target := (*proyekList)[targetIndex]
	fmt.Printf("\nData proyek yang akan dihapus:")
	fmt.Printf("\nID: %s", target.ID)
	fmt.Printf("\nJudul: %s", target.Judul)
	fmt.Printf("\nKlien: %s", target.Klien)
	fmt.Printf("\nStatus: %s", target.Status)
	fmt.Printf("\nTanggal Terima: %s", target.TanggalTerima)
	fmt.Printf("\nDeadline: %s", target.Deadline)
	fmt.Printf("\nCatatan: %s\n", target.Catatan)

	fmt.Print("\nApakah Anda yakin ingin menghapus proyek ini? (y/n): ")
	konfirmasi, _ := reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(strings.ToLower(konfirmasi))

	if konfirmasi == "y" {
		(*proyekList)[targetIndex] = (*proyekList)[len(*proyekList)-1]
		*proyekList = (*proyekList)[:len(*proyekList)-1]
		fmt.Println("Proyek berhasil dihapus.")
	} else {
		fmt.Println("Proyek batal dihapus.")
	}
}

// Fungsi seqsearch
func seqSearch(proyekList []Proyek, id string) (Proyek, bool) {
	for _, proyek := range proyekList {
		if proyek.ID == id {
			return proyek, true
		}
	}
	return Proyek{}, false
}

// Fungsi binsearch
func binSearch(proyekList []Proyek, id string) (Proyek, bool) {
	kiri := 0
	kanan := len(proyekList) - 1

	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if proyekList[mid].ID == id {
			return proyekList[mid], true
		}

		if proyekList[mid].ID < id {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}

	return Proyek{}, false
}

func seqSearchByID(proyekList []Proyek, id string) (Proyek, bool) {
	for _, proyek := range proyekList {
		if proyek.ID == id {
			return proyek, true
		}
	}
	return Proyek{}, false
}

func binSearchByName(proyekList []Proyek, nama string) (Proyek, bool) {
	kiri := 0
	kanan := len(proyekList) - 1

	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		banding := strings.Compare(strings.ToLower(proyekList[mid].Judul), strings.ToLower(nama))

		if banding == 0 {
			return proyekList[mid], true
		}

		if banding < 0 {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}

	return Proyek{}, false
}

// Fungsi Cari Proyek
func cariProyek(proyekList []Proyek) {
	if len(proyekList) == 0 {
		fmt.Println("Belum ada proyek yang bisa dicari.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nPilih metode pencarian:")
	fmt.Print("\n1. Cari berdasarkan ID")
	fmt.Print("\n2. Cari berdasarkan nama proyek")
	fmt.Print("\nPilihan (1/2): ")

	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.TrimSpace(pilihan)

	var hasil Proyek
	var found bool
	var searchInput string

	switch pilihan {
	case "1":
		fmt.Print("Masukkan ID proyek yang dicari: ")
		searchInput, _ = reader.ReadString('\n')
		searchInput = strings.TrimSpace(searchInput)
		hasil, found = seqSearchByID(proyekList, searchInput)
	case "2":
		fmt.Print("Masukkan nama proyek yang dicari: ")
		searchInput, _ = reader.ReadString('\n')
		searchInput = strings.TrimSpace(searchInput)

		proyekTerurut := make([]Proyek, len(proyekList))
		copy(proyekTerurut, proyekList)
		bubbleSortByJudul(&proyekTerurut)
		hasil, found = binSearchByName(proyekTerurut, searchInput)
	default:
		fmt.Println("Pilihan tidak valid!")
		return
	}

	if found {
		fmt.Println("\nProyek ditemukan:")
		fmt.Printf("ID: %s\n", hasil.ID)
		fmt.Printf("Judul: %s\n", hasil.Judul)
		fmt.Printf("Klien: %s\n", hasil.Klien)
		fmt.Printf("Status: %s\n", hasil.Status)
		fmt.Printf("Tanggal Terima: %s\n", hasil.TanggalTerima)
		fmt.Printf("Deadline: %s\n", hasil.Deadline)
		fmt.Printf("Catatan: %s\n", hasil.Catatan)
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

// Fungsi mengurutkan proyek berdasarkan Judul
func bubbleSortByJudul(proyekList *[]Proyek) {
	n := len(*proyekList)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if strings.ToLower((*proyekList)[j].Judul) > strings.ToLower((*proyekList)[j+1].Judul) {
				(*proyekList)[j], (*proyekList)[j+1] = (*proyekList)[j+1], (*proyekList)[j]
			}
		}
	}
}

// Fungsi mengurutkan proyek berdasarkan ID
func insertionSortByID(proyekList *[]Proyek) {
	n := len(*proyekList)
	for i := 1; i < n; i++ {
		key := (*proyekList)[i]
		j := i - 1
		for j >= 0 && (*proyekList)[j].ID > key.ID {
			(*proyekList)[j+1] = (*proyekList)[j]
			j--
		}
		(*proyekList)[j+1] = key
	}
}

// Fungsi mengurutkan proyek berdasarkan status
func selectionSortByStatus(proyekList *[]Proyek) {
	n := len(*proyekList)
	urutanStatus := func(status string) int {
		switch status {
		case "Pending":
			return 1
		case "Ongoing":
			return 2
		case "Selesai":
			return 3
		default:
			return 4
		}
	}

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if urutanStatus((*proyekList)[j].Status) < urutanStatus((*proyekList)[minIdx].Status) {
				minIdx = j
			}
		}
		(*proyekList)[i], (*proyekList)[minIdx] = (*proyekList)[minIdx], (*proyekList)[i]
	}
}

func initDummyData() []Proyek {
	return []Proyek{
		{
			ID:            "001",
			Judul:         "Website E-commerce",
			Klien:         "Toko Online ABC",
			Status:        "Ongoing",
			TanggalTerima: "01/01/2024",
			Deadline:      "01/03/2024",
			Catatan:       "Fokus pada fitur keranjang belanja",
		},
		{
			ID:            "002",
			Judul:         "Aplikasi Mobile Banking",
			Klien:         "Bank XYZ",
			Status:        "Pending",
			TanggalTerima: "15/01/2024",
			Deadline:      "15/04/2024",
			Catatan:       "Prioritas keamanan",
		},
		{
			ID:            "003",
			Judul:         "Sistem Manajemen Hotel",
			Klien:         "Hotel Sejahtera",
			Status:        "Selesai",
			TanggalTerima: "01/12/2023",
			Deadline:      "01/02/2024",
			Catatan:       "Sudah diimplementasi",
		},
		{
			ID:            "004",
			Judul:         "Aplikasi Fitness Tracker",
			Klien:         "Gym FitPro",
			Status:        "Ongoing",
			TanggalTerima: "10/01/2024",
			Deadline:      "10/05/2024",
			Catatan:       "Integrasi dengan wearable devices",
		},
		{
			ID:            "005",
			Judul:         "Website Portfolio",
			Klien:         "Fotografer Studio",
			Status:        "Selesai",
			TanggalTerima: "05/12/2023",
			Deadline:      "05/01/2024",
			Catatan:       "Gallery foto responsif",
		},
		{
			ID:            "006",
			Judul:         "Sistem POS Restoran",
			Klien:         "Resto Sederhana",
			Status:        "Pending",
			TanggalTerima: "20/01/2024",
			Deadline:      "20/03/2024",
			Catatan:       "Integrasi dengan kitchen display",
		},
		{
			ID:            "007",
			Judul:         "Aplikasi Delivery Service",
			Klien:         "Kurir Express",
			Status:        "Ongoing",
			TanggalTerima: "01/02/2024",
			Deadline:      "01/06/2024",
			Catatan:       "Real-time tracking",
		},
		{
			ID:            "008",
			Judul:         "Website Sekolah",
			Klien:         "SMA Negeri 1",
			Status:        "Selesai",
			TanggalTerima: "10/11/2023",
			Deadline:      "10/12/2023",
			Catatan:       "Portal informasi siswa",
		},
		{
			ID:            "009",
			Judul:         "Aplikasi Task Manager",
			Klien:         "Startup Tech",
			Status:        "Pending",
			TanggalTerima: "15/02/2024",
			Deadline:      "15/05/2024",
			Catatan:       "Kolaborasi tim",
		},
		{
			ID:            "010",
			Judul:         "Sistem Inventory",
			Klien:         "Toko Bangunan",
			Status:        "Ongoing",
			TanggalTerima: "01/03/2024",
			Deadline:      "01/07/2024",
			Catatan:       "Stok otomatis",
		},
	}
}

func main() {
	proyekList := initDummyData()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== SISTEM MANAJEMEN PROYEK FREELANCING ===")
		fmt.Println("1. Lihat Proyek")
		fmt.Println("2. Edit Proyek")
		fmt.Println("3. Hapus Proyek")
		fmt.Println("4. Cari Proyek")
		fmt.Println("5. Keluar")
		fmt.Print("\nPilih menu (1-5): ")

		pilihan, _ := reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)

		switch pilihan {
		case "1":
			lihatProyek(proyekList)
		case "2":
			editProyek(&proyekList)
		case "3":
			hapusProyek(&proyekList)
		case "4":
			cariProyek(proyekList)
		case "5":
			fmt.Println("Terima kasih telah menggunakan sistem manajemen proyek.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu 1-5.")
		}
	}
}
