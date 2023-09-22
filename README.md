# healpix

A Go-lang implementation of the HEALPix projection and pixelization of a sphere. This package is **early stage** and not yet recommended for integration. However, features incorporated into the master branch generally have good test coverage.

The HEALPix pixelization is 'equal-area' for each pixel: a pixel at the equator of the sphere contains the same resolution of data as a pixel at the poles (unlike many pixelizations based on more common projections). The resolution increments as a scaled, squared power of 2, and this package supports spheres divided into up to 12*((2^29)^2) pixels, i.e. 29 different levels of resolution.

HEALPix supports two basic pixel numbering schemes:

1. **Nested** - useful for efficient querying of nearest neighbors for individual pixels
2. **Ring** - useful for efficient spherical harmonic computations

A third numbering scheme, **Nested Unique**, allows for the selection of pixels at different scales of resolution, allowing for 'sparse' HEALPix queries and data set representations in addition to the default 'dense' representation.

## Roadmap

- [x] - Querying nearest neighbors
- [ ] - Support 'Nested Unique' pixel numbering (for multiresolution)
- [ ] - Support Cartesian 3-vector 'positions'
- [ ] - Querying discs
- [ ] - Querying polygons
- [ ] - Multiresolution pixel range sets

## References

The following prior works were studied carefully to aid the implementation of this package.

Gorski, et al. "HEALPix: A Framework for High-Resolution Discretization and Fast Analysis of Data Distributed on the Sphere" *The Astrophysical Journal*, 622:759–771, 2005 April 1. [Link](https://iopscience.iop.org/article/10.1086/427976/pdf)

I. Martinez-Castellanos, et al. "Multiresolution HEALPix Maps for Multiwavelength and Multimessenger Astronomy" *The Astronomical Journal*, vol. 63, no. 6, 2022 May 11. [Link](https://iopscience.iop.org/article/10.3847/1538-3881/ac6260#ajac6260s2)

Reinecke & Hivon. "Efficient data structures for masks on 2D grids" *Astronomy and Astrophysics*, vol. 580, 2015 August 19. [Link](https://www.aanda.org/articles/aa/full_html/2015/08/aa26549-15/aa26549-15.html#:~:text=Efficient%20data%20structures%20for%20masks%20on%202D%20grids,...%207%207.%20Generation%20from%20analytical%20shapes%20)