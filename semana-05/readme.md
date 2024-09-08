# Ponderada Semana 5 - Teste de aplicação React com Cypress

Rafael Techio

Essa ponderada é guiada pelo tutorial: https://www.browserstack.com/guide/how-to-test-react-using-cypress.

Através dessa ponderada, ampliando por aqui os comentários de código, pude aplicar conceitos de TDD através de testes de interface automatizados. Dessa forma, podemos extender os conceitos discutidos nas duas últimas ponderadas também para testes de frontend. 

A seguir, detalho como iniciei um projeto cypress, configurei o primeiro teste e executei dois testes que falharam conforme o tutorial. Por fim, executo um terceiro teste na wikipedia que retorna sucesso.


## Criando package.json: 

![image](https://github.com/user-attachments/assets/c42151ba-e0b0-4eb1-8615-8d9206d7a556)


## Instalando cypress:

![image](https://github.com/user-attachments/assets/70606cf8-5d95-4f15-90f5-c1474de2a1f8)

## Abrindo cypress

```
npx cypress open
```

![image](https://github.com/user-attachments/assets/2a430180-0235-4c59-b7a9-ca16f591dcc7)

## Configurando Teste

Selecione E2E testing:

![image](https://github.com/user-attachments/assets/bf657977-cdce-498a-b77b-b73aeb335b1a)

Caso necessário, as configurações gerais do projeto cypress podem ser editadas aqui. Clique em Continue.

![image](https://github.com/user-attachments/assets/a37afe96-3c1f-48d3-9a0a-1b9c832f83e0)

Selecione o browser que deseja, no caso, selecionei o chrome. E clique no botão "Start E2E Testing in Chrome".

![image](https://github.com/user-attachments/assets/ddfaa8b7-5d7b-4399-b435-54a3b54c76bc)

Na tela que irá abrir, Selecione "Create new spec":

![image](https://github.com/user-attachments/assets/c6449730-696a-4ed1-801e-bc35843e91ba)

Edite o nome:

![image](https://github.com/user-attachments/assets/663d52fc-b28c-46cb-890d-9c8cb2812e49)

Caso desejar, aqui já poderá ser editado o teste do spec. Clicque em "Okay, run the spec":

![image](https://github.com/user-attachments/assets/2bd11adf-def1-45a9-94df-d516c306dfde)

## Rodando Teste de Exemplo

Ao selecionar para rodar o spec, a tela de loading aparecerá:

![image](https://github.com/user-attachments/assets/34864a24-dec8-4afb-af8d-da6e442be009)

Ao final o resultado ficará disponível, junto com a tela onde os testes pararam:

![image](https://github.com/user-attachments/assets/3475d4fe-dc14-4b2b-a485-b6df0edb5987)

## Editando Teste

O teste pode ser editado diretamente por IDEs na pasta configurada do cypress:

![image](https://github.com/user-attachments/assets/e595540f-2a3d-4783-9ac1-845fe3e01a4a)

## Executando teste por linha de comando

![image](https://github.com/user-attachments/assets/ff1e3da7-f09f-495a-8fef-535a0a4ddc8c)


## Segundo teste

![image](https://github.com/user-attachments/assets/b8dbf8c6-4f09-4bc9-97a0-e257c84ddf48)

![image](https://github.com/user-attachments/assets/65d050a7-7bd5-4862-a3b7-fd52b70e04da)


## Teste autoral

Devo entrar em uma página na wikipedia e verificar seu título. Escolhi a página dos Jogos Paralímpicos de Verão de 2024 em Paris.

![image](https://github.com/user-attachments/assets/6c37a8d0-2bc0-47ec-a8f4-99d0f4bb4238)


Código:

![image](https://github.com/user-attachments/assets/6b37f2f6-c0b4-4b5b-a491-0f9f709edd2a)

Resultado:

![image](https://github.com/user-attachments/assets/421a484f-399a-4a97-8927-8f8a0644b230)

