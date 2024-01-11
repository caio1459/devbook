package uploads

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

//UploadHandler

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Obtém o arquivo enviado
	file, fh, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Cria o caminho completo para o arquivo
	filename := filepath.Join("images", fh.Filename)

	// Abre o arquivo para leitura
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Copia o conteúdo do arquivo para o disco
	io.Copy(f, file)
	// Fecha o arquivo
	f.Close()
	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Arquivo enviado com sucesso!")
}
