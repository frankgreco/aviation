# aviation
> open aviation data from the faa

![build and deploy project](https://github.com/frankgreco/aviation/workflows/build%20and%20deploy%20project/badge.svg?branch=master)


## download
```
INFO[0000] retrieving archive file from url              url="http://registry.faa.gov/database/ReleasableAircraft.zip"
INFO[0000] successfully retrieved archive file from url  fields.time=388.4231ms url="http://registry.faa.gov/database/ReleasableAircraft.zip"
INFO[0004] unzipping archive
INFO[0004] found file in archive                         file=ardata.pdf
INFO[0004] found file in archive                         file=ACFTREF.txt
INFO[0004] found file in archive                         file=DEALER.txt
INFO[0004] found file in archive                         file=DEREG.txt
INFO[0006] found file in archive                         file=DOCINDEX.txt
INFO[0006] found file in archive                         file=ENGINE.txt
INFO[0006] found file in archive                         file=MASTER.txt
INFO[0006] found file in archive                         file=RESERVED.txt
INFO[0007] uploading files to aws s3
INFO[0032] successfully wrote file to aws s3             file=DEREG.txt key=6-20-2020/DEREG.txt
INFO[0032] successfully wrote file to aws s3             file=DOCINDEX.txt key=6-20-2020/DOCINDEX.txt
INFO[0032] successfully wrote file to aws s3             file=ENGINE.txt key=6-20-2020/ENGINE.txt
INFO[0039] successfully wrote file to aws s3             file=MASTER.txt key=6-20-2020/MASTER.txt
INFO[0040] successfully wrote file to aws s3             file=RESERVED.txt key=6-20-2020/RESERVED.txt
INFO[0040] successfully wrote file to aws s3             file=ardata.pdf key=6-20-2020/ardata.pdf
INFO[0041] successfully wrote file to aws s3             file=ACFTREF.txt key=6-20-2020/ACFTREF.txt
INFO[0041] successfully wrote file to aws s3             file=DEALER.txt key=6-20-2020/DEALER.txt
INFO[0041] successfully uploaded all files to aws s3     fields.time=34.933225009s
```