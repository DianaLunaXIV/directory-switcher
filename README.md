# OpenMW Directory Switcher

A Go application for switching the active `config` and `data` directories between base OpenMW and Starwind on Steam Deck.

# Prerequisites, Motivations, and Caveats

This application is intended to be used with the Flatpak version of OpenMW on Steam Deck, but could theoretically be 
used in any scenario where OpenMW's `config` and `data` directories are collocated. 

The application requires that you have followed Starwind's installation instructions and isolated its `config` and `data` folders in such a way that they can be copied. Rename those folders `starwindconfig` and `starwinddata` and copy them into the OpenMW Flatpak's directory (default `$HOME/.var/app/org.openmw.OpenMW`). 

On a typical non-Steam Deck installation, it's probably easier to set separate launchers up by pointing the `XDG_CONFIG_HOME` and `XDG_DATA_HOME` environment variables to Starwind's `config` and `data` directories. This may be possible to do using the Flatpak installation of OpenMW, but I was unable to pass these environment variables correctly in my initial attempts, hence the genesis of this application.

Go 1.23 is currently required to build the executable for this application. Releases will be available soon. 


# Installation

1. Clone the repository and run `go build`.
2. Copy the resulting program (`directory-switcher`) into the OpenMW Flatpak directory (`$HOME/.var/app/org.openmw.OpenMW`). It should be next to `config`, `data`, `starwindconfig`, and `starwinddata`.
3. Execute `directory-switcher` either in the terminal or by double clicking on it. This should rename `config` and `data` to `mwconfig` and `mwdata`, and rename `starwindconfig` and `starwinddata` to `config` and `data` respectively. 
4. Repeat step 3 to switch back and forth between Morrowind and Starwind. 
5. (Optional) Right click on `directory-switcher` and click `Add to Steam` to execute the application from Gaming Mode. 
