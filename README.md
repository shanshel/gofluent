# Go Fluent
##### Extremely simple and lightweight package for multi language support to use in your app

Why ? 
- I don't like the candies that comes with most of i18n packages 
- I want something with very simple interface can be used in any API application with no hessle


## Installation
```golang
go get github.com/shanshel/gofluent
```
## How to use
Create en.yaml anywhere in your project
```yaml
try_again: Something went wrong, please try again later 
profile:
    greeting: Hello %s 
    short: My name is %s and I'm %d years old
```
Note: make sure that all the keys are in lowercase

now in your go file
```golang
func anything() {
    languages := gofluent.Lang{}
    dir, _ := filepath.Abs("./langs")
    _ = languages.Setup(dir, "en", true) //true = will preload all langagues from yaml files inside the dir
    
    //example usage
    languages.Switch("en") //you can switch to other lang using it (also it will load the file if it's not loaded yet)
    
    languages.Get("", "try_again")
    languages.Get("profile", "greeting", "Alex")
    languages.Get("profile", "short", "Alex", 28)
}
```
Ideally you want to make languages variable as public. 
and don't forget to check for errors, I Ignored them to make the examples more readable.
