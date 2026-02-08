# App Web - Loja Go

Sistema web de gerenciamento de produtos desenvolvido em Go.

## 🚀 Desenvolvimento Rápido

### Iniciar servidor com hot reload (recomendado)

```powershell
.\dev.ps1
```

Isso iniciará o servidor na porta 8080 com **hot reload automático**. Sempre que você salvar alterações nos arquivos `.go` ou `.html`, o servidor será recompilado e reiniciado automaticamente - **muito mais rápido** que `go run main.go`!

### Modo tradicional (sem hot reload)

```powershell
go run main.go
```

## 📦 Tecnologias

- Go 1.25.4
- MySQL
- HTML Templates
- Bootstrap 4.3.1

## 🛠️ Estrutura do Projeto

```
.
├── config/
│   ├── db/          # Configuração do banco de dados
│   └── routes/      # Configuração de rotas
├── controllers/     # Controladores (handlers HTTP)
├── model/          # Modelos de dados
├── templates/      # Templates HTML
├── tmp/           # Builds temporários (ignorado no git)
└── .air.toml      # Configuração do Air (hot reload)
```

## 💡 Dicas

- **Hot reload**: O Air detecta mudanças em arquivos `.go`, `.html`, `.tpl` e `.tmpl`
- **Performance**: Compilações incrementais são ~10x mais rápidas
- **Logs de erro**: Erros de build são salvos em `build-errors.log`
- Para parar o servidor: `Ctrl+C`

## 🔧 Configuração do Banco

Edite `config/db/database.go` para ajustar a conexão MySQL.

## 🧩 Funcionalidades

- **Listar produtos**: Mostra todos os produtos na página inicial. A aplicação consulta o banco de dados via camadas de `controllers` e `model` e renderiza a lista usando os templates em `templates/`.
- **Adicionar produto**: Formulário em `new.html` permite criar um novo produto; ao submeter, o servidor recebe a requisição HTTP POST, valida os dados e insere o registro no banco.
- **Editar produto**: Formulário de edição pré-populado permite alterar campos de um produto existente; o servidor atualiza o registro no banco usando a rota de atualização correspondente.
- **Remover produto**: Ação que exclui um produto por ID; o controlador executa a remoção no banco e redireciona para a listagem.
- **Páginas e templates**: Templates HTML em `templates/` usam Bootstrap para layout; os controllers passam dados para os templates que geram o HTML final.
- **Rotas**: As rotas HTTP são configuradas em `config/routes/` e mapeiam URLs para funções nos `controllers/`.
- **Conexão com banco**: `config/db/database.go` gerencia a conexão MySQL usada por models e controllers.
- **Hot reload**: O script `dev.ps1` inicia o Air para hot reload — ao salvar `.go` ou `.html`, a aplicação recompila e reinicia automaticamente.
- **Logs de build**: Erros de compilação são registrados em `build-errors.log` para depuração rápida.
