start saturatedfat.exe

$reg = "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings"
Set-ItemProperty -Path $reg -Name ProxyServer -Value "localhost:8080"
Set-ItemProperty -Path $reg -Name ProxyEnable -Value 1
