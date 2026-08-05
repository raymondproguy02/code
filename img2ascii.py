from PIL import Image
import argparse
import sys

ASCII_CHARS = ['@', '#', 'S', '%', '?', '*', '+', ';', ':', ',', '.']

def resize(image, new_width=100):
    width, height = image.size
    ratio = height / width
    new_height = int(new_width * ratio * 0.55)
    return image.resize((new_width, new_height))

def to_grayscale(image):
    return image.convert('L')

def pixels_to_ascii(image):
    pixels = image.getdata()
    ascii_str = ''.join(ASCII_CHARS[pixel * len(ASCII_CHARS) // 256] for pixel in pixels)
    return ascii_str

def convert(image_path, width=100):
    try:
        image = Image.open(image_path)
        image = resize(image, width)
        image = to_grayscale(image)
        
        ascii_str = pixels_to_ascii(image)
        img_width = image.width
        
        ascii_art = '\n'.join(ascii_str[i:i+img_width] for i in range(0, len(ascii_str), img_width))
        print(ascii_art)
        
    except Exception as e:
        print(f"❌ Error: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("image", help="Image file path")
    parser.add_argument("-w", "--width", type=int, default=100)
    args = parser.parse_args()
    
    convert(args.image, args.width)
