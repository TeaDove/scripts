import settings
from tortoise import Tortoise, run_async
from tournaments import Event, Team, Tournament


async def main():
    await Tortoise.init(config=settings.TORTOISE_ORM)
    # Create instance by save
    tournament = Tournament(name="Trackmania Nations Forever")
    await tournament.save()

    # Or by .create()
    await Event.create(name="World cup 2008", tournament=tournament)
    world_cup_event = await Event.create(name="World cup 2020", tournament=tournament)
    participants = []
    for i in range(2):
        team = await Team.create(name="Team {}".format(i + 1), size=i * 5)
        participants.append(team)

    # M2M Relationship management is quite straightforward
    # (also look for methods .remove(...) and .clear())
    await world_cup_event.participants.add(*participants)

    # You can query related entity just with async for
    print("Teams of {}".format(world_cup_event))  # noqa: T201
    async for team in world_cup_event.participants:
        print(team)  # noqa: T201

    # After making related query you can iterate with regular for,
    # which can be extremely convenient for using with other packages,
    # for example some kind of serializers with nested support
    for _team in world_cup_event.participants:
        pass

    # Or you can make preemptive call to fetch related objects
    await Event.filter(participants=participants[0].id).prefetch_related("participants", "tournament")

    # Tortoise supports variable depth of prefetching related entities
    # This will fetch all events for team and in those events tournaments will be prefetched
    teams = await Team.all().prefetch_related("events__tournament")
    print(teams)  # noqa: T201
    print(teams[0].events[0].tournament)  # noqa: T201

    # You can filter and order by related models too
    await Tournament.filter(events__name__in=["Test", "Prod"]).order_by("-events__participants__name").distinct()


if __name__ == "__main__":
    run_async(main())
