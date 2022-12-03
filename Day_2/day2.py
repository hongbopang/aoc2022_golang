# -*- coding: utf-8 -*-
"""
Created on Thu Nov 17 18:00:47 2022

@author: hongb
"""
f = open("input.txt")
f = ["A Y", "B X", "C Z"]
ans1 = 0

for line in f:
    symbols = line.split()
    op, me = ord(symbols[0]) - ord("A"), ord(symbols[1]) - ord("X")
    result = me - op
    
    ans1 += 3 * (result + 1) + me + 1
    print(ans1)
print(ans1)
        
        