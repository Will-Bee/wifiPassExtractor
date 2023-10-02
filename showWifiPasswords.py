
try:
    from colorama import init, Fore
    import subprocess
except:
    import os
    os.system('pip3 install colorama')
    os.system('pip3 install subprocess')
    try:
        from colorama import init, Fore
        import subprocess
    except:
        print('[-] Failed to install dependencies')
        print('[-] Please install dependencies manually: pip3 install colorama subprocess{Fore.RESET}')
        exit()



failList = []
init(convert=True)

print(Fore.GREEN + "Wifi passwords saved in this computer")
print(" ")

data = subprocess.check_output(['netsh', 'wlan', 'show', 'profiles']).decode('cp1252', errors="backslashreplace").split('\n')
profiles = [i.split(":")[1][1:-1] for i in data if "All User Profile" in i]
for i in profiles:
    try:
        results = subprocess.check_output(['netsh', 'wlan', 'show', 'profile', i, 'key=clear']).decode('iso8859-1', errors="backslashreplace").split('\n')
        results = [b.split(":")[1][1:-1] for b in results if "Key Content" in b]
        try:
            print(Fore.CYAN + "{:<30}|  {:<}".format(i, results[0]))
        except IndexError:
            print(Fore.GREEN + "{:<30}|  {:<}".format(i, "[NO PASS]"))
    except subprocess.CalledProcessError:
        print(Fore.RED + "{:<30}|  {:<}".format(i, "ENCODING ERROR"))
        failList.append(i)


print(f"\n{Fore.GREEN}SUCCESS: {len(profiles) - len(failList)}/{len(profiles)}")

print(f"\n{Fore.RED}ENCODING ERRORS: {Fore.YELLOW}{len(failList)}")
print(f"{Fore.RED}ENCODING ERRORS: {Fore.YELLOW}{failList}")

print(Fore.RESET + "")
