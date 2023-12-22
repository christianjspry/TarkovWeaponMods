#Tarkov Weapon Mods Project

### About:
This is a project that will be a web application that fills a gap in the weapon building functionality in Escape From Tarkov: customizing weapon builds based off of trader levels.


### Todo: 
Front end
Figure out the rest of the DB models, such as how to link all the components.
- Current thought is to have a linking table between components that *could* be linked.
- How to handle when one item is incompatible with another item in the build?
Presets
Schema. PostgreSQL schema is making the models more unwieldy than they should be
More fields for upcoming and existing DB models:
- bool questRestricted or int || null questRestricted
- trader cost (account for the intellegence center? have that in the UI? send in post requests. or handle in state in UI?)
- trader Loyalty level
- Black Market average price? could this be doable? outside api connection?
Efficiency Rating for gear
- Rouble per ergo/recoil stat. maybe some consideration for weight

