# Домашнее задание к занятию "4.3. Языки разметки JSON и YAML"

### 1. Мы выгрузили JSON, который получили через API запрос к нашему сервису:

````
{ "info" : "Sample JSON output from our service\t",
    "elements" :[
        { "name" : "first",
        "type" : "server",
        "ip" : 7175 
        },
        { "name" : "second",
        "type" : "proxy",
        "ip : 71.78.22.43
        }
    ]
}
Нужно найти и исправить все ошибки, которые допускает наш сервис
````
Ответ:
````json
{ "info" : "Sample JSON output from our service\t",
    "elements" : [
        {
          "name" : "first",
          "type" : "server",
          "ip" : 7175
        },
        {
          "name" : "second",
          "type" : "proxy",
          "ip" : "71.78.22.43"
        }
    ]
}
````

### 2. В прошлый рабочий день мы создавали скрипт, позволяющий опрашивать веб-сервисы и получать их IP. К уже реализованному функционалу нам нужно добавить возможность записи JSON и YAML файлов, описывающих наши сервисы. Формат записи JSON по одному сервису: { "имя сервиса" : "его IP"}. Формат записи YAML по одному сервису: - имя сервиса: его IP. Если в момент исполнения скрипта меняется IP у сервиса - он должен так же поменяться в yml и json файле.

Ответ:

````python
#!/usr/bin/env python3


import os
import socket
from ast import literal_eval
import json
import yaml


last_state = {}
current_state = {}
state_file = "./last_state.txt"
services = ["drive.google.com", "mail.google.com", "google.com"]

# открываем файл c результатами последней проверки
if os.path.exists(state_file):
    with open(state_file, 'r') as f:
        last_state = literal_eval(f.readline().strip())

# Проверяем текущее состояние
for service in services:
    ips = socket.gethostbyname_ex(service)
    print(f'{service} - {ips[2]}')
    current_state[service] = ips[2]
    if len(last_state) > 0:
        old = set(last_state[service])
        new = set(ips[2])
        diff = new - old
        if len(diff) > 0:
            # Печатаем различия
            print(f'[ERROR] <{service}> IP mismatch: <{old}> <{new}>')

# Записываем результаты последней проверки в файл
with open(state_file, 'w') as f:
    f.write(str(current_state))

for service in current_state:
    content = {}
    content[service] = current_state[service]

    # Записываем json
    with open(service + '.json', 'w') as f:
        f.write(json.dumps(content))
    # Записываем yaml
    with open(service + '.yaml', 'w') as f:
        f.write(yaml.dump(content))
````
