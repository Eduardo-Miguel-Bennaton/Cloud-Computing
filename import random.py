import os
import random
import string

# Define the directory where the test files will be saved
test_files_directory = r'C:\Users\eduar\Desktop\Final Project\mockfiles'

# Function to ensure the directory exists
def create_directory(directory):
    print(f"Checking if directory '{directory}' exists...")
    if not os.path.exists(directory):
        print(f"Creating directory '{directory}'...")
        os.makedirs(directory)
    else:
        print(f"Directory '{directory}' already exists.")

# Functions to generate content for each file type
def generate_type_a(n):
    equations = []
    for _ in range(n):
        num1 = random.randint(1, 100)
        num2 = random.randint(1, 100)
        operator = random.choice(['+', '-', '*', '/'])
        equation = f"{num1} {operator} {num2}\n"
        equations.append(equation)
    return equations

def generate_type_b(n):
    words = []
    for _ in range(n):
        word = ''.join(random.choice(string.ascii_letters) for _ in range(random.randint(3, 10)))
        words.append(f"{word[::-1]}\n")
    return words

def generate_type_c(n):
    binaries = []
    for _ in range(n):
        binary = ''.join(random.choice(['0', '1']) for _ in range(8))  # 8-bit binary numbers
        binaries.append(f"{binary}\n")
    return binaries

# Function to call the appropriate content generator based on file type
def generate_files(n_files, file_type):
    if file_type == 'Type A':
        return generate_type_a(n_files)
    elif file_type == 'Type B':
        return generate_type_b(n_files)
    elif file_type == 'Type C':
        return generate_type_c(n_files)
    else:
        raise ValueError(f"Invalid file type: '{file_type}'")

# Function to write content to a file
def write_to_file(file_path, content):
    with open(file_path, 'w') as file:
        file.writelines(content)
    print(f"Written to file '{file_path}'.")

# Main function to generate the specified number of files for each file type
def generate_all_test_cases(n_files_per_type):
    create_directory(test_files_directory)  # Ensure the directory exists
    for file_type in ['Type A', 'Type B', 'Type C']:
        file_type_directory = os.path.join(test_files_directory, file_type)
        create_directory(file_type_directory)  # Create subdirectory for each file type
        for i in range(1, n_files_per_type + 1):
            print(f"Generating file {i} of type '{file_type}'...")
            content = generate_files(1, file_type)  # Generate content for one file
            file_name = f"{file_type}_file_{i}.txt"
            file_path = os.path.join(file_type_directory, file_name)
            write_to_file(file_path, content)  # Write the content to the file
    print("All files have been generated successfully.")

# Generate 100 files for each type A, B, and C
generate_all_test_cases(100)