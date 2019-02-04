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

def bubble_sort_v2(l):
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

        if exchn_c == 0:
            break

    print("比較回数(合計)：{}".format(cmpr_sum))
    print("交換回数(合計)：{}".format(exchn_sum))

if __name__=='__main__':
    '''
    l = [9, 2, 7, 4, 5]
    bubble_sort(l)
    print(l)
    '''

    '''
    print("--bubble_sort--")
    l1 = [5, 7, 9, 2, 4]
    bubble_sort(l1)
    print(l1)
    print("--bubble sort--")
    print()
    '''

    '''
    print("--bubble_sort_v2--")
    l2 = [5, 7, 9, 2, 4]
    bubble_sort_v2(l2)
    print(l2)
    print("--bubble sort_v2--")
    print()
    '''

    print("--5.3 1--")
    l3 = [8, 6, 5, 3, 1]
    bubble_sort_v2(l3)
    print(l3)

    print("--5.3 2--")
    l4 = [1, 3, 8, 6, 5]
    bubble_sort_v2(l4)
    print(l4)
