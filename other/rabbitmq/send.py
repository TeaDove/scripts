#!/usr/bin/env python
import sys

import pika

with pika.BlockingConnection(pika.ConnectionParameters("localhost")) as connection:
    channel = connection.channel()
    channel.queue_declare(queue="hello")
    message = " ".join(sys.argv[1:]) or "Hello World!"
    channel.basic_publish(exchange="", routing_key="hello", body=message)
    print(" [x] Sent %r" % (message,))  # noqa: T201
