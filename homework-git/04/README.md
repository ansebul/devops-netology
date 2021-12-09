# Домашнее задание к занятию «2.4. Инструменты Git»

Для выполнения задания был клонирован репозиторий 
терраформа https://github.com/hashicorp/terraform 

    $ git clone https://github.com/hashicorp/terraform.git terraform
    Клонирование в «terraform»…
    remote: Enumerating objects: 253270, done.
    remote: Counting objects: 100% (416/416), done.
    remote: Compressing objects: 100% (163/163), done.
    remote: Total 253270 (delta 280), reused 341 (delta 248), pack-reused 252854
    Получение объектов: 100% (253270/253270), 186.12 MiB | 1.69 MiB/s, готово.
    Определение изменений: 100% (157048/157048), готово.`


В результате работы с репозиторием были получены ответы на вопросы: 

###  1. Найдите полный хеш и комментарий коммита, хеш которого начинается на `aefea`.
    
    Полный хэш: aefead2207ef7e2aa5dc81a34aedf0cad4c32545
    Текст комментария: Update CHANGELOG.md

  Ответ был получен 2-мя способами:

  1)

    $ git log aefea -n1
    commit aefead2207ef7e2aa5dc81a34aedf0cad4c32545
    Author: Alisdair McDiarmid <alisdair@users.noreply.github.com>
    Date:   Thu Jun 18 10:29:58 2020 -0400

    Update CHANGELOG.md


  2)
    
    $ git show --quiet aefea
    commit aefead2207ef7e2aa5dc81a34aedf0cad4c32545
    Author: Alisdair McDiarmid <alisdair@users.noreply.github.com>
    Date:   Thu Jun 18 10:29:58 2020 -0400

    Update CHANGELOG.md


### 2. Какому тегу соответствует коммит `85024d3`?

  Ответ:

    v0.12.23
 
  Ответ был получен 2-мя способами:
 
  1)

    $ git log 85024d3 --oneline -n1
    85024d310 (tag: v0.12.23) v0.12.23

  2)

    $ git show --quiet 85024d3
    commit 85024d3100126de36331c6982bfaac02cdab9e76 (tag: v0.12.23)
    Author: tf-release-bot <terraform@hashicorp.com>
    Date:   Thu Mar 5 20:56:10 2020 +0000

    v0.12.23
    
 
### 3. Сколько родителей у коммита `b8d720`? Напишите их хеши.

  Ответ:

    2 родителя. Хеши:
    56cd7859e05c36c06b56d013b55a252d0bb7e158
    9ea88f22fc6269854151c571162c5bcf958bee2b

  Ответ был получен 3-мя способами:
 
  1)
  
    $ git log --pretty=%P -n1 b8d720
    56cd7859e05c36c06b56d013b55a252d0bb7e158 9ea88f22fc6269854151c571162c5bcf958bee2b

  2)
 
    $ git show --oneline --quiet b8d720^@
    56cd7859e Merge pull request #23857 from hashicorp/cgriggs01-stable
    9ea88f22f add/update community provider listings
 
 
  3)
    $ git show --oneline --quiet b8d720^1
    56cd7859e Merge pull request #23857 from hashicorp/cgriggs01-stable

    $ git show --oneline --quiet b8d720^2
    9ea88f22f add/update community provider listings


### 4. Перечислите хеши и комментарии всех коммитов которые были сделаны между тегами  v0.12.23 и v0.12.24.

 Ответ:
    
    b14b74c49 [Website] vmc provider links
    3f235065b Update CHANGELOG.md
    6ae64e247 registry: Fix panic when server is unreachable
    5c619ca1b website: Remove links to the getting started guide's old location
    06275647e Update CHANGELOG.md
    d5f9411f5 command: Fix bug when using terraform login on Windows
    4b6d06cc5 Update CHANGELOG.md
    dd01a3507 Update CHANGELOG.md
    225466bc3 Cleanup after v0.12.23 release
 
 Ответ был получен:
  
    $ git log --oneline v0.12.23..v0.12.24^


### 5. Найдите коммит в котором была создана функция `func providerSource`, ее определение в коде выглядит так `func providerSource(...)` (вместо троеточего перечислены аргументы).

 Ответ:

    commit 8c928e83589d90a031f811fae52a81be7153e82f
    Author: Martin Atkins <mart@degeneration.co.uk>
    Date:   Thu Apr 2 18:04:39 2020 -0700


 Ответ был получен:

    git log  -S "func providerSource(" --patch



### 6. Найдите все коммиты в которых была изменена функция `globalPluginDirs`.

Ответ:

    commit 78b12205587fe839f10d946ea3fdc06719decb05
    Author: Pam Selle <204372+pselle@users.noreply.github.com>
    Date:   Mon Jan 13 16:50:05 2020 -0500

    Remove config.go and update things using its aliases
    --
    commit 52dbf94834cb970b510f2fba853a5b49ad9b1a46
    Author: James Bardin <j.bardin@gmail.com>
    Date:   Wed Aug 9 17:46:49 2017 -0400

    keep .terraform.d/plugins for discovery
    --
    commit 41ab0aef7a0fe030e84018973a64135b11abcd70
    Author: James Bardin <j.bardin@gmail.com>
    Date:   Wed Aug 9 10:34:11 2017 -0400

    Add missing OS_ARCH dir to global plugin paths
    --
    commit 66ebff90cdfaa6938f26f908c7ebad8d547fea17
    Author: James Bardin <j.bardin@gmail.com>
    Date:   Wed May 3 22:24:51 2017 -0400

    move some more plugin search path logic to command
    --
    commit 8364383c359a6b738a436d1b7745ccdce178df47
    Author: Martin Atkins <mart@degeneration.co.uk>
    Date:   Thu Apr 13 18:05:58 2017 -0700

    Push plugin discovery down into command package


Ответ был получен:

Установлен git version 2.17.1, поэтому:

    $ git log -L:globalPluginDirs:plugins.go | grep 'commit \w' -A 4

Для git version 2.22 и выше можно:

    $ git log -s -L:globalPluginDirs:plugins.go 



### 7. Кто автор функции `synchronizedWriters`? 

  Ответ: 

    Martin Atkins <mart@degeneration.co.uk>

    commit 5ac311e2a91e381e2f52234668b49ba670aa0fe5
    Author: Martin Atkins <mart@degeneration.co.uk>
    Date:   Wed May 3 16:25:41 2017 -0700

  Ответ был получен:

    $ git log  -S "func synchronizedWriters(" --patch
