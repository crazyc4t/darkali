#!/bin/bash

# ______           _   _     _   _          _
# |  _  \         | | | |   | \ | |        | |
# | | | |__ _ _ __| |_| |__ |  \| | ___  __| | ___ _ __
# | | | / _` | '__| __| '_ \| . ` |/ _ \/ _` |/ _ \ '__|
# | |/ / (_| | |  | |_| | | | |\  |  __/ (_| |  __/ |
# |___/ \__,_|_|   \__|_| |_\_| \_/\___|\__,_|\___|_|


# Links to wget from

Hack="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/Hack.zip"

RobotoMono="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/RobotoMono.zip"

UbuntuMono="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/UbuntuMono.zip"

FiraCode="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/FiraCode.zip"

FiraMono="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/FiraMono.zip"

DejaVuSansMono="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/DejaVuSansMono.zip"

JetBrainsMono="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/JetBrainsMono.zip"

LiberationMono="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/LiberationMono.zip"

Noto="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/Noto.zip"

Iconsolata="https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/Iconsolata.zip" 

Download='---Downloading-Fonts---'

function banner() {

                            
echo $'                 ('         
echo $'         (  (    )\ ) ('     
echo $'  (     ))\ )(  (()/( )\ )'  
echo $'  )\ ) /((_|()\  ((_)|()/('  
echo $' _(_/((_))  ((_) _| | )(_))'
echo $'|   \)) -_)|  _/ _` || || |'
echo $'|_||_|\___||_| \__,_| \_, |' 
echo $'                      |__/'  
echo -e '\nMade By DarthNeder'
echo 'https://saidneder.tech/'
sleep 0.5
}                               

banner

fonts () {
echo -e '\nThe following type of fonts can be installed'
echo -e '- Hack\n- RobotoMono\n- UbuntuMono\n- FiraCode\n- FiraMono\n- DejaVu\n- JetBrains\n- LiberationMono\n- Noto\n- Iconsolata'
echo -e '\nRun the script with the desired font! "./nerdy.sh Hack"'
}

errorhand () {
	echo -e "\nNot able to install font, check the syntax of your argument and try again\nTry './nerdy.sh -h' for help"
}

help () {
	fonts
	echo -e 'Not sure what type of font you want to use?'
	echo -e 'You can go this website and get your test ride!'
	echo -e 'https://www.programmingfonts.org/'
}

if [[ $1 = '-h' ]] || [[ $1 = '--help' ]]; then
	help
elif [[ $1 = 'Hack' ]]; then 
	echo $Download
	echo $Hack
	wget $Hack
	echo 'Extracting to /usr/share/fonts'
	sudo unzip Hack.zip -d /usr/share/fonts
	rm -rf Hack.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'RobotoMono' ]]; then
	echo $Download
	echo $RobotoMono
	wget $RobotoMono
	echo 'Extracting to /usr/share/fonts'
	sudo unzip RobotoMono.zip -d /usr/share/fonts
	rm -rf RobotoMono.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'UbuntuMono' ]]; then
	echo $Download
	echo $UbuntuMono
	wget $UbuntuMono
	echo 'Extracting to /usr/share/fonts'
	sudo unzip UbuntuMono.zip -d /usr/share/fonts
	rm -rf UbuntuMono.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'FiraCode' ]]; then
	echo $Download
	echo $FiraCode
	wget $FiraCode
	echo 'Extracting to /usr/share/fonts'
	sudo unzip FiraCode.zip -d /usr/share/fonts
	rm -rf FiraCode.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'FiraMono' ]]; then
	echo $Download
	echo $FiraMono
	wget $FiraMono
	echo 'Extracting to /usr/share/fonts'
	sudo unzip FiraMono.zip -d /usr/share/fonts
	rm -rf FiraMono.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'DejaVu' ]]; then
	echo $Download
	echo $DejaVuSansMono
	wget $DejaVuSansMono
	echo 'Extracting to /usr/share/fonts'
	sudo unzip DejaVuSansMono.zip -d /usr/share/fonts
	rm -rf DejaVuSansMono.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'JetBrains' ]]; then
	echo $Download
	echo $JetBrainsMono
	wget $JetBrainsMono
	echo 'Extracting to /usr/share/fonts'
	sudo unzip JetBrainsMono.zip -d /usr/share/fonts
	rm -rf JetBrainsMono.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'LiberationMono' ]]; then
	echo $Download
	echo $LiberationMono
	wget $LiberationMono
	echo 'Extracting to /usr/share/fonts'
	sudo unzip LiberationMono.zip -d /usr/share/fonts
	rm -rf LiberationMono.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'Noto' ]]; then
	echo $Download
	echo $Noto
	wget $Noto
	echo 'Extracting to /usr/share/fonts'
	sudo unzip Noto.zip -d /usr/share/fonts
	rm -rf Noto.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
elif [[ $1 = 'Iconsolata' ]]; then
	echo $Download
	echo $Iconsolata
	wget $Iconsolata
	echo 'Extracting to /usr/share/fonts'
	sudo unzip Iconsolata.zip -d /usr/share/fonts
	rm -rf Iconsolata.zip
	fc-cache -fv
	echo 'Done!'
	exit 0
else
	errorhand
	exit 2
fi
