# Tortoise
- [Туториал](https://tortoise-orm.readthedocs.io/en/latest/)
- [Гитхаб](https://github.com/tortoise/tortoise-orm)

# Aertich(ПО для миграции)
[полный туториал](https://tortoise-orm.readthedocs.io/en/latest/migration.html)
```
aerich init -t settings.TORTOISE_ORM # Инициализация конфига
aerich init-db # Инициализация ДБ
aerich migrate --name drop_column # Миграция в drop_column
aerich upgrade # Миграция в latest
```
> Aerich не поддерижвает sqlite3 адекватно!

# Структура
```
├── __init__.py
├── aerich.ini # Файл инициализации aerich, автогенериться
├── data
│  └── db.sqlite3 # ДБ
├── init_db.py # Инициализация дб
├── migrations # Файлы миграции
│  └── models
│     └── 0_20210516200403_init.sql
├── README.md
├── requirements.txt
├── settings.py # Настройки тортойза
├── tournaments.py # Модель, если храниться в папке, записывать в settings через точку
└── worker_db.py # рабочий
```

# Extra
Лучший сайт для рисования схем ДБ: dbdiagram.io
