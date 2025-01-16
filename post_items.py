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
    requests.post(url="http://localhost:8080/item", json=item)
    #print(requests.get(url="http://192.168.0.132:8080/types").text)

main()