// Copyright (c) 2016 ECS Team, Inc. - All Rights Reserved
// https://github.com/ECSTeam/cloudfoundry-top-plugin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package appView

const HelpText = `
**App Stats View**

App list view shows a list of all deployed applications regardless of
state. The full set of stats may not be available until the warm-up
period is complete.  After the warm-up period, if any applications are
found not to be in the desired state (e.g., instances set to 4 but
only 3 are running) an alert will be displayed and the application will
be colored red.

**Header information:**

  Evnts        - Total number of events received by the foundation.
  Warm-up      - It can take up to 30 seconds to receive all event
                 information before stats are accurate.
  Duration     - Amount of time stats have been collected.
  Target       - The target URL of monitored foundation.
  CPU (Used)   - Amount of CPU consumed by all app instances.
  CPU (Max)    - Sum of CPU capacity across all cells.
  Mem (Used)   - Amount of memory actually in use by all app instances.
  Mem (Max)    - Sum of memory capacity across all cells.
  Mem (Rsrvd)  - Amount of requested memory for all application
                 marked as started (app mem * instance count).
  Dsk (Used)   - Amount of disk actually in use by all app
                 instances.
  Dsk (Max)    - Sum of disk capacity across all cells.
  Dsk (Rsrvd)  - Amount of requested disk for all applications
                 marked as started (app mem * instance count).
  Apps (total) - Number of applications deployed to foundation.
  Cntrs        - Number of reporting containers.
 
**Application list stats:**

  APPLICATION - Application name
  SPACE - Space name
  ORG - Organization name
  RCR - Total reporting Containers
  CPU%% - Total CPU percent consumed by all containers
  MEM - Total memory used by all containers
  DISK - Total disk used by all containers
  RESP - Avg response time in milliseconds over last 60 seconds
  LOGS - Total number of log events for all instance of app
  REQ/1 - Number of HTTP(S) request/responses in last 1 second
  REQ/10 - Number of HTTP(S) request/responses in last 10 seconds
  REQ/60 - Number of HTTP(S) request/responses in last 60 seconds
  TOT-REQ - Count of all of the HTTP(S) request/responses
  2XX - Count of HTTP(S) responses with status code 200-299
  3XX - Count of HTTP(S) responses with status code 300-399
  4XX - Count of HTTP(S) responses with status code 400-499
  5XX - Count of HTTP(S) responses with status code 500-599

NOTE: The HTTP counters are based on traffic through the 
go-router.  Applications that talk directly container-to-
container will not show up in the REQ/TOT-REQ/nXX counters.

**Display: **
Press 'd' to select data view.

**Order / Sort display: **
Press 'o' to show the sort order window allowing multi-column
sorting of any column.

**Clear stats: **
Press shift-C to clear the statistics counters.

**Clipboard menu: **
Press 'c' when a row is selected to open the clipboard menu.
This will copy to clipboard a command you can paste in 
terminal window later.

**Pause display update:**
Press 'p' to toggle pause display update.  When display update is
paused top will continue to capture statstics and display updated
values when unpaused.

**Filter display: **
Press 'f' to show the filter window which allows for filtering
which rows should be displayed

**Reload metadata: **
Press 'r' to force a reload of metadata for app/space/org.  The
metadata is loaded at startup and attempts to stay current by
recognizing when specific data needs to be reloaded. However there
can be circumstances were data becomes stale.

**Refresh screen interval: **
Press 's' to set the sleep time between refreshes. Default
is 1 second.  Valid values are 0.1 - 60.  The refresh interval only
effects how often the client screen is refreshed, it has no effect
on frequency the foundation delivers events. Top uses passive
monitoring for stats, a faster refresh interval will not introduce
additonal load on the CF foundation.

**Select application detail: **
Press UP arrow or DOWN arrow to highlight an application row.
Press ENTER to select the highlighted application and show
additional detail.

**Scroll columns into view: **
Press RIGHT or LEFT arrow to scroll the columns into view if the
window is not wide enough to view all columns.  You can also resize
terminal window to show more columns/rows (resize of cmd.exe window
is not supported on windows while top is running).

**Refresh: **
Press SPACE to force an immediate screen refresh.

**Quit: **
Press 'q' to quit application.

**Log Window: **
Press shift-D to open log window.  This shows internal top
logging messages.  This window will open automatically if any error
message is logged (e.g., connection timeouts).
`
