from tortoise import fields
from tortoise.models import Model


# Описание модели ДБ
class Stadium(Model):
    id = fields.IntField(pk=True)
    name = fields.TextField()
    # Поле один к одному
    current_tournament = fields.OneToOneField("models.Tournament", related_name="statium", null=True)

    def __str__(self):
        return self.name


class Tournament(Model):
    id = fields.IntField(pk=True)
    name = fields.TextField()

    def __str__(self):
        return self.name


class Event(Model):
    id = fields.IntField(pk=True)
    name = fields.TextField()
    # auto_now_add - добавляет Datetime.now() при первой записи. auto_now - при обновление
    date = fields.DatetimeField(auto_now_add=True)
    # первое поле - таблица к которой ключ идёт, related_name - обратной имя,
    # though для Много-к-много - название таблицы через которую идёт Много-к-много
    # По умолчанию: {имя первой}_{имя второй}, i.g. event_team
    tournament = fields.ForeignKeyField("models.Tournament", related_name="events")
    participants = fields.ManyToManyField("models.Team", related_name="events")

    def __str__(self):
        return '"{}" "{}" at {}'.format(self.name, self.tournament, self.date)


class Team(Model):
    id = fields.IntField(pk=True)
    name = fields.TextField()
    size = fields.IntField()

    def __str__(self):
        return "{}: {}".format(self.name, self.size)
