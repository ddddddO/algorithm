def fib_recursive(n):
    if n < 2:
        return n
    
    print("fib_recursive")

    f = fib_recursive(n-2) + fib_recursive(n-1)
    return f


def fib_dynamic_programing(n):
    if n < 2:
        return n
    
    fib_l = [0 for _ in range(n+1)]
    fib_l[0], fib_l[1] = 0, 1
    cnt = 0

    for i in range(2, n+1):
        fib_l[i] = fib_l[i-2] + fib_l[i-1]
        cnt+=1

    return fib_l[n], cnt

if __name__=='__main__':
    '''
    print(fib_recursive(4)) # 0, 1, 1, 2, 3(=第4項)
    print(fib_recursive(7)) # 0, 1, 1, 2, 3, 5, 8, 13(=第7項)

    print(fib_dynamic_programing(4))
    print(fib_dynamic_programing(7))
    '''

    print("6.4 1")
    n = fib_recursive(5)
    print("第5項：{}".format(n))
    print()

    print("6.4 2")
    n2, cnt = fib_dynamic_programing(5)
    print("第5項：{}".format(n2))
    print("足し算の回数：{}".format(cnt))
    print()