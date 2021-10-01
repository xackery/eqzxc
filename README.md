# EQzxc

EverQuest Zip Extract/Compressor is a go-based program for working with s3d, pfs, eq

## Features

[Working on first release](https://github.com/xackery/eqzxc/issues/5) (pfs loading, wld parsing, gtlf exporting)


## Goals
- run eqzxc, target a pfs archive (*.eqg, *.s3d, *.pak, or *.pfs)
- parse wld data, convert to a raw format
	- .map for quake 3 map editing?
	- .gltf for 3ds/maya/blender editing?
	- .toml for meta data tweaking that isn't supported by 3rd party programs
		- lights.wld data
		- objects.wld data refs    
- run eqzxc, target a extracted pfs archive (_<name>.<ext>/ dir)
	- parse .map files, toml, etc, and repack to eqg/s3d/pak
	
## Inspiration

https://github.com/alimalkhalifa/visual-eq-gltf-exporter/blob/master/src/loaders/s3d.js

https://github.com/qmuntal/gltf

https://github.com/danwilkins/LanternExtractor/tree/development/0.2.0

https://github.com/danwilkins/LanternExtractor/blob/development/0.2.0/LanternExtractor/EQ/Wld/WldFile.cs

https://github.com/Zaela/EQGZoneImporter/blob/master/src/wld.cpp
