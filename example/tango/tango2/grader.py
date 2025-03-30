import json

scores = {
    "q1": 10,
    "q2": 10,
    "q3": 10,
}

if __name__ == '__main__':

    with open("settings.json", "r") as f:
        print(f.readline())

    print(json.dumps({"scores": scores}))
