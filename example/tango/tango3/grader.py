import json

scores = {
    "q1": 0,
    "q2": 0,
    "q3": 0,
}


def finish():
    print()
    print(json.dumps({"scores": scores}))
    exit(0)


if __name__ == '__main__':

    try:
        with open("handin.json", "r", encoding="utf-8") as f:
            submission = json.load(f)
    except Exception:
        print("There was a problem parsing your submission.")
        finish()

    print("Removing non-ASCII characters from submission.")

    for k, v in submission.items():
        submission[k] = str(v).encode("ascii", "ignore").decode("ascii")
    print(f"Your submission: {submission}")

    print("\nGrading question 1")
    q1: str = submission.get("q1", None)
    print(f"Your answer to question 1 was: {q1}")
    if q1 is None:
        print("You did not answer question 1.")
    else:
        if q1.startswith("a"):
            scores["q1"] = 10
            print("Good! That starts with the character 'a'.")
        elif q1.startswith("A"):
            scores["q1"] = 5
            print("Whoops! That starts with the character 'A'. Remember that Python is case-sensitive. Partial credit.")
        else:
            print("Whoops! That does not start with the character 'a'.")

    print("\nGrading question 2")
    q2: str = submission.get("q2", None)
    print(f"Your answer to question 2 was: {q2}")
    if q2 is None:
        print("You did not answer question 2.")
    else:
        if q2 == "Yes":
            scores["q2"] = 10
            print("Good! You confirmed you were in lecture.")
        elif q2 == "No":
            scores["q2"] = 5
            print("Whoops! You said you were not in lecture. Partial credit for honesty.")
        else:
            scores["q2"] = 9
            print("You submitted an option that wasn't provided. Partial credit for creativity.")

    print("\nGrading question 3")
    q3: str = submission.get("q3", None)
    print(f"Your answer to question 3 was: {q3}")
    if q3 is None:
        print("You did not answer question 3.")
    else:
        try:
            q3_int = int(q3)
            print("Good! We set your score for q3 to the number you submitted.")
            scores["q3"] = q3_int
        except ValueError:
            print("Whoops! That's not a number.")

    finish()
