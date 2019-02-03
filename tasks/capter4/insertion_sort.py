def insert_sort(l):
    if len(l) < 1:
        return

    cmpr_sum = 0
    for i in range(len(l)-1):
        tmp = l[i+1] # 整列済み部分の１つ後ろの要素
        cmpr_n = 0

        # 整列済み部分の後ろからtmpを挿入するところをサーチ
        for j in range(i, -1, -1):
            cmpr_n+=1

            if l[j] >= tmp:
                l[j], l[j+1] = tmp, l[j]
            else:
                # 残りの整列済み部分のサーチを飛ばす
                break

        cmpr_sum+= cmpr_n
    print("比較回数(合計)：{}".format(cmpr_sum))

if __name__ == '__main__':
    print("--insertion sort start--")
    l = [6, 3, 8, 5, 7]
    insert_sort(l)
    print(l) # 参照渡しだから
    print("--insertion sort end--")
    print()

    print("--練習4.3 1--")
    l1 = [5, 8, 3, 6, 1]
    insert_sort(l1)
    print(l1)
    print()

    print("--練習4.3 2--")
    l2 = [8, 6, 5, 3, 1]
    insert_sort(l2)
    print(l2)
    print()

    print("--練習4.3 3--")
    l3 = [1, 3, 5, 6, 8]
    insert_sort(l3)
    print(l3)
    print()
