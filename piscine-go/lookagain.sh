
find .  -name '*.sh*' | sort -r | sed 's/.sh//g' | rev | cut -d '/' -f1 | rev 
