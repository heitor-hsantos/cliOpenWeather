#!/bin/bash

# Define o diretório e o arquivo de configuração
# Usar ~/.config/ é uma convenção padrão no Linux
CONFIG_DIR="$HOME/.config/weatherapp"
CONFIG_FILE="$CONFIG_DIR/config"

echo "--- Configuração do Weather App ---"

# Pede a chave da API ao usuário
echo -n "Por favor, insira sua chave da API do OpenWeatherMap e pressione [ENTER]: "
read API_KEY

# Verifica se a chave foi inserida
if [ -z "$API_KEY" ]; then
    echo "Nenhuma chave de API foi fornecida. A instalação foi cancelada."
    exit 1
fi

# Cria o diretório de configuração, se não existir
mkdir -p "$CONFIG_DIR"

# Escreve as variáveis no arquivo de configuração
# O '>' cria/sobrescreve o arquivo, e o '>>' adiciona ao final
echo "OPENWEATHER_API_KEY=${API_KEY}" > "$CONFIG_FILE"
echo "OPENWEATHER_API_URL=https://api.openweathermap.org/data/3.0/onecall" >> "$CONFIG_FILE"

echo ""
echo "✅ Configuração concluída com sucesso!"
echo "O arquivo de configuração foi salvo em: $CONFIG_FILE"