# K-means

This module is an implementation of sequential k-means clustering algorithm as presented in [Princeton university materials](http://www.cs.princeton.edu/courses/archive/fall08/cos436/Duda/C/sk_means.htm)

If you have some batch data and expect some more on the fly, this is algorithm just for you.

## Usage

```
	k := kmeans.New(5, kmeans.CanberraDistance, 0.1)

	// Seed gets initial clusters from batch of data
	k.Seed(batch)

	// Add adds new point to clusters, remembering all the previous data
	k.Add(o)

	// Alternatively you can use Addf which "forgets" old data. It uses second approach
	// and constant alpha (0.1) passed during creation of k to update cluster centers.
	k.Addf(o)

	// You can access clusters like this:
	for _, c := range k.C {
		// some code...
	}
```
