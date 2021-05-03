# Better KQB Scoreboard

Better KQB Scoreboard helps casters keep track of scores and display them beautifully within OBS. It is primarily targeted at IGL match casters, but will be adding support for custom matches as well. It allows you to directly import team data (teams, standings, avatars, etc.) directly from IGL so there is no room for error when setting up your stream. It will prompt you to record map wins and will tally the results for you and they will be displayed properly within the widget. Thus allowing you to focus on streaming and let the widget handle the legwork for you. 

KQB Scoreboard was written in Go and provides a browser widget interface for use with OBS. 

## Instructions
1. Download the new GUI release from [Github Releases Page](https://github.com/achhabra2/kqb-scoreboard/releases)
2. Unzip the file and run the App
3. Follow the on screen instructions for IGL or Custom Match Type
4. Create a browser source in OBS pointed at http://localhost:8080/static/
* Recommended dimensions are 1760x90 to get the look in the screen shot below
* Then add a green chroma filter to make it transparent
5. Profit


## Screen Shots and Demo
See the [Demo Video](https://www.youtube.com/watch?v=ZEOmpLE7CRs) for an idea of how it works. 

New Theme!
![Screen Shot 4](/screenshots/scoreboard-ss-4.PNG)

New Screenshot!
![Screen Shot 3](/screenshots/scoreboard-ss-3.png)


Theme Support 05/03/21
![Theme Support](/screenshots/theme-support.PNG)