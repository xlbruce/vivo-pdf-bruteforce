package main
import (
	"fmt"
	"log"
	"os"

	"rsc.io/pdf"
)

func GeneratePasswords(result *string) func() string {
	i := 0
	return func() string {
		if i == 999 {
			log.Fatal("Máximo número de tentativas. Não foi possível descriptografar o arquivo. (999)\n")
			return ""
		}

		i++
		p := fmt.Sprintf("%03d", i)
		*result = p
		return p
	}
}

func usage() string {
  //TODO improve binary execution
  return fmt.Sprintf("Usage: %s <pdf_filename>", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(usage())
	}

	var password string

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		msg := fmt.Sprintf("Erro ao abrir o arquivo: %s\n", err.Error())
		log.Fatal(msg)
	}

	stat, _ := f.Stat()
	_, err = pdf.NewReaderEncrypted(f, stat.Size(), GeneratePasswords(&password))
	if err != nil {
		msg := fmt.Sprintf("Arquivo PDF inválido: %s", err.Error())
		log.Fatal(msg)
	}

	if password == "" {
		fmt.Printf("O arquivo não é protegido por senha.\n")
	} else {
		fmt.Printf("A senha do PDF %s é: %s\n", filename, password)
	}

}
