import asyncio
import logging

logging.basicConfig(
    level=logging.INFO, format="%(asctime)s %(name)-12s %(levelname)-8s %(message)s", datefmt="%y-%m-%d %H:%M:%S"
)


async def foo1():
    logging.info("Started foo1")
    await asyncio.sleep(2)
    logging.info("Ended foo1")
    return 2


async def foo2(dummy):
    logging.info("Started foo2")
    await asyncio.sleep(dummy)
    logging.info("Ended foo2")
    return dummy


async def foo3(dummy):
    logging.info("Started foo3")
    await asyncio.sleep(dummy)
    1 / 0
    logging.info("Started foo3")
    return 0


async def multiple_tasks(dummy):
    input_coroutines = [foo1(), foo2(dummy), foo2(dummy * 2), foo3(dummy)]
    res = await asyncio.gather(*input_coroutines, return_exceptions=True)
    return res


if __name__ == "__main__":
    dummy = 5
    result = asyncio.get_event_loop().run_until_complete(multiple_tasks(dummy))
    logging.info(result)
