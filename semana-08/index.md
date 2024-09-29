# Ponderada Semana 8
Rafael Techio

## Telemetria e TDD

O TDD, conforme já discutido nas ponderadas anteriores, foca no ciclo de escrita de código e testagem do mesmo, especificando cenários e comportamentos antes da implementação. Sistemas de telemetria podem ser utilizados nesse contexto para garantir critérios de qualidade e realizar debugg durante a execução, como parte da análise dinâmica. Visto que para seu uso, primeiro há de ter uma implementação, o uso de ferramentas de telemetria pode ser posto em prática a partir do segundo ciclo do TDD.

A telemetria pode ser usada para verificar quantidade de passagens em um trecho de código, execução de escopos e uso de recursos computacionais como CPU, memória, acesso a disco e rede. Dessa forma, pode-se definir critérios, métricas e fundamentar melhorias de implementações.

Por fim, nota-se que telemetria e tdd são conceitos distintos, mas que podem ser usados juntos.

A seguir, de acordo com o artigo, como utilizar o signoz para telemetria em uma aplicação go:

## Usando telemetria em aplicações Go com Signoz

Clonar repositório do signoz:

![image](https://github.com/user-attachments/assets/f93aa36f-5ad0-4e21-b0e5-75f017309f65)

Clonar repositório de exemplo:

![image](https://github.com/user-attachments/assets/91ddaadf-78fc-4079-bdee-7858b3a59612)

O repositório deve ficar assim:

![image](https://github.com/user-attachments/assets/0bc58bc6-40ea-4b4a-ba1b-b9d2dd743824)

Deve-se iniciar o signoz utilizando uma das abordagens presentes na [documentação](https://signoz.io/docs/install/):

Vá até o projeto de exemplo e instale as dependências:

![image](https://github.com/user-attachments/assets/0d5a9578-b232-49d3-a7b8-c7d221359e25)

Certifique-se que o código instalado contem esses scripts:

![image](https://github.com/user-attachments/assets/85d6a232-00cc-4979-81f0-f070f475f8a0)
![image](https://github.com/user-attachments/assets/9de36939-edc8-49d8-8701-9fec092c0f2e)

Sete as variáveis de ambiente necessárias:

![image](https://github.com/user-attachments/assets/5fee49da-0e1f-40da-8074-35afa0bd74dc)

E agora o projeto está configurado com a telemetria de signoz.
