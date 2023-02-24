import asyncio
import time

from loguru import logger

RPC = 100


async def handle(j: int):
    # logger.debug(j)
    ...


def run_with_rpc(items):
    indexes = list(range(len(items) // RPC + 1))
    batches = [items[i * RPC : (i + 1) * RPC] for i in indexes]
    loop = asyncio.get_event_loop()
    len()
    for idx, i in enumerate(batches):
        start = time.perf_counter()
        group = asyncio.gather(*[handle(j) for j in i])
        loop.run_until_complete(group)
        # if log_counter >
        logger.info(idx)
        dev = time.perf_counter() - start
        if dev < 1:
            time.sleep(1 - dev)


def main():
    items = list(range(1_000_000))
    run_with_rpc(items)


if __name__ == "__main__":
    main()
