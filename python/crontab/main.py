from pathlib import Path
from typing import Tuple

from crontab import CronTab

# https://crontab.guru/


def setup_crontab(  # noqa: CCR001
    files: Tuple[str], setting: str, comment: str, venv_folder: str = None, directory: str = None
):

    with CronTab(user=True) as cron:
        cron.remove_all(comment=comment)

    new_files = []
    for file in files:
        if isinstance(file, str):
            path_of_file = Path(file)
            del file
            new_files.append(path_of_file)
        else:
            new_files.append(file)
    files = new_files

    if directory:
        directory = Path(directory)

    with CronTab(user=True) as cron:
        for file in files:
            if venv_folder:
                command = (  # noqa: ECE001
                    f"cd {directory.absolute() if directory else file.parent.absolute()} && "
                    f"source {venv_folder}/bin/activate && "
                    f"python3 {str(file.absolute())} "
                    f">>{directory.absolute() if directory else file.parent.parent.absolute()}/crontab.log 2>&1"
                )
            else:
                command = (
                    f"cd {directory.absolute() if directory else file.parent.absolute()} && "
                    f"python3 {str(file.absolute())} "
                    f">>{directory.absolute() if directory else file.parent.parent.absolute()}/crontab.log 2>&1"
                )
            print(command)  # noqa: T201
            job = cron.new(command=command, comment=comment)
            job.setall(setting)


if __name__ == "__main__":
    setup_crontab([Path("./main.py")], "*/1 * * * *", "crontab-setup-example")
