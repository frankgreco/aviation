# aviation
![deploy](https://github.com/frankgreco/aviation/workflows/deploy/badge.svg)


## download

The download app will retrieve the [database](https://www.faa.gov/licenses_certificates/aircraft_certification/aircraft_registry/releasable_aircraft_download/) of registered aircraft from the FAA and place it into S3 for further processing. The FAA updates this file every day at 04:30 UTC and this application will run every day, as an AWS Lambda function, at 06:00 UTC.

The following is example an example log after a successful exectuion.

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

## load

The load app will update a PostgreSQL database every day from the updated FAA database. This will run as an AWS Lambda function and will be triggered after the completion of the _download_ application.