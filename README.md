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
$ cd api/cmd/

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
		"Status": "Vivo(a)",
		"Name": "A.O",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/9/94/A.O_Anime_Infobox.png/revision/latest?cb=20211117163925&path-prefix=pt",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/5/56/A.O_Manga_Infobox.png/revision/latest?cb=20150916190642&path-prefix=pt",
		"SplitedName": [
			"A.O"
		],
		"Bounty": "",
		"Appearance": "A.O é um homem muito alto com um rosto comprido, em forma de retângulo, olhos fundos sob as sobrancelhas proeminentes, costeletas escuras que percorrem todo o caminho até o queixo e cabelo ondulado longo esverdeada. Em sua cabeça ele usa um amarelo e marrom listrado bandana e sobre ele um chapéu de capitão tricorn com o que parece ser o seu Roger alegre nele, um crânio branco vestindo um capuz animalesca. Ele também usa um azul escuro, camisa aberta com as mangas arregaçadas e uma gola de babados, e brincos elaborados composta por três pequenas jóias cada um em ambas as orelhas.\n"
	},
	{
		"Status": "Vivo(a)",
		"Name": "Abdullah",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/f/f1/Abdullah_Anime_Infobox.png/revision/latest?cb=20211001025009&path-prefix=pt",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/4/44/Abdullah_Manga_Infobox.png/revision/latest?cb=20211001025320&path-prefix=pt",
		"SplitedName": [
			"Abdullah"
		],
		"Bounty": "",
		"Appearance": "Abdullah tem uma cabeça pontiaguda que é careca na parte superior com cabelo preto em tranças presas para os lados. Sua testa alta tem três cicatrizes verticais sobre ela. Ele ostenta óculos retangulares finas e, juntamente com grandes calças roxas presas por suspensórios castanhos. Ele também tem uma armadura sobre o braço esquerdo.\n"
	}
 ]
 ```
## Get: `/v1/personagens/specific?name={name || part of name}`
- **Get** returns a specific character if the name is complete, or similar characters if part of the name matches the query

**Search with full name**
```
request:
query = /specific?name=Monkey D. Luffy

response:
[
	{
		"Status": "Estado:Vivo(a)",
		"Name": "Monkey D. Luffy",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/6/6d/Monkey_D._Luffy_Anime_Post_Timeskip_Infobox.png/revision/latest?cb=20190303115209&path-prefix=pt",
		"Image_manga": "",
		"SplitedName": [
			"Monkey",
			"D.",
			"Luffy"
		],
		"Bounty": "3,000,000,000",
		"Appearance": "Luffy é conhecido como \"Luffy Chapéu de Palha\" por sua marca registrada que é o Chapéu de Palha, que ele ganhou emprestado ainda jovem pelo lendário Yonkou, Shanks \"O Ruivo\", que, por sua vez o recebeu de Gol D. Roger. Luffy veste um short azul, sandálias e um colete vermelho com mangas. Luffy também tem uma cicatriz com dois pontos embaixo do olho esquerdo (que ele ganhou por se esfaquear usando uma faca em seu olho para mostrar para Shanks que ele era um homem) e cabelo preto curto e bagunçado. Ele foi gravemente ferido por Akainu na Batalha de Marineford, deixando uma grande cicatriz em forma de \"X\" em seu peito. Parecendo ser magro debaixo de sua camisa, ele tem um físico surpreendentemente bem construído. Ele é aparentemente muito baixo, pois sua altura é ofuscada pela maioria dos indivíduos na série.\n",
	}
]

```
**Search with part of the name**
```
request:
query = /specific?name=Monkey

response:
[
	{
		"Status": "",
		"Name": "Monkey D. Dragon",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/f/f5/Monkey_D._Dragon_Anime_Infobox.png/revision/latest?cb=20240127214459&path-prefix=pt",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/5/52/Monkey_D._Dragon_Manga_Infobox.png/revision/latest?cb=20240127214515&path-prefix=pt",
		"SplitedName": [
			"Monkey",
			"D.",
			"Dragon"
		],
		"Bounty": "",
		"Appearance": "Dragon é um homem alto, de meia-idade, com cabelos pretos espetados com um topete, e uma tatuagem no lado esquerdo do rosto, que é vermelho escuro no anime e vermelho claro no mangá. Ele tem um pouco de barba no queixo. Em geral, ele veste um longo manto verde, por baixo do qual usa o traje laranja de um revolucionário. Na maioria das vezes, ele é visto com um sorriso. Apenas em raras ocasiões, como na execução de Roger, Dragon não é visto sorrindo.\n"
	},
	{
		"Status": "Desconhecido",
		"Name": "Monkey D. Garp",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/e/e1/Monkey_D._Garp_Anime_Infobox.png/revision/latest?cb=20240109223317&path-prefix=pt",
		"Image_manga": "https://static.wikia.nocookie.net/onepiece/images/4/45/Monkey_D._Garp_Manga_Infobox.png/revision/latest?cb=20240109223310&path-prefix=pt",
		"SplitedName": [
			"Monkey",
			"D.",
			"Garp"
		],
		"Bounty": "(3.000.000.000)",
		"Appearance": "Garp é um homem de idade, alto, moreno, de peito largo e musculoso. Ele tem barba e uma cicatriz no olho esquerdo. No anime, a cor de seus olhos é azul e seu cabelo é cinza, enquanto no mangá é branco. Garp e Tsuru são os dois únicos vice-almirantes que têm ombreiras especiais; Garp é dourado e azul (preto e vermelho no anime), enquanto Tsuru é roxo com pontos brancos, sendo que a cor padrão é azul e vermelho. Durante a Batalha de Edd War, Garp tinha duas listras pretas em suas braçadeiras, e durante sua aparição em Water 7, aumentou para três listras; na época da Guerra de Marineford, aumentou novamente para quatro listras.\n"
	},
	{
		"Status": "Vivo(a)",
		"Name": "Monkey D. Luffy",
		"Image": "https://static.wikia.nocookie.net/onepiece/images/6/6d/Monkey_D._Luffy_Anime_Post_Timeskip_Infobox.png/revision/latest?cb=20190303115209&path-prefix=pt",
		"Image_manga": "",
		"SplitedName": [
			"Monkey",
			"D.",
			"Luffy"
		],
		"Bounty": "3,000,000,000",
		"Appearance": "Luffy é conhecido como \"Luffy Chapéu de Palha\" por sua marca registrada que é o Chapéu de Palha, que ele ganhou emprestado ainda jovem pelo lendário Yonkou, Shanks \"O Ruivo\", que, por sua vez o recebeu de Gol D. Roger. Luffy veste um short azul, sandálias e um colete vermelho com mangas. Luffy também tem uma cicatriz com dois pontos embaixo do olho esquerdo (que ele ganhou por se esfaquear usando uma faca em seu olho para mostrar para Shanks que ele era um homem) e cabelo preto curto e bagunçado. Ele foi gravemente ferido por Akainu na Batalha de Marineford, deixando uma grande cicatriz em forma de \"X\" em seu peito. Parecendo ser magro debaixo de sua camisa, ele tem um físico surpreendentemente bem construído. Ele é aparentemente muito baixo, pois sua altura é ofuscada pela maioria dos indivíduos na série.\n"
	}
]
```


&#xa0;

<a href="#top">Back to top</a>
