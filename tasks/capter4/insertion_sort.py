# TODO:見直し必須
def insert_sort(l):
    if len(l) < 1:
        return

    for i in range(len(l)-1):
        tmp = l[i+1] # 整列済み部分の１つ後ろの要素

        # 整列済み部分の後ろからtmpを挿入するところをサーチ
        for j in range(i, -1, -1):
            if l[j] >= tmp:
                l[j], l[j+1] = tmp, l[j]

if __name__ == '__main__':
    print("--insertion sort start--")
    l = [6, 3, 8, 5, 7]
    insert_sort(l)
    print(l) # 参照渡しだから
    print("--insertion sort end--")