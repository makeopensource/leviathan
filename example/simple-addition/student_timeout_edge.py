import time

# test timeout should be set to 10 secs
# and this should fail

def add(a, b):
    time.sleep(5) # add delay for test
    return a + b

def subtract(a, b):
    time.sleep(5) # add delay for test
    return a - b

def multiply(a, b):
    time.sleep(1) #
    return a * b

def divide(a, b):
    return a / b
