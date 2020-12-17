import cv2
from PIL import Image

def find_face():
    face_cascade = cv2.CascadeClassifier('frontalface_default.xml')

    img = cv2.imread('test.png')

    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

    faces = face_cascade.detectMultiScale(gray, 1.1, 4)

    # for (x, y, w, h) in faces:
    #     cv2.rectangle(img, (x, y), (x+w, y+h), (255, 0, 0), 2)
    face = Image.open('test.png')

    x, y, w, h = faces[0]
    mask = Image.open('mask.png')
    resize_mask = mask.resize((w, h))

    face.paste(resize_mask, (x, y),  resize_mask)
    face.save('img_with_mask.png')

    #cv2.imshow('img', face)
    #cv2.waitKey()

    return faces


find_face()
