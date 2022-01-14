# Домашнее задание к занятию «3.2. Работа в терминале, лекция 2»


###  1. Какого типа команда cd? Попробуйте объяснить, почему она именно такого типа; опишите ход своих мыслей, если считаете что она могла бы быть другого типа.

Ответ: Это встроенная команда оболочки

    $ type cd
    cd is a shell builtin


### 2. Какая альтернатива без pipe команде `grep <some_string> <some_file> | wc -l`? `man grep` поможет в ответе на этот вопрос. Ознакомьтесь с документом о других подобных некорректных вариантах использования pipe.

    grep -c <some_string> <some_file>

### 3. Какой процесс с PID 1 является родителем для всех процессов в вашей виртуальной машине Ubuntu 20.04?

 Ответ: `systemd`

    $ ps -axu | awk '($2=="1"){print $0}'
    root           1  0.6  1.1 167416 11264 ?        Ss   15:39   0:02 /sbin/init

 Посмотрим, что это за /sbin/init (а это символическая ссылка на systemd):

    $ ls -luF /sbin/init
    lrwxrwxrwx 1 root root 20 Jan 14 15:39 /sbin/init -> /lib/systemd/systemd*

### 4. Как будет выглядеть команда, которая перенаправит вывод stderr ls на другую сессию терминала?

    ls 2>/dev/pts/3

### 5. Получится ли одновременно передать команде файл на stdin и вывести ее stdout в другой файл? Приведите работающий пример.

 С прошлой домашней работы осталась директория с 100000 файлов:

    $ ls -1 testfolder/ | wc -l
    100000

 Используем её:

    $ ls -1 ./testfolder/ > filenames.txt
    $ cat filenames.txt | wc -l > count_lines.txt
    $ cat count_lines.txt 
    100000

 Ответ на вопрос:

    $ cat filenames.txt | wc -l > count_lines.txt

    # Можно сделать короче:

    wc -l < filenames.txt > count_lines.txt


### 6. Получится ли находясь в графическом режиме, вывести данные из PTY в какой-либо из эмуляторов TTY? Сможете ли вы наблюдать выводимые данные?

 Да, получилось. Предварительно переключился в консоль `Ctr+Alt+F2`, авторизовался. Выполнил команду `tty` и узнал устройство (/dev/tty2)
 Переключился в графический режим и выполнил:

    $ tty
    /dev/pts/5
    $ echo "Hello!" > /dev/tty2

 Переключился в консоль и увидел [результат](https://postimg.cc/VJqBVBQ5)


### 7. Выполните команду `bash 5>&1`. К чему она приведет?

К запуску командной оболочки bash c перенаправленным файловым дескриптором 5 в стандартный вывод.

### Что будет, если вы выполните `echo netology > /proc/$$/fd/5`? Почему так происходит?

Увидим "netology" в терминале, потому-что мы пишем в файловый дескриптор 5, который является симлинком для текущей сессии терминала:

    $ tty
    /dev/pts/0
    vagrant@vagrant:~$ ll /proc/$$/fd
    total 0
    dr-x------ 2 vagrant vagrant  0 Nov 15 08:38 ./
    dr-xr-xr-x 9 vagrant vagrant  0 Nov 15 08:38 ../
    lrwx------ 1 vagrant vagrant 64 Nov 15 08:43 0 -> /dev/pts/0
    lrwx------ 1 vagrant vagrant 64 Nov 15 08:43 1 -> /dev/pts/0
    lrwx------ 1 vagrant vagrant 64 Nov 15 08:43 2 -> /dev/pts/0
    lrwx------ 1 vagrant vagrant 64 Nov 15 08:43 255 -> /dev/pts/0
    lrwx------ 1 vagrant vagrant 64 Nov 15 08:38 5 -> /dev/pts/0


### 8. Получится ли в качестве входного потока для pipe использовать только stderr команды, не потеряв при этом отображение stdout на pty? Напоминаем: по умолчанию через pipe передается только stdout команды слева от `|` на stdin команды справа. Это можно сделать, поменяв стандартные потоки местами через промежуточный новый дескриптор, который вы научились создавать в предыдущем вопросе.

 Ответ (1-е - получаем листинг домашней директории, т.е. не теряем stdout. И 2-е - сообщение об ошибке отправляем на вход awk):

    $ ls nonexist_dir ~  3>&2 2>&1 1>&3 | awk '{print $2, $3, $4}'
    cannot access 'nonexist_dir':
    /home/vagrant:
    count_lines.txt  error.txt  filenames.txt  testfolder


### 9. Что выведет команда cat `/proc/$$/environ`? Как еще можно получить аналогичный по содержанию вывод?

Выводит переменные окружения:

    $ cat /proc/$$/environ
    LANG=en_US.UTF-8USER=vagrantLOGNAME=vagrantHOME=/home/vagrantPATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/binSHELL=/bin/bashTERM=xterm-256colorXDG_SESSION_ID=8XDG_RUNTIME_DIR=/run/user/1000DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/busXDG_SESSION_TYPE=ttyXDG_SESSION_CLASS=userMOTD_SHOWN=pamLANGUAGE=en_US:SSH_CLIENT=10.0.2.2 41938 22SSH_CONNECTION=10.0.2.2 41938 10.0.2.15 22SSH_TTY=/dev/pts/0

 Аналогичный по содержанию (только больше) вывод можно получить с помощью команды env или printenv:

    SHELL=/bin/bash
    LANGUAGE=en_US:
    PWD=/home/vagrant
    LOGNAME=vagrant
    XDG_SESSION_TYPE=tty
    MOTD_SHOWN=pam
    HOME=/home/vagrant
    LANG=en_US.UTF-8
    LS_COLORS=rs=0:di=01;34:ln=01;36:mh=00:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:mi=00:su=37;41:sg=30;43:ca=30;41:tw=30;42:ow=34;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arc=01;31:*.arj=01;31:*.taz=01;31:*.lha=01;31:*.lz4=01;31:*.lzh=01;31:*.lzma=01;31:*.tlz=01;31:*.txz=01;31:*.tzo=01;31:*.t7z=01;31:*.zip=01;31:*.z=01;31:*.dz=01;31:*.gz=01;31:*.lrz=01;31:*.lz=01;31:*.lzo=01;31:*.xz=01;31:*.zst=01;31:*.tzst=01;31:*.bz2=01;31:*.bz=01;31:*.tbz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.war=01;31:*.ear=01;31:*.sar=01;31:*.rar=01;31:*.alz=01;31:*.ace=01;31:*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.cab=01;31:*.wim=01;31:*.swm=01;31:*.dwm=01;31:*.esd=01;31:*.jpg=01;35:*.jpeg=01;35:*.mjpg=01;35:*.mjpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.svg=01;35:*.svgz=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.webm=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:*.avi=01;35:*.fli=01;35:*.flv=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:*.cgm=01;35:*.emf=01;35:*.ogv=01;35:*.ogx=01;35:*.aac=00;36:*.au=00;36:*.flac=00;36:*.m4a=00;36:*.mid=00;36:*.midi=00;36:*.mka=00;36:*.mp3=00;36:*.mpc=00;36:*.ogg=00;36:*.ra=00;36:*.wav=00;36:*.oga=00;36:*.opus=00;36:*.spx=00;36:*.xspf=00;36:
    SSH_CONNECTION=10.0.2.2 41938 10.0.2.15 22
    LESSCLOSE=/usr/bin/lesspipe %s %s
    XDG_SESSION_CLASS=user
    TERM=xterm-256color
    LESSOPEN=| /usr/bin/lesspipe %s
    USER=vagrant
    SHLVL=1
    XDG_SESSION_ID=8
    XDG_RUNTIME_DIR=/run/user/1000
    SSH_CLIENT=10.0.2.2 41938 22
    XDG_DATA_DIRS=/usr/local/share:/usr/share:/var/lib/snapd/desktop
    PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin
    DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus
    SSH_TTY=/dev/pts/0
    _=/usr/bin/env
    OLDPWD=/home/vagrant/testfolder

### 10. Используя man, опишите что доступно по адресам /proc/<PID>/cmdline, /proc/<PID>/exe.

    /proc/[pid]/cmdline -  Файл только для чтения, содержит полную командную
                           строку для процесса, кроме зомби-процессов. 
                           Аргументы командной строки записаны как строки.

    /proc/[pid]/exe -      Это символическая ссылка, содержащая полный путь 
                           к исполняемой программе.


### 11. Узнайте, какую наиболее старшую версию набора инструкций SSE поддерживает ваш процессор с помощью /proc/cpuinfo

 Ответ: `sse4.2`

    $ grep -i SSE /proc/cpuinfo 
    flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni ssse3 cx16 sse4_1 sse4_2 x2apic popcnt hypervisor lahf_lm pti flush_l1d
    flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni ssse3 cx16 sse4_1 sse4_2 x2apic popcnt hypervisor lahf_lm pti flush_l1d

### 12. При открытии нового окна терминала и vagrant ssh создается новая сессия и выделяется pty. Это можно подтвердить командой tty, которая упоминалась в лекции 3.2. Однако:

    vagrant@netology1:~$ ssh localhost 'tty'
    not a tty
### Почитайте, почему так происходит, и как изменить поведение.

 Выполнение команд на удалённой машине через ssh происходит без терминала:

    vagrant@vagrant:~$ ssh  localhost 'tmux'
    open terminal failed: not a terminal

    $ ssh  localhost 'screen'
    Must be connected to a terminal.


 Если для работы программы нужен терминал, то используем опцию -t

    $ ssh -t localhost 'tty'
    /dev/pts/1

 ### 13. Бывает, что есть необходимость переместить запущенный процесс из одной сессии в другую. Попробуйте сделать это, воспользовавшись `reptyr`. Например, так можно перенести в screen процесс, который вы запустили по ошибке в обычной SSH-сессии.

 В первой консоли запускаем top, во второй консоли запускаем screen.
 Пробуем переместить процесс top (PID 1889) в screen.
 Получаем ошибку:

    $ reptyr 1889
    Unable to attach to pid 1889: Operation not permitted
    The kernel denied permission while attaching. If your uid matches
    the target's, check the value of /proc/sys/kernel/yama/ptrace_scope.
    For more information, see /etc/sysctl.d/10-ptrace.conf

 Смотрим `/etc/sysctl.d/10-ptrace.conf` , а там с мотивировкой "PTRACE не нужен на обычной Убунте" kernel.yama.ptrace_scope выставлена в 1. Т.е. разрешён только дочерних процессов, поэтому делаем:

    $ sudo sysctl kernel.yama.ptrace_scope=0
    kernel.yama.ptrace_scope = 0

 Запускаем top, его PID теперь 1960. Возвращаемся во вторую консоль в screen и пробуем снова:

    $reptyr 1960

 Теперь у нас top работает в screen

### 14. `sudo echo string > /root/new_file` не даст выполнить перенаправление под обычным пользователем, так как перенаправлением занимается процесс shell'а, который запущен без sudo под вашим пользователем. Для решения данной проблемы можно использовать конструкцию `echo string | sudo tee /root/new_file`. Узнайте что делает команда `tee` и почему в отличие от `sudo echo` команда с `sudo tee` будет работать.

`tee` отправляет данные со стандартного ввода в файл и стандартный вывод. 

`echo string | sudo tee /root/new_file` срабатывает без ошибки, потому-что sudo применяется к `tee`, прав для записи хватает.

Ещё по-умолчанию `sudo` запрашивает пароль с терминала пользователя, а не со стандартного ввода, поэтому string прозрачно передаётся `tee`.


