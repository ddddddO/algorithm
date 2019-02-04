def fib_recursive(n):
    if n < 2:
        return n
    
    return fib_recursive(n-2) + fib_recursive(n-1)


def fib_dynamic_programing(n):
    if n < 2:
        return n
    
    fib_l = [0 for _ in range(n+1)]
    fib_l[0], fib_l[1] = 0, 1

    for i in range(2, n+1):
        fib_l[i] = fib_l[i-2] + fib_l[i-1]

    return fib_l[n]

if __name__=='__main__':
    print(fib_recursive(4)) # 0, 1, 1, 2, 3(=第4項)
    print(fib_recursive(7)) # 0, 1, 1, 2, 3, 5, 8, 13(=第7項)

    print(fib_dynamic_programing(4))
    print(fib_dynamic_programing(7))