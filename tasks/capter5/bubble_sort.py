def bubble_sort(l):
    length = len(l)
    cmpr_sum, exchn_sum = 0, 0
    
    for i in range(length-1):
        cmpr_c, exchn_c = 0, 0
        for j in range(length-1, i, -1):
            cmpr_c+=1
            if l[j-1] >= l[j]:
                l[j-1], l[j] = l[j], l[j-1]
                exchn_c+=1

        print("比較回数：{}".format(cmpr_c))
        print("交換回数：{}".format(exchn_c))
        print()

        cmpr_sum+=cmpr_c
        exchn_sum+=exchn_c

    print("比較回数(合計)：{}".format(cmpr_sum))
    print("交換回数(合計)：{}".format(exchn_sum))

if __name__=='__main__':
    l = [9, 2, 7, 4, 5]
    bubble_sort(l)
    print(l)