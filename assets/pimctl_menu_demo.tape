Set Shell zsh
Set Framerate 60
Set WindowBar Colorful
Set Margin 20
Set MarginFill "#674EFF"
Set BorderRadius 10
Set Width 1920
Set Height 1024

Output pimctl_menu_demo.gif
Output pimctl_menu_demo.mp4
Output pimctl_menu_demo.webm

Hide
Type "export AZURE_CONFIG_DIR=$ALEXW_AZURE_CONFIG_DIR"
Enter
Type "export PIMCTL_PROFILE=alexw"
Enter
Show

Type "pimctl account show"
Enter
Sleep 2s
Type "pimctl group eligible menu"
Enter
Sleep 4s
Down
Sleep 1s
Enter
Sleep 2s
Type "I need access to update an important system."
Enter
Sleep 10s

Hide
Type "unset AZURE_CONFIG_DIR"
Enter
Type "unset PIMCTL_PROFILE"
Enter
Show

Type "pimctl account show"
Enter
Sleep 2s
Type "pimctl group approval menu"
Enter
Sleep 5s
Enter
Sleep 5s
Enter
Sleep 2s
Type "It's fine, I've asked him to."
Enter
Sleep 10s

