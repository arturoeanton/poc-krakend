# poc-krakend

## Description

This POC  is for test dinamic (Hot reload) routes in krakend.

Allows you to create, modify and delete enpoints in "configuration.json" without restart the application. 

At the bottomm of the document, you can view some examples about it. 

## Demo Video

The next video shows how to run and modify the endpoints 

[![asciicast](https://asciinema.org/a/464252.svg)](https://asciinema.org/a/464252)


## Installation

* For run the POC, enter the following commands. It is neccesary to create a new folder to clone the different forks for the corresponding projects. 

```
mkdir POC
cd POC
git clone https://github.com/arturoeanton/poc-krakend.git
git clone https://github.com/arturoeanton/lura.git
git clone https://github.com/arturoeanton/gin.git

cd poc-krakend
go mod vendor
go run .
```

In the case of the gin pull request, the modification is the following. This modification allows to reset search tree at gin level. 


```
func (e *Engine) ResetTrees() {
	e.trees = make(methodTrees, 0, 9)
}
```

### Hot reload

```
touch configuration.json
```

## Examples

* Based on this initial configuration file we will make some examples. 

```
{
	"version": 2,
	"name": "Express API Gateway",
	"endpoints": [
		{
			"endpoint": "/pokemon",
			"method": "GET", 
			"backend": [
				{ 
					"host": ["https://pokeapi.co"],
					"method": "GET",
					"url_pattern": "/api/v1/pokemon/ditto"
				}
			]
		},
		{
			"endpoint": "/type/{id}",
			"method": "GET", 
			"backend": [
				{ 
					"host": ["https://pokeapi.co"],
					"method": "GET",
					"url_pattern": "/api/v1/type/{id}"
				}
			]
		}
	]
}
```

* Once the application is running, modify the file and change the "url_pattern" like this for the pokemon endpoint. 

```
{
	"version": 2,
	"name": "Express API Gateway",
	"endpoints": [
		{
			"endpoint": "/pokemon",
			"method": "GET", 
			"backend": [
				{ 
					"host": ["https://pokeapi.co"],
					"method": "GET",
					"url_pattern": "/api/v1/pokemon"
				}
			]
		},
		{
			"endpoint": "/type/{id}",
			"method": "GET", 
			"backend": [
				{ 
					"host": ["https://pokeapi.co"],
					"method": "GET",
					"url_pattern": "/api/v1/type/{id}"
				}
			]
		}
	]
}

```

* The endpoint change on hot while the application is running. 


