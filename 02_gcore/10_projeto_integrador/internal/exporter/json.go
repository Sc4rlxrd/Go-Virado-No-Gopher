package exporter

import (
	"encoding/json"
	"os"
)

// Interface para exportação de dados

type Exportador interface{
	Exportar(caminho string , conteudo any ) error
}

type ExportadorJSON struct{}

func (e ExportadorJSON) Exportar (caminho string , conteudo any) error{
	dados,err:= json.MarshalIndent(conteudo,""," ")
	if err!= nil{
		return err
	}
	return os.WriteFile(caminho,dados,0644)
}