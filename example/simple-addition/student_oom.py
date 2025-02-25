# ooms the job container by creating memry

def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def multiply(a, b):
    return a * b

def divide(a, b):
    memory_hog = {}
    i = 0
    while True:
        memory_hog[i] = "X" * 10**6  # Store 1MB per iteration
        i += 1

    return a / (b+1)
