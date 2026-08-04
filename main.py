import secrets
import string
import argparse

def generate_password(length=16, use_symbols=True):
    chars = string.ascii_letters + string.digits
    if use_symbols:
        chars += string.punctuation
    
    password = ''.join(secrets.choice(chars) for _ in range(length))
    return password

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-l", "--length", type=int, default=16)
    parser.add_argument("-n", "--no-symbols", action="store_true")
    args = parser.parse_args()
    
    pwd = generate_password(args.length, not args.no_symbols)
    print(f"🔑 {pwd}")
    print(f"📏 Length: {len(pwd)} characters")
