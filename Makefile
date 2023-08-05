O = out
.PHONY: all
all: show-greetings install-dependence update-code build
show-greetings:
	@printf "\033[1;38;2;254;228;208m"
	@printf "                  _________\n"
	@printf "                 /        /\\ \n"
	@printf "                /        /  \\ \n"
	@printf "               /        /    \\ \n"
	@printf "              /________/      \\ \n"
	@printf "              \\        \\      /\n"
	@printf "               \\        \\    /\n"
	@printf "                \\        \\  /\n"
	@printf "                 \\________\\/\n"
	@printf "═╔╝╔═╝╔═║╔╔ ║ ║║ ║  ╔═╝╔═║╔═ ═╔╝╔═║╝╔═ ╔═╝╔═║\n"
	@printf " ║ ╔═╝╔╔╝║║║║ ║ ╝ ═╝║  ║ ║║ ║ ║ ╔═║║║ ║╔═╝╔╔╝\n"
	@printf " ╝ ══╝╝ ╝╝╝╝══╝╝ ╝  ══╝══╝╝ ╝ ╝ ╝ ╝╝╝ ╝══╝╝ ╝\n\n"
install-dependence:
	@printf "\033[1;38;2;254;228;208m[+] Install dependents.\033[0m\n"&&sleep 1s
	@pkg install ndk-multilib-native-static tsu coreutils p7zip gettext tar unzip zip git wget dpkg curl nano proot axel termux-tools util-linux pv gawk whiptail clang ndk-sysroot ndk-multilib libc-client-static libcap-static binutils	
build :update-code
	make -C src
	@mkdir -pv $(O)
	@mkdir -pv $(O)/doc/termux-container
	@mkdir -pv $(O)/share/termux-container
	@cp -rv doc/* $(O)/doc/termux-container/
	@cp -rv share $(O)/share/termux-container
	@mkdir -pv $(O)/bin/
	@cp -v src/container $(O)/bin/container
	@cp -v src/rootfstool/rootfstool $(O)/bin/rootfstool
	@mv -v src/container-console/container-console $(O)/bin/container-console
	@mv -v src/ruri/ruri $(O)/bin/ruri
update-code:
	@git submodule update --init
install:build
	@printf "\033[1;38;2;254;228;208m[+] Install.\033[0m\n"&&sleep 1s
	@cp -v $(O)/bin/* /data/data/com.termux/files/usr/bin/
	@cp -rv $(O)/share/termux-container /data/data/com.termux/files/usr/share/
	@cp -rv $(O)/doc/* /data/data/com.termux/files/usr/share/doc
pack-deb:
	@mkdir -v $(O)/deb
	@mkdir -pv $(O)/deb/data/data/com.termux/files/usr
	@mkdir -pv $(O)/deb/data/data/com.termux/files/usr/bin/
	@mkdir -pv $(O)/deb/data/data/com.termux/files/usr/share/
	@mkdir -pv $(O)/deb/data/data/com.termux/files/usr/share/doc
	@cp -v $(O)/bin/* $(O)/deb/data/data/com.termux/files/usr/bin/
	@cp -rv $(O)/share/termux-container $(O)/deb/data/data/com.termux/files/usr/share/
	@cp -rv $(O)/doc/* $(O)/deb/data/data/com.termux/files/usr/share/doc
	@cp -rv dpkg-conf $(O)/deb/DEBIAN
	@printf "\033[1;38;2;254;228;208m[+] Build packages.\033[0m\n"&&sleep 1s
	@cd $(O)/deb&&chmod -Rv 755 DEBIAN&&chmod -Rv 777 data/data/com.termux/files/usr/bin
	@cd $(O)/deb&&dpkg -b . ../../termux-container.deb
	@printf "\033[1;38;2;254;228;208m    .^.   .^.\n"
	@printf "    /⋀\\_ﾉ_/⋀\\ \n"
	@printf "   /ﾉｿﾉ\\ﾉｿ丶)|\n"
	@printf "  |ﾙﾘﾘ >   )ﾘ\n"
	@printf "  ﾉノ㇏ Ｖ ﾉ|ﾉ\n"
	@printf "        ⠁⠁\n"
	@printf "\033[1;38;2;254;228;208m[*] Build done,package: termux-container.deb\033[0m\n"
clean:
	@printf "\033[1;38;2;254;228;208m[+] Clean.\033[0m\n"&&sleep 1s
	@rm -rfv $(O)
	@printf "\033[1;38;2;254;228;208m    .^.   .^.\n"
	@printf "    /⋀\\_ﾉ_/⋀\\ \n"
	@printf "   /ﾉｿﾉ\\ﾉｿ丶)|\n"
	@printf "  |ﾙﾘﾘ >   )ﾘ\n"
	@printf "  ﾉノ㇏ Ｖ ﾉ|ﾉ\n"
	@printf "        ⠁⠁\n"
	@printf "\033[1;38;2;254;228;208m[*] Cleaned Up.\033[0m\n"
gnu-dev :
	sudo mkdir -pv /data/data/com.termux/files
	sudo mkdir -pv /data/data/com.termux/files/usr/bin
	sudo mkdir -pv /data/data/com.termux/files/usr/etc
	sudo mkdir -pv /data/data/com.termux/files/usr/var
	sudo mkdir -pv /data/data/com.termux/files/home
	sudo ln -svf /usr/bin/bash /data/data/com.termux/files/usr/bin/bash
	sudo chmod -v 777 /data/data/com.termux/files/usr/bin/bash
	sudo chmod -Rv 777 /data
	cp -frav `realpath .` /data/data/com.termux/files/home/
	printf "\033[1;38;2;254;228;208m[*] Go to /data/data/com.termux/files/home/termux-container for dev environment\n"
help :
	@echo "Available commands:"
	@echo "make build    :Just build"
	@echo "make install  :Build and install"
	@echo "make pack-deb :Build .deb package"
