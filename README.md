# EQzxc

EverQuest Zip Extract/Compressor is a go-based program for working with s3d, pfs, eq

## Features

[Working on first release](https://github.com/xackery/eqzxc/issues/5) (pfs loading, wld parsing, gtlf exporting)


## Goals
- load a pfs (*.eqg, *.s3d, or *.pfs)
- load a wld
- convert wld to gltf
    - (convert lights.wld to gltf too!)
    - (convert objects.wld to gltf too even if just meta info since it's a world - object reference)
    - (extract wld region data to a markdown file you can edit or maybe use meta data points in blender if i wanted to be slick)
- open gltf in blender, edit content, save
- convert gltf to wld
- save modified wld back into pfs
- have eq load the modified zone

## Inspiration

https://github.com/alimalkhalifa/visual-eq-gltf-exporter/blob/master/src/loaders/s3d.js

https://github.com/qmuntal/gltf

https://github.com/danwilkins/LanternExtractor/tree/development/0.2.0

https://github.com/danwilkins/LanternExtractor/blob/development/0.2.0/LanternExtractor/EQ/Wld/WldFile.cs

https://github.com/Zaela/EQGZoneImporter/blob/master/src/wld.cpp
