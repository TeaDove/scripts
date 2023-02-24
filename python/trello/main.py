from trello import TrelloClient

client = TrelloClient(api_key="api_key", api_secret="api_secret", token="token")  # Подключение Трелло Апи # noqa: S106

all_boards = client.list_boards()

print([board.name for board in all_boards])  # noqa: T201
