from capter3 import selection_sort as sls
from capter4 import insertion_sort as ins

l1, _, _, = sls.sel_sort([5, 4, 3, 2, 1])
print(l1)

l2 = [5, 4, 3, 2, 1]
ins.insert_sort(l2)
print(l2)