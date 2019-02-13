field = [[0 for _ in range(7)] for _ in range(6)]
field[1][1] = 1

for row in range(1, 6):
    for colm in range(1, 7):
        if (row==1 and colm==1) or (row==1 and colm==4) or (row==4 and colm==4):
            continue
        field[row][colm] = field[row-1][colm] + field[row][colm-1]

if __name__=='__main__':
    print("最短経路の数え上げ問題")
    print(field)
    print("最短経路数：{}".format(field[-1][-1]))