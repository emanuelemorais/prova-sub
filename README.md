# Prova substitutiva - Emanuele Lacerda Morais Martins

## Como rodar o sistema

- Para essa atividade foi utilizado um broker local, sendo assim, para executa-lo entre em `prova-sub/config` e rode o comando abaixo:

```
 mosquitto -c mosquitto.conf;  
```

- Em seguida, para iniciar o simulador dos sensores da estufa, entre em `prova-sub/cmd/greenhouse-main` e rode o comando abaixo:

```
go run greenhouse-main.go
```

- Por último para realizar o recebimento dos dados enviado, entre em `prova-sub/cmd/subscriber` e rode o comando abaixo:

```
go run subscriber.go
```

Após seguir esses passos espera-se que o output das informações seja no seguinte modelo:

`<identificado da estufa> : <identificador do sensor> | <valor de umidade> %  Umidade <possível alerta>`

Ficará da seguinte forma:

```
estufa-3 : hortalicas-0 | 59%  Umidade 
estufa-2 : hortalicas-2 | 89%  Umidade [ALERTA: Umidade ALTA] 
estufa-3 : flores-0 | 69%  Umidade 
estufa-2 : flores-0 | 96%  Umidade [ALERTA: Umidade ALTA] 
estufa-3 : flores-1 | 24%  Umidade [ALERTA: Umidade BAIXA] 
estufa-3 : flores-2 | 72%  Umidade 
estufa-1 : flores-2 | 67%  Umidade 
estufa-2 : hortalicas-1 | 45%  Umidade 
estufa-1 : hortalicas-2 | 38%  Umidade 
estufa-2 : flores-1 | 70%  Umidade 
estufa-2 : flores-2 | 20%  Umidade [ALERTA: Umidade BAIXA] 
estufa-3 : hortalicas-1 | 31%  Umidade 
estufa-1 : flores-0 | 10%  Umidade [ALERTA: Umidade BAIXA] 
estufa-1 : hortalicas-0 | 5%  Umidade [ALERTA: Umidade BAIXA] 
estufa-1 : hortalicas-1 | 29%  Umidade [ALERTA: Umidade BAIXA] 
```

## Como rodar o teste

Para executar o teste entre em `prova-sub/internal/greenhouse` e rode o comando abaixo:

```
go test
```

O retorno esperado será como o do modelo abaixo:

```
Publishing message to /flores
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-1","type":"flores","humidity":63,"timestamp":"2024-04-10 09:52:17.495227 -0300 -03 m=+0.006737668"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-3","sensor_id":"flores-2","type":"flores","humidity":52,"timestamp":"2024-04-10 09:52:18.320651 -0300 -03 m=+1504.552531251"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-1","type":"flores","humidity":94,"timestamp":"2024-04-10 09:52:18.320624 -0300 -03 m=+1504.552504085"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-2","sensor_id":"flores-1","type":"flores","humidity":81,"timestamp":"2024-04-10 09:52:18.320714 -0300 -03 m=+1504.552593710"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-2","sensor_id":"flores-2","type":"flores","humidity":47,"timestamp":"2024-04-10 09:52:18.320775 -0300 -03 m=+1504.552654543"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-3","sensor_id":"flores-0","type":"flores","humidity":90,"timestamp":"2024-04-10 09:52:18.320989 -0300 -03 m=+1504.552868293"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-2","type":"flores","humidity":70,"timestamp":"2024-04-10 09:52:18.320927 -0300 -03 m=+1504.552806876"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-2","sensor_id":"flores-0","type":"flores","humidity":87,"timestamp":"2024-04-10 09:52:18.321269 -0300 -03 m=+1504.553148960"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-0","type":"flores","humidity":79,"timestamp":"2024-04-10 09:52:18.321058 -0300 -03 m=+1504.552938001"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-3","sensor_id":"flores-1","type":"flores","humidity":13,"timestamp":"2024-04-10 09:52:18.321403 -0300 -03 m=+1504.553282501"} 
Publishing message to /flores
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-1","type":"flores","humidity":38,"timestamp":"2024-04-10 09:52:20.496506 -0300 -03 m=+3.008034626"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-3","sensor_id":"flores-2","type":"flores","humidity":5,"timestamp":"2024-04-10 09:52:21.322558 -0300 -03 m=+1507.554455751"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-1","type":"flores","humidity":57,"timestamp":"2024-04-10 09:52:21.322597 -0300 -03 m=+1507.554494418"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-2","sensor_id":"flores-1","type":"flores","humidity":18,"timestamp":"2024-04-10 09:52:21.322554 -0300 -03 m=+1507.554452043"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-3","sensor_id":"flores-0","type":"flores","humidity":78,"timestamp":"2024-04-10 09:52:21.32262 -0300 -03 m=+1507.554517585"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-2","type":"flores","humidity":46,"timestamp":"2024-04-10 09:52:21.322534 -0300 -03 m=+1507.554431585"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-2","sensor_id":"flores-2","type":"flores","humidity":6,"timestamp":"2024-04-10 09:52:21.322412 -0300 -03 m=+1507.554309251"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-2","sensor_id":"flores-0","type":"flores","humidity":97,"timestamp":"2024-04-10 09:52:21.322631 -0300 -03 m=+1507.554528585"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-1","sensor_id":"flores-0","type":"flores","humidity":70,"timestamp":"2024-04-10 09:52:21.323005 -0300 -03 m=+1507.554903043"} 
New message with qos 1 on topic /estufa/flores: {"estufa_id":"estufa-3","sensor_id":"flores-1","type":"flores","humidity":10,"timestamp":"2024-04-10 09:52:21.323174 -0300 -03 m=+1507.555071918"} 
Publishing message to /hortalicas
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-1","sensor_id":"hortalicas-1","type":"hortalicas","humidity":28,"timestamp":"2024-04-10 09:52:21.497896 -0300 -03 m=+4.009430293"} 
Publishing message to /flores
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-3","sensor_id":"hortalicas-0","type":"hortalicas","humidity":48,"timestamp":"2024-04-10 09:52:24.324125 -0300 -03 m=+1510.556040501"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-2","sensor_id":"hortalicas-1","type":"hortalicas","humidity":13,"timestamp":"2024-04-10 09:52:24.324272 -0300 -03 m=+1510.556187085"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-1","sensor_id":"hortalicas-0","type":"hortalicas","humidity":78,"timestamp":"2024-04-10 09:52:24.324327 -0300 -03 m=+1510.556243001"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-1","sensor_id":"hortalicas-2","type":"hortalicas","humidity":90,"timestamp":"2024-04-10 09:52:24.324323 -0300 -03 m=+1510.556238126"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-2","sensor_id":"hortalicas-2","type":"hortalicas","humidity":13,"timestamp":"2024-04-10 09:52:24.324275 -0300 -03 m=+1510.556191001"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-1","sensor_id":"hortalicas-1","type":"hortalicas","humidity":57,"timestamp":"2024-04-10 09:52:24.324374 -0300 -03 m=+1510.556289085"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-2","sensor_id":"hortalicas-0","type":"hortalicas","humidity":37,"timestamp":"2024-04-10 09:52:24.324261 -0300 -03 m=+1510.556177001"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-3","sensor_id":"hortalicas-1","type":"hortalicas","humidity":22,"timestamp":"2024-04-10 09:52:24.324357 -0300 -03 m=+1510.556273001"} 
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-3","sensor_id":"hortalicas-2","type":"hortalicas","humidity":38,"timestamp":"2024-04-10 09:52:24.324308 -0300 -03 m=+1510.556223293"} 
Publishing message to /hortalicas
New message with qos 1 on topic /estufa/hortalicas: {"estufa_id":"estufa-1","sensor_id":"hortalicas-1","type":"hortalicas","humidity":69,"timestamp":"2024-04-10 09:52:24.499164 -0300 -03 m=+7.010716835"} 
PASS
ok      prova-sub/internal/greenhouse   8.453s
```

Esse teste realiza a validação do QOS = 1 enviado e reebido e também verifica a integridade das mensagens recebidas

## Video do funcionamento completo

https://drive.google.com/file/d/1emgNTZrV1grfXftvoYmj2SbPqSPDFwkH/view?usp=sharing