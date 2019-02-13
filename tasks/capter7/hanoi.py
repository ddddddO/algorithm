def hanoi(k, start, yobi, end):
    if k >= 2:
        hanoi(k-1, start, end, yobi)
    print("{}軸の円盤を{}軸へ移動".format(start, end))
    if k >= 2:
        hanoi(k-1, yobi, start, end)

if __name__=='__main__':
    print("ハノイの塔")
    hanoi(2, 1, 2, 3) 