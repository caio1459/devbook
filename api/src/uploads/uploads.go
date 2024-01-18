package uploads

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func UploadHandler(id uint64, img string) {
	//Verifica se a pasta images existe para fazer o upload
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		err := os.Mkdir("images", 0755)
		if err != nil {
			log.Fatalf("Falha ao criar local do arquivo: %v", err)
		}
	}

	// Baixar a imagem a partir da URL fornecida
	response, err := http.Get(img)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Criar um arquivo local para salvar a imagem
	localFileName := path.Join("images", fmt.Sprintf("user_%v_%v", id, path.Base(img)))
	localFile, err := os.Create(localFileName)
	if err != nil {
		log.Fatalf("Falha ao criar arquivo: %v", err)
	}
	defer localFile.Close()

	// Copiar o conteúdo da resposta HTTP para o arquivo local
	_, err = io.Copy(localFile, response.Body)
	if err != nil {
		log.Fatalf("Falha o conteúdo da resposta HTTP para o arquivo local: %v", err)
	}
}
