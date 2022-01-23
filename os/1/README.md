# Домашнее задание к занятию "3.3. Операционные системы, лекция 1"


###  1. Какой системный вызов делает команда cd? В прошлом ДЗ мы выяснили, что cd не является самостоятельной программой, это shell builtin, поэтому запустить strace непосредственно на cd не получится. Тем не менее, вы можете запустить strace на /bin/bash -c 'cd /tmp'. В этом случае вы увидите полный список системных вызовов, которые делает сам bash при старте. Вам нужно найти тот единственный, который относится именно к cd.

Ответ: `chdir()`

Решение:

    $ strace /bin/bash -c 'cd /tmp'
    В выводе видим:
    
    ...
    chdir("/tmp")
    ...

    Читаем man chdir и убеждаемся, что нашли ответ:

    chdir()  changes  the current working directory of the calling process to the directory specified in path.


### 2. Используя strace выясните, где находится база данных `file` на основании которой она делает свои догадки.

   Ответ: `/usr/share/misc/magic.mgc`

   Решение:
   
    Логично предположить, что нам надо найти все файлы, которые file открывает на чтение, но которые не являются системными библиотеками и переданным ему аргументом.
   
    Запускаем:
    $ strace -f /usr/bin/file /bin/id

    И видим попытки обращения к

    stat("/home/vagrant/.magic.mgc", 0x7ffe9dcd9da0) = -1 ENOENT (No such file or directory)
    stat("/home/vagrant/.magic", 0x7ffe9dcd9da0) = -1 ENOENT (No such file or directory)
    openat(AT_FDCWD, "/etc/magic.mgc", O_RDONLY) = -1 ENOENT (No such file or directory)
    stat("/etc/magic", {st_mode=S_IFREG|0644, st_size=111, ...}) = 0
    openat(AT_FDCWD, "/etc/magic", O_RDONLY) = 3
    ...
    openat(AT_FDCWD, "/usr/share/misc/magic.mgc", O_RDONLY) = 3

    Посмотрим, что находится в /etc/magic и находим там отсылку к man magic. Читаем man и узнаём, что основная база сигнатур обычно расположена в /usr/share/misc/magic.mgc, а предыдущие вызовы - пользовательские базы.

    Собственно, ниже в выводе strace видим обращение:
    openat(AT_FDCWD, "/usr/share/misc/magic.mgc", O_RDONLY) = 3


### 3. Предположим, приложение пишет лог в текстовый файл. Этот файл оказался удален (deleted в lsof), однако возможности сигналом сказать приложению переоткрыть файлы или просто перезапустить приложение – нет. Так как приложение продолжает писать в удаленный файл, место на диске постепенно заканчивается. Основываясь на знаниях о перенаправлении потоков предложите способ обнуления открытого удаленного файла (чтобы освободить место на файловой системе).

 Ответ: `cat /dev/null > /proc/1712/fd/4`

 Решение:

    Эмулируем ситуацию:
    $ exec 4> ping.log
    $ ping localhost >&4

    В другой консоли:
    $ rm ~/ping.log
    $ ps -ax | grep ping
     1712 pts/0    S+     0:00 ping localhost
   
    # lsof -p1712
    COMMAND  PID    USER   FD   TYPE DEVICE SIZE/OFF    NODE NAME
    ping    1712 vagrant  cwd    DIR  253,0     4096 1051845 /home/vagrant
    ping    1712 vagrant  rtd    DIR  253,0     4096       2 /
    ping    1712 vagrant  txt    REG  253,0    72776 1835881 /usr/bin/ping
    ping    1712 vagrant  mem    REG  253,0    51832 1841607 /usr/lib/x86_64-linux-gnu/libnss_files-2.31.so
    ping    1712 vagrant  mem    REG  253,0  3035952 1835290 /usr/lib/locale/locale-archive
    ping    1712 vagrant  mem    REG  253,0   137584 1841525 /usr/lib/x86_64-linux-gnu/libgpg-error.so.0.28.0
    ping    1712 vagrant  mem    REG  253,0  2029224 1841468 /usr/lib/x86_64-linux-gnu/libc-2.31.so
    ping    1712 vagrant  mem    REG  253,0   101320 1841650 /usr/lib/x86_64-linux-gnu/libresolv-2.31.so
    ping    1712 vagrant  mem    REG  253,0  1168056 1835853 /usr/lib/x86_64-linux-gnu/libgcrypt.so.20.2.5
    ping    1712 vagrant  mem    REG  253,0    31120 1841471 /usr/lib/x86_64-linux-gnu/libcap.so.2.32
    ping    1712 vagrant  mem    REG  253,0   191472 1841428 /usr/lib/x86_64-linux-gnu/ld-2.31.so
    ping    1712 vagrant    0u   CHR  136,0      0t0       3 /dev/pts/0
    ping    1712 vagrant    1w   REG  253,0    68456 1048591 /home/vagrant/ping.log (deleted)
    ping    1712 vagrant    2u   CHR  136,0      0t0       3 /dev/pts/0
    ping    1712 vagrant    3u  icmp             0t0   30895 00000000:0001->00000000:0000
    ping    1712 vagrant    4w   REG  253,0    68456 1048591 /home/vagrant/ping.log (deleted)
    ping    1712 vagrant    5u  sock    0,9      0t0   30896 protocol: PINGv6

    # ls -l /proc/1712/fd  
    total 0
    lrwx------ 1 root root 64 Jan 23 12:37 0 -> /dev/pts/0
    l-wx------ 1 root root 64 Jan 23 12:37 1 -> '/home/vagrant/ping.log (deleted)'
    lrwx------ 1 root root 64 Jan 23 12:37 2 -> /dev/pts/0
    lrwx------ 1 root root 64 Jan 23 12:37 3 -> 'socket:[30895]'
    l-wx------ 1 root root 64 Jan 23 12:37 4 -> '/home/vagrant/ping.log (deleted)'
    lrwx------ 1 root root 64 Jan 23 12:37 5 -> 'socket:[30896]'

   Основываясь на полученных знаниях пытаемся обнулить лог:
   `# cat /dev/null > /proc/1712/fd/4`

  К сожалению, файл не обнуляется.  Проблему, похоже, надо решать другим способом.
  Но моё дело - предложить (с). Если у Вас есть рабочее решение, мне было бы интересно его узнать.

 


### 4. Занимают ли зомби-процессы какие-то ресурсы в ОС (CPU, RAM, IO)?

  Ответ: Нет.

  Процесс при завершении (как нормальном, так и в результате не обрабатываемого сигнала) освобождает все свои ресурсы и становится «зомби» — пустой записью в таблице процессов, хранящей статус завершения, предназначенный для чтения родительским процессом.
  Может только переполнить таблицу процессов.

### 5. В iovisor BCC есть утилита opensnoop. На какие файлы вы увидели вызовы группы open за первую секунду работы утилиты?

 Зададим промежуток времени 1 секунда:

    # opensnoop-bpfcc -d1
    PID    COMM               FD ERR PATH
    1077   vminfo              4   0 /var/run/utmp
    610    dbus-daemon        -1   2 /usr/local/share/dbus-1/system-services
    610    dbus-daemon        18   0 /usr/share/dbus-1/system-services
    610    dbus-daemon        -1   2 /lib/dbus-1/system-services
    610    dbus-daemon        18   0 /var/lib/snapd/dbus-1/system-services/


### 6. Какой системный вызов использует uname -a? Приведите цитату из man по этому системному вызову, где описывается альтернативное местоположение в /proc, где можно узнать версию ядра и релиз ОС.

Находим вызовы:
     
    $ strace uname -a
    ...
    uname({sysname="Linux", nodename="vagrant", ...}) = 0
    ...
    uname({sysname="Linux", nodename="vagrant", ...}) = 0
    uname({sysname="Linux", nodename="vagrant", ...}) = 0
    ...

В man uname ничего полезного нет, кроме отсылки к uname(2). Ставим дополнительные маны: 

  `# apt install manpages-dev`

Читаем man 2 uname и находим, что часть информации можно найти в файловой системе: /proc/sys/kernel/{ostype, hostname, osrelease, version, domainname}. И проверим это:

    $ cat /proc/sys/kernel/{ostype,hostname,osrelease,version,domainname}
    Linux
    vagrant
    5.4.0-91-generic
    #102-Ubuntu SMP Fri Nov 5 16:31:28 UTC 2021
    (none)

Ещё версию ОС можно узнать:

    $ cat /etc/issue
    Ubuntu 20.04.3 LTS \n \l

    $ cat /etc/lsb-release
    DISTRIB_ID=Ubuntu
    DISTRIB_RELEASE=20.04
    DISTRIB_CODENAME=focal
    DISTRIB_DESCRIPTION="Ubuntu 20.04.3 LTS"


### 7. Чем отличается последовательность команд через ; и через && в bash? Есть ли смысл использовать в bash &&, если применить set -e?

Команды, разделённые ";" выполняются последовательно, одна за другой.
Команды, разделённые через "&&" выполняются только при успешном выполнении предыдущей.
Есть смысл использовать в bash && если применить set -e. При даной конструкции оболочка не видит ненулевого статуса и не завершает работу.





### 8. Из каких опций состоит режим bash set -euxo pipefail и почему его хорошо было бы использовать в сценариях?

    -e указывает оболочке выйти, если команда дает ненулевой статус выхода.
    -u обрабатывает неустановленные или неопределенные переменные, за исключением специальных параметров, таких как подстановочные знаки (*) или «@», как ошибки во время раскрытия параметра
    -x печатает аргументы команды во время выполнения
    -o pipefail если какая то команда в скрипте завершится кодом ошибки то этот код будет присвоен всему скрипту, иначе скрипт получает код последней выполненной команды

 Выглядит полезным при написании сценариев, т.к. заставляет писать код правильно и меньше возиться с дебагом.


### 9. Используя `-o stat` для `ps`, определите, какой наиболее часто встречающийся статус у процессов в системе. В man ps ознакомьтесь (/PROCESS STATE CODES) что значат дополнительные к основной заглавной буквы статуса процессов. Его можно не учитывать при расчете (считать S, Ss или Ssl равнозначными).

    # ps -a -o stat
    STAT
    S
    S
    S+
    R+

Самый частый - спящий процесс, который ожидает какого-то события для завершения (S), работающий процесс на втором месте (R)

