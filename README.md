# Previsão do tempo em Cli 


### Pequeno projetinho em GO visando aprender mais sobre a linguagem GO usando o padrão de API Rest 

### Utilizando a API OneCall 3.0 da OpenWHeather para extrair os dados do clima com Golang
## O uso da API é simples:

os comandos são iniciados com weatherapp e depois o comando ->

get weather(se o JSON ja estiver configurado ou utilizando as configurações default)

get coordinates -> repassa as coodenadas sem separação por virgula apenas espaços

show para exibir as configurações do JSON 

set coordinates -> repassa as coordenadas para salvar no JSON

set excluded -> repassa os campos a serem excluidos do JSON

Exemplo de uso -> weatherapp get weather -> resposta esperada é o JSON com informações de Temperatura,humidade e chance de chuva

a api funciona apenas em ambiente Linux, para usa-lo é necessario gerar uma API key no site da OpenWeather para o OneCall 3.0
execure o setup.sh e insira a sua api key para o app funcionar corretamente.

<html>
  <div> 
   <img src="https://github.com/egonelbre/gophers/blob/master/vector/fairy-tale/witch-learning.svg">
  </div>
 
</html>
