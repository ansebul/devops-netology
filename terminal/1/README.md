# Домашнее задание к занятию «3.1. Работа в терминале, лекция 1»


###  1. Установлено средство виртуализации Oracle VirtualBox

    $ virtualbox -h | head -n3
    Oracle VM VirtualBox Manager 5.2.42_Ubuntu
    (C) 2005-2020 Oracle Corporation
    All rights reserved.



### 2. Установлено средство автоматизации Hashicorp Vagrant.


    $ whereis vagrant
    vagrant: /usr/bin/vagrant /opt/vagrant/bin/vagrant
    $ vagrant -v
    Vagrant 2.2.19


### 3. Подготовлен терминал

    Konsole
    Версия 17.12.3


### 4. С помощью базового файла конфигурации запущен Ubuntu 20.04 в VirtualBox посредством Vagrant

   Создана директория, в которой будут храниться конфигурационные файлы Vagrant. В ней выполнена команда `vagrant init` :

    $ cd devops/homeworks/
    $ mkdir vagrant
    $ cd vagrant/
    $ vagrant init
    A `Vagrantfile` has been placed in this directory. You are now
    ready to `vagrant up` your first virtual environment! Please read
    the comments in the Vagrantfile as well as documentation on
    `vagrantup.com` for more information on using Vagrant.


 Выполнена в этой же директории `vagrant up` :

    $ vagrant up
    Bringing machine 'default' up with 'virtualbox' provider...
    ==> default: Box 'bento/ubuntu-20.04' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
    ==> default: Loading metadata for box 'bento/ubuntu-20.04'
    default: URL: https://vagrantcloud.com/bento/ubuntu-20.04
    ==> default: Adding box 'bento/ubuntu-20.04' (v202112.19.0) for provider: virtualbox
    default: Downloading: https://vagrantcloud.com/bento/boxes/ubuntu-20.04/versions/202112.19.0/providers/virtualbox.box
    ==> default: Successfully added box 'bento/ubuntu-20.04' (v202112.19.0) for 'virtualbox'!
    ==> default: Importing base box 'bento/ubuntu-20.04'...
    ==> default: Matching MAC address for NAT networking...
    ==> default: Checking if box 'bento/ubuntu-20.04' version '202112.19.0' is up to date...
    ==> default: Setting the name of the VM: vagrant_default_1636633792727_76752
    Vagrant is currently configured to create VirtualBox synced folders with
    the `SharedFoldersEnableSymlinksCreate` option enabled. If the Vagrant
    guest is not trusted, you may want to disable this option. For more
    information on this option, please refer to the VirtualBox manual:

    https://www.virtualbox.org/manual/ch04.html#sharedfolders

    This option can be disabled globally with an environment variable:

    VAGRANT_DISABLE_VBOXSYMLINKCREATE=1

    or on a per folder basis within the Vagrantfile:

    config.vm.synced_folder '/host/path', '/guest/path', SharedFoldersEnableSymlinksCreate: false
    ==> default: Clearing any previously set network interfaces...
    ==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    ==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
    ==> default: Booting VM...
    ==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
    ==> default: Machine booted and ready!
    ==> default: Checking for guest additions in VM...
    default: The guest additions on this VM do not match the installed version of
    default: VirtualBox! In most cases this is fine, but in rare cases it can
    default: prevent things such as shared folders from working properly. If you see
    default: shared folder errors, please make sure the guest additions within the
    default: virtual machine match the version of VirtualBox you have installed on
    default: your host and reload your VM.
    default: 
    default: Guest Additions Version: 6.1.24
    default: VirtualBox Version: 5.2
    ==> default: Mounting shared folders...
    default: /vagrant => /home/abs/devops/homeworks/vagrant

 Попробовал обе команды: `vagrant suspend` и `vagrant halt`

### 5. Какие ресурсы выделены по-умолчанию?


    Память: 1024Мб
    2 CPU
    64 Гб HDD
    Сетевая карта INTEL


### 6. Как добавить оперативной памяти или ресурсов процессора виртуальной машине?

 Прописать в файле Vagrantfile необходимые параметры, например:
 
    config.vm.provider "virtualbox" do |v|
      v.memory = 16384
      v.cpus = 64
    end


### 7. С помощью `vagrant ssh` из директории, в которой содержится Vagrantfile, получен доступ в виртуальную машину

### 8. Ознакомился с разделами man bash, почитал о настройках самого bash

  Какой переменной можно задать длину журнала `history`, и на какой строчке manual это описывается?

    Длинну журнала истории команд, отображаемую через history,
    можно задать через переменную окружения HISTSIZE (man line 846),
    но также следует учитывать переменную окружения HISTFILESIZE (line 862)
    в которой задаётся максимальное число строк, 
    содержащихся в файле истории. 
    По-умолчанию в HISTFILESIZE = 2000, а HISTSIZE = 1000.

  Что делает директива `ignoreboth` в bash?

    ignoreboth - заменяет собой две директивы:
    ignorespace  и  ignoredups и предписывает не сохранять
    в истории последующие повторяющиеся команды и команды,
    предварённые пробелом.


### 9. В каких сценариях использования применимы скобки {} и на какой строчке man bash это описано?

    списки  { list; } сторка 257
    функции, строка 403
    фича Баша - раскрытие скобок, строка 1090
    раскрытие переменных, строка 1167

### 10. Как создать однократным вызовом touch 100000 файлов? А получилось ли создать 300000? Если нет, то почему?

 Вот так: `touch f{1..100000}`

    $ mkdir newfolder
    $ cd newfolder
    $ touch f{1..100000}
    $ ls -1 | wc -l
    100000
    $ touch f{1..300000}
    -bash: /usr/bin/touch: Argument list too long

### 11. Что делает конструкция `[[ -d /tmp ]]`
 
  Проверяет наличие каталога /tmp
  Эта запись аналогична `test`

### 12. Добейтесь в выводе `type -a` bash в виртуальной машине наличия первым пунктом в списке: `/bash is /tmp/new_path_directory/bash`

 Это можно сделать, например, так:

    $ mkdir /tmp/new_path_directory
    $ ln -s /bin/bash /tmp/new_path_directory/
    $ export PATH="/tmp/new_path_directory:${PATH}"
    $ echo $PATH
    /tmp/new_path_directory:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin

    $ type -a bash
    bash is /tmp/new_path_directory/bash
    bash is /usr/bin/bash
    bash is /bin/bash

 ### 13. Чем отличается планирование команд с помощью `batch` и `at`?

 `at` - выполнение команды в заданное время 

 `batch` - выполнение команды, когда позволяет нагрузка (system load level)

### 14. Завершена работа виртуальной машины через `vagrant halt`

