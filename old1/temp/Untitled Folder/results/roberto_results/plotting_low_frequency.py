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
file = open("resultsDC3_noCompress_sample2.data", "r") 
for line in file: 
	j = json.loads(line)
	xValues.append(xController)
	yValues.append(j['metrics'][0]['value']/j['metrics'][0]['count'])
	xController = xController + 5

file2 = open("resultsDC4_compressedData_sample2.data", "r") 
for line2 in file2: 
	j2 = json.loads(line2)
	xValues2.append(xController2)
	yValues2.append(j2['metrics'][0]['value']/j2['metrics'][0]['count'])
	xController2 = xController2 + 5

#plotting graph
fig, ax = plt.subplots()
ax.plot(xValues, yValues, label='No compression')
ax.plot(xValues2, yValues2, label='Compression')
ax.legend()
ax.set(xlabel='Time (s)', ylabel='Response time (ms)', title='Low-Frequency Workload')
ax.grid()
fig.savefig("low_frequency.png")
plt.show()
