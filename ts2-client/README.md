# TS2 Client (Version 0.7)

This is a clone of the [TS2 client](https://github.com/ts2/ts2), the graphical user interface, which might need to be extended as part of the BahnTS2 project.

## Overview
**Train Signalling Simulator (TS2)** is a railway simulation game where you have to dispatch trains across an area and keep them on schedule. 

## Links
* TS2 Homepage - [ts2.github.io](http://ts2.github.io/)

## Status
TS2 is beta software, meaning it is playable, but still lacks many features that one would expect from a real simulation, will it ever be finished.
When starting TS2, you will be able to download simulations from the [ts2-data](https://github.com/ts2/ts2-data) server.

New simulations can be created with the editor provided with TS2.

## Installation
* Source installation:
    - Download and install Python v3 or above at [www.python.org](http://www.python.org).
    - Download and install PyQt v5 or above at [http://www.riverbankcomputing.co.uk](http://www.riverbankcomputing.co.uk).
    - Grab the sources from [GitHub](https://github.com/ts2/ts2/releases/tag/v0.7.0).
    - Run start-ts2.py with Python3
* If Python packages are missing, you can install them in a virtual environment (`venv`) (which does not require admin priviledges).
  Refer to the Python tutorial at https://docs.python.org/3/tutorial/venv.html

## Playing (QuickStart)
* Load a simulation from the _simulation_ directory (or the _data_ directory if you have installed from sources).
    If you want to load a simulation from a previous version of TS2, you will need to open it with the editor
    first and save it before loading it again in the main window.
* Route setting:
    - To turn a signal from red to green, you need to set a route from this signal to the next one.
    - To set a route left click on the signal and then to the next one. If you can create a route
        between these signals, the track between both signals is highlighted in white, the points are
        turned automatically for this route and the first signal color turn to yellow (or green if
        the second signal is already yellow or green).
    - To cancel a route, right-click on its first signal.
    - Routes are automatically cancelled by the first train passing through. However, you can set a
        persistent route by holding the shift key before clicking on the second signal. Persistent
        routes have a little white square next to their first signal.
    - Forcing route setting: It is possible to force a route setting by pressing _ctrl_ and _alt_ while
        clicking on the second signal. Beware as this will not check other conflicting routes and may result
        in train crashes or other unknown behaviour.
* Train information:
    - Click on a train code on the scene or on the train list to see its information on the
        "Train Information" window. The "Service information" window will also update.
* Station information:
    - Click on a platform on the scene to display the station timetable in the "Station information"
        window.
* Interact with trains:
    - Right click on a train code on the scene or on the "Train information" or on the "Train List"
    window to display the train menu. This menu enables you to:
        + Assign a new service to this train. Select the service in the popup window and click "Ok".
        + Reset the service. i.e. tell the train driver to stop at the first station again.
        + Reverse train. i.e. change the train direction.
    - Trains automatically change service when it is over (on drain demo BW01 becomes WB01 when reaching
    depot)
* You should see trains run, stop at red signals and at scheduled stations. They should depart at the
    departure time or after some time after the arrival time if the departure time is past.
* Scoring:
    Each time a train arrives late at a station, stops at the wrong platform or is routed to a wrong direction
    penalty points are added to the score.

## Developing

Whether you want to write your own simulation with the editor or to develop a new client to interact with TS2, 
head over to the 
[Technical Manual](https://gitlab.rz.uni-bamberg.de/swt/teaching/2020-ws/swt-pr1-2-m/group-c/-/blob/master/ts2-server/docs/ts2-technical-manual.pdf)

## Change log

### Version 0.7:
- New client-server architecture:
    - Multiplayer by connecting several players on the same simulation
    - Websocket API to interact with the simulation

### Version 0.6:
- New release that is PyQt5 = Python3 only
- Announce move and infrastructure to github
- New home page, et all
- Also Data is now json and in ts2-data
- Ability to download sims / signals from ts2-data
- Trains shunting
- Trains splitting (joining postponed)
- New parametric signals
- French BAL signals (except blinking)
- UK Shunt signals
- UI improvements (toolbars, etc.)
- Catch exceptions when loading sims

### Version 0.5:
- Last PyQt4 version
- Improved editor including the following features 
    - Multi-selection
    - Copy/Paste
    - Mass setting of properties
    - Resizing of platform items with mouse
- New signals with :
    - Short length 
    - Freely positionable berth
    - New signal types, including UK 4 aspects signals
