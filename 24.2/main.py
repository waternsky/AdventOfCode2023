import sympy

fp = open("../24.1/input.24.1.txt", 'r')

class hailstone:
    def __init__(self, x, y, z, vx, vy, vz) -> None:
        self.x = x 
        self.y = y 
        self.z = z 
        self.vx = vx 
        self.vy = vy 
        self.vz = vz 


hails = []
for line in fp:
    position = line.split("@")[0]
    velocity = line.split("@")[1]
    parr = list(map(int, position.split(",")))
    varr = list(map(int, velocity.split(",")))
    h = hailstone(parr[0], parr[1], parr[2], varr[0], varr[1], varr[2])
    hails.append(h)


x, y, z, p, q, r = sympy.symbols("x y z p q r")

equations = []
for hail in hails:
    equations.append((x - hail.x) * (q - hail.vy) - (y - hail.y) * (p - hail.vx))
    equations.append((y - hail.y) * (r - hail.vz) - (z - hail.z) * (q - hail.vy))

ans = sympy.solve(equations)
print(ans)
print(ans[0][x] + ans[0][y] + ans[0][z])
