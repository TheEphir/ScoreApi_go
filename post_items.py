import requests

item = {
    "item_type": "test",
    "name": "test_item",
	"score": 10,
	"image": "",
	"description": "",
}

url = "http://localhost:8080/item"

def main():
    print(requests.post(url="http://localhost:8080/item", json=item))
    # print(requests.delete(url="http://localhost:8080/film/avatar").text)

main()