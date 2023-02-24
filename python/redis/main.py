# quickstart редиса

import os
from pathlib import Path

import redis

BASE = Path(os.path.realpath(__file__))
os.chdir(BASE.parent)

r = redis.Redis(host="localhost", port=10001)

default_dict = {b"a": 1, b"b": 2}
r.hset("0", mapping=default_dict)
print(r.hget("0", "a"))  # noqa: T201
print(r.hget("0", "b"))  # noqa: T201

r.set("hello", "world")
print(r.get("hello"))  # noqa: T201
