# tfym-error-helper
**TFYM** is probably not your *solution* but a **friend** you are looking for to find it together.

- Tfym uses Gemini API, so you will need to have an API key.


# Installing 

Clone the repo, compile, move the executable to usr/local/bin/, put your API key and you are good to **go**!.


```bash 
git clone https://github.com/Ceyhun-Tufan/tfym-error-helper.git
```
```bash 
cd tfym-error-helper
```
```bash 
go build .
```
```bash 
sudo mv ./tfym /usr/local/bin
```
```bash 
mkdir -p ~/.config/tfym-cli
```
```bash 
echo 'API_KEY=your_key_here' > ~/.config/tfym-cli/.env
```
```bash 
chmod 600 ~/.config/tfym-cli/.env
```

# Usage
 
> tfym explain me what is a segment fault


TFYM:  

>  Aight, homie, listen up. A segmentation fault is basically when your program tries to access memory it ain't got permission to touch, like crashing a party you weren't invited to. Boom, error!
