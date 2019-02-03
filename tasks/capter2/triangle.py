def triangle(x, y, z):
    if (x+y>z) and (y+z>x) and (z+x>y):
        if (x==y) and (y==z) and (z==x):
            return "正三角形"
        elif (x==y) or (y==z) or (z==x):
            return "二等辺三角形"
        else:
            return "三角形"
    else:
        return "三角形ではない"

print(triangle(4, 4, 4))
print(triangle(4, 4, 7))
print(triangle(4, 5, 6))
print(triangle(2, 3, 6))