<div align="center" id="top"> 

  &#xa0;

  <!-- <a href="https://project_onepiece.netlify.app">Demo</a> -->
</div>

<h1 align="center">Project_onepiece</h1>

<p align="center">


  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/{{YOUR_GITHUB_USERNAME}}/project_onepiece?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/{{YOUR_GITHUB_USERNAME}}/project_onepiece?color=56BEB8" /> -->

  <!-- <img alt="Github stars" src="https://img.shields.io/github/stars/{{YOUR_GITHUB_USERNAME}}/project_onepiece?color=56BEB8" /> -->
</p>

<!-- Status -->

<!-- <h4 align="center"> 
	🚧  Project_onepiece 🚀 Under construction...  🚧
</h4> 

<hr> -->


<br>

## :dart: About ##

This is a project made to practice

## :sparkles: Features ##

:heavy_check_mark: Scrapping;\
:heavy_check_mark: Api;\
:heavy_check_mark: Web;

## :rocket: Technologies ##

The following tools were used in this project:

- [Go](https://go.dev/)
- [Python](https://www.python.org/)
- [React](https://pt-br.reactjs.org/)
- [MongoDB](https://www.mongodb.com/pt-br)
- [Docker](https://www.docker.com)

## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com), [Go](https://go.dev/), [Python](https://www.python.org/) and [Docker](https://www.docker.com) installed.

# :checkered_flag: Starting #

## To start you can follow this steps ##
```bash
# Clone this project
$ git clone https://github.com/marcosmrs8/project_onepiece

# Make your .env file in the root of your project
Example:
MONGODB=mongodb://admin:admin@localhost:27017
DATABASE=test
COLLECTION=one_piece

# Access
$ cd project_onepiece/frontend/onepiece_frontend

# Run docker file to frontend
$ docker build -t frontend .

# Return to root
$ cd ...

# Run docker to provide MONGODB and the front-end will run on port 3000
$ docker-compose up

# Access
$ cd scrapping/

# Install dependencies
$ pip install -r requirements.txt

# Run the project
$ python main.py

# Remember: This will take a long time to complete and can get some errors

# Back to root
$ cd ..

# You cant in paralalel run the api
$ cd api/

# Install dependencies
$ go mod tidy

# And run
$ go run main.go


```
### Now you can access the http://localhost:3000/ to see the frontend

# Rotas
## Get: `v1/personagens/`
- **Get** return all characters from the one-piece collection
 ``` 
 [
	{
		"Status": "Unknown",
		"Name": "A O",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/8/8c/A_O_Anime_Infobox.png/revision/latest?cb=20160102105316",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/0/0a/A_O_Manga_Infobox.png/revision/latest?cb=20130920143634",
		"SplitedName": [
			"A",
			"O"
		],
		"Bounty": ""
	},
	{
		"Status": "Alive",
		"Name": "Abdullah",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/f/f1/Abdullah_Anime_Infobox.png/revision/latest?cb=20151004140309",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/4/44/Abdullah_Manga_Infobox.png/revision/latest?cb=20151004140342",
		"SplitedName": [
			"Abdullah"
		],
		"Bounty": ""
	},
	{
		"Status": "Deceased",
		"Name": "Absalom",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/5/56/Absalom_Anime_Infobox.png/revision/latest?cb=20230101154942",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/d/d6/Absalom_Manga_Infobox.png/revision/latest?cb=20131008193307",
		"SplitedName": [
			"Absalom"
		],
		"Bounty": ""
	},
  ]
 ```
## Get: `/v1/personagens/specific`
- **Get** returns a specific character if the name is complete, or similar characters if part of the name matches the query

**Search with full name**
```
request:
body:
{
	"name":"Monkey D. Luffy"
}

response:
[
	{
		"Status": "Alive",
		"Name": "Monkey D. Luffy",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/6/6d/Monkey_D._Luffy_Anime_Post_Timeskip_Infobox.png/revision/latest?cb=20200429191518",
		"Image_manga": "",
		"SplitedName": [
			"Monkey",
			"D.",
			"Luffy"
		],
		"Bounty": "3,000,000,000"
	}
]

```
**Search with part of the name**
```
request:
body:
{
	"name": "D."
}

response:
[
	{
		"Status": "Deceased",
		"Name": "Gol D. Roger",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/2/24/Gol_D._Roger_Anime_Infobox.png/revision/latest?cb=20230207184122",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/c/c9/Gol_D._Roger_Manga_Infobox.png/revision/latest?cb=20190325013442",
		"SplitedName": [
			"Gol",
			"D.",
			"Roger"
		],
		"Bounty": "5,564,800,000"
	},
	{
		"Status": "Alive",
		"Name": "Jaguar D. Saul",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/c/cc/Jaguar_D._Saul_Anime_Infobox.png/revision/latest?cb=20221113011440",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/c/c7/Jaguar_D._Saul_Manga_Infobox.png/revision/latest?cb=20220721201047",
		"SplitedName": [
			"Jaguar",
			"D.",
			"Saul"
		],
		"Bounty": ""
	},
]
```


&#xa0;

<a href="#top">Back to top</a>
