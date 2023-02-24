import settings
from tortoise import Tortoise, run_async


# ! Инициализация ДБ, лучше через aerich!
# ! Aerich не поддерживает sqlite3!
async def init():
    await Tortoise.init(config=settings.TORTOISE_ORM)
    await Tortoise.generate_schemas()


if __name__ == "__main__":
    run_async(init())
