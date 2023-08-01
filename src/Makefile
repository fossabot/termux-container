all :
	[ -e ruri/ruri ]||(cd ruri&&make static-bionic)
	[ -e container-console/container-console ]||(cd container-console&&make)
format :
	shfmt -i 2 -w container-rewrite
