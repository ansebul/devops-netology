# Домашнее задание к занятию "4.2. Использование Python для решения типовых DevOps задач"

### 1. Есть скрипт:

```python
#!/usr/bin/env python3
a = 1
b = '2'
c = a + b
````
* Какое значение будет присвоено переменной c?
````
Никакое не присвоится. Будет ошибка, строку с числом складываем.
````

* Как получить для переменной c значение 12?
````
Привести 'a' к типу 'str'
c = str(a) + b
````

* Как получить для переменной c значение 3?
````
Привести 'b' к типу 'int'
c = a + int(b)
````

### 2. Мы устроились на работу в компанию, где раньше уже был DevOps Engineer. Он написал скрипт, позволяющий узнать, какие файлы модифицированы в репозитории, относительно локальных изменений. Этим скриптом недовольно начальство, потому что в его выводе есть не все изменённые файлы, а также непонятен полный путь к директории, где они находятся. Как можно доработать скрипт ниже, чтобы он исполнял требования вашего руководителя?

```python
import os

bash_command = ["cd ~/netology/sysadm-homeworks", "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
is_change = False
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(prepare_result)
        break
```

Ответ:

````python
#!/usr/bin/env python3

import os

# Выносим директорию с репозиторием в переменную.
repo_dir = "/devops/_DEVOPS-15/homeworks/my_repos/devops-netology"

# По-умолчанию установлена русская раскладка ru_RU.UTF-8 и все сообщения на кириллице,
# поэтому переключаемся в английскую.
bash_command = ["LANG=\"en_En.UTF-8\"", "cd ~"+repo_dir, "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
is_change = False
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print("~" + repo_dir + "/" + prepare_result)  # Выполнили второе пожелание руководства: полный путь до файла
        # break - Выполнили первое пожелание руководства: из-за этой команды выводился только первый изменённый файл.


````
 
### 3. Доработать скрипт выше так, чтобы он мог проверять не только локальный репозиторий в текущей директории, а также умел воспринимать путь к репозиторию, который мы передаём как входной параметр. Мы точно знаем, что начальство коварное и будет проверять работу этого скрипта в директориях, которые не являются локальными репозиториями.

Ответ:

````python
#!/usr/bin/env python3

import os
import sys

# Выносим директорию с репозиторием в переменную.
repo_dir = "/devops/_DEVOPS-15/homeworks/my_repos/devops-netology"

if (len(sys.argv) - 1) > 0:
    if os.path.exists(os.path.expanduser("~" + sys.argv[1] + "/.git")):
        print('Работаем с репозиторием в директории: ~' + sys.argv[1])
        repo_dir = sys.argv[1]
    else:
        print('Ошибка! В указанной директории репозитория нет!')
        exit(1)

# По-умолчанию установлена русская раскладка ru_RU.UTF-8 и все сообщения на кириллице,
# поэтому переключаемся в английскую.
bash_command = ["LANG=\"en_En.UTF-8\"", "cd ~"+repo_dir, "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
is_change = False
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print("~" + repo_dir + "/" + prepare_result)  # Выполнили второе пожелание руководства: полный путь до файла
        # break - Выполнили первое пожелание руководства: из-за этой команды выводился только первый изменённый файл.

````

### 4. Наша команда разрабатывает несколько веб-сервисов, доступных по http. Мы точно знаем, что на их стенде нет никакой балансировки, кластеризации, за DNS прячется конкретный IP сервера, где установлен сервис. Проблема в том, что отдел, занимающийся нашей инфраструктурой очень часто меняет нам сервера, поэтому IP меняются примерно раз в неделю, при этом сервисы сохраняют за собой DNS имена. Это бы совсем никого не беспокоило, если бы несколько раз сервера не уезжали в такой сегмент сети нашей компании, который недоступен для разработчиков. Мы хотим написать скрипт, который опрашивает веб-сервисы, получает их IP, выводит информацию в стандартный вывод в виде: <URL сервиса> - <его IP>. Также, должна быть реализована возможность проверки текущего IP сервиса c его IP из предыдущей проверки. Если проверка будет провалена - оповестить об этом в стандартный вывод сообщением: [ERROR] <URL сервиса> IP mismatch: <старый IP> <Новый IP>. Будем считать, что наша разработка реализовала сервисы: drive.google.com, mail.google.com, google.com.

Ответ:

````python
#!/usr/bin/env python3


import os
import socket
from ast import literal_eval

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

# Записываем результаты текущей проверки в файл
with open(state_file, 'w') as f:
    f.write(str(current_state))

````





