# Desafio

Como desenvolvedor preciso de um endpoint onde posso obter dados par alimentar minha maquina de aprendizagem a treinar um bot de transações financeiras:
- Disponibilizar um metodo GET e com os parametros de moeda(code e codein ex: BTCUSD) e uma data especifica retornar os dados financeiros e as 10 noticias mais populares para essa cotação de moeda

## Precisamos ter como resposta caso sucesso:

```
Simbolo da cotação(ex: btcusd)
    nome da cotação
    maior valor
    menor valor
    lista de noticias com as 10 noticias mais populares
        autor
        titulo
        descrição
        conteudo resumido
```

Consumir as seguintes API publicas

- Para buscar dados das moedas
https://docs.awesomeapi.com.br/api-de-moedas

- Para buscar notificas 
https://newsapi.org/
