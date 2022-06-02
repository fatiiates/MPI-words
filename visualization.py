import subprocess
import matplotlib.pyplot as plt

getcmd = lambda np : "mpirun --oversubscribe -np " + str(np) + " generator/bin/gen_data 1000000 100 100"

def generation_visualization():
    plt.ylabel('generation time(seconds)')

def counting_visualization():
    plt.ylabel('counting time(seconds)')
    global getcmd
    getcmd = lambda np : "cd counter && ./bin/counter " + str(np)

x = [i+1 for i in range(100)]
y = []



# setting the axes at the centre
fig = plt.figure()
ax = fig.add_subplot(1, 1, 1)
ax.spines['right'].set_color('none')
ax.spines['top'].set_color('none')
ax.xaxis.set_ticks_position('bottom')
ax.yaxis.set_ticks_position('left')


counting_visualization()
#generation_visualization()

for i in x:
    a = subprocess.Popen(getcmd(i), shell=True, stdout=subprocess.PIPE, stderr=subprocess.DEVNULL).stdout.read()
    y.append(float(a))
    print(str(i) + ". bitti")

# plot the function
plt.plot(x,y, 'r')
plt.xlabel('process number')

plt.show()
fig.savefig('count.png')
