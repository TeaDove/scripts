import configparser
from pathlib import Path

from pyrogram import Client, filters, types


class Bot:
    def __init__(self, folder_name=Path("secret_data")):
        self.folder_name = folder_name
        self.config = configparser.ConfigParser()
        self.config.read(folder_name / "config.ini")

    def start(self):
        app = Client(
            session_name=str(self.folder_name / "my_account"),
            api_id=self.config["credentials"]["pyrogram_api_id"],
            api_hash=self.config["credentials"]["pyrogram_api_hash"],
        )

        app.start()
        self.my_id = app.get_users("me").id
        app.stop()

        @app.on_message(filters.text)
        def my_handler(client, message: types.Message):
            message.reply("Hi!")

        app.run()


if __name__ == "__main__":
    my_bot = Bot()
    my_bot.start()
