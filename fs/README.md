# Домашнее задание к занятию "3.5. Файловые системы"


###  1. Узнайте о sparse (разряженных) файлах.

Прочитал, осмыслил. В разряженных файлах пустое место не забивается нулями, а информация о нём хранится в метаданных файловой системы.



### 2. Могут ли файлы, являющиеся жесткой ссылкой на один объект, иметь разные права доступа и владельца? Почему?
   
Не могут, т.к.владелец и права доступа выставляются для объекта (файла), а жесткая ссылка является "псевдонимом", т.е. ещё одним именем и не имеет таких аттрибутов. В терминологии ЯП - это указатель (ссылка).



### 3. Сделайте vagrant destroy на имеющийся инстанс Ubuntu. Замените содержимое Vagrantfile. Данная конфигурация создаст новую виртуальную машину с двумя дополнительными неразмеченными дисками по 2.5 Гб.

````
$ vagrant up 
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'bento/ubuntu-20.04'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'bento/ubuntu-20.04' version '202112.19.0' is up to date...
==> default: Setting the name of the VM: vagrant_default_1643383213443_51738
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
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
    default: Guest Additions Version: 6.1.30
    default: VirtualBox Version: 5.2
==> default: Mounting shared folders...
    default: /vagrant => /home/abs/devops/_DEVOPS-15/homeworks/vagrant
````  
Посмотрим, что получилось:
````
vagrant@vagrant:~$ lsblk
NAME                      MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
loop0                       7:0    0 55.4M  1 loop /snap/core18/2128
loop1                       7:1    0 70.3M  1 loop /snap/lxd/21029
loop2                       7:2    0 32.3M  1 loop /snap/snapd/12704
loop3                       7:3    0 55.5M  1 loop /snap/core18/2284
loop4                       7:4    0 43.4M  1 loop /snap/snapd/14549
loop5                       7:5    0 61.9M  1 loop /snap/core20/1328
loop6                       7:6    0 67.2M  1 loop /snap/lxd/21835
sda                         8:0    0   64G  0 disk 
├─sda1                      8:1    0    1M  0 part 
├─sda2                      8:2    0    1G  0 part /boot
└─sda3                      8:3    0   63G  0 part 
  └─ubuntu--vg-ubuntu--lv 253:0    0 31.5G  0 lvm  /
sdb                         8:16   0  2.5G  0 disk 
sdc                         8:32   0  2.5G  0 disk
````
Да, два диска по 2,5Гб присутствуют.



### 4. Используя `fdisk`, разбейте первый диск на 2 раздела: 2 Гб, оставшееся пространство.
````
root@vagrant:~# fdisk /dev/sdb

Welcome to fdisk (util-linux 2.34).                                                                                                                       
Changes will remain in memory only, until you decide to write them.
Be careful before using the write command.


Command (m for help): n
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-5242879, default 2048): 
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-5242879, default 5242879): +2G

Created a new partition 1 of type 'Linux' and of size 2 GiB.

Command (m for help): n
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 
First sector (4196352-5242879, default 4196352): 
Last sector, +/-sectors or +/-size{K,M,G,T,P} (4196352-5242879, default 5242879): 

Created a new partition 2 of type 'Linux' and of size 511 MiB.

Command (m for help): w
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.
````
Получилось так:
````
root@vagrant:~# fdisk -l /dev/sdb
Disk /dev/sdb: 2.51 GiB, 2684354560 bytes, 5242880 sectors
Disk model: VBOX HARDDISK   
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: dos
Disk identifier: 0x0d3d31ef

Device     Boot   Start     End Sectors  Size Id Type
/dev/sdb1          2048 4196351 4194304    2G 83 Linux
/dev/sdb2       4196352 5242879 1046528  511M 83 Linux
````

### 5. Используя `sfdisk`, перенесите данную таблицу разделов на второй диск.

````
root@vagrant:~# sfdisk -d /dev/sdb | sfdisk /dev/sdc
Checking that no-one is using this disk right now ... OK

Disk /dev/sdc: 2.51 GiB, 2684354560 bytes, 5242880 sectors
Disk model: VBOX HARDDISK   
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes

>>> Script header accepted.
>>> Script header accepted.
>>> Script header accepted.
>>> Script header accepted.
>>> Created a new DOS disklabel with disk identifier 0x0d3d31ef.
/dev/sdc1: Created a new partition 1 of type 'Linux' and of size 2 GiB.
/dev/sdc2: Created a new partition 2 of type 'Linux' and of size 511 MiB.
/dev/sdc3: Done.

New situation:
Disklabel type: dos
Disk identifier: 0x0d3d31ef

Device     Boot   Start     End Sectors  Size Id Type
/dev/sdc1          2048 4196351 4194304    2G 83 Linux
/dev/sdc2       4196352 5242879 1046528  511M 83 Linux

The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.
````


### 6. Соберите mdadm RAID1 на паре разделов 2 Гб.

````
root@vagrant:~# mdadm --create --verbose /dev/md0 --level=1 --raid-devices=2 /dev/sdb1 /dev/sdc1
mdadm: Note: this array has metadata at the start and
    may not be suitable as a boot device.  If you plan to
    store '/boot' on this device please ensure that
    your boot-loader understands md/v1.x metadata, or use
    --metadata=0.90
mdadm: size set to 2094080K
Continue creating array? y
mdadm: Defaulting to version 1.2 metadata
mdadm: array /dev/md0 started.
root@vagrant:~# cat /proc/mdstat
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md0 : active raid1 sdc1[1] sdb1[0]
      2094080 blocks super 1.2 [2/2] [UU]
      [==================>..]  resync = 92.4% (1937664/2094080) finish=0.0min speed=215296K/sec
      
unused devices: <none>
````


### 7. Соберите mdadm RAID0 на второй паре маленьких разделов.
````
root@vagrant:~# mdadm --create --verbose /dev/md1 --level=0 --raid-devices=2 /dev/sdb2 /dev/sdc2 
mdadm: chunk size defaults to 512K
mdadm: Defaulting to version 1.2 metadata
mdadm: array /dev/md1 started.
````
````
root@vagrant:~# cat /proc/mdstat
Personalities : [linear] [multipath] [raid0] [raid1] [raid6] [raid5] [raid4] [raid10] 
md1 : active raid0 sdc2[1] sdb2[0]
      1042432 blocks super 1.2 512k chunks
      
md0 : active raid1 sdc1[1] sdb1[0]
      2094080 blocks super 1.2 [2/2] [UU]
      
unused devices: <none>
````
````
root@vagrant:~# mdadm --detail --scan | sudo tee -a /etc/mdadm/mdadm.conf
ARRAY /dev/md0 metadata=1.2 name=vagrant:0 UUID=e75cbe5f:1ba1451b:60b44128:b2bbd3da
ARRAY /dev/md1 metadata=1.2 name=vagrant:1 UUID=5f2b65b8:78f198ac:d2fad029:a0789860
````

### 8. Создайте 2 независимых PV на получившихся md-устройствах.

````
root@vagrant:~# pvcreate /dev/md0
  Physical volume "/dev/md0" successfully created.
root@vagrant:~# pvcreate /dev/md1
  Physical volume "/dev/md1" successfully created.
````  
  
### 9. Создайте общую volume-group на этих двух PV

````
root@vagrant:~# vgcreate first_group /dev/md0 /dev/md1
  Volume group "first_group" successfully created
````
  
  
### 10. Создайте LV размером 100 Мб, указав его расположение на PV с RAID0

````
root@vagrant:~# lvcreate -L 100 -ntestlv first_group /dev/md1
  Logical volume "testlv" created.
````  

В итоге получили вот такую конструкцию:

````
root@vagrant:~# vgdisplay
  --- Volume group ---
  VG Name               ubuntu-vg
  System ID             
  Format                lvm2
  Metadata Areas        1
  Metadata Sequence No  2
  VG Access             read/write
  VG Status             resizable
  MAX LV                0
  Cur LV                1
  Open LV               1
  Max PV                0
  Cur PV                1
  Act PV                1
  VG Size               <63.00 GiB
  PE Size               4.00 MiB
  Total PE              16127
  Alloc PE / Size       8064 / 31.50 GiB
  Free  PE / Size       8063 / <31.50 GiB
  VG UUID               aK7Bd1-JPle-i0h7-5jJa-M60v-WwMk-PFByJ7
   
  --- Volume group ---
  VG Name               first_group
  System ID             
  Format                lvm2
  Metadata Areas        2
  Metadata Sequence No  2
  VG Access             read/write
  VG Status             resizable
  MAX LV                0
  Cur LV                1
  Open LV               0
  Max PV                0
  Cur PV                2
  Act PV                2
  VG Size               <2.99 GiB
  PE Size               4.00 MiB
  Total PE              765
  Alloc PE / Size       25 / 100.00 MiB
  Free  PE / Size       740 / 2.89 GiB
  VG UUID               rN4nlW-Vh9u-gd3n-CZdZ-Gh3F-UR1H-nJE45y
   
root@vagrant:~# lvdisplay
  --- Logical volume ---
  LV Path                /dev/ubuntu-vg/ubuntu-lv
  LV Name                ubuntu-lv
  VG Name                ubuntu-vg
  LV UUID                ftN15m-3lML-YH5x-R5P2-kLCd-kzW3-32dlqO
  LV Write Access        read/write
  LV Creation host, time ubuntu-server, 2021-12-19 19:37:44 +0000
  LV Status              available
  # open                 1
  LV Size                31.50 GiB
  Current LE             8064
  Segments               1
  Allocation             inherit
  Read ahead sectors     auto
  - currently set to     256
  Block device           253:0
   
  --- Logical volume ---
  LV Path                /dev/first_group/testlv
  LV Name                testlv
  VG Name                first_group
  LV UUID                vmkSV8-2J3Q-ESHp-Dlx2-bKVL-U1qq-CQ31xl
  LV Write Access        read/write
  LV Creation host, time vagrant, 2022-01-28 16:41:03 +0000
  LV Status              available
  # open                 0
  LV Size                100.00 MiB
  Current LE             25
  Segments               1
  Allocation             inherit
  Read ahead sectors     auto
  - currently set to     4096
  Block device           253:1
   
root@vagrant:~# pvdisplay
  --- Physical volume ---
  PV Name               /dev/sda3
  VG Name               ubuntu-vg
  PV Size               <63.00 GiB / not usable 0   
  Allocatable           yes 
  PE Size               4.00 MiB
  Total PE              16127
  Free PE               8063
  Allocated PE          8064
  PV UUID               sDUvKe-EtCc-gKuY-ZXTD-1B1d-eh9Q-XldxLf
   
  --- Physical volume ---
  PV Name               /dev/md0
  VG Name               first_group
  PV Size               <2.00 GiB / not usable 0   
  Allocatable           yes 
  PE Size               4.00 MiB
  Total PE              511
  Free PE               511
  Allocated PE          0
  PV UUID               Yqpnkl-kQuR-Hnn3-ejuv-L4Ei-VBNr-iYuCn1
   
  --- Physical volume ---
  PV Name               /dev/md1
  VG Name               first_group
  PV Size               1018.00 MiB / not usable 2.00 MiB
  Allocatable           yes 
  PE Size               4.00 MiB
  Total PE              254
  Free PE               229
  Allocated PE          25
  PV UUID               vANs3f-Xkyp-c72e-P2k8-42P6-cxki-5cOd7V
````
  
### 11. Создайте `mkfs.ext4` ФС на получившемся LV

````
root@vagrant:~# mkfs.ext4 /dev/first_group/testlv
mke2fs 1.45.5 (07-Jan-2020)
Creating filesystem with 25600 4k blocks and 25600 inodes

Allocating group tables: done                            
Writing inode tables: done                            
Creating journal (1024 blocks): done
Writing superblocks and filesystem accounting information: done
````


### 12. Смонтируйте этот раздел в любую директорию, например, /tmp/new
````
root@vagrant:~# mkdir /tmp/new
root@vagrant:~# mount /dev/first_group/testlv /tmp/new
root@vagrant:~# mount | grep lv
/dev/mapper/ubuntu--vg-ubuntu--lv on / type ext4 (rw,relatime)
/dev/mapper/first_group-testlv on /tmp/new type ext4 (rw,relatime,stripe=256)
````

### 13. Поместите туда тестовый файл, например wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz

````
root@vagrant:~# cd /tmp/new
root@vagrant:/tmp/new# wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz
--2022-01-28 16:51:48--  https://mirror.yandex.ru/ubuntu/ls-lR.gz
Resolving mirror.yandex.ru (mirror.yandex.ru)... 213.180.204.183, 2a02:6b8::183
Connecting to mirror.yandex.ru (mirror.yandex.ru)|213.180.204.183|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 22056717 (21M) [application/octet-stream]
Saving to: ‘/tmp/new/test.gz’

/tmp/new/test.gz                       100%[==========================================================================>]  21.03M  8.18MB/s    in 2.6s    

2022-01-28 16:51:51 (8.18 MB/s) - ‘/tmp/new/test.gz’ saved [22056717/22056717]
````

### 14. Прикрепите вывод `lsblk`

````
root@vagrant:/tmp/new# lsblk
NAME                      MAJ:MIN RM  SIZE RO TYPE  MOUNTPOINT
loop0                       7:0    0 55.4M  1 loop  /snap/core18/2128
loop1                       7:1    0 70.3M  1 loop  /snap/lxd/21029
loop2                       7:2    0 32.3M  1 loop  /snap/snapd/12704
loop3                       7:3    0 55.5M  1 loop  /snap/core18/2284
loop4                       7:4    0 43.4M  1 loop  /snap/snapd/14549
loop5                       7:5    0 61.9M  1 loop  /snap/core20/1328
loop6                       7:6    0 67.2M  1 loop  /snap/lxd/21835
sda                         8:0    0   64G  0 disk  
├─sda1                      8:1    0    1M  0 part  
├─sda2                      8:2    0    1G  0 part  /boot
└─sda3                      8:3    0   63G  0 part  
  └─ubuntu--vg-ubuntu--lv 253:0    0 31.5G  0 lvm   /
sdb                         8:16   0  2.5G  0 disk  
├─sdb1                      8:17   0    2G  0 part  
│ └─md0                     9:0    0    2G  0 raid1 
└─sdb2                      8:18   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
    └─first_group-testlv  253:1    0  100M  0 lvm   /tmp/new
sdc                         8:32   0  2.5G  0 disk  
├─sdc1                      8:33   0    2G  0 part  
│ └─md0                     9:0    0    2G  0 raid1 
└─sdc2                      8:34   0  511M  0 part  
  └─md1                     9:1    0 1018M  0 raid0 
    └─first_group-testlv  253:1    0  100M  0 lvm   /tmp/new
    
````

### 15. Протестируйте целостность файла

````
root@vagrant:/tmp/new# gzip -t /tmp/new/test.gz
root@vagrant:/tmp/new# echo $?
0
````

### 16. Используя `pvmove`, переместите содержимое PV с RAID0 на RAID1

````
root@vagrant:/tmp/new# pvmove -b /dev/md1 /dev/md0
root@vagrant:/tmp/new# pvmove -i5 /dev/md0
  /dev/md0: Moved: 16.00%
  /dev/md0: Moved: 100.00%
````

### 17. Сделайте --fail на устройство в вашем RAID1 md

````
root@vagrant:/tmp/new# mdadm /dev/md0 --fail /dev/sdb1
mdadm: set /dev/sdb1 faulty in /dev/md0
````


### 18. Подтвердите выводом `dmesg`, что RAID1 работает в деградированном состоянии.

````
[ 6005.918356] md/raid1:md0: Disk failure on sdb1, disabling device.
               md/raid1:md0: Operation continuing on 1 devices.
````
               
### 19. Протестируйте целостность файла, несмотря на "сбойный" диск он должен продолжать быть доступен
               
````
root@vagrant:/tmp/new# gzip -t /tmp/new/test.gz; echo $?
0          
````

### 20. Погасите тестовый хост, vagrant destroy

````
$ vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Forcing shutdown of VM...
==> default: Destroying VM and associated drives...
````
               
