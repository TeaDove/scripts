import requests

from bs4 import BeautifulSoup

session = requests.Session()  # Используется для работы, например, с кукиз
login_page = session.get(
    "https://www.123rf.com/photo_125721283_ground-parking-cars-after-"
    "snowfall-in-winter-view-from-above-top-or-aerial-view-automobiles-covered-.html"
)
soup = BeautifulSoup(login_page.content, "html.parser")
# select_one - поиск по CSS-selector, выдаётся словарь атрибутов
images = soup.select(".imgSrcClass")
for image in images:
    r = requests.get(image["href"])
