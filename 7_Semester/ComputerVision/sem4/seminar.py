import cv2
import numpy as np

img = cv2.imread("flowers.jpeg", 0)

cv2.namedWindow("Original", cv2.WINDOW_NORMAL)
cv2.imshow("Original", img)
cv2.waitKey(0)

cv2.imwrite("flowers.png", img)
cv2.imwrite('flowers1gray.bmp',img)

jpeg_quality = 1  ## A value between 0 and 100 (higher means better quality, but larger file size)

cv2.imwrite('flowers1_jpg.jpg', img, [cv2.IMWRITE_JPEG_QUALITY, jpeg_quality])
flowers1_jpg = cv2.imread("flowers1_jpg.jpg")



cv2.namedWindow("Compressed_jpg", cv2.WINDOW_NORMAL)
cv2.imshow("Compressed_jpg", flowers1_jpg)
cv2.waitKey(1)
cv2.imwrite('flowers1_jp2.jp2', img,[cv2.IMWRITE_JPEG2000_COMPRESSION_X1000,0])
flowers1_jp2 = cv2.imread("flowers1_jp2.jp2")

cv2.namedWindow("Compressed_jp2", cv2.WINDOW_NORMAL)
cv2.imshow("Compressed_jp2", flowers1_jp2)
cv2.waitKey(0)