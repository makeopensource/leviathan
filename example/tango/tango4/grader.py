import json

scores = {
    "q1": 0,
}


def finish():
    print()
    print(json.dumps({"scores": scores}))
    exit(0)


if __name__ == '__main__':

    try:
        with open("settings.json", "r", encoding="utf-8") as f:
            settings = json.load(f)
    except Exception as e:
        print("There was a problem parsing your form as JSON.")
        print(e)
        finish()

    print("Removing non-ASCII characters from submission.")

    for k, v in settings.items():
        settings[k] = str(v).encode("ascii", "ignore").decode("ascii")
    print(f"Your settings.json: {settings}")

    print("\nGrading question 1")
    q1: str = settings.get("q1", None)
    print(f"Your input string was: '{q1}'")
    print(f"Calling your 'give_me_string' function with argument '{q1}'...")
    try:
        from handin import give_me_string

        output = give_me_string(q1)
        print(f"Your function returned: '{output}'")
        if output == q1:
            print("Good! Your function returned the same string you passed in.")
            scores["q1"] = 10
        else:
            print("Whoops! Your function did not return the same string you passed in.")
    except Exception as e:
        print(f"Your function raised an exception: {e}")

    finish()
