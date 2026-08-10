import time
import sys

def timer(seconds):
    try:
        print(f"⏱️ Timer set for {seconds}s")
        while seconds > 0:
            mins, secs = divmod(seconds, 60)
            print(f"\r⏳ {mins:02d}:{secs:02d}", end='')
            time.sleep(1)
            seconds -= 1
        print("\n🔔 Time's up!")
    except KeyboardInterrupt:
        print("\n⏹️ Timer cancelled")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python timer.py <seconds>")
        sys.exit(1)
    
    timer(int(sys.argv[1]))
