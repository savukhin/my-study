import cv2
import numpy as np

l = cv2.imread("bird.jpeg")
cv2.namedWindow("original", cv2.WINDOW_NORMAL)
cv2.imshow("original", l)

row,col,ch = l.shape

gauss = np.random.normal(0, 25.0, (row,col,ch))
gauss = gauss.reshape(row, col, ch)
# cv2.namedWindow("original", cv2.WINDOW_NORMAL)
# cv2.imshow("original", l)

ll = l.astype("float")

noisy = ll + gauss

noisy = cv2.min((255,), cv2.max(noisy, (0)))
noisy = noisy.astype("uint8")

cv2.namedWindow("gauss", cv2.WINDOW_NORMAL)
cv2.imshow("gauss", gauss)

cv2.namedWindow("noisy", cv2.WINDOW_NORMAL)
cv2.imshow("noisy", noisy)

img_blur_7 = cv2.GaussianBlur(noisy, (7, 7), 4, 4)
cv2.namedWindow("result", cv2.WINDOW_NORMAL)
cv2.imshow("result", img_blur_7)

J = cv2.ximgproc.amFilter(noisy, noisy, 3.9, 0.2)
cv2.namedWindow("adaptive manifold filter", cv2.WINDOW_NORMAL)
cv2.imshow("adaptive manifold filter", J)

cv2.waitKey(0)

