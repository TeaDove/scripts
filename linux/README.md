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
ssh -R <откуда-адресс>:<откуда-порт>:<куда-адресс>:<куда-порт> -N <юзер сервера>@<адресс сервера>
```
Запускать на машине Х
Будет проксировать запросы из машины У `<откуда-адресс>:<откуда-порт>` в машину Х `<куда-адресс>:<куда-порт>`

Не забыть вставить в /etc/ssh/sshd_config
```shell
AllowTcpForwarding yes
GatewayPorts yes
X11Forwarding yes
ClientAliveInterval 15
ClientAliveCountMax 4
```
И выполнить
`systemctl restart sshd`
На клиенте в `/etc/ssh/ssh_config`
```shell
ClientAliveInterval 15
ClientAliveCountMax 4
```

Или экспозинг через нат:
`ssh ssh-j.com`

#### Например
``` bash
ssh -R root@116.203.245.151 -N 0.0.0.0:8000:localhost:8000
```
Будет форвардить все запросы в машину (root@116.203.245.151) в клиенскую машину с порта 8000


#### Systemd
```shell
vim /etc/systemd/system/amazing.service

# >
[Unit]
Description=amazing

[Service]
ExecStart=echo "nice"

[Install]
WantedBy=multi-user.target

systemctl daemon-reload
systemctl restart amazing
systemctl status amazing
```

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

Запустить new_service.sh

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

## Nets

### Scan for devices in network

```bash
sudo nmap -sn 192.168.1.0/24
```

### Scan for open ports of this machine

```bash
netstat -tulpn # linux

lsof -PiTCP -sTCP:LISTEN # darwin
```

### Scan for open ports by ip-address

```bash
sudo nmap -n -PN -sT -sU -p- localhost
```

### Resolve domain using specific DNS

```bash
host ya.ru 8.8.8.8
```
