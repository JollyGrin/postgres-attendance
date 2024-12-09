DCL Attendance Idea:

- create a script that can be loaded into any DCL scene by the owner
- when a scene has a tracking script enabled, immediately starts tracking to attendance table
- backend will chronjob Decentraland Events page to scrape all visible events, add to db in events table

TODO 
- [ ]  api request to dcl events, and populate events table
- [x] migration, add 137,-2 to data
- [ ] improve docs how to instantiate fresh install with the `dump.sql`


# DB updates
the duration call is working to give durations, yet i need to combine timeslots if multiple enters are available with overlapping times. Right now i get multiple 10 minutes if no exit was found, but within 10 mins of each other.


Needed APIs

- [ ] Get attendance records for location between start/end time



## Dedupe the attendance log
I already have a way to record enter/exit, but this will be triggered by multiple people. I need a way to check if the record already exists within a 5 second window before adding. 

