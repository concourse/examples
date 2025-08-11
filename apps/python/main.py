#!/usr/bin/env python3
import sys

def add_six(number):
    return number + 6

def main():
    # Check if there's a command-line argument
    if len(sys.argv) != 2:
        print("Usage: python3 main.py <number>")
        sys.exit(1)

    try:
        number = float(sys.argv[1])

        result = add_six(number)

        if result.is_integer():
            result = int(result)

        print(f"{result}")

    except ValueError:
        print(f"Error: '{sys.argv[1]}' is not a valid number.")
        sys.exit(1)

if __name__ == "__main__":
    main()
