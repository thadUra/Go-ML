import pandas as pd
import numpy as np
from sklearn.cluster import KMeans
import time

# Reference for this code: https://www.askpython.com/python/examples/principal-component-analysis

# PCA function
def PCA(X , num_components):
    X_meaned = X - np.mean(X , axis = 0)
    cov_mat = np.cov(X_meaned , rowvar = False)
    eigen_values , eigen_vectors = np.linalg.eigh(cov_mat)
    sorted_index = np.argsort(eigen_values)[::-1]
    sorted_eigenvalue = eigen_values[sorted_index]
    sorted_eigenvectors = eigen_vectors[:,sorted_index]
    eigenvector_subset = sorted_eigenvectors[:,0:num_components]
    X_reduced = np.dot(eigenvector_subset.transpose() , X_meaned.transpose() ).transpose()
    return X_reduced
 
# Get the Iris dataset
data = pd.read_csv("../../tests/misc/iris_data.csv", names=['sepal length','sepal width','petal length','petal width','target'])
x = data.iloc[:,0:4]
target = data.iloc[:,4]
 
# Run PCA
mat_reduced = PCA(x , 2)

# Start timer
start = time.time()

kmeans = KMeans(n_clusters = 3, init = 'k-means++', max_iter = 500)
kmeans.fit(mat_reduced)


# End timer
end = time.time()
print("K-Means: %s seconds" % (end - start))
 