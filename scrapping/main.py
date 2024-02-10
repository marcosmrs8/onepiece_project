import requests
from bs4 import BeautifulSoup
from pymongo import MongoClient
from os.path import join, dirname
import os
from dotenv import load_dotenv

baseUrl = 'https://onepiece.fandom.com'
html_content = requests.get(f'{baseUrl}/pt/wiki/Lista_de_Personagens_Canônicos').text

soup = BeautifulSoup(html_content, "lxml")
print(soup.title.text)
namesToIgnore = ['List of Canon Characters/Names A-M', 'List of Canon Characters/Names N-Z']
dotenv_path = join(dirname(__file__), '../.env')
load_dotenv(dotenv_path)

mongoConnection = os.environ.get("MONGODB")
client = MongoClient(mongoConnection)


def updatebd(image_anime, image_manga, titulo, data):
    try:
        separetedName = titulo.split(' ')
        dadosToUpdate = {'image': image_anime, 'image_manga': image_manga, 'name': titulo, 'splitedName': separetedName}
        dadosToUpdate.update(data)
        database = client.test.one_piece
        database.update_one({'name': titulo}, {"$set": dadosToUpdate}, upsert=True)
        print('adicionado: ', titulo)
    except Exception as e:
        print(e)


def getCharacterAndUpdate(link, title):
    new_content = requests.get(link).text
    
    soup2 = BeautifulSoup(new_content, "lxml")
    aparencia_heading = soup2.find('span', {'id': 'Aparência'})
    data = {}
    if aparencia_heading:
        next_sibling = aparencia_heading.find_next('p')
        if next_sibling:
            data['appearance'] = next_sibling.text
    image_anime = soup2.find('a', {'class': 'image'})
    image_manga = soup2.find('a', {'title': 'Mangá'})
    
    try:
        description = soup2.find('div', {'data-source': 'status'}).text
        data['Status'] = description.replace('\n', '').replace('Status:', '')
    except Exception:
        print('erro no status')
        pass
    try:
        bounty = soup2.find('div', {'data-source': 'bounty'}).text
        bountyTreated = bounty.split('\n')
        newBounty = bountyTreated[2].split('[')
        data['bounty'] = newBounty[0]
        print('bounty', newBounty[0])
    except Exception:
        pass
    if image_manga:
        updatebd(image_anime.attrs['href'], image_manga.attrs['href'], title, data)
    else:
        updatebd(image_anime.attrs['href'], '', title, data)


gdp_table = soup.findAll("table", attrs={"class": "wikitable sortable"})
allCharacters = []
for table in gdp_table:
    gdp_table_data = table.tbody.find_all("tr")
    for item in gdp_table_data:
        try:
            if item.a.attrs['title'] not in namesToIgnore:
                link = baseUrl + item.a.attrs['href']
                title = item.a.attrs['title']
                newTitle = title.replace("/", " ").replace("-", " ").replace("Vegapunk", "").strip()
                if("Capítulo" not in newTitle):
                    getCharacterAndUpdate(link, newTitle)
                    allCharacters.append({"link": link, "title": newTitle})
        except Exception as e:
            print(f'error: {e}', item.text)
