# Desafio

Como desenvolvedor preciso de um endpoint onde eu possa obter dados para alimentar minha máquina de aprendizagem e treinar um robo de transações financeiras:
- Disponibilizar um método GET e os parâmetros de moeda(code e codein ex: BTCUSD) e de data específica retornar os dados financeiros e as 10 notícias mais populares para essa cotação de moeda

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
        conteúdo resumido
```

Consumir as seguintes API publicas

- Para buscar dados das moedas
https://docs.awesomeapi.com.br/api-de-moedas

- Para buscar notificas 
https://newsapi.org/

