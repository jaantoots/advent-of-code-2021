z = (a[0] + 15) * 26 + a[1] + 5
x = (a[2] + 6) % 26 - 14
w = a[3]
if x == w { x = 0 } else { x = 1 }
z *= 25*x + 1
z += (w + 7) * x
x = (a[4] + 9) % 26 - 7
w = a[5]
if x == w { x = 0 } else { x = 1 }
z *= 25*x + 1
z += (w + 6) * x
z *= 26
z += a[6] + 14
z *= 26
z += a[7] + 3
x = (a[8] + 1) % 26 - 7
w = a[9]
if x == w { x = 0 } else { x = 1 }
z *= 25*x + 1
z += (w + 3) * x
x = z % 26 - 8
z /= 26
w = a[10]
if x == w { x = 0 } else { x = 1 }
z *= 25*x + 1
z += (w + 4) * x
x = z % 26 - 7
z /= 26
w = a[11]
if x == w { x = 0 } else { x = 1 }
z *= 25*x + 1
z += (w + 6) * x
w = a[12]
x = z % 26 - 5
z /= 26
if x == w { x = 0 } else { x = 1 }
z *= 25*x + 1
z += (w + 7) * x
w = a[13]
x = z % 26 - 10
z /= 26
if x == w { x = 0 } else { x = 1 }
z *= 25*x + 1
z += (w + 1) * x
