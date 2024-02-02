import cv2
import numpy as np

img = cv2.imread("flowers.bmp")

gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)

cv2.namedWindow("Image", cv2.WINDOW_NORMAL)
cv2.imshow("Image", img)
cv2.waitKey(0)

Gx = cv2.Sobel(gray, cv2.CV_32F, 1, 0)
Gy = cv2.Sobel(gray, cv2.CV_32F, 0, 1)

G = abs(Gx) + abs(Gy)
G = cv2.min(255,G)

G = np.uint8(255*(G>50))

cv2.namedWindow('Sobel',cv2.WINDOW_NORMAL)
cv2.imshow('Sobel', G)
cv2.waitKey(0)

canny = cv2.Canny(gray, 30, 50)
cv2.namedWindow('Canny',cv2.WINDOW_NORMAL)
cv2.imshow('Canny', canny)
cv2.waitKey(0)