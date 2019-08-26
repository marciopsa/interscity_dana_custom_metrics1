import matplotlib
import matplotlib.pyplot as plt
import json

#var prepping
list = []
yValues = []
xValues = []
xController = 0

yValues2 = []
xValues2 = []
xController2 = 0

#getting data
file = open("resultsDC2_sgbd_sample2.data", "r") 
for line in file: 
	j = json.loads(line)
	xValues.append(xController)
	yValues.append(j['metrics'][0]['value']/j['metrics'][0]['count'])
	xController = xController + 5

file2 = open("resultsDC3_cache_sample2.data", "r") 
for line2 in file2: 
	j2 = json.loads(line2)
	xValues2.append(xController2)
	yValues2.append(j2['metrics'][0]['value']/j2['metrics'][0]['count'])
	xController2 = xController2 + 5

#plotting graph
fig, ax = plt.subplots()
ax.plot(xValues, yValues, label='No cache')
ax.plot(xValues2, yValues2, label='Cache')
ax.legend()
ax.set(xlabel='Time (s)', ylabel='Response time (ms)', title='High-Frequency Workload')
ax.grid()
fig.savefig("high_frequency.png")
plt.show()
