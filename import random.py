import os
import random

base_directory = r'.\mockfiles'

def create_directory(directory):
    if not os.path.exists(directory):
        os.makedirs(directory)

common_words = ["dog", "cat", "cow", "bird", "fish", "tree", "car", "house", "boat", "bike", "sun", "moon", "star", "book", "pen", "computer", "ocean", "mountain", "flower", "rain", "cloud", "music", "phone", "shoe", "friend", "sky"]

def generate_type_a(n):
    equations = []
    for _ in range(n):
        num1 = random.randint(1, 100)
        num2 = random.randint(1, 100)
        operator = random.choice(['+', '-', '*', '/'])
        equation = f"{num1} {operator} {num2}\n"
        equations.append(equation)
    return equations

def generate_reversed_common_words(n):
    reversed_words = []
    for _ in range(n):
        word = random.choice(common_words)
        reversed_word = word[::-1]
        reversed_words.append(f"{reversed_word}\n")
    return reversed_words

def generate_type_c(n):
    ascii_characters = []
    for _ in range(n):
        binary_string = ''.join(random.choice(['0', '1']) for _ in range(7))
        ascii_character = chr(int(binary_string, 2)) + '\n'
        ascii_characters.append(ascii_character)
    return ascii_characters

def generate_files(n_files, file_type):
    if file_type == 'Type A':
        return generate_type_a(n_files)
    elif file_type == 'Type B':
        return generate_reversed_common_words(n_files)
    elif file_type == 'Type C':
        return generate_type_c(n_files)
    else:
        raise ValueError(f"Invalid file type: '{file_type}'")

def write_to_file(file_path, content):
    with open(file_path, 'w') as file:
        file.writelines(content)

def generate_all_test_cases(n_files_per_type):
    create_directory(base_directory)

    for i in range(1, n_files_per_type + 1):

        a_content = generate_files(1, 'Type A')
        a_file_path = os.path.join(base_directory, f"TypeAfile{i}.txt")
        write_to_file(a_file_path, a_content)

        b_content = generate_files(1, 'Type B')
        b_file_path = os.path.join(base_directory, f"TypeBfile{i}.txt")
        write_to_file(b_file_path, b_content)

        c_content = generate_files(1, 'Type C')
        c_file_path = os.path.join(base_directory, f"TypeCfile{i}.txt")
        write_to_file(c_file_path, c_content)

    print("All files have been generated successfully.")

generate_all_test_cases(33)
