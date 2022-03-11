# Домашнее задание к занятию "4.1. Командная оболочка Bash: Практические навыки"

## Обязательные задания

### 1. Есть скрипт:

```bash
a=1
b=2
c=a+b
d=$a+$b
e=$(($a+$b))
````
* Какие значения переменным c,d,e будут присвоены?
* Почему?

Ответ:

````
a=1             # Присвоено "1" (строковое)
b=2             # Присвоено "2" (строковое)
c=a+b           # Присвоено "a+b" (потому-что это строка)
d=$a+$b         # Присвоено "1+2" (В строку подставляются значения переменных)
e=$(($a+$b))    # Присвоено "3" (Выполнили операцию с целыми числами)
````


### 2. На нашем локальном сервере упал сервис и мы написали скрипт, который постоянно проверяет его доступность, записывая дату проверок до тех пор, пока сервис не станет доступным. В скрипте допущена ошибка, из-за которой выполнение не может завершиться, при этом место на Жёстком Диске постоянно уменьшается. Что необходимо сделать, чтобы его исправить:

```bash
while ((1==1)
do
	curl https://localhost:4757
	if (($? != 0))
	then
		date >> curl.log
	fi
done
```

Ответ:

````
Значение $? надо сначала присвоить переменной, а потом уже проверять,
т.к. перезаписывается при проверке в if.
````
 
### 3. Необходимо написать скрипт, который проверяет доступность трёх IP: 192.168.0.1, 173.194.222.113, 87.250.250.242 по 80 порту и записывает результат в файл log. Проверять доступность необходимо пять раз для каждого узла.

Скрипт:

````bash
#!/usr/bin/env bash


iplist=("192.168.0.1" "173.194.222.113" "87.250.250.242")
logfile="./checker_v1.log"

for ip in ${iplist[@]}
do
    a=1
    while (($a < 6))
    do
        status=$(nc -zv -w 2  ${ip} 80 2>&1)
        echo "Check#${a} ip: ${ip} result: ${status}" >> $logfile
        let "a += 1"
        sleep 1
    done
done
````

Лог:
````
Check#1 ip: 192.168.0.1 result: nc: connect to 192.168.0.1 port 80 (tcp) timed out: Operation now in progress
Check#2 ip: 192.168.0.1 result: nc: connect to 192.168.0.1 port 80 (tcp) timed out: Operation now in progress
Check#3 ip: 192.168.0.1 result: nc: connect to 192.168.0.1 port 80 (tcp) timed out: Operation now in progress
Check#4 ip: 192.168.0.1 result: nc: connect to 192.168.0.1 port 80 (tcp) timed out: Operation now in progress
Check#5 ip: 192.168.0.1 result: nc: connect to 192.168.0.1 port 80 (tcp) timed out: Operation now in progress
Check#1 ip: 173.194.222.113 result: Connection to 173.194.222.113 80 port [tcp/http] succeeded!
Check#2 ip: 173.194.222.113 result: Connection to 173.194.222.113 80 port [tcp/http] succeeded!
Check#3 ip: 173.194.222.113 result: Connection to 173.194.222.113 80 port [tcp/http] succeeded!
Check#4 ip: 173.194.222.113 result: Connection to 173.194.222.113 80 port [tcp/http] succeeded!
Check#5 ip: 173.194.222.113 result: Connection to 173.194.222.113 80 port [tcp/http] succeeded!
Check#1 ip: 87.250.250.242 result: Connection to 87.250.250.242 80 port [tcp/http] succeeded!
Check#2 ip: 87.250.250.242 result: Connection to 87.250.250.242 80 port [tcp/http] succeeded!
Check#3 ip: 87.250.250.242 result: Connection to 87.250.250.242 80 port [tcp/http] succeeded!
Check#4 ip: 87.250.250.242 result: Connection to 87.250.250.242 80 port [tcp/http] succeeded!
Check#5 ip: 87.250.250.242 result: Connection to 87.250.250.242 80 port [tcp/http] succeeded!
````

### 4. Необходимо дописать скрипт из предыдущего задания так, чтобы он выполнялся до тех пор, пока один из узлов не окажется недоступным. Если любой из узлов недоступен - IP этого узла пишется в файл error, скрипт прерывается

Скрипт:

````bash
#!/usr/bin/env bash


iplist=("173.194.222.113" "87.250.250.242" "192.168.0.1")
logfile="./error"

while ((1 == 1))
do
    for ip in ${iplist[@]}
    do
        status=$(nc -zv -w 2  ${ip} 80 2>&1 | grep succeeded| wc -l)
        if [ "$status" -eq "0" ]; then
            echo "Хост недоступен! ip: ${ip}" >> $logfile
            exit
        fi
    done
done

````

Лог:

````
Хост недоступен! ip: 192.168.0.1
````



## Дополнительное задание (со звездочкой*) - необязательно к выполнению

Мы хотим, чтобы у нас были красивые сообщения для коммитов в репозиторий. Для этого нужно написать локальный хук для git, который будет проверять, что сообщение в коммите содержит код текущего задания в квадратных скобках и количество символов в сообщении не превышает 30. Пример сообщения: \[04-script-01-bash\] сломал хук.


*Один из вариантов может быть таким:*
````bash
#!/usr/bin/env bash

commitRegex='^(\[[0-9]{1,5}\].{0,25})'
if ! grep -qE "$commitRegex" "$1"; then
    echo "Aborting according commit message policy. Please, use specify issue code: [XXX] message"
    exit 1
fi

````

