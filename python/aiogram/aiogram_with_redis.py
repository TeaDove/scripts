# quickstart для Aiogram

import configparser
import logging
import os
from pathlib import Path

import redis
from aiogram import Bot, Dispatcher, executor, types, utils

BASE = Path(os.path.realpath(__file__))
os.chdir(BASE.parent)


config = configparser.ConfigParser()
config.read("secret_data/config.ini")
logging.basicConfig(
    level=logging.WARNING, format="%(asctime)s %(name)-12s %(levelname)-8s %(message)s", datefmt="%y-%m-%d %H:%M:%S"
)
bot = Bot(token=config["credentials"]["telegram-api"])
dp = Dispatcher(bot)

r = redis.Redis(host="localhost", port=10001)
APP_NAME = config["settings"]["app_name"]
RES_DICT = utils.get_res_dict()


async def get_chat_dict(chat_id: int) -> dict:
    chat_dict = r.hgetall(f"{NAME}:{chat_id}")  # noqa: F821
    default_dict = {b"a": 1, b"b": 2}
    if not chat_dict:
        r.hset(f"{APP_NAME}:{chat_id}", mapping=default_dict)
        chat_dict = default_dict
    elif set(chat_dict.keys()) != set(default_dict.keys()):
        r.delete(f"{APP_NAME}:{chat_id}")
        r.hset(f"{APP_NAME}:{chat_id}", mapping=default_dict)
    return chat_dict


@dp.message_handler()
async def send(message: types.Message):
    await message.reply(RES_DICT["start"], parse_mode="html")


if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)
