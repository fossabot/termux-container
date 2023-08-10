all :
	[ -e ruri/ruri ]||(cd ruri&&make static-bionic)
	[ -e container-console/container-console ]||(cd container-console&&make)
format :
	shfmt -i 2 -w container
gettext :
	xgettext -o container.pot -L Shell --keyword=po_getmsg --no-wrap container
	msginit -i container.pot -l zh_CN.UTF-8 --no-wrap -o container.po
	rm container.pot
	mv container.po ../lang/container-newlang.po
