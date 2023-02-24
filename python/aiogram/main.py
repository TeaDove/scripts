# quickstart для Aiogram

import configparser
import logging

from aiogram import Bot, Dispatcher, executor, types

config = configparser.ConfigParser()
config.read("secret_data/config.ini")
logging.basicConfig(
    level=logging.WARNING,
    format="%(asctime)s %(name)-12s %(levelname)-8s %(message)s",
    datefmt="%y-%m-%d %H:%M:%S",
)
bot = Bot(token=config["credentials"]["telegram-api"])
dp = Dispatcher(bot)


@dp.message_handler()
async def send(message: types.Message):
    await message.reply(RES_DICT["start"], parse_mode="html")  # noqa: F821


if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)
