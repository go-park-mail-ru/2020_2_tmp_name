import cv2
from PIL import Image


def get_mask_path(mask_name):
    return 'masks/' + mask_name + '.png'


def get_photo_mask_name(path, mask):
    return path + '_' + mask + '.png'


def find_faces(path):
    face_cascade = cv2.CascadeClassifier('cascade/frontalface_default.xml')
    img = cv2.imread(path)
    faces = face_cascade.detectMultiScale(cv2.cvtColor(img, cv2.COLOR_BGR2GRAY), 1.1, 4)

    return faces


def overlay_mask(path, mask):
    faces = find_faces(path)

    photo = Image.open(path)

    for face in faces:
        x, y, w, h = face

        mask_path = get_mask_path(mask)
        mask_photo = Image.open(mask_path).resize((w, h))

        photo.paste(mask_photo, (x, y),  mask_photo)

    photo_path = get_photo_mask_name(path, mask)
    photo.save(photo_path)

    return photo, photo_path
