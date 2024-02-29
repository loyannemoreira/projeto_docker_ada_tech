<<<<<<< HEAD
<<<<<<< HEAD
# projeto_docker_ada_tech
Este projeto consiste em uma aplicação de formulário de cadastro desenvolvida em Golang, utilizando MongoDB como banco de dados, e Mongo Express como ferramenta de administração do banco de dados. O principal objetivo é demonstrar como executar e gerenciar uma aplicação em um ambiente de contêineres Docker.
=======
# Formulário de Cadastro em Golang com MongoDB

## Descrição
Este projeto consiste em uma aplicação de formulário de cadastro desenvolvida em Golang, utilizando MongoDB como banco de dados, e Mongo Express como ferramenta de administração do banco de dados. O principal objetivo é demonstrar como executar e gerenciar uma aplicação em um ambiente de contêineres Docker.
=======
# Formulário de Cadastro em Golang com MongoDB

## Descrição
Este projeto consiste em uma aplicação de formulário de cadastro desenvolvida em Golang, utilizando MongoDB como banco de dados, e Mongo Express como ferramenta de administração do banco de dados. O principal objetivo é demonstrar como executar e gerenciar uma aplicação em um ambiente de contêineres Docker.
>>>>>>> origin/main

## Funcionalidades da Aplicação

A aplicação permite aos usuários preencher um formulário de cadastro com campos de Nome e Endereço, armazenando os dados no banco de dados MongoDB.

## Imagens Docker

- Golang
- MongoDB
- Mongo Express

## Justificativa das Ferramentas Escolhidas

### Golang (Go):

- **Eficiência e Desempenho:** Golang é conhecido por sua eficiência e desempenho, tornando-o ideal para ambientes de DevOps onde a velocidade é essencial.

- **Facilidade de Compilação e Distribuição:** A capacidade de compilar binários independentes facilita a distribuição e implantação de aplicativos em diferentes ambientes.

- **Concorrência e Paralelismo:** Golang oferece suporte nativo à concorrência e ao paralelismo, essenciais para sistemas de DevOps escaláveis.

- **Ecossistema de Ferramentas:** Golang possui um ecossistema robusto de ferramentas e bibliotecas, incluindo muitas ferramentas populares para automação, testes e implantação contínua.

### MongoDB:

- **Flexibilidade e Escalabilidade:** MongoDB é altamente flexível e escalável, adequado para ambientes de DevOps com necessidades variáveis de armazenamento e processamento de dados.

- **Modelo de Dados Semiestruturado:** O modelo semiestruturado do MongoDB facilita a adaptação a mudanças nos requisitos de dados, comuns em ambientes ágeis de desenvolvimento.

- **Integração com Ferramentas de DevOps:** MongoDB oferece integrações sólidas com várias ferramentas de DevOps, facilitando sua incorporação em pipelines de CI/CD e ambientes de orquestração de contêineres.

- **Documentação e Comunidade Ativa:** MongoDB possui uma documentação abrangente e uma comunidade ativa, fornecendo suporte valioso para equipes de DevOps.


## Uso e Instalação

1. Clone este repositório:
``` git
git clone https://github.com/seu-usuario/seu-projeto.git
```

2. Construa as imagens Docker e construa os containers:
``` bash
docker-compose up -d --build
```

3. Acesse a aplicação web em seu navegador:
``` 
http://localhost:8080
```

## Desafios e Bugs Encontrados

Durante o desenvolvimento deste projeto, enfrentamos os seguintes desafios e bugs:

**Desafio**: Integrar as versões do MongoDB e do Golang. 
**Descrição e resolução**: A primeira tentativa foi declarar as imagens com a tag "latest". Porém, ao tentar construir os containers, surgiram alguns conflitos. A resolução foi realizar testes com as versões mais estáveis, até encontrar a melhor combinação.

## Recursos Adicionais

- [Repositório do projeto no Docker Hub] (https://hub.docker.com/repository/docker/loyanne/projeto_final_conteinerizacao_ada_tech/general)
- [Documentação do Docker] (https://docs.docker.com/)
- [Documentação do Golang] (https://golang.org/doc/)
- [Documentação do MongoDB] (https://docs.mongodb.com/)
- [Documentação do Mongo Express] https://github.com/mongo-express/mongo-express

## Observações

- Este projeto é um exemplo simplificado e educativo. Em ambientes de produção, recomenda-se adotar práticas de segurança adicionais, como autenticação de usuários e criptografia de dados.

## Contribuições
Contribuições são sempre muito bem-vindas! Sinta-se à vontade para enviar sugestões e pull requests para este repositório.
<<<<<<< HEAD
>>>>>>> cdcf7d5 (Projeto do módulo de Conteinerização do programa Deva (Ada + b3))
=======
>>>>>>> origin/main
