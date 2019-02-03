def sel_sort(l):
    cmpr_sum, exchn_sum = 0, 0
    j, length = 1, len(l)
    for i in range(length-1):
        cmpr_c, exchn_c = 0, 0
        min = l[i]
        tmp_n = i
        for k in range(j, length):
            cmpr_c+=1
            if min >= l[k]:
                min = l[k]
                tmp_n = k

        tmp = l[i]
        l[i] = min
        l[tmp_n] = tmp
        exchn_c+=1

        print("比較回数：{}".format(cmpr_c)); cmpr_sum+=cmpr_c
        print("交換回数：{}".format(exchn_c)); exchn_sum+=exchn_c
        print()

        j+=1
    
    return l, cmpr_sum, exchn_sum

def sel_sort_at_any_time(l):
    cmpr_sum, exchn_sum = 0, 0
    j, length = 1, len(l)
    for i in range(length-1):
        cmpr_c, exchn_c = 0, 0
        for k in range(j, length):
            cmpr_c+=1
            if l[i] >= l[k]:
                #tmp = l[k]
                #l[k] = l[i]
                #l[i] = tmp
                l[i], l[k] = l[k], l[i]

                exchn_c+=1

        print("比較回数：{}".format(cmpr_c)); cmpr_sum+=cmpr_c
        print("交換回数：{}".format(exchn_c)); exchn_sum+=exchn_c
        print()

        j+=1
    
    return l, cmpr_sum, exchn_sum


print("--例題3.1--")
sorted_l, cmpr_sum, exchn_sum = sel_sort([8, 4, 3, 9, 6])
print(sorted_l)
print("比較回数(合計)：{}".format(cmpr_sum))
print("交換回数(合計)：{}".format(exchn_sum))
print()

print("--例題3.2--")
sorted_any_l, cmpr_any_sum, exchn_any_sum = sel_sort_at_any_time([8, 4, 3, 9, 6])
print(sorted_any_l)
print("比較回数(合計)：{}".format(cmpr_any_sum))
print("交換回数(合計)：{}".format(exchn_any_sum))
print()

print("--練習問題3.4 1--")
s1, c1, e1 = sel_sort([9, 2, 7, 4, 5])
print(s1)
print("比較回数(合計)：{}".format(c1))
print("交換回数(合計)：{}".format(e1))
print()

print("--練習問題3.4 2--")
s2, c2, e2 = sel_sort_at_any_time([9, 2, 7, 4, 5])
print(s2)
print("比較回数(合計)：{}".format(c2))
print("交換回数(合計)：{}".format(e2))
print()

print("--練習問題3.4 3--")
s3, c3, e3 = sel_sort_at_any_time([2, 4, 5, 7, 9])
print(s3)
print("比較回数(合計)：{}".format(c3))
print("交換回数(合計)：{}".format(e3))
print()
