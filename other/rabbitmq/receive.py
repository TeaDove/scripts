#!/usr/bin/env python
import time

import pika

with pika.BlockingConnection(pika.ConnectionParameters("localhost")) as connection:
    channel = connection.channel()

    def callback(ch, method, properties, body):
        print(" [x] Received %r" % (body,))  # noqa: T201
        time.sleep(body.count(b"."))
        print(" [x] Done")  # noqa: T201
        ch.basic_ack(delivery_tag=method.delivery_tag)

    channel.basic_qos(prefetch_count=1)
    channel.basic_consume("hello", callback)
    print(" [*] Waiting for messages. To exit press CTRL+C")  # noqa: T201
    channel.start_consuming()
