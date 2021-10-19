#!/usr/bin/env python3

# imports
import os
import time

Y = set(['yes', 'y', 'YES', 'Y'])
N = set(['no', 'n', 'NO', 'N'])

def banner():
    print("WEBSTREAM INSTALLER SCRIPT")

# installing script
def installwebstream():
    agree = input("Are you sure you want to install webstream Y \ N ")
    if agree in Y:
        print('Installing Go Modules')
        time.sleep(1)
        os.system('go mod tidy')
        time.sleep(1)
        os.system('go build -o webstream cmd/main.go')
        time.sleep(1)
        print('Running webstream')
        time.sleep(1)
        os.system('./cmd/webstream')
    elif agree in N:
        print('OK BYE')


# call all the def function into main block
def main():
    banner()
    time.sleep(2)
    installwebstream()

# and call the main block into main function
if __name__ == "__main__":
    main()