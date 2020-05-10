# Better KQB Scoreboard

Better KQB Scoreboard helps casters keep track of the score and display it beautifully. KQB Scoreboard was written in Go and provides a front end interface for use with OBS. It will import the teams and standings from IGL to display within your cast. 

## Instructions
1. Download the entire release zip file from github
2. Unzip the directory and run the kqb-scoreboard executable (Mac and Windows compatible)
3. Create a browser source in OBS pointed at http://localhost:8080
* Within the browser source set the dimensions to 1920 x 400
* Then add a green chroma filter to make it transparent
4. Profit

### IMPORTANT NOTE FOR WINDOWS USERS
Unless you have the developer terminal on windows some parts of the app will not display 100% correctly, this is purely cosmetic and will have no affect on what the viewers see. On windows instead of using the arrow keys you must use the J K H L keys to navigate the menus. 

#### Windows Key Bindings (Will hopefully be improved)
- *J = Down*
- *K = Up*
- *H = Page Down*
- *L = Page Up*
- *Enter = Select*

### ALPHA LIMITATIONS
Note during this alpha phase there may be bugs and minor issues you run into, here is a list of current known issues: 
1. No way to subtract from scores - so please make sure you record map results correctly
2. Dark mode not working - currently the score boxes have a white background, looking to correctly implement a dark mode
3. Windows keybinding and display issues as per above


## Screen Shots and Demo
See the [Demo Video](https://www.youtube.com/watch?v=em6KXidrXVI) for an idea of how it works. 

![Screen Shot 1](/screenshots/scoreboard-ss-1.png)
![Screen Shot 2](/screenshots/scoreboard-ss-2.png)