## Translation:
For example, create a Chinese translation file:      
```
xgettext -o container.pot -L Shell --keyword=po_getmsg --no-wrap container
msginit -i container.pot -l zh_CN.UTF-8 --no-wrap -o container.po
```
Do not compile the *.po file to *.mo, because termux does not support multiple languages.       
Translation files are in lang directory.      
### (有兴趣的话可以一起写中二文翻译，因为作者不够中二)
## Adding a new mirror:
See [rootfstool](https://github.com/Moe-hacker/rootfstool/blob/main/CONTRIBUTING.md)