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

yValues3 = []
xValues3 = []
xController3 = 0

#getting data
file = open("NewResults_LF_Compression.data", "r") 
for line in file: 
	j = json.loads(line)
	xValues.append(xController)
	yValues.append(j['metrics'][0]['value']/j['metrics'][0]['count'])
	xController = xController + 5

file2 = open("NewResults_LF_NoCompression.data", "r") 
for line2 in file2: 
	j2 = json.loads(line2)
	xValues2.append(xController2)
	yValues2.append(j2['metrics'][0]['value']/j2['metrics'][0]['count'])
	xController2 = xController2 + 5

file3 = open("NewResults_LF_CacheAndCompression.data", "r") 
for line3 in file3: 
	j3 = json.loads(line3)
	xValues3.append(xController3)
	yValues3.append(j3['metrics'][0]['value']/j3['metrics'][0]['count'])
	xController3 = xController3 + 5

#plotting graph
fig, ax = plt.subplots()
ax.plot(xValues, yValues, label='Compression')
ax.plot(xValues2, yValues2, label='NoCompression')
ax.plot(xValues3, yValues3, label='CacheAndCompression')
ax.legend()
ax.set(xlabel='Time (s)', ylabel='Response time (ms)', title='High-Frequency Workload')
ax.grid()
fig.savefig("high_frequency.png")
plt.show()
