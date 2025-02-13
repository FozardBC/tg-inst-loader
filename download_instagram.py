import os
import sys
import instaloader
from pathlib import Path
import shutil

# Инициализация Instaloader без авторизации
L = instaloader.Instaloader()

# Папка для скачивания
DOWNLOAD_DIR = "downloads"

def clear_directory(directory_path):
    """Очищает указанную директорию перед скачиванием."""
    folder_path = Path(directory_path)
    if not folder_path.exists():
        os.makedirs(directory_path, exist_ok=True)
        return

    try:
        for item in folder_path.iterdir():
            if item.is_file() or item.is_symlink():
                item.unlink()  # Удаляем файлы
            elif item.is_dir():
                shutil.rmtree(item)  # Удаляем подкаталог
    except Exception as e:
        print(f"Ошибка при очистке директории {directory_path}: {e}")

def normalize_instagram_url(url):
    """Приводит ссылку Instagram к стандартному виду."""
    url = url.split("?")[0]  # Удаляем GET-параметры
    url = url.replace("reels/", "reel/")  # Приводим reels → reel
    return url

def download_instagram_post(post_url):
    """Скачивает контент из Instagram-поста или Reels."""
    try:
        post_url = normalize_instagram_url(post_url)  # Приводим ссылку в нормальный вид
        shortcode = post_url.split("/")[-2]  # Извлекаем shortcode
        print(f" downloading content from shortcode: {shortcode}")

        # Очищаем папку перед скачиванием
        clear_directory(DOWNLOAD_DIR)

        # Загружаем контент
        post = instaloader.Post.from_shortcode(L.context, shortcode)
        L.download_post(post, target=DOWNLOAD_DIR)
        print(f" Content downloaded {DOWNLOAD_DIR}")

        return True
    except Exception as e:
        print(f" Error load: {e}")
        return False

def find_downloaded_files(directory_path):
    """Возвращает список скачанных файлов (изображения и видео)."""
    files = []
    for file_name in os.listdir(directory_path):
        if file_name.lower().endswith((".jpg", ".mp4")):
            files.append(os.path.join(directory_path, file_name))
    return files

if __name__ == "__main__":
    
    if len(sys.argv) < 2:
        print("not instagram posr or reel link")
        print("Example: python instagram_downloader.py https://www.instagram.com/reel/EXAMPLE/")
        sys.exit(1)

    post_url = sys.argv[1]
    print(post_url)

    if "instagram.com/p/" not in post_url and "instagram.com/reel/" not in post_url:
        print(" Err: Invaling URL.")
        sys.exit(1)

    success = download_instagram_post(post_url)

    if success:
        files = find_downloaded_files(DOWNLOAD_DIR)
        if files:
            print(f" Downloaded files\n" + "\n".join(files))
        else:
            print("Files not found after download.")
