import cv2
import numpy as np
import pywt
import pywt.data
import matplotlib.pyplot as plt


gray = cv2.imread("flowers.png", 0)

cv2.namedWindow("Original", cv2.WINDOW_NORMAL)
cv2.imshow("Original", gray)
cv2.waitKey(0)

L = 3
wavelettype = 'db4'
percent = 95

c = pywt.wavedec2(gray, wavelettype, level = L)

a = c[0].flatten()

out = a.tolist()

for j1 in range(1, L):
    for j2 in range(0,3):
        a = c[j1][j2].flatten()
        for j3 in range(0,len(a)):
            out.append(a[j3])

absout = np.abs(out)
absout.sort()

k = int(percent * len(absout) / 100)
threshold = absout[k]


for j1 in range(0, len(c[0])):
    for j2 in range(0,len(c[0][j1])):
        if abs(c[0][j1][j2]) <= threshold:
            c[0][j1][j2] = 0


for j1 in range(1, L):
    for j2 in range(0,3):
        for j3 in range(0,len(c[j1][j2])):
            for j4 in range(0,len(c[j1][j2][j3])):
                 if abs(c[j1][j2][j3][j4]) <= threshold:
                        c[j1][j2][j3][j4] = 0

compressed_img = pywt.waverec2(c, wavelettype)

compressed_img = cv2.max(0,cv2.min(255,compressed_img))

cv2.namedWindow('Our_Compressed_image', cv2.WINDOW_NORMAL)
cv2.imshow("Our_Compressed_image",compressed_img.astype(np.uint8))
cv2.waitKey(0)

cv2.imwrite("+Our_jpg_image.jpg", gray)
cv2.imwrite("Original_image.jpg", gray)
cv2.imwrite("+Our_compressed_img.jpg", compressed_img)

cv2.destroyAllWindows()
