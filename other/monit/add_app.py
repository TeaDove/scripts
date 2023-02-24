#!/usr/bin/python3
# flake8: noqa
# -*- coding: utf-8 -*- #

import argparse
import os
import shutil
from pathlib import Path

SCRIPT = Path(os.path.realpath(__file__))
os.chdir(SCRIPT.parent)
BASE = SCRIPT.parent

template = """
check process {1} with pidfile /var/run/{1}.pid
    start program = "{0}/monitor.sh {1} start"
    stop program  = "{0}/monitor.sh {1} stop"
    onreboot laststate
    # CPU checking
    if total cpu > 60% for 3 times within 5 cycles then exec "/etc/monit/monit_alert.py"
    if total cpu > 80% for 3 times within 5 cycles then restart
    if total cpu > 80% for 3 times within 5 cycles then exec "/etc/monit/monit_alert.py" # Потому что монит не умеет "and"
    if total cpu > 90% for 3 times within 5 cycles then stop
    if total cpu > 90% for 3 times within 5 cycles then exec "/etc/monit/monit_alert.py"
    # MEM check
    if totalmem > 500.0 MB for 3 times within 5 cycles then exec "/etc/monit/monit_alert.py"
    if totalmem > 800.0 MB for 3 times within 5 cycles then restart
    if totalmem > 800.0 MB for 3 times within 5 cycles then exec "/etc/monit/monit_alert.py"
    if totalmem > 1024.0 MB for 3 times within 5 cycles then stop
    if totalmem > 1024.0 MB for 3 times within 5 cycles then exec "/etc/monit/monit_alert.py"
    # Restart check
    if 2 restarts within 5 cycles then exec "/etc/monit/monit_alert.py"
    if 5 restarts within 10 cycles then stop
    if 5 restarts within 10 cycles then exec "/etc/monit/monit_alert.py"
"""


def main():
    parser = argparse.ArgumentParser(description="generate config for monit and copy monitor.sh and start.sh to app")
    parser.add_argument("location", action="store", help="location of project")
    args = parser.parse_args()

    location = Path(args.location)
    process_name = location.name
    result = template.format(location, process_name)
    with open(f"/etc/monit/conf.d/{process_name}.conf", "w") as f:
        f.write(result)
    print(f'Generated config for "{process_name}"')
    shutil.copyfile("./app-example/start-tmpl.sh", location / "start.sh")
    os.chmod(location / "start.sh", 0o744)
    shutil.copyfile("./app-example/monitor.sh", location / "monitor.sh")
    os.chmod(location / "monitor.sh", 0o744)
    print(f'Copied "monitor.sh" and "start.sh" to {location}. Do not forget to change "start.sh"')


if __name__ == "__main__":
    main()
