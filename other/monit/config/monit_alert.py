#!/usr/bin/python3
# -*- coding: utf-8 -*- #

"""

 This script is used to send Telegram notifications
 when a Monit alert is raised.

 Author = Peter Ibragimov
 Forked from = https://github.com/Matei-Ciobotaru/Monit-Telegram-Alerts
"""


import os

import requests

# Monit alert bot info
BOT_TOKEN = "***:***"  # noqa: S105
CHAT_ID = "***"

# Monit Telegram message information fields
ALERT_FIELDS = ["HOST", "DATE", "EVENT", "SERVICE", "DESCRIPTION"]

# Telegram url
BASE_URL = "https://api.telegram.org"


def alert_message(alert_fields):
    """
    Obtain Monit alert details via environment variables
    """

    header = "🔔 <b>MONIT ALERT</b> 🔔\n\n"
    message_lines = []

    for field in alert_fields:

        variable = "MONIT_" + field
        field_value = os.environ.get(variable, "N/A")

        field_name = "<b>{0}: </b>".format(field.title())
        line = field_name + field_value.upper()

        message_lines.append(line)

    message = header + "\n".join(message_lines)

    return message


def send_alert(token, chat_id, message):
    """
    Send Telegram alert message
    """
    requests.get(
        f"{BASE_URL}/bot{token}/sendMessage",
        params={"text": "{}".format(message), "chat_id": chat_id, "parse_mode": "html"},
    )


def main():
    """
    Script excecution
    """
    message = alert_message(ALERT_FIELDS)
    send_alert(BOT_TOKEN, CHAT_ID, message)


if __name__ == "__main__":
    main()
