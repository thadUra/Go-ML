package dataframe

// Print prints the data from the dataframe type.
func (df *Dataframe) Print() {

}

// At returns a value given a label for the row and column.
func (df *Dataframe) At(row int, label string) {

}

// Iat returns a value given an index for the row and column.
func (df *Dataframe) Iat(row, col int) {

}

// Loc returns a row of data given a value.
func (df *Dataframe) Loc(value interface{}) {

}

// Iloc returns a row of data given an index value.
func (df *Dataframe) Iloc(idx int) {

}

// Iloc returns rows of data given a range of indices.
func (df *Dataframe) IlocRange(start, end int) {

}

// IsNull prints booleans to see if a value in the dataframe is NaN or not.
func (df *Dataframe) IsNull() {

}

// Iterrows returns something to allow iterating over the dataframe
func (df *Dataframe) Iterrows() {

}

// Shape returns the dimensionality (rows and columns) of the dataframe.
func (df *Dataframe) Shape() (int, int) {
	return df.rows, df.cols
}

// Count returns the number of non-null values in each column.
func (df *Dataframe) Count() {

}

// Unique returns the number of non-null values in each column.
// If 'axis' is '1', it grabs unique values at each row instead.
func (df *Dataframe) Nunique(axis int) {

}
