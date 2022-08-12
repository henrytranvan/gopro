# gopro
1. Go is very simple language, it doesn't have many built in functions, keywords...
2. If no initial value is given to a variable, the Go compiler will automatically initialise that variable to the zero value of its data type
3. Go functions can return multiple values.
4. There is only one way to format curly braces in Go.
5. There is a global Go rule that also applies to function and variable names and is valid for all packages except main: everything that begins with a lowercase letter is considered private and is accessible in the current package only
6. The official name for := is short assignment statement and it is very frequently used in Go, especially for getting the return values from functions and for loops with the range keyword
7. Every statement that exists outside of the code of a function must begin with a keyword such as func or var
8. When you pass an array to a function, what is happening is that Go creates a copy of that array and passes that copy to that function—therefore any changes you make to an array inside a function are lost when the function returns.
9. Array and slide alway indexes by positive integer, map can indexes by string but value always one type aMap : = map([string] int), structure can indexes by string and values are multiple types(same records at C), structure need to define a type, then use new(aStructureType) to create a new structure value. A slide of structures are popular data type
10. aMap : = make(map[string] string),  aSlide = make([]int), use make to initialise var in memory
11. Array, slide and map use square bracket[] but structure use curly bracket {} Ex, states["WA"] = "Washington"
12. Go use '\n'  for a character but "This is a string" for string same C language.
13. No order in Map, hash table, associate table, key value table Ex. for k,v := range states { ... }
14. len() func use to array, slide, map, string ...
15. initial character Var, func type lowercase mean private, non export, same private property, method on java, c++, Uppercase mean public
16. type User struct { }
17. It usually convert map key to slide for sorting map by key
18. Do not need break after each switch case after 1 case is true, Go will break out of the Switch statement.
19. There is no While statement in Go
20. defer - run others statements, before finish the func run statement in defer wait to any thing else has been done, implement this command -  defer file.Close() - file, err := os.Create("./fileString.txt")
21. if err != nil { panic(err) } will stop and get out of this application
22. go use git to install third party packages.
23. main package and main function is special meaning in go, it entry point to application
24. Use func in same package but at different files do not need to import
25. []string(var)  []byte(var) string(var) int(var) ... convert type, type cast
26. alex := person{firstName: “Alex”, lastName: “Anderson”}  Struct data
27. Io.Exit(1) exit app with error.
28. https://github.com/StephenGrider/GoCasts   https://github.com/mactsouk/mastering-Go-3rd https://github.com/PacktPublishing/
