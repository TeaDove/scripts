from celery.schedules import crontab

from .celery import app


@app.task  # noqa: PIE783
def test(arg):
    print(arg)  # noqa: T201


@app.task  # noqa: PIE783
def add(x, y):
    z = x + y
    print(z)  # noqa: T201


@app.on_after_configure.connect
def setup_periodic_tasks(sender, **kwargs):
    # Calls test('hello') every 10 seconds.
    sender.add_periodic_task(10.0, test.s("hello"), name="add every 10")

    # Calls test('world') every 30 seconds
    sender.add_periodic_task(30.0, test.s("world"), expires=10)

    # Executes every Monday morning at 7:30 a.m.
    sender.add_periodic_task(
        crontab(hour=7, minute=30, day_of_week=1),
        test.s("Happy Mondays!"),
    )
