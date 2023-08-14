import numpy as np
from keras.models import Sequential
from keras.layers.core import Dense
import time

# Start timer
start = time.time()

# NN model generation and training
training_data = np.array([[0,0],[0,1],[1,0],[1,1]], "float32")
target_data = np.array([[0],[1],[1],[0]], "float32")

model = Sequential()
model.add(Dense(16, input_dim=2, activation='relu'))
model.add(Dense(1, activation='sigmoid'))
model.compile(loss='mean_squared_error')
model.fit(training_data, target_data, epochs=1000, verbose=0)

# End timer
end = time.time()
print("NN: %s seconds" % (end - start))