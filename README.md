THIS IS A WOWSIM FORK, I AM NOT AFFILIATED WITH THEM IN ANY WAY. JUST A FUN PERSONAL PROJECT. THE ORIGINAL SIM IS HERE:  https://github.com/wowsims/classic

----------------------------------------------------------------------------------------------------------------------------------------------------------
INSTRUCTIONS TO DOWNLOAD USING BASH:

1. INSTALL TOOLS

 Git: https://git-scm.com/download/win

 Go 1.23.x (not 1.24+): https://go.dev/dl/ — pick go1.23.x.windows-amd64.msi

 Node 20 or 22 LTS, not Node 24: https://nodejs.org/

2. CHECK VERSIONS

 git --version

 go version

 node -v

 npm -v

Should be Go1.23, node v20 or v22

3. CLONE REPO

 cd ~

 git clone https://github.com/Sanctumx/classic.git

 cd classic

4. INSTALL DEPENDENCIES

 npm install

5. BUILD SIM

 make wasm

6. START UI

 npx vite --port 8080 --host 127.0.0.1

7. OPEN SIM

 http://127.0.0.1:8080/classic/hunter/ in your browser

8. HOW TO UPDATE

 cd ~/classic
 
 git remote set-url origin https://github.com/Sanctumx/classic.git
 
 git fetch origin
 
 git checkout master
 
 git pull origin master
 
 make wasm

------------------------------------------------------------------------------------------------------------------------------------------


