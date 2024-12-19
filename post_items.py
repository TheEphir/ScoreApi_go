import requests

item = {
    "item_type": "beer",
    "name": "kozel",
	"score": 6,
	"image": "",
	"description": "",
}

url = "http://localhost:8080/item"

def main():
    print(requests.post(url=url, json=item).text)

main()