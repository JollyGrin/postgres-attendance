DCL Attendance Idea:

- create a script that can be loaded into any DCL scene by the owner
- when a scene has a tracking script enabled, immediately starts tracking to attendance table
- backend will chronjob Decentraland Events page to scrape all visible events, add to db in events table

TODO 
- [ ]  api request to dcl events, and populate events table
- [x] migration, add 137,-2 to data


Needed APIs

- [ ] Get attendance records for location between start/end time



## Dedupe the attendance log
I already have a way to record enter/exit, but this will be triggered by multiple people. I need a way to check if the record already exists within a 5 second window before adding. 

