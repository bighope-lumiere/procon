from math import gcd

K = int(input())

ans = 0
for a in range(1, K+1):
    for b in range(1, K+1):
        res = gcd(a, b)
        for c in range(1, K+1):
            ans += gcd(res, c)

print(ans)
