N = int(input())
C = list(map(int, input().split()))
Q = int(input())

sell = 0

z = 0
s = 0

min_s_C = 1000000000
min_z_C = 1000000000

for i in range(0, N):
  if i % 2 == 0:
    min_s_C = min(C[i], min_s_C)
  else:
    min_z_C = min(C[i], min_z_C)

for _ in range(0, Q):
  query = list(map(int, input().split()))
  if query[0] == 1:
    x = query[1] - 1
    a = query[2]
    if x % 2 == 0:
      card_x = C[x] - z - s
    else:
      card_x = C[x] - z
    if card_x >= a:
      C[x] -= a
      sell += a
      if x % 2 == 0:
        min_s_C = min(C[x], min_s_C)
      else:
        min_z_C = min(C[x], min_z_C)
  elif query[0] == 2:
    a = query[1]
    if min_s_C - s - z >= a:
      s += a
  elif query[0] == 3:
    a = query[1]
    if min(min_s_C - s - z, min_z_C - z) >= a:
      z += a

for i in range(0, N):
  if i % 2 == 0:
    sell += s

sell += z * N

print(sell)


