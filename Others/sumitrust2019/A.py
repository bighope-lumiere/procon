m1, d1 = map(int, input().split())
m2, d2 = map(int, input().split())
if (m2 - m1) % 12 == 1:
    print(1)
else:
    print(0)
