import cv2
import numpy as np

I = cv2.imread('F.png')


__,I = cv2.threshold(I,70,255,cv2.THRESH_BINARY)

cv2.namedWindow("Original", cv2.WINDOW_NORMAL)
cv2.imshow("Original",I)
cv2.waitKey(0)

def skeletonize(J):
    J1= J.copy() 
    skel = J.copy()

    n = 0
    skel[:,:] = 0
    kernel = cv2.getStructuringElement(cv2.MORPH_CROSS, (3,3))

    while True:
        eroded = cv2.morphologyEx(J1, cv2.MORPH_ERODE, kernel)
        temp = cv2.morphologyEx(eroded, cv2.MORPH_DILATE, kernel)
        temp  = cv2.subtract(J1, temp)
        skel = cv2.bitwise_or(skel, temp)
        J1[:,:] = eroded[:,:]
        n = n + 1
        if np.sum(J1 ==255) == 0:
            break
    print(n)
    return skel


Res = skeletonize(I)

cv2.namedWindow("Result", cv2.WINDOW_NORMAL)
cv2.imshow("Result",Res)
cv2.waitKey(0)

kernel = np.ones((5, 5), np.uint8)
E = cv2.morphologyEx(Res, cv2.MORPH_CLOSE, kernel)

cv2.namedWindow("E", cv2.WINDOW_NORMAL)
cv2.imshow("E",E)
cv2.waitKey(0)

cv2.destroyAllWindows()
