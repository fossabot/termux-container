[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FMoe-hacker%2Ftermux-container.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FMoe-hacker%2Ftermux-container?ref=badge_shield)

<p align="center">May The Force Be With You</p>
<p align="center">(>_×)</p>         

-----------  
<p align="center">
    <img src="https://stars.medv.io/Moe-hacker/termux-container.svg", title="commits" width="50%"/>
</p>
<p align="center">
 「この胸が脈打つうちは君をまだ守っていたい」
 </p>
 <p align="center">
    &emsp;
 </p>
<p align="center">
Run linux with chroot&unshare/proot on your Android phone,safely and easy 
</p>

![](https://img.shields.io/github/stars/Moe-hacker/termux-container?style=for-the-badge&color=fee4d0&logo=instatus&logoColor=fee4d0)
![](https://img.shields.io/github/forks/Moe-hacker/termux-container?style=for-the-badge&color=fee4d0&logo=git&logoColor=fee4d0)
![](https://img.shields.io/github/license/Moe-hacker/termux-container?style=for-the-badge&color=fee4d0&logo=apache&logoColor=fee4d0)
![](https://img.shields.io/github/repo-size/Moe-hacker/termux-container?style=for-the-badge&color=fee4d0&logo=files&logoColor=fee4d0)
![](https://img.shields.io/github/last-commit/Moe-hacker/termux-container?style=for-the-badge&color=fee4d0&logo=codeigniter&logoColor=fee4d0)
![](https://img.shields.io/badge/language-shell\&c-green?style=for-the-badge&color=fee4d0&logo=sharp&logoColor=fee4d0)

You can read this doc in :

**[<kbd> <br> 简体中文 <br> </kbd>](README-ZH.md)**&emsp;**[<kbd> <br> English <br> </kbd>](README.md)**&emsp;**[<kbd> <br> Portuguese <br> </kbd>](README-PT.md)**&emsp;

### WARNING:      
```
* Your warranty is now void.
* I am not responsible for anything that may happen to your device by using this script.
* You do it at your own risk and take the responsibility upon yourself.
```
### About:      
Termux-container is a script that runs Linux containers on your android phone. It's safe, easy to use and well-designed.      
### About ruri:      
ruri runs a linux container with unshare namespaces and dropped capabilities in your system.      
### Why does its development seem so slow？      
The existing chroot/unshare command didn't meet the needs of the new version termux-container, so I needed to rewrite a more powerful container implementation. [ruri v2.0](https://github.com/Moe-hacker/ruri) is coming, which will bring a more free and secure container implementation, and termux-container needs to wait for it to finish to continue development.      
It's like `chroot` and `unshare`, but it can reduce the container's capability set for more security.       
It's published under the MIT license, everyone is free to use it.      
### WARNING:      
Here's only the new version of termux-container and it might not be stable.      
For any bugs, please report at Discussions.      
### Getting start:      
The new version hasn't been released yet, if you want to try it, please remove the old version and run:      
```sh
git clone https://github.com/Moe-hacker/termux-container
cd termux-container
pkg install make
make
make install
```
### Usage:     
See `container -h` and `help` in container console.
### About me:            
A noob, want to be a geek.            
If my work helps, please give me a star !             

--------
<p align="center">「生きてて良かった、そんなこと思える日を,</p>
<p align="center">願ってしまうんだ」</p>         


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FMoe-hacker%2Ftermux-container.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FMoe-hacker%2Ftermux-container?ref=badge_large)