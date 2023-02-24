# крайне полезный шаблон для sqlite3

import logging
import sqlite3

logging.basicConfig(level=logging.DEBUG)


class Database:
    def __init__(self, file_name: str = ":memory:"):
        self.conn = None
        self.cursor = None
        if file_name:
            self.open(file_name)

    def open(self, file_name):
        try:
            self.conn = sqlite3.connect(file_name)
            self.cursor = self.conn.cursor()
        except sqlite3.Error:
            print("Error connecting to database!")  # noqa: T201

    def close(self):
        if self.conn:
            self.conn.commit()
            self.cursor.close()
            self.conn.close()

    def recreate(self):
        temp_conn = sqlite3.connect(":memory:")
        temp_cursor = temp_conn.cursor()
        temp_cursor.execute(
            """create table place(
                                            id integer primary key autoincrement,
                                            bank_name text,
                                            type text,
                                            lat float,
                                            lon float,
                                            city text,
                                            address text)"""
        )
        temp_conn.commit()
        temp_conn.backup(self.conn)
        temp_conn.close()

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, traceback):
        self.close()

    def get(self, table, columns: str = "*", limit=None):
        query = f"SELECT {columns} from {table}"
        self.cursor.execute(query)
        # fetch data
        rows = self.cursor.fetchall()
        return rows[len(rows) - limit if limit else 0 :]

    def get_where(self, table: str, columns: str, where: str, limit=None):
        query = f"SELECT {columns} from {table} where {where}"
        logging.debug(query)
        self.cursor.execute(query)
        # fetch data
        rows = self.cursor.fetchall()
        return rows[len(rows) - limit if limit else 0 :]

    def getLast(self, table, columns):
        return self.get(table, columns, limit=1)[0]

    @staticmethod
    def to_csv(data, filename="output.csv"):
        with open(filename, "a") as file:  # noqa: SCS109
            file.write(",".join([str(j) for i in data for j in i]))

    def write(self, table, columns, data):
        query = f"INSERT INTO {table} ({columns}) VALUES ({data})"
        self.cursor.execute(query)

    def query(self, sql):
        self.cursor.execute(sql)
