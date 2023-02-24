"""
Лучше использовать ассинхронность!
"""

from multiprocessing import Pool
from pathlib import Path

import requests


def job(list_to_pool, check_doubles: bool = True):
    idx, url, folder = list_to_pool
    file_name = str(idx) + ".png"
    if not check_doubles or file_name not in list(folder.iterdir()):
        try:
            a = requests.get(url)
        except Exception:
            print(f"{file_name}\tError: download error!")  # noqa: T201
        else:
            open(folder / file_name, "wb").write(a.content)  # noqa: SIM115, SCS109
            print(f"{file_name}\tDownloaded!")  # noqa: T201
    else:
        print(f"{file_name}\tError: already downloaded!")  # noqa: T201


def download_pool(urls: list, folder=Path() / "data"):
    list_to_pool = [[idx, url, folder] for idx, url in enumerate(urls)]

    pool = Pool()
    pool.map(job, list_to_pool)
