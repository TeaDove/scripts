# Установка
1. Установить зависимости
```
apt install monit
apt install apache2-utils # Для htpasswd
```
2. Скопировать конфиги
```
cp config/monitrc /etc/monit
cp config/make_pem.sh /etc/monit
cp config/monit_alert.py /etc/monit
```
3. В `/etc/monit`
```
./make_pem.sh <домен сайта, например tesseract.club>
```
В monitrc поменять конфигурацию на нужную и создать `htpasswd`:
```
htpasswd -n monit > htpasswd
```
В `monit_alert.py` добавить token бота и chat id
## Добавление нового проекта:
```
./add_app.py <расположение проекта>
```
В папке проекта изменить `start.sh` чтобы он запускал проект
```
monit reload
```
# Monit
Конфиги и шаблоны для монита<br>
`config` - шаблон конфигурации<br>
`app-example` - пример приложения

# app-example
```
├── main.py  - fastapi сервер(для примера)
├── monitor.sh  - скрипт для запуска сервера через монит
├── out.log - логи
├── README.md - этот ридми
└── start.sh - запуск uvicorn(для пример)
```
## monitor.sh
Использование: <br>
`monitor.sh process_name {start|stop}`<br>
start - запуск процесса<br>
stop - остановка процесса
## start.sh
В start.sh положите файл по шаблону:<br>
```
#!/bin/bash
exec <команда для запуска процесса>
```
# Config
Положить в /etc/monit
```
├── conf.d/ - конфиги приложения
├── make_pem.sh - скрипт для компиляции ssl сертификата от letsencrypt
├── monit_alert.py - уведомления через телеграм
└── monitrc - основной конфиг
```
Вместо `alert` используйте `exec "/etc/monit/monit_alert.sh"`
## conf.d
`test-uvicorn.conf` - пример конфигурации
`make_conf.sh` - генерация конфига, в аргументе указать локацию проекта, например:<br>
`./make_conf.sh /home/teadove/Documents/projects/fast-helper`

# Дополнительно
[Статья](https://habr.com/en/post/73506/) на хабре о конфигурации и настройке
