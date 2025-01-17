import requests

item = {
    "item_type": "beer",
    "name": "kozel 123",
	"score": 10,
	"image": "",
	"description": "",
}

url = "http://localhost:8080/item"

def main():
    # requests.post(url="http://localhost:8080/item", json=item)
    print(requests.get(url="http://localhost:8080/filsm/avatar").text)

main()