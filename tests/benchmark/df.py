import pandas as pd
import numpy as np
import time

# Start timer
start = time.time()

# Get the Iris dataset
df = pd.read_csv("../../tests/misc/iris_data.csv", names=['sepal length','sepal width','petal length','petal width','target'])

shape = df.shape
val = df.at[67, 'sepal width']
df.sort_values(by=['sepal length'])
df.pop('petal width')

# End timer
end = time.time()
print("Dataframe: %s seconds" % (end - start))
 