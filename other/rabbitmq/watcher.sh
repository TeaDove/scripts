#!/bin/sh

watch -n 1 -t 'sudo /usr/sbin/rabbitmqctl list_queues name messages_unacknowledged messages_ready messages durable auto_delete consumers | tail -n +3 |  column -t;'
