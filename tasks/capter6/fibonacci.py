def fib_recursive(n):
    if n < 2:
        return n
    
    return fib_recursive(n-2) + fib_recursive(n-1)


if __name__=='__main__':
    print(fib_recursive(4)) # 0, 1, 1, 2, 3(=第4項)
    print(fib_recursive(7)) # 0, 1, 1, 2, 3, 5, 8, 13(=第7項)