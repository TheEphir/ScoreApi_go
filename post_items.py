import requests

item = {
    "item_type": "film",
    "name": "kozel",
	"score": 6,
	"image": "",
	"description": "",
}

url = "http://localhost:8080/item"

def main():
    print(requests.get(url="http://localhost:8080/types").text)

main()