from pathlib import Path


def get_res_dict() -> dict:
    res_dict = {}
    for file in Path("res").iterdir():
        res_dict[file.stem] = open(file, "r").read()  # noqa: SIM115
    return res_dict
