#! /usr/bin/python3
import re
import math

a1 = input()

answer = []
for aaa in range(0, int(a1)):
    list = []
    list2 = []
    input()
    for x in range(0, 8):
        c = input()
        tmp = []
        k = 0
        for y in c:
            if y == '#':
                tmp.append(k)
            k = k+1
        list.append(c)
        list2.append(tmp)

    tmpX = []
    a = 0
    b = 0
    ggg = 0
    for x in list2:
        if a == 0:
            tmpX = x
            a = len(x)
            b = len(x)
        else:
            a = b
            b = len(x)

        if a == 2 and b == 1:
            break
        else:
            tmpX = x
            ggg = ggg + 1

    answer.append([ggg+1, tmpX[0] + 2])

for x in answer:
    print(x[0], x[1])