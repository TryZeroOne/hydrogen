# Hydrogen Project

Hydrogen is a utility to protect your code.  
Hydrogen uses [garble](https://github.com/burrowers/garble) for compilation and [upx](https://github.com/upx/upx) for compression


# Installation:
```
go install github.com/TryZeroOne/hydrogen@latest
```

# Usage:
```
hydrogen <BUILD/PROTECT> <FILE TO PROTECT/OUTPUT FILE> -f <FLAGS>
```

### Flags:

<details><summary>Build Flags</summary>

```
  o   Obfuscate files                   
  j   Junk code                         
  g   Garbage                        
  p   Custom entry point
  c   Compress
  e   Encrypt
```

</details>

<details><summary>Protect Flags</summary>

```
  c   Compress
  e   Encrypt
  g   Garbage
```

</details>

### Examples:

<details><summary>Build Examples</summary>

```
hydrogen build build.exe
hydrogen build build.exe -f jgc
```

</details>

<details><summary>Protect Examples</summary>

```
hydrogen protect f.bin	
hydrogen protect f.bin -f cg
```

</details>


# Donations:
<details><summary>Ton</summary>
EQAmUr0NqEz6nnfUc2GeeGbUhOmd7Wh1zvIVQWWdj_MN6wlY

</details>
<details><summary>Litecoin</summary>
LMtj3jCFjgvDSCP1jqoE5AdbSbSevVxRJg
</details>
<details><summary>Monero</summary> 
429o1bxqyhs83hozpwbEZJitPcX8W73Nz86YRvyiWFkHAfnMk2ZA1VjeNnduKLKcFw45U2VAsQTFs7S5Ac1E16roKhnP777

</details>
