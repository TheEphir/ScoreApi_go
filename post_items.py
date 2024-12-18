import requests
import json

item = {
    "name": "test1",
	"score": -1,
	"image": "",
	"description": "",
}

url = "http://localhost:8080/beer"

def main():
    print(requests.post(url=url, json=item).text)

main()