package gen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ServiceData struct {
	ServiceName string
	MethodName  string
	DomainPkg   string
	HandlerPkg  string
	ProtoFile   string
}

func AutoGeneration() {
	// Direktori tempat file .proto berada
	protoDir := "proto"

	// Menelusuri direktori dan mencari semua file .proto
	err := filepath.Walk(protoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Jika file adalah file .proto
		if strings.HasSuffix(info.Name(), ".proto") {
			// Misalnya, kita dapat mengekstrak nama service dan method dari file ini
			data, err := parseProtoFile(path)
			if err != nil {
				log.Printf("Error parsing file %s: %v\n", path, err)
				return nil // Lanjutkan ke file berikutnya
			}

			// Generate kode berdasarkan data yang diekstrak dari file .proto
			generateCode("gen/templates/handler.tpl", data, fmt.Sprintf("gen/proto/%s_handler.go", data.ServiceName))
			generateCode("gen/templates/domain.tpl", data, fmt.Sprintf("gen/domain/%s_domain.go", data.ServiceName))
			generateCode("gen/templates/container.tpl", data, fmt.Sprintf("gen/di/%s_container.go", data.ServiceName))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generation completed.")
}

// parseProtoFile membaca dan mengekstrak data dari file .proto
// Fungsi ini bisa ditingkatkan lebih lanjut untuk parsing nama service, method, dan tipe data dari file .proto.
func parseProtoFile(filePath string) (ServiceData, error) {
	// Di sini Anda akan membaca file .proto dan memparse-nya untuk mendapatkan data yang dibutuhkan
	// Misalnya menggunakan `protogen` atau menulis custom parser untuk menangkap nama service dan method
	// Untuk saat ini kita asumsikan service dan method dapat diambil langsung dari nama file
	// Anda bisa sesuaikan dengan struktur yang lebih canggih

	// Contoh sederhana:
	baseName := filepath.Base(filePath)
	serviceName := strings.TrimSuffix(baseName, ".proto") // Anggap nama service sesuai dengan nama file .proto

	// Menghasilkan data service dari file
	return ServiceData{
		ServiceName: serviceName,
		MethodName:  "MyMethod", // Anda bisa menambahkan logic untuk mengambil metode dari file .proto
		DomainPkg:   "pkg/domain",
		HandlerPkg:  "pkg/handler",
		ProtoFile:   filePath,
	}, nil
}

func generateCode(templatePath string, data ServiceData, outputFilePath string) {
	// Membaca template dari file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}

	// Membuat file output
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}
	defer outputFile.Close()

	// Menghasilkan kode dari template dan data
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}

	fmt.Printf("File generated at %s\n", outputFilePath)
}
