import cv2
import numpy as np
import math

img = cv2.imread("Fugures.png")

cv2.namedWindow("Image", cv2.WINDOW_NORMAL)
cv2.imshow("Image", img)
cv2.waitKey(0)

dst = cv2.Canny(img, 50, 200, None,3)

cv2.namedWindow('Countors', cv2.WINDOW_NORMAL)
cv2.imshow('Countors',dst)
cv2.waitKey(0)

lines = cv2.HoughLines(dst, 1, np.pi / 180, 150, None, 0, 0)

result = np.zeros(np.shape(img),"uint8")

if lines is not None:
    for i in range(0, len(lines)):
        r = lines[i][0][0]
        theta = lines[i][0][1]
        a = np.cos(theta)
        b = np.sin(theta)
        x0 = a * r
        y0 = b * r

        pt1 = (int(x0 + 2000*(-b)), int(y0 + 2000*(a)))
        pt2 = (int(x0 - 2000*(-b)), int(y0 - 2000*(a)))
        cv2.line(result, pt1, pt2, (150,0,0), 1, cv2.LINE_AA)

cv2.namedWindow('Result', cv2.WINDOW_NORMAL)
cv2.imshow('Result', cv2.min((255,255,255),result + img))
cv2.waitKey(0)

cv2.cos

cv2.destroyAllWindows()