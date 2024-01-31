package uploads

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func UploadHandler(img string, id uint64, entityType string) error {
	var localFileName string

	//Verifica se a pasta images existe para fazer o upload
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		err := os.Mkdir("images", 0755)
		if err != nil {
			return errors.New("falha ao criar local do arquivo")
		}
	}

	// Baixar a imagem a partir da URL fornecida
	response, err := http.Get(img)
	if err != nil {
		return errors.New("falha ao baixar a imagem da requisição")
	}
	defer response.Body.Close()

	// Criar um arquivo local para salvar a imagem
	switch entityType {
	case "user":
		localFileName = path.Join("images", fmt.Sprintf("%v_%v", id, path.Base("user.jpg")))
	case "publication":
		localFileName = path.Join("images", fmt.Sprintf("%v_%v", id, path.Base("publication.jpg")))
	}

	//localFileName := path.Join("images", path.Base("user.jpg"))
	localFile, err := os.Create(localFileName)
	if err != nil {
		return errors.New("falha ao criar arquivo")
	}
	defer localFile.Close()

	// Copiar o conteúdo da resposta HTTP para o arquivo local
	_, err = io.Copy(localFile, response.Body)
	if err != nil {
		return errors.New("falha o conteúdo da resposta HTTP para o arquivo local")
	}
	return nil
}
