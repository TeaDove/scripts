# Lifehacks
### Пробелы в табы:
``` bash
sed -e 's/    /\t/g' main.py > tmp.py
```

### Смена рабочей директории на директорию скрипта<br>
``` python
from pathlib import Path
import os

SCRIPT = Path(os.path.realpath(__file__))
os.chdir(SCRIPT.parent)
```
### Создание requirements для проекта
А лучше юзай poetry и `python3 -m venv venv`
``` bash
pip3 install pipreqs && pipreqs --print
```
Добавление в poetry из requirements
``` bash
cat requirements.txt|xargs poetry add
```
### Настройка логера
``` python
import logging

logging.basicConfig(level=logging.WARNING, format='%(asctime)s %(name)-12s %(levelname)-8s %(message)s', datefmt='%y-%m-%d %H:%M:%S')

# Или просто:
# pip3 install loguru
from loguru import logger
```
### Мониторинг прохождения в цикле
`pip3 install tqdm`
``` python
sum_ = 0
for i in tqdm(range(10000000)):
    sum_+=1
print(sum_)
```
### Кодировка utf-8(вставить в начало файла)
```
# -*- coding: utf-8 -*-
```
### Классовые декораторы, [pep8](https://www.python.org/dev/peps/pep-0008/), [pep257](https://www.python.org/dev/peps/pep-0257/), приватные методы и тд
см в [python/general/README.md](https://gitlab.com/TeaDove/scripts/-/tree/master/python/general/)
### venv(для fish)
``` bash
virtualenv .venv -p=(which python3)  # Создание .venv с python3
source .venv/bin/activate.fish  # Активация .venv
deactivate # Деактивация .venv
pip install --upgrade pip # Обновление pip'a
```
### Профайлеры:
- filprofiler
- memory_profiler
- cProfile
### Словарь полей
``` python
fields_dict = {field: getattr(self, field) for field in dir(self) if not field.startswith('_')}
str(fields_dict)
```
