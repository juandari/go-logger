/*
Package pocketlog exposes an API to log your work.

First, instantiate a logger with pocketlog.New, and giving it a threshold level.
Nessages of lesser critically won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to log messages on three levels:
- Debug: mostly used to debug code, follow step-by-step processes
- Info: valuable messages providing instigts to the milestones of the a process
- Error: error messagese to understand what went wrong
*/
package pocketlog
