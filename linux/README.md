# Lifehacks
### Смена рабочей директории на директорию скрипта
``` bash
cd "$(dirname "$(readlink -f "${BASH_SOURCE[0]}")")"
```
### Информация по имени процесса
(лучше использовать более удобную [версию](https://gitlab.com/TeaDove/dotfiles "ps aux | head -n 1 & ps aux | grep -v grep --color=auto | grep $argv"), которую можно найти [тут](https://gitlab.com/TeaDove/dotfiles) под "Алиасы и функции в fish")
`ps aux | grep <proc name>`
#### Например
``` bash
root@TechnoTesseract$ ps aux | grep ssh
root       731  0.0  0.2  15852  4608 ?        Ss   Jan24   1:25 /usr/sbin/sshd -D
root     25403  0.0  0.4  16900  8328 ?        Ss   20:11   0:00 sshd: root@pts/0
root     26335  0.0  0.0   6144   888 pts/0    S+   20:18   0:00 grep --color=auto ssh
```
### Порт-форвардинг через ssh(не очень стабильный)
#### TCP
Более качественный [туториал](https://robotmoon.com/ssh-tunnels/)
``` bash
ssh -f -N <юзер сервера>@<адресс сервера> -L <откуда-адресс>:<откуда-порт>:<куда-адресс>:<куда-порт>
```
Или экспозинг через нат:
`ssh ssh-j.com`
#### Например
``` bash
ssh -f -N root@116.203.245.151 -L tesseract.club:8002:tesseract.club:8000
```
Будет форвардить с tesseract.club:8002 на tesseract.club:8000, настройка производилась для сервера 116.203.245.151 с юзером root. Для удаления, найдите процесс через ps aux и убейте.
#### UDP
``` bash
mkfifo pipe_for_udp
nc -ul <адрес отправителя> <порт отправителя> < pipe_for_udp | nc -u <адрес получателя> <порт получается> > fifo_test
```
#### Например
``` bash
mkfifo pipe_for_udp
nc -ul localhost 9999 < pipe_for_udp | nc -u localhost 10000 > pipe_for_udp
# Подключение:
nc -u localhost 9999
nc -lu localhost 10000
```
Будет форврадить из localhost:9999 в localhost:10000, ещё можно закинуть в pipe_for_udp какую-нибудь строко, это помогает ошибки исправить
### Новый сервер
``` bash
# обновление
sudo apt update
sudo apt upgrade
# Питон, тмукс, фиш, htop и другие утилы
sudo apt install python3 python3-pip python3-dev python3-setuptools htop neofetch tmux fish git curl wget vim mc
pip3 install setuptools
# gcc
sudo apt install build-essential
# супервизор
sudo apt install supervisor
# Включить backport репу
printf "%s\n" "deb http://ftp.de.debian.org/debian buster-backports main" | tee /etc/apt/sources.list.d/buster-backports.list
# Или сразу всё:
curl https://gitlab.com/TeaDove/scripts/-/raw/master/linux/new_server.sh | bash
```
### Vim и терминальные шоркаты
Vim<br>
`:W` - запись через sudo <br>
`? или /` - поиск по файлу <br>
`N/shift + N` - следующий/предыдущий в поиске <br>
`:<num> / <num>gg` - прыгнуть к строке с номером `<num>`<br>
Terminal<br>
`ctrl + U/K` - вырезать до конца / до начала <br>
`ctrl + W` - вырезать слово <br>
`ctrl + Y` - вставить <br>
`ctrl + A/E` - прыгнуть в начало/конец строки <br>
### Мониторинг сервера
- #### открытые tcp порты
`lsof -i -P -n | grep LISTEN`
- #### **неудачные** попытки входа по ssh
`grep "Failed password" /var/log/auth.log`
- #### **удачные** попытки входа по ssh(как по RSA, так и по паролю)
`grep "Accepted" /var/log/auth.log`
- #### последние логины
`last -a`
- #### аналогично, но **не**удачные
`lastb`
- #### текущие юзеры
`w`
### Текст шеринг
`echo just testing!  | nc termbin.com 9999`<br>
Удобнее всего пихнуть в alias, как в моём [конфиге](https://gitlab.com/TeaDove/dotfiles/-/blob/master/.config/fish/config.fish)
### Поиск устройств в сети
```fping -agq <ip>/<mask>```<br>
Например<br>
```fping -agq 192.168.1.1/24``` - пропингует всех в 192.168.1.1/24 и выдаст доступные айпи адреса, что ответили на пинг.
### Docker
``` bash
docker ps -a --size # все контейнейры
docker image prune -a # удаление картинок без привязанных контейнеров
docker system df -v
```
