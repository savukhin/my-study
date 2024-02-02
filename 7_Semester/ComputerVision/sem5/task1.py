import cv2
import numpy as np

I = cv2.imread('Jorig.png')
Inoise = cv2.imread('JwithPepper.png')


__,I = cv2.threshold(I,70,255,cv2.THRESH_BINARY)

cv2.namedWindow("Inoise", cv2.WINDOW_NORMAL)
cv2.imshow("Inoise",Inoise)

cv2.namedWindow("Original", cv2.WINDOW_NORMAL)
cv2.imshow("Original",I)
cv2.waitKey(0)

kernel = np.ones((3, 3), np.uint8)
# R = cv2.erode(Inoise,kernel)
R = cv2.dilate(Inoise,kernel)

cv2.namedWindow("Result", cv2.WINDOW_NORMAL)
cv2.imshow("Result",R)
cv2.waitKey(0)

kernel1 = np.ones((5, 5), np.uint8)
# RR = cv2.dilate(R,kernel)
RR = cv2.erode(R,kernel)

cv2.namedWindow("Result", cv2.WINDOW_NORMAL)
cv2.imshow("Result",RR)
cv2.waitKey(0)

cv2.destroyAllWindows()
