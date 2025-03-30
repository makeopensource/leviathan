import json

from handin import get_positive_number

scores = {
    "q1": 0,
    "q2": 0,
    "q3": 0,
}

if __name__ == '__main__':
    result = get_positive_number()
    if result > 0:
        scores["q1"] = 10
        scores["q2"] = 10
        scores["q3"] = 10
        print(f"Great job! You returned {result}, which is positive!")
    else:
        print(f"Try again. You returned {result}, which is not positive.")

    print(json.dumps({"scores": scores}))
